package core

import (
	"fmt"
	"encoding/xml"
	"golang.org/x/net/html/charset"
	"bytes"
	"time"
	// "os"
	"log"
	"github.com/olekukonko/tablewriter"
)

type MonitInst struct {
	XMLName xml.Name `xml:"monit"`
	ID string `xml:"id,attr"`
	Server Server `xml:"server"`
	Services Services `xml:"services"`
	Event Event `xml:"event"`
}

type Server struct {
	XMLName xml.Name `xml:"server"`
	Uptime uint `xml:"uptime"`
	Poll uint `xml:"poll"`
	Hostname string `xml:"localhostname"`
}

type Services struct {
	XMLName xml.Name `xml:"services"`
	ServiceArr []Service `xml:"service"`
}

type Service struct {
	XMLName xml.Name `xml:"service"`
	Name string `xml:"name,attr"`
	Type uint `xml:"type"`
	Monitor uint `xml:"monitor"`
	CollectedSec uint `xml:"collected_sec"`
	Status uint  `xml:"status"`
	// MonitorMode uint `xml:"monitormode"`

	// Type 5
	System System `xml:"system"`
}

type System struct {
	XMLName xml.Name `xml:"system"`
	Cpu Cpu `xml:"cpu"`
	Memory Memory `xml:"memory"`
}

type Cpu struct {
	XMLName xml.Name `xml:"cpu"`
	User float32 `xml:"user"`
	System float32 `xml:"system"`
}

type Memory struct {
	XMLName xml.Name `xml:"memory"`
	Percent float32 `xml:"percent"`
}

type Event struct {
	XMLName xml.Name `xml:"event"`
	Service string `xml:"service"`
	Type uint `xml:"type"`
	ID string `xml:"id"`
	State uint `xml:"state"`
	Action uint `xml:"action"`
	Message string `xml:"message"`
}

func descServiceType (serviceType uint) string {
	switch serviceType {
	case 1:
		return "Directory"
	case 2:
		return "File"
	case 3:
		return "Daemon"
	case 4:
		return "Connection"
	case 5:
		return "System"
	default:
		return "Unknown"
	}
}

func descMonitorStatus (status uint) string {
	switch status {
	case 1:
		return "MONITORED"
	case 0:
		return "UNMONITORED"
	case 2:
		return "INIT"
	default:
		log.Printf("WARN\tMonitor status [%d] is unknown\n", status)
		return "__UNK__"
	}
}

func descServiceStatus (status uint) string {
	switch status {
	case 0:
		return "OK"
	case 32:
		return "Connection failed"
	case 1073741824:
		return "Does exist"
	default:
		log.Printf("WARN\tService status [%d] is unknown\n", status)
		return "NOT OK"
	}
}

func descTimestamp (ts int64) time.Time {
	tm := time.Unix(ts, 0)
	return tm
}

func _testPrint1(monitInst MonitInst) {
	fmt.Printf("Monit inst [%s] - Host server [%s]\n", monitInst.ID, monitInst.Server.Hostname)
	var cpu, memory float32
	for i := 0; i < len(monitInst.Services.ServiceArr); i++ {
		tmpService := monitInst.Services.ServiceArr[i]
		fmt.Printf("* Service [%s]\n", tmpService.Name)

		fmt.Printf("\t- Type: %s\n", descServiceType(tmpService.Type))
		fmt.Printf("\t- Monitor: %s\n", descMonitorStatus(tmpService.Monitor))
		fmt.Printf("\t- CollectedSec: %s\n", descTimestamp(int64(tmpService.CollectedSec)))
		fmt.Printf("\t- Status: %s\n", descServiceStatus(tmpService.Status))

		if tmpService.Type == 5 {
			memory = tmpService.System.Memory.Percent
			cpu = tmpService.System.Cpu.User + tmpService.System.Cpu.System
		}
	}
	fmt.Printf("MEMORY %.2f%% - CPU %.2f%%\n", memory, cpu)
}

func processReportEvent(monitInst MonitInst, monitHostsMap *map[string]MonitHost) {
	log.Printf("INFO\tNew event - Service: %s - ID: %s - STATE: %d - ACTION: %d - MSG: %s\n",
		monitInst.Event.Service,
		monitInst.Event.ID,
		monitInst.Event.State,
		monitInst.Event.Action,
		monitInst.Event.Message)

	eventType := "error"
	if monitInst.Event.ID == "65536" {
		if monitInst.Event.Action == 6 {
			// monit started event
			eventType = "success"
		} else /* action=3 */ {
			// monit stopped event
			eventType = "error"

			// clear host report
			(*monitHostsMap)[monitInst.ID] = MonitHost{
				ID: monitInst.ID,
				Hostname: monitInst.Server.Hostname,
				AlertMessage: monitInst.Event.Message,
			}
			// log.Println("Item after clear host: ", (*monitHostsMap)[monitInst.ID])
			// log.Printf("Map after clear host: %+v", (*monitHostsMap))
			// log.Println("Param pointer address: ", monitHostsMap)
		}
	} else if monitInst.Event.ID == "131072" {
		// Monitor/Unmonitor event
		eventType = "warning"
	} else if monitInst.Event.State == 0 {
		eventType = "success"
	}

	if len(socketConnections) > 0 {
		socketMsg := SocketEventMessage{
			Type: eventType,
			Host: monitInst.Server.Hostname,
			Message: monitInst.Event.Message,
			Service: monitInst.Event.Service,
			ServiceTypeDesc: descServiceType(monitInst.Event.Type),
		}
		for _, socket := range socketConnections {
			socket.WriteMessage(1, []byte(fmt.Sprintf(socketMsg.StringValue())))
		}
	} else {
		log.Println("INFO\tSocket arrays was empty, skip push event message to clients")
	}

	//
	// Enqueue event for Alert Worker
	//
	PublishEvent(EventMessage{
		ID: RandomID(),
		Host: monitInst.Server.Hostname,
		ServiceName: monitInst.Event.Service,
		ServiceType: descServiceType(monitInst.Event.Type),
		EventMessage: monitInst.Event.Message,
		Status: eventType,
	})
}

