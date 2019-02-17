package json_response

type EventJson struct {
	ID        int    `json:"-"`
	UserId    int    `json:"user_id"`
	EventDate string `json:"event_date"`
	EventTime string `json:"event_time"`
	EventCode string `json:"event_code"`
	DeviceId  string `json:"device_id"`
	FuncKey   string `json:"func_key"`
}
