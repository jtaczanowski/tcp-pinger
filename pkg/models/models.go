package models

type Ping struct {
	Host        string
	TimeElapsed float64
	Err         error
}

type PingsCollection []Ping

type HostPingCollection struct {
	TimeElapsed []float64
	Err         []error
}

type HostPingsCollection map[string]*HostPingCollection

type SenderCollection map[string]float64
