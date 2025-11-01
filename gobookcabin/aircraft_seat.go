package gobookcabin

import (
	"fmt"
	"strconv"
)

const (
	AircraftSeatTableName = "aircraft_seat"
)

type AircraftSeat struct {
	BaseEntity
	AircraftType string `gorm:"not null"`
	Row          int    `gorm:"not null"`
	SeatsInRow   string `gorm:"not null"`
}

func (AircraftSeat) TableName() string {
	return AircraftSeatTableName
}

type SingleAircraftSeat struct {
	BaseEntity
	AircraftType string `gorm:"not null"`
	Row          int    `gorm:"not null"`
	SeatAlphabet string `gorm:"not null"`
}

func (SingleAircraftSeat) TableName() string {
	return AircraftSeatTableName
}

func (s SingleAircraftSeat) GetSeat() string {
	rowStr := strconv.Itoa(s.Row)
	return fmt.Sprintf("%s%s", rowStr, s.SeatAlphabet)
}
