package jobs

import "log"

const (
	GetLogJobInterval = "0 0 * * * *"
)

func GetLogJob() {
	log.Println("=======Cron function========")
}
