package model

type Schedule struct {
	ID           string `json:"id"`            // uuid
	MaxAvailable uint   `json:"max_available"` // eg) 100 people can reserve at 9:20
	Stock        uint   `json:"stock"`         // eg) 100 - 20 people reserved = 80
	ScheduleDate
}

type ScheduleDate struct {
	Year  uint `json:"year"`
	Month uint `json:"month"`
	Day   uint `json:"day"`
	Hour  uint `json:"hour"`
	Min   uint `json:"min"` // eg) hour 9, min 20 is 9:20
}

// 管理者による予約枠登録
func NewSchedule(date *ScheduleDate, maxAvailable uint, uuidGen UUIDGenerator) *Schedule {
	return &Schedule{
		ID:           uuidGen.New().String(),
		MaxAvailable: maxAvailable,
		Stock:        maxAvailable,
		ScheduleDate: *date,
	}
}
