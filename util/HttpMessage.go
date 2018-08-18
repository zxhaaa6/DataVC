package util

type SimpleMessage struct {
	Success bool
	Status  int
	Message string
	Data    interface{}
}

func GenSimpleJson(data interface{}) SimpleMessage {
	return SimpleMessage{Success: true, Status: 200, Message: "success", Data: data}
}

func GenErrorJson(err Error) SimpleMessage {
	return SimpleMessage{Success: false, Status: err.Code, Message: err.Message}
}
