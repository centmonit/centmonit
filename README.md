# CentMonit

A centralized Monit approach!

# Development

Build
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o centmonit
```

Deploy test
```
CONTAINER=test_monit
docker cp ./centmonit $CONTAINER:/home/
```

# XML Parsing

service::type
- 0 => "Filesystem"
- 1 => "Directory"
- 2 => "File"
- 3 => "Daemon" / "Process"
- 4 => "Connection" / "Host"
- 5 => "System"

action:
- 1: alert
- 3: stop
- 6: start

Base props
- name
- type
- status
- monitored
- collected_sec

Extend props:
- System
    - load
    - cpu
    - memory
    - swap
    - uptime
- Filesystem
    - percent
    - usage
    - total
- Process
    - pid
    - uptime
    - memory
    - cpu
- Host
    - response_time