package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const EXAMPLE_CONFIG = `
{
	"GraphiteHost": "localhost",
	"GraphitePort": 2003,
	"GraphiteProtocol": "tcp",
    "GraphiteIntervalSecond": 60,
    "GraphitePrefix": "tcp-pinger",
	"AppendHostnameToGraphitePrefix": true,
	"ShowPingsOnConsole": true,
    "Hosts": [
        {
            "Host": "google.com:443",
            "TimeoutSecond": 10,
            "CheckIntervalMiliSecond": 1000
        },
        {
            "Host": "bing.com:443",
            "TimeoutSecond": 10,
            "CheckIntervalMiliSecond": 1000
        }
    ]
}
`

type HostConfig struct {
	Host                    string
	TimeoutSecond           int
	CheckIntervalMiliSecond int
}

type HostsConfigs []HostConfig

type Config struct {
	GraphiteHost                   string
	GraphitePort                   int
	GraphiteProtocol               string
	GraphiteIntervalSecond         int
	GraphitePrefix                 string
	AppendHostnameToGraphitePrefix bool
	ShowPingsOnConsole             bool
	Hosts                          HostsConfigs
}

func GetConfig() (*Config, error) {
	name, err := os.Executable()
	if err != nil {
		return nil, fmt.Errorf("Can't get current executable path: %s\nFor exit ctrl+c", err)
	}
	dir, filename := filepath.Split(name)
	filenameWithoutExtension := strings.Split(filename, "\\.")[0]

	c := flag.String("c", dir+filenameWithoutExtension+"-config"+".json", "Specify the configuration file.")
	flag.Parse()
	filePath, err := filepath.Abs(*c)
	if err != nil {
		return nil, fmt.Errorf("Can't open config file: %s\nExample configuration file:\n %s", err, EXAMPLE_CONFIG)
	}
	log.Printf("Reading configuration file from path: %s", filePath)
	file, err := os.Open(*c)
	if err != nil {
		return nil, fmt.Errorf("Can't open config file: %s\nExample configuration file:\n %s", err, EXAMPLE_CONFIG)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := &Config{}
	err = decoder.Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("Can't decode config JSON: %s\nExample configuration file:\n %s", err, EXAMPLE_CONFIG)
	}
	if config.AppendHostnameToGraphitePrefix {
		config.GraphitePrefix = config.GraphitePrefix + "." + getHostname()
	}
	return config, nil
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown-hostname"
	}
	return strings.ReplaceAll(strings.ReplaceAll(hostname, ":", "-"), ".", "_")
}
