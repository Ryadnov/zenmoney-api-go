package entities

import (
	"time"
)

type (
	Timestamp string
	ObjectId string

	InstrumentId int
	CompanyId int
	UserId int
)

func GetCurrentTimestamp() Timestamp {
	return Timestamp(time.Now().UTC().Format("20060102150405"))
}