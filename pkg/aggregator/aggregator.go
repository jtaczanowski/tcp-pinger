package aggregator

import (
	"strings"

	"github.com/jtaczanowski/tcp-pinger/pkg/models"
)

func Start(graphitePrefix string, collectorToAggregatorChan chan models.PingsCollection, aggregatorToSenderChan chan models.SenderCollection) {
	for pingCollection := range collectorToAggregatorChan {
		collectionAggregated := models.HostPingsCollection{}
		senderCollection := models.SenderCollection{}
		for _, ping := range pingCollection {
			if _, ok := collectionAggregated[ping.Host]; !ok {
				collectionAggregated[ping.Host] = &models.HostPingCollection{}
			}
			hostPingCollection := collectionAggregated[ping.Host]
			hostPingCollection.TimeElapsed = append(hostPingCollection.TimeElapsed, ping.TimeElapsed)
			hostPingCollection.Err = append(hostPingCollection.Err, ping.Err)
		}
		for host, values := range collectionAggregated {
			var sumPingTime float64
			var avgPingTime float64
			var maxPingTime float64
			var errorsCount float64
			var pingCount float64
			for idx := range values.TimeElapsed {
				sumPingTime += values.TimeElapsed[idx]
				if maxPingTime < values.TimeElapsed[idx] {
					maxPingTime = values.TimeElapsed[idx]
				}
			}
			avgPingTime = sumPingTime / float64(len(values.TimeElapsed))
			for idx := range values.Err {
				if values.Err[idx] != nil {
					errorsCount += 1.0
				}
			}
			pingCount = float64(len(values.TimeElapsed))

			hostMetricsString := strings.ReplaceAll(strings.ReplaceAll(host, ":", "-"), ".", "_")
			senderCollection[hostMetricsString+"."+"ping-count"] = pingCount
			senderCollection[hostMetricsString+"."+"avg-pings-time"] = avgPingTime
			senderCollection[hostMetricsString+"."+"max-pings-time"] = maxPingTime
			senderCollection[hostMetricsString+"."+"errors"] = errorsCount

		}
		aggregatorToSenderChan <- senderCollection
	}
}
