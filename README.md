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

# Refs
- tutorialedge.net/golang/go-websocket-tutorial
- tutorialedge.net/golang/parsing-xml-with-golang
- golangexample.com/a-powerful-url-router-and-dispatcher-for-golang
- dev.to/koddr/let-s-write-config-for-your-golang-web-app-on-right-way-yaml-5ggp
- github.com/MasslessParticle/GoQ
- github.com/vmihailenco/taskq
- medium.com/@alain.drolet.0/how-to-unmarshal-an-array-of-json-objects-of-different-types-into-a-go-struct-10eab5f9a3a2
- golangcode.com/json-encode-an-array-of-objects/
- github.com/tidwall/gjson
- github.com/mattn/go-sqlite3/blob/master/_example/simple/simple.go
- appdividend.com/2019/12/02/golang-http-example-get-post-http-requests-in-golang
- ednsquare.com/story/date-and-time-manipulation-golang-with-examples------cU1FjK
