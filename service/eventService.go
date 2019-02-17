package service

import (
	"RestService/config"
	"RestService/domain"
	"RestService/json_response"
	"github.com/jinzhu/gorm"
	"time"
)

type EventService struct {
	db            *gorm.DB
	configuration *config.Config
}

func NewEventService(orm *gorm.DB, configuration *config.Config) *EventService {
	return &EventService{db: orm, configuration: configuration}
}

func (service EventService) GetEvents(fromDate time.Time, toDate time.Time) []json_response.EventJson {
	//fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
	//	t.Year(), t.Month(), t.Day(),
	//	t.Hour(), t.Minute(), t.Second())
	//service.db.Debug().Preload("Biometric").Find(&events)
	var events []domain.Event
	service.db.Debug().Where("Verify_Time BETWEEN ? AND ?", formatTime(fromDate), formatTime(toDate)).Find(&events)
	return convert(events, service.configuration.DeviceSerial)
}

func convert(events []domain.Event, deviceSerial string) (ret []json_response.EventJson) {
	for _, event := range events {
		ret = append(ret, json_response.NewJsonEvent(event, deviceSerial))
	}
	return
}

func formatTime(date time.Time) string {
	return date.Format("2006-01-02T15:04:05")
}
