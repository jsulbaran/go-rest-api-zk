package service

import (
	"RestService/domain"
	"github.com/jinzhu/gorm"
	"time"
)

type EventService struct {
	db *gorm.DB
}

func NewEventService(orm *gorm.DB) *EventService {
	return &EventService{db: orm}
}

func (service EventService) GetEvents(fromDate time.Time, toDate time.Time) []domain.Event {
	//fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
	//	t.Year(), t.Month(), t.Day(),
	//	t.Hour(), t.Minute(), t.Second())
	//service.db.Debug().Preload("Biometric").Find(&events)
	var events []domain.Event
	service.db.Debug().Where("Verify_Time BETWEEN ? AND ?", formatTime(fromDate), formatTime(toDate)).Find(&events)
	return events
}

func formatTime(date time.Time) string {
	return date.Format("2006-01-02T15:04:05")
}
