package json_response

import (
	"RestService/domain"
	"RestService/util"
	"strings"
)

func NewJsonEvent(event domain.Event, deviceSerial string) EventJson {
	result := EventJson{}
	result.UserId = util.StringToInt(event.UserId)
	//result.EventCode=
	result.EventDate = strings.Replace(event.EventDateTime[:10], "-", "", -1)
	result.EventTime = strings.Replace(event.EventDateTime[11:], ":", "", -1)
	if event.Status == 0 || event.Status == 1 {
		result.EventCode = "01010102"
	} else {
		result.EventCode = "01020102"
	}
	result.DeviceId = deviceSerial

	result.FuncKey = getFunctionKey(event.Status)

	return result
}
func getFunctionKey(status int) string {
	funckey := "255"
	if status == 0 {
		funckey = "10"
	} else if status == 1 {
		funckey = "30"
	}

	return funckey
}
