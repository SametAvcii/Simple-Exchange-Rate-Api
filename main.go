package main

import (
	Cron_Job "github.com/sametavcii/Exchange-Rate-Api/Cron-Job"
	"github.com/sametavcii/Exchange-Rate-Api/RabbitMq"
)

func main() {

	Cron_Job.Cron()
	RabbitMq.Send()
	RabbitMq.Receive()
}
