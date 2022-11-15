package crontask

const (
	msgCronSucc = "success"
)

//{"status":"success"}
//{"data":{"message":"Backgroundjobs are using system cron!"},"status":"error"}
type CronMessage struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
