package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
)

const (
	userEnv = "PAGE_USER"
	pwEnv   = "PAGE_PASSWORD"
	dateEnv = "DATE"
)

var (
	dateFormat, _ = regexp.Compile("[0-9]{8}")
)

func main() {
	username := parseEnvMandatory(userEnv)
	password := parseEnvMandatory(pwEnv)

	baseURL := *flag.String("baseUrl", "https://www.stundenplan24.de/10124219/vplanle/vdaten/VplanLe", "base url ultil date")
	flag.Parse()

	date := os.Getenv(dateEnv)
	if date == "" {
		today := time.Now()
		date = fmt.Sprintf("%d%02d%02d", today.Year(), int(today.Month()), today.Day())
	}
	if !dateFormat.MatchString(date) {
		log.Fatalln("incorrect date format. Please format as YYYYMMDD")
	}
	config := ClientConfig{User: username, Password: password, BaseURL: baseURL, Date: date}
	xmlResponse, _ := config.RequestXML()
	//Ab hier XML
	var document Vp
	xml.Unmarshal(xmlResponse, &document)
	log.Println(document.Head.Titel)
}

func parseEnvMandatory(variableKEy string) string {
	variableValue := os.Getenv(variableKEy)
	if variableValue == "" {
		log.Fatalln("Environment variable: " + variableKEy + " is empty")
	}
	return variableValue
}
