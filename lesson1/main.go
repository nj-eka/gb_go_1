package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
	"github.com/nj-eka/gb_go_1/lesson1/greeting"
	"github.com/nj-eka/gb_go_1/lesson1/version"
)

//https://www.pool.ntp.org/zone/ru
const ntpMSKHost = "ntp1.stratum1.ru"
const ntpUTCHost = "0.ru.pool.ntp.org"

func main() {
	fmt.Printf("Привет! Это версия [%s] коммита [%s] от [%s]\n", version.Version, version.Commit, version.BuildTime)
	fmt.Println("Сверим наши часы: \n	- местное время: ", time.Now())
	ntpTime, err := ntp.Time(ntpMSKHost)
	if err != nil {
		log.Fatalf("Error: %s", err)
	} else {
		fmt.Println("	- Московское время: ", ntpTime)
	}
	response, err := ntp.Query(ntpUTCHost)
	if err != nil {
		log.Fatalf("Error: %s", err)
	} else {
		fmt.Printf("	- в Гринвиче: %v\n	- с поправкой: %v\n", response.Time, time.Now().Add(response.ClockOffset))
	}
	err = response.Validate()
	if err == nil {
		fmt.Println("Кстати, возможна синхронизация.")
	}
	fmt.Println(greeting.SayGo())
}
