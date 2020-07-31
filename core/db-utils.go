package core

import (
	"log"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/xid"
)

const dbFile = "./db/core.db"

func __init_db__() {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		file, err := os.Create(dbFile)
		if err != nil {
			log.Println("ERROR\tCreate DB file error:", err.Error())
		}
		file.Close()
		log.Println("INFO\tDB file are ready")
	} else {
		log.Println("INFO\tDB file created")
	}
}

const (
	DB_CHANNEL_TYPE_SLACK_WEBHOOK = 1
	DB_CHANNEL_TYPE_SMTP_EMAIL = 2
)

func __generate_row_id__() string {
	return xid.New().String()
}

func __init_tables__() {
	db, _ := sql.Open("sqlite3", dbFile)
	defer db.Close()

	// type: 1: slack incoming webhook; 2: smtp email server
	channelSQL := `
		create table if not exists channel (
			"id" varchar(64) NOT NULL PRIMARY KEY,
			"name" varchar(256) NOT NULL,
			"type" integer NOT NULL,
			"webhook_url" varchar(1024),
			"host" varchar(256),
			"port" integer,
			"ssl" integer,
			"user" varchar(256),
			"passwd" varchar(256)
		);`

	statement, err := db.Prepare(channelSQL)
	if err != nil {
		log.Println("ERROR\tCreate CHANNEL table error:", err.Error())
	}
	statement.Exec()
	log.Println("INFO\tCreate CHANNEL table success")
}

func __insert_slack_channel__(channelName string, webhookURL string) (bool, string) {
	db, _ := sql.Open("sqlite3", dbFile)
	defer db.Close()

	sql := `INSERT INTO channel(id, name, type, webhook_url) VALUES (?, ?, ?, ?)`
	statement, _ := db.Prepare(sql)

	channelType := DB_CHANNEL_TYPE_SLACK_WEBHOOK
	row_id := __generate_row_id__()
	_, err := statement.Exec(row_id, channelName, channelType, webhookURL)
	if err != nil {
		return false, ""
	}
	return true, row_id
}

func __insert_smtp_channel__(
		channelName string,
		host string,
		port uint,
		ssl bool,
		user string,
		passwd string) (bool, string) {
	db, _ := sql.Open("sqlite3", dbFile)
	defer db.Close()

	sql := `INSERT INTO channel(id, name, type, host, port, ssl, user, passwd) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	statement, _ := db.Prepare(sql)

	channelType := DB_CHANNEL_TYPE_SMTP_EMAIL
	row_id := __generate_row_id__()
	_, err := statement.Exec(row_id, channelName, channelType, host, port, ssl, user, passwd)
	if err != nil {
		return false, ""
	}
	return true, row_id
}

type ChannelObject struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Type uint `json:"type"`
	WebhookURL string `json:"webhook_url,omitempty"`
	Host string `json:"host,omitempty"`
	Port uint `json:"port,omitempty"`
	SSL bool `json:"ssl,omitempty"`
	User string `json:"user,omitempty"`
	Passwd string `json:"passwd,omitempty"`
}

func ChannelGetAll() []ChannelObject {
	db, _ := sql.Open("sqlite3", dbFile)
	defer db.Close()

	rows, _ := db.Query(`
		SELECT
			id, name, type, webhook_url, host, port, ssl, user, passwd
		FROM channel
		order by ID desc
	`)
	defer rows.Close()

	result := make([]ChannelObject, 0)

	log.Println("INFO\tChannel get all row")
	for rows.Next() {
		obj := ChannelObject{}
		var _webhook sql.NullString
		var _host sql.NullString
		var _port sql.NullInt32
		var _ssl sql.NullBool
		var _user sql.NullString
		var _passwd sql.NullString

		err := rows.Scan(&obj.ID, &obj.Name, &obj.Type, &_webhook, &_host, &_port, &_ssl, &_user, &_passwd)
		if err != nil {
			log.Println("ERROR\t - scan row error:", err)
		} else {
			log.Printf("INFO\t - row item:%+v\n", obj)
		}
		// Extra checking null field
		obj.WebhookURL = _webhook.String
		obj.Host = _host.String
		obj.Port = uint(_port.Int32)
		obj.SSL = bool(_ssl.Bool)
		obj.User = _user.String
		obj.Passwd = _passwd.String

		result = append(result, obj)
	}
	return result
}

func ChannelSave(payload []byte) (success bool, rowID string) {
	ch := ChannelObject{}
	json.Unmarshal(payload, &ch)
	log.Printf("INFO\tChannelSave() get payload object: %+v\n", ch)
	if ch.Type == DB_CHANNEL_TYPE_SLACK_WEBHOOK {
		return __insert_slack_channel__(ch.Name, ch.WebhookURL)
	} else {
		return __insert_smtp_channel__(ch.Name, ch.Host, ch.Port, ch.SSL, ch.User, ch.Passwd)
	}
}

func ChannelRemove(id string) bool {
	db, _ := sql.Open("sqlite3", dbFile)
	defer db.Close()

	sql := `delete from channel where id = ?`
	statement, _ := db.Prepare(sql)
	_, err := statement.Exec(id)

	if err != nil {
		return false
	}
	return true
}

func PrepareDB() {
	__init_db__()
	__init_tables__()
}

func DBTest() {
	fmt.Println("Test DB")
	// __init_db__()
	// __init_tables__()
	__insert_slack_channel__("default channel 1", "http://abc.com/def")
	__insert_smtp_channel__("channel 2", "smtp.gmail.com", 587, false, "", "")
}