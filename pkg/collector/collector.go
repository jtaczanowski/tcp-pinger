package collector

import (
	"time"

	"github.com/jtaczanowski/tcp-pinger/pkg/config"
	"github.com/jtaczanowski/tcp-pinger/pkg/models"
)

func Start(config *config.Config, pingerToCollectorChan chan models.Ping, collectorToAggregatorChan chan models.PingsCollection) {
	collection := models.PingsCollection{}
	ticker := time.NewTicker(time.Second * time.Duration(config.GraphiteIntervalSecond))
	for {
		select {
		case ping := <-pingerToCollectorChan:
			collection = append(collection, ping)
		case <-ticker.C:
			collectorToAggregatorChan <- collection
			collection = models.PingsCollection{}
		}
	}

}
