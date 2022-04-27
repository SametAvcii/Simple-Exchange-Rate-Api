package Cron_Job

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"github.com/sametavcii/Exchange-Rate-Api/Api"
	"time"
)

func GetRequest() {
	url := "https://v6.exchangerate-api.com/v6/8384bc788cb52678281ba006/latest/USD"
	Api.GetRequest(url)
}
func Log() {
	fmt.Println(time.Now())
}

func Cron() {

	s := gocron.NewScheduler()
	s.Every(2).Second().Do(Log)
	s.Every(2).Second().Do(GetRequest)
	<-s.Start()

}
