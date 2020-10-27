# tcp-pinger
tcp-pinger measures time of opening new TCP connection to remote host. Metrics are sent to Graphite.
<img src="/_example/tcp-pinger-grafana.png" alt="tcp-pinger" height="100%" width="100%">

## Example usage

### Run tcp-pinger
You can use already builded binaries from [bin](https://github.com/jtaczanowski/tcp-pinger/tree/main/bin) directory. Example configuration file is included in [_example](https://github.com/jtaczanowski/tcp-pinger/tree/main/_example) directory.

```
./tcp-pinger -c tcp-pinger-config.json
```

Example configuration file [tcp-pinger-config.json](https://raw.githubusercontent.com/jtaczanowski/tcp-pinger/main/_example/tcp-pinger-config.json):
```
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
```

When runing from default configuration tcp-pinger will show on console times of opening TCP connection:

<img src="/_example/tcp-pinger.png" alt="tcp-pinger" height="100%" width="100%">

It can be disabled from configuration file:

```
"ShowPingsOnConsole": false,
```

### Import Grafana dashboard
Example Grafana dashboard file [tcp-pinger-grafana-dashboard.json](https://raw.githubusercontent.com/jtaczanowski/tcp-pinger/main/_example/tcp-pinger-grafana-dashboard.json)

### To build new tcp-builder binaries
```make build```