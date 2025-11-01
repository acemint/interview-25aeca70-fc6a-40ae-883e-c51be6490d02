package gobookcabin

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type GormVoucherService struct {
	db *gorm.DB
}

func NewGormVoucherService(db *gorm.DB) *GormVoucherService {
	return &GormVoucherService{db: db}
}

func (s *GormVoucherService) Check(ctx context.Context, request *CheckVoucherRequest) (*Voucher, error) {
	var voucher Voucher
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.Where("flight_date = ? AND flight_number = ?", request.FlightDate, request.FlightNumber).
			First(&voucher)
		if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return Errorf(ErrCodeInternal, "failed to get voucher data")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &voucher, nil
}

func (s *GormVoucherService) Generate(ctx context.Context, request *GenerateVoucherRequest) (*Voucher, error) {
	var voucher Voucher
	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var existingVoucher Voucher
		result := tx.Where("flight_date = ? AND flight_number = ?", request.FlightDate, request.FlightNumber).
			First(&existingVoucher)
		if result.RowsAffected == 1 {
			return Errorf(ErrCodeInvalid, "voucher has been generated")
		}
		if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return Errorf(ErrCodeInternal, "failed to obtain voucher data")
		}

		seatsSelection := make([]SingleAircraftSeat, SeatAmount)
		result = tx.Raw(`
			SELECT 
				aircraft_type,
				row,
				json_each.value as seat_alphabet
			FROM aircraft_seat,
			json_each(seats_in_row)
			WHERE aircraft_type = ?
			ORDER BY RANDOM()
			LIMIT ?
		`, request.AircraftType, SeatAmount).Scan(&seatsSelection)
		if result.Error != nil || result.RowsAffected != SeatAmount {
			return Errorf(ErrCodeInternal, "failed to find aircraft seats")
		}

		voucher = Voucher{
			CrewName:     request.CrewName,
			CrewID:       request.CrewID,
			FlightNumber: request.FlightNumber,
			FlightDate:   request.FlightDate,
			AircraftType: request.AircraftType,
			Seat1:        seatsSelection[0].GetSeat(),
			Seat2:        seatsSelection[1].GetSeat(),
			Seat3:        seatsSelection[2].GetSeat(),
		}
		result = tx.Create(&voucher)
		if result.Error != nil {
			return Errorf(ErrCodeInternal, "failed to insert voucher data")
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return &voucher, nil
}
