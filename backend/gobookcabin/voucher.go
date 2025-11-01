package gobookcabin

const (
	SeatAmount       = 3
	VoucherTableName = "voucher"
)

type Voucher struct {
	BaseEntity
	CrewName     string `gorm:"not null"`
	CrewID       string `gorm:"not null"`
	FlightNumber string `gorm:"not null"`
	FlightDate   string `gorm:"not null"`
	AircraftType string `gorm:"not null"`
	Seat1        string
	Seat2        string
	Seat3        string
}

func (Voucher) TableName() string {
	return VoucherTableName
}

type GenerateVoucherRequest struct {
	CrewName     string `json:"name" validate:"required"`
	CrewID       string `json:"id" validate:"required"`
	FlightNumber string `json:"flightNumber" validate:"required"`
	FlightDate   string `json:"date" validate:"required,datetime=2006-01-02"`
	AircraftType string `json:"aircraft" validate:"required"`
}

type CheckVoucherRequest struct {
	FlightNumber string `json:"flightNumber" validate:"required"`
	FlightDate   string `json:"date" validate:"required,datetime=2006-01-02"`
}

type GenerateVoucherResponse struct {
	Success bool     `json:"success"`
	Seats   []string `json:"seats"`
}

func NewGenerateVoucherResponse(voucher *Voucher) *GenerateVoucherResponse {
	seats := make([]string, 0)
	seats = append(seats, voucher.Seat1, voucher.Seat2, voucher.Seat3)
	return &GenerateVoucherResponse{
		Success: true,
		Seats:   seats,
	}
}

type CheckVoucherResponse struct {
	Exists bool `json:"exists"`
}

func NewCheckVoucherResponse(voucher *Voucher) *CheckVoucherResponse {
	return &CheckVoucherResponse{
		Exists: voucher.ID != 0,
	}
}
