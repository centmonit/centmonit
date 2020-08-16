package main

import (
	"log"
	"io"
	"os"
	"local.domain/CentMonit/core"
	"gopkg.in/natefinch/lumberjack.v2"
	"net/smtp"
	gomail "gopkg.in/mail.v2"
	"text/template"
	"bytes"
	"fmt"
)

func __do_work__() {
	cfg, err := core.GetConfig("./config.yml")
	if err != nil {
			log.Fatal("ERROR\t", err)
	}

	mw := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   "./logs/log.txt",
		MaxSize:    cfg.Log.MaxFiles, // MB
		MaxBackups: cfg.Log.MaxFileSize,
		MaxAge:     cfg.Log.RetentionDays, //days
		Compress:   false, // disabled by default
	})
	log.SetOutput(mw)

	// Stop if exceed 5 day trial

	log.Printf("INFO\tCentMonit starting...\n")
	log.Printf("INFO\tListenning at %s - web port %s - api port %s", cfg.Net.ApiHost, cfg.Net.WebPort, cfg.Net.ApiPort)

	core.PrepareDB()
	core.ConfigWebServer(cfg.Net.ApiHost, cfg.Net.ApiPort, cfg.Auth)

	go core.StartWebServer(cfg.Net.WebPort)
	go core.StartApiServer(cfg.Net.ApiPort)
	go core.InitWorkers()

	log.Printf("INFO\tCentMonit started")
	for {}
}

func __do_test_xml__() {
	// xml := `<?xml version="1.0" encoding="ISO-8859-1"?><monit id="cd1c5796162f389db974659f975fa007" incarnation="1594092205" version="5.26.0"><server><uptime>5</uptime><poll>30</poll><startdelay>0</startdelay><localhostname>monit-master</localhostname><controlfile>/etc/monitrc</controlfile><httpd><address>127.0.0.1</address><port>2812</port><ssl>0</ssl></httpd><credentials><username>admin</username><password>monit</password></credentials></server><platform><name>Linux</name><release>4.4.0-177-generic</release><version>#207-Ubuntu SMP Mon Mar 16 01:16:10 UTC 2020</version><machine>x86_64</machine><cpu>8</cpu><memory>16203564</memory><swap>0</swap></platform><services><service name="check_tmp_dir"><type>1</type><collected_sec>1594092209</collected_sec><collected_usec>749744</collected_usec><status>0</status><status_hint>0</status_hint><monitor>1</monitor><monitormode>0</monitormode><onreboot>0</onreboot><pendingaction>0</pendingaction><mode>7777</mode><uid>-1</uid><gid>-1</gid><timestamps><access>0</access><change>0</change><modify>0</modify></timestamps></service><service name="monit-master"><type>5</type><collected_sec>1594092209</collected_sec><collected_usec>749746</collected_usec><status>0</status><status_hint>0</status_hint><monitor>1</monitor><monitormode>0</monitormode><onreboot>0</onreboot><pendingaction>0</pendingaction><system><load><avg01>0.30</avg01><avg05>0.43</avg05><avg15>0.48</avg15></load><cpu><user>0.0</user><system>0.0</system><wait>0.0</wait></cpu><memory><percent>33.1</percent><kilobyte>5364392</kilobyte></memory><swap><percent>0.0</percent><kilobyte>0</kilobyte></swap></system></service></services><servicegroups></servicegroups></monit>`
	// xml = `<?xml version="1.0" encoding="ISO-8859-1"?><monit id="cd1c5796162f389db974659f975fa007" incarnation="1594193421" version="5.26.0"><server><uptime>1456</uptime><poll>30</poll><startdelay>0</startdelay><localhostname>monit-master</localhostname><controlfile>/etc/monitrc</controlfile><httpd><address>127.0.0.1</address><port>2812</port><ssl>0</ssl></httpd><credentials><username>admin</username><password>monit</password></credentials></server><platform><name>Linux</name><release>4.4.0-177-generic</release><version>#207-Ubuntu SMP Mon Mar 16 01:16:10 UTC 2020</version><machine>x86_64</machine><cpu>8</cpu><memory>16203564</memory><swap>0</swap></platform><services><service name="check_tmp_dir"><type>1</type><collected_sec>1594194719</collected_sec><collected_usec>352984</collected_usec><status>0</status><status_hint>0</status_hint><monitor>1</monitor><monitormode>0</monitormode><onreboot>0</onreboot><pendingaction>0</pendingaction><mode>7777</mode><uid>-1</uid><gid>-1</gid><timestamps><access>0</access><change>0</change><modify>0</modify></timestamps></service><service name="monit-master"><type>5</type><collected_sec>1594194692</collected_sec><collected_usec>211069</collected_usec><status>0</status><status_hint>0</status_hint><monitor>2</monitor><monitormode>0</monitormode><onreboot>0</onreboot><pendingaction>0</pendingaction></service></services><servicegroups></servicegroups><event><collected_sec>1594194722</collected_sec><collected_usec>273739</collected_usec><service>monit-master</service><type>5</type><id>131072</id><state>2</state><action>1</action><message><![CDATA[monitor action done]]></message></event></monit>`
	// core.TestParse(xml)
}

func __do_test_smtp1__() {
	// Set up authentication information.
	auth := smtp.PlainAuth("", "alert@pingcom.vn", "123456789!", "smtp.gmail.com")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"tuanha@pingcom.vn","tuanha@conera.io",}
	msg := []byte("From: alert@pingcom.vn\r\n" +"To: tuanha@pingcom.vn,tuanha@conera.io\r\n" +
		"Subject: Go language email send using smpt package\r\n" +
		"\r\n" +
		"Hi Team,\nThis is sample email using Go language executable file.\nIn this we will be learning how to send email using go lang.\n\nRegards,\nEducateSpace support\r\n")

	err := smtp.SendMail("smtp.gmail.com:587", auth, "alert@pingcom.vn", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}

func __do_test_smtp2__() {
	m := gomail.NewMessage()
	m.SetHeader("From", "alert@pingcom.vn")
	m.SetHeader("To", "tuanha@pingcom.vn", "cora@example.com")
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	// m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.gmail.com", 587, "alert@pingcom.vn", "123456789!")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
			panic(err)
	}
}

func __do_test_smtp3__() {
	m := gomail.NewMessage()
	m.SetHeader("From", "CentMonit <alert@pingcom.vn>")
	m.SetHeader("To", "tuanha@pingcom.vn")
	// m.SetHeader("Subject", "✅ Success email")
	m.SetHeader("Subject", "❗️ Warning email") // ⚠️

	t, _ := template.ParseFiles("./template/alert-email.html")
	m.AddAlternativeWriter("text/html", func(w io.Writer) error {
		return t.Execute(w, struct {
			Name    string
			Message string
		} {"JoeHart", "I had played for ManC"})
	})

	m.Embed("./html/LogoMakr_4rwcqQ.png")

	d := gomail.NewDialer("smtp.gmail.com", 587, "alert@pingcom.vn", "123456789!")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
			panic(err)
	}
}

func __do_test_slack__() {
	t, _ := template.ParseFiles("./template/alert-slack.json")
	var body bytes.Buffer
	// t.Execute(os.Stdout, struct {} {})
	t.Execute(&body, struct {
		Status string
	} {"Error"})
	fmt.Println(string(body.Bytes()[:]))
}

func main() {
	__do_work__()
	// core.DBTest()
	// __do_test_smtp3__()
	// __do_test_slack__()
}
