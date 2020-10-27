package sender

import (
	"log"

	"github.com/jtaczanowski/go-graphite-client"
	"github.com/jtaczanowski/tcp-pinger/pkg/config"
	"github.com/jtaczanowski/tcp-pinger/pkg/models"
)

func Start(config *config.Config, aggregatorToSenderChan chan models.SenderCollection) {
	for senderCollection := range aggregatorToSenderChan {
		if config.GraphiteHost != "" {
			graphiteClient := graphite.NewClient(config.GraphiteHost, config.GraphitePort, config.GraphitePrefix, config.GraphiteProtocol)
			if err := graphiteClient.SendData(senderCollection); err != nil {
				log.Printf("Error sending metrics: %v", err)
			}
		}
	}
}
