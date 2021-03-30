package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
	"github.com/nj-eka/gb_go_1/lesson1/greeting"
)

//https://www.pool.ntp.org/zone/ru
const ntpMSKHost = "ntp1.stratum1.ru"
const ntpUTCHost = "0.ru.pool.ntp.org"

func main() {
	fmt.Println("Привет! Сверим часы:")
	fmt.Println("Местное время: ", time.Now())
	ntpTime, err := ntp.Time(ntpMSKHost)
	if err != nil {
		log.Fatalf("Error: %s", err)
	} else {
		fmt.Println("Московское время: ", ntpTime)
	}
	response, err := ntp.Query(ntpUTCHost)
	if err != nil {
		log.Fatalf("Error: %s", err)
	} else {
		fmt.Printf("В Гринвиче: %v\nС уточнениями: %v\n", response.Time, time.Now().Add(response.ClockOffset))
	}
	err = response.Validate()
	if err == nil {
		fmt.Println("Кстати, возможна синхронизация.")
	}
	fmt.Println(greeting.SayGo())

}
