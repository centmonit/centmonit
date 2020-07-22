package core

import "fmt"

type SocketEventMessage struct {
	// Channel string = 'EVENT'
	Type string    // mandatory: error, info...
	Host string  	 // opt
	Message string // mandatory
	Service string // opt
	ServiceTypeDesc string // opt
}

func (sm SocketEventMessage) StringValue() string {
	return fmt.Sprintf(
		`{
			"channel": "%s",
			"type": "%s",
			"host": "%s",
			"message": "%s",
			"service": "%s",
			"serviceTypeDesc": "%s"
		}`,
		"EVENT",
		sm.Type,
		sm.Host,
		sm.Message,
		sm.Service,
		sm.ServiceTypeDesc,
	)
}

type SocketHostMessage struct {
	// Channel string = 'HOST'
	Host MonitHost  // mandatory
}

func (sm SocketHostMessage) StringValue() string {
	return fmt.Sprintf(
		`{
			"channel": "%s",
			"id": "%s",
			"poll": %d,
			"hostname": "%s",
			"uptime": %d,
			"ram": %f,
			"cpu": %f,
			"services": %d,
			"goodServices": %d,
			"failServices": %d,
			"skipServices": %d
		}`,
		"HOST",
		sm.Host.ID,
		sm.Host.Poll,
		sm.Host.Hostname,
		sm.Host.Uptime,
		sm.Host.RAM,
		sm.Host.CPU,
		sm.Host.Services,
		sm.Host.GoodServices,
		sm.Host.FailServices,
		sm.Host.SkipServices,
	)
}

// ID string `json:"id"`
// 	Hostname string `json:"hostname"`
// 	Uptime uint `json:"uptime"`
// 	RAM float32 `json:"ram"`
// 	CPU float32 `json:"cpu"`
// 	Services uint `json:"services"`
// 	GoodServices uint `json:"goodServices"`
// 	FailServices uint `json:"failServices"`
// 	SkipServices uint `json:"skipServices"`
