package core

import (
	"fmt"
	"log"
	"time"
	"math/rand"
	// "runtime"
	// "github.com/masslessparticle/goq"
	// "github.com/masslessparticle/goq/testhelpers"
	// "github.com/masslessparticle/goq/pubsub"
	"encoding/json"
	"strings"
	"github.com/tidwall/gjson"
	"bytes"
	"io/ioutil"
	"net/http"
	// "net/smtp"
	"io"
	gomail "gopkg.in/mail.v2"
	"text/template"
)

const TOTAL_WORKERS = 5
const MESSAGE_BUFFER = 10  // for each worker
var roundRobinIndex uint
var QueueChannel [TOTAL_WORKERS]chan EventMessage

type EventMessage struct {
	ID string
	Host string
	ServiceName string
	ServiceType string
	EventMessage string
	Status string // success, warning, error
}

type EventFilter struct {
	Host string `json:"host"`
	Services []string `json:"services"`
}

type EventFilters []EventFilter

func processingEvent_v1(listener string, event EventMessage) {
	log.Printf("INFO\t[%s] - I got event to process: %+v\n", listener, event)
	sleepTime := rand.Intn(5 - 2) + 2
	sleepTime = 5
	// fmt.Print(" - I will sleep ", sleepTime)
	time.Sleep(time.Duration(sleepTime) * time.Second)
	log.Printf("INFO\t[%s] - Finish event #%s\n", listener, event.ID)
}

func processingEvent(listener string, event EventMessage) {
	log.Printf("INFO\t[%s] - I got event to process: %+v\n", listener, event)

	if event.ServiceName == "Monit" {
		// transform event for Monit service itself
		event.ServiceName = event.Host
	}

	alertRules := AlertRuleGetAll()
	for _, rule := range alertRules {
		// log.Printf("INFO\t[%s] - Filter: %+v - Type: %T", listener, rule.Filters, rule.Filters)
		efs := EventFilters{}
		json.Unmarshal([]byte(rule.Filters), &efs)
		log.Printf("INFO\t[%s] - Filters: %+v - Type: %T", listener, efs, efs)

		// check whether the rule was matched with event
		matched := isFiltersMatchedWithEvent(listener, efs, event)
		log.Printf("INFO\t[%s] - Rule matched: %t\n", listener, matched)
		if matched {
			alertType := gjson.Get(rule.AlertConf, "type")
			log.Printf("INFO\t[%s] - Alert type %s\n", listener, alertType)
			if alertType.String() == "SMTP" {
				go processAlertViaSMTP(
					gjson.Get(rule.AlertConf, "value").String(),
					gjson.Get(rule.AlertConf, "smtpReceivers").String(),
					event)
			} else {
				// Slack Webhook
				go processAlertViaSlackWebhook(gjson.Get(rule.AlertConf, "value").String(), event)
			}
		}
	}
	log.Printf("INFO\t[%s] - Finish event #%s\n", listener, event.ID)
}

func isFiltersMatchedWithEvent(listener string, eventFilters EventFilters, eventMessage EventMessage) bool {
	for _, filter := range eventFilters {
		log.Printf("INFO\t[%s] - Check filter: %+v", listener, filter)
		if strings.EqualFold(filter.Host, "all") {
			return true
		} else if filter.Host == eventMessage.Host &&
			(isItemInArray("All", filter.Services) || isItemInArray(eventMessage.ServiceName, filter.Services)) {
			return true
		}
	}
	return false
}

func isItemInArray(item string, array []string) bool {
	for _, val := range array {
		if val == item {
			return true
		}
	}
	return false
}

func processAlertViaSlackWebhook(alertChannelID string, event EventMessage) {
	log.Printf("INFO\tProcess slack webhook alert #%s\n", alertChannelID)
	// Get channel detail
	obj := ChannelGet(alertChannelID)
	if obj.ID != "" {
		url := obj.WebhookURL
		// log.Printf("INFO\tSlack webhook URL:%s\n", url)

		emoIcon := ":heavy_check_mark:"
		if event.Status != "success" {
			emoIcon = ":warning:"
			// emoIcon = ":exclamation:"
		}

		t, _ := template.ParseFiles("./template/alert-slack.json")
		var payload bytes.Buffer
		t.Execute(&payload, struct {
			Sender string
			Status string
			Icon string
			DateTime string
			Host    string
			Service string
			ServiceType string
			Description string
		} {
			obj.CustomText,
			strings.Title(event.Status), // Camel case transforming
			emoIcon,
			time.Now().Format(time.RFC850),
			event.Host,
			event.ServiceName,
			event.ServiceType,
			event.EventMessage,
		})

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload.Bytes()))
		if err != nil {
			log.Println("ERROR\tPost to slack webhook error:", err)
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("ERROR\tPost to slack response read error:", err)
		} else {
			log.Println("INFO\tPost to slack response:", string(body))
		}
	} // else not found channel by id
}

func processAlertViaSMTP(alertChannelID string, receivers string, event EventMessage) {
	log.Printf("INFO\tProcess smtp alert #%s - receivers: %s\n", alertChannelID, strings.Split(receivers, "\n"))
	// Get channel detail
	obj := ChannelGet(alertChannelID)
	// log.Printf("INFO\tSMTP info - host: %s, port: %d, ssl: %t, user: %s, passwd: %s\n", obj.Host, obj.Port, obj.SSL, obj.User, obj.Passwd)

	m := gomail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("%s <%s>", obj.CustomText, obj.User))

	toMap := map[string][]string{"To": strings.Split(receivers, "\n")}
	m.SetHeaders(toMap)

	mailSubject := "✅"
	if event.Status != "success" {
		// mailSubject = "❗️"
		mailSubject = "⚠️"
	}
	mailSubject += fmt.Sprintf(" %s • %s", event.Host, event.ServiceName)
	m.SetHeader("Subject", mailSubject)

	t, _ := template.ParseFiles("./template/alert-email.html")
	m.AddAlternativeWriter("text/html", func(w io.Writer) error {
		return t.Execute(w, struct {
			DateTime string
			Host    string
			Service string
			ServiceType string
			Description string
			Status string
		} {
			time.Now().Format(time.RFC850),
			event.Host,
			event.ServiceName,
			event.ServiceType,
			strings.Replace(event.EventMessage, "\n", "<br/>", -1),
			strings.Title(event.Status), // Camel case transforming
		})
	})

	m.Embed("./html/LogoMakr_4rwcqQ.png")

	var d *gomail.Dialer
	if obj.User != "" && obj.Passwd != "" {
		d = gomail.NewDialer(obj.Host, int(obj.Port), obj.User, obj.Passwd)
	} else {
		d = &gomail.Dialer{Host: obj.Host, Port: int(obj.Port)}
	}

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		log.Println("ERROR\tSendMail error:", err)
	} else {
		log.Println("INFO\tSendMail success")
	}
}

func startListeners() {
	for idx, qc := range QueueChannel {
		go startChannelListener(idx, qc)
	}
}

func startChannelListener(index int, qc chan EventMessage) {
	log.Printf("INFO\tStart event worker #%d\n", index)
	for msg := range qc {
		processingEvent(fmt.Sprintf("worker%d", index), msg)
	}
}

func PublishEvent(event EventMessage) {
	QueueChannel[roundRobinIndex] <- event
	roundRobinIndex = (roundRobinIndex+1) % TOTAL_WORKERS
}

func InitWorkers() {
	for idx := range QueueChannel {
		QueueChannel[idx] = make(chan EventMessage, MESSAGE_BUFFER)
	}
	roundRobinIndex = 0
	startListeners()
}
