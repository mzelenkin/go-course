package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

const ntpServer = "pool.ntp.org"

func main() {
	ntptime, err := ntp.Time(ntpServer)
	if err != nil {
		_, _ = os.Stderr.WriteString("NTP error: " + err.Error())
		os.Exit(1)
	}

	localtime := time.Now()

	fmt.Println(localtime.Format("Локальное время: 2 Jan 2006 15:04:05"))
	fmt.Println(ntptime.Format("Время NTP сервера: 2 Jan 2006 15:04:05"))
}
