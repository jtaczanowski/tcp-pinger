package factory

import (
	"log"

	"github.com/jtaczanowski/tcp-pinger/pkg/aggregator"
	"github.com/jtaczanowski/tcp-pinger/pkg/collector"
	"github.com/jtaczanowski/tcp-pinger/pkg/config"
	"github.com/jtaczanowski/tcp-pinger/pkg/models"
	"github.com/jtaczanowski/tcp-pinger/pkg/pinger"
	"github.com/jtaczanowski/tcp-pinger/pkg/sender"
)

type App struct {
	config                    *config.Config
	pingerToCollectorChan     chan models.Ping
	collectorToAggregatorChan chan models.PingsCollection
	aggregatorToSenderChan    chan models.SenderCollection
}

func (app *App) Initialize() {
	var err error
	app.config, err = config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	app.pingerToCollectorChan = make(chan models.Ping, 10)
	app.collectorToAggregatorChan = make(chan models.PingsCollection, 10)
	app.aggregatorToSenderChan = make(chan models.SenderCollection, 10)
}

func (app *App) Run() {
	pinger.Start(app.config, app.pingerToCollectorChan)
	go collector.Start(app.config, app.pingerToCollectorChan, app.collectorToAggregatorChan)
	go aggregator.Start(app.config.GraphitePrefix, app.collectorToAggregatorChan, app.aggregatorToSenderChan)
	sender.Start(app.config, app.aggregatorToSenderChan)
}
