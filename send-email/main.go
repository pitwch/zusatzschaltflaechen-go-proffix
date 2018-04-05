package main

import (
	"os"
	"fmt"
	"net/smtp"
	"log"
	"encoding/json"
	"gopkg.in/ini.v1"
)

type Configuration struct {
	From     string
	Host     string
	Port     string
	Password string
	Subject  string
	Message  string
}

func main() {

	//Read PROFFIX Zusatzfeld
	//pxzu, err := ini.Load(os.Args[1])
	pxzu, err := ini.Load("F:\\test.ini")

	if err != nil {
		fmt.Printf("Error reading Zusatzfeld: %v", err)
		os.Exit(1)
	}

	email := pxzu.Section("Fields").Key("dfsEMail").String()
	anrede := pxzu.Section("Fields").Key("cmbBriefanrede").String()

	send(email, anrede)

}

func send(email string, anrede string) {
	//Read config.json
	file, _ := os.Open("C:\\go-work\\src\\zusatzschaltflaechen-go-proffix\\zusatzschaltflaechen-go-proffix\\send-email\\config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Printf("Error reading Config File: %v", err)
		os.Exit(1)
	}
	from := configuration.From
	pass := configuration.Password
	host := configuration.Host
	port := configuration.Port
	subject := configuration.Subject
	to := email
	body := configuration.Subject + "\n" + Configuration{}.Message

	fmt.Println(configuration.Port)
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n" +
		body

	err = smtp.SendMail(host+":"+port,
		smtp.PlainAuth("", from, pass, host),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("SMTP Error: %s", err)
		return
	}

	log.Print("Email sent to ", to)
}