func processReportStats(monitInst MonitInst, monitHostsMap *map[string]MonitHost) {
	var cpu, memory float32
	var TOTAL_SERVICES = len(monitInst.Services.ServiceArr)
	var OK_SERVICES, UNMONITORED_SERVICES, FAIL_SERVICES = 0, 0, 0

	data := make([][]string, TOTAL_SERVICES)

	var xxx = make([]MonitHostService, TOTAL_SERVICES)

	for i := 0; i < TOTAL_SERVICES; i++ {
		tmpService := monitInst.Services.ServiceArr[i]

		xxx[i] = MonitHostService{
			tmpService.Name,
			descServiceType(tmpService.Type),
			tmpService.Monitor,
			tmpService.Status,
		}

		data[i] = make([]string, 5)
		data[i][0] = tmpService.Name
		data[i][1] = descServiceType(tmpService.Type)
		data[i][2] = descMonitorStatus(tmpService.Monitor)
		data[i][3] = descServiceStatus(tmpService.Status)

		if descMonitorStatus(tmpService.Monitor) == "MONITORED" && tmpService.Status == 0 {
			OK_SERVICES++
		} else if descMonitorStatus(tmpService.Monitor) == "MONITORED" && tmpService.Status != 0 {
			FAIL_SERVICES++
		} else if descMonitorStatus(tmpService.Monitor) == "UNMONITORED" {
			UNMONITORED_SERVICES++
		}

		data[i][4] = descTimestamp(int64(tmpService.CollectedSec)).String() // Format("2006-01-02 15:04:05")

		if tmpService.Type == 5 {
			memory = tmpService.System.Memory.Percent
			cpu = tmpService.System.Cpu.User + tmpService.System.Cpu.System
		}
	}
	log.Printf("INFO\tSTATUS: %d/%d are OK - %d are skipped\n", OK_SERVICES, TOTAL_SERVICES, UNMONITORED_SERVICES)
	log.Printf("INFO\tMEMORY %.2f%% - CPU %.2f%%\n", memory, cpu)

	if len(*monitHostsMap) <= MAX_AGENTS {
		host := MonitHost {
			ID: monitInst.ID,
			Poll: monitInst.Server.Poll,
			Hostname: monitInst.Server.Hostname,
			Uptime: monitInst.Server.Uptime,
			RAM: memory,
			CPU: cpu,
			Services: uint(TOTAL_SERVICES),
			GoodServices: uint(OK_SERVICES),
			FailServices: uint(FAIL_SERVICES),
			SkipServices: uint(UNMONITORED_SERVICES),
		}

		(*monitHostsMap)[monitInst.ID] = host

		hostsServicesMap[monitInst.ID] = MonitHostServices{xxx}

		if len(socketConnections) > 0 {
			log.Println("INFO\tSockets broadcast")
			socketMsg := SocketHostMessage{host}
			for _, socket := range socketConnections {
				socket.WriteMessage(1, []byte(fmt.Sprintf(socketMsg.StringValue())))
			}
		} else {
			// log.Println("INFO\tSocket was nil, skip push host message to client")
		}
	} else {
		log.Printf("WARN\tSkip agent [%s] because quota over\n", monitInst.Server.Hostname)
	}

	// table := tablewriter.NewWriter(os.Stdout)
	table := tablewriter.NewWriter(log.Writer())
	table.SetHeader([]string{"Service", "Type", "Monitor", "Status", "CollectedSec"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}

func processReportResult(monitInst MonitInst, monitHostsMap *map[string]MonitHost) {
	log.Printf("INFO\tMonit inst [%s] - Host server [%s]\n", monitInst.ID, monitInst.Server.Hostname)

	processReportStats(monitInst, monitHostsMap)

	if monitInst.Event != (Event{}) {
		// process event must come after
		processReportEvent(monitInst, monitHostsMap)
	}
}

func ParseAgentReport(xmlInput string, monitHostsMap *map[string]MonitHost) {
	var monitInst MonitInst

	// Non UTF-8
	reader := bytes.NewReader([]byte(xmlInput))
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	decoder.Decode(&monitInst)

	// UTF-8
	// xml.Unmarshal([]byte(xmlInput), &monitInst)

	// fmt.Println(monitInst)
	// _testPrint1(monitInst)
	processReportResult(monitInst, monitHostsMap)
}
