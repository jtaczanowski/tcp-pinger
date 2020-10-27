package pinger

import (
	"log"
	"net"
	"time"

	"github.com/jtaczanowski/tcp-pinger/pkg/config"
	"github.com/jtaczanowski/tcp-pinger/pkg/models"
)

func Start(conf *config.Config, pingerToCollectorChan chan models.Ping) {
	for _, hostConfig := range conf.Hosts {
		log.Println("Start tcp pinging host: " + hostConfig.Host)
		go func(hostConfig config.HostConfig) {
			for range time.Tick(time.Millisecond * time.Duration(hostConfig.CheckIntervalMiliSecond)) {
				go func(hostConfig config.HostConfig) {
					var ping models.Ping
					timeStart := time.Now()
					conn, err := net.DialTimeout("tcp", hostConfig.Host, time.Second*time.Duration(hostConfig.TimeoutSecond))
					ping.TimeElapsed = float64(time.Since(timeStart).Microseconds()) / 1000
					ping.Host = hostConfig.Host
					if err != nil {
						log.Printf("Tcp ping for host: %s error: %s time elapsed: %vms\n", hostConfig.Host, err, ping.TimeElapsed)
						ping.Err = err
						pingerToCollectorChan <- ping
						return
					}
					ping.Host = hostConfig.Host
					conn.Close()
					if conf.ShowPingsOnConsole {
						log.Printf("Tcp ping for host: %s time elapsed: %vms\n", hostConfig.Host, ping.TimeElapsed)
					}
					pingerToCollectorChan <- ping
				}(hostConfig)
			}
		}(hostConfig)
	}
}
