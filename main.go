package main 

import (
	"context"
	"log"
	"fmt"
	"os"
	"regexp"
	"time"
	"github.com/aws/aws-lambda-go/lambda"
)

const (
	userEnv = "PAGE_USER"
	pwEnv   = "PAGE_PW"
	dateEnv = "DATE"
	codeEnv = "CODE"
	baseURLEnv = "BASE_URL"
)

var (
	dateFormat, _ = regexp.Compile("[0-9]{8}")
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	code := parseEnvMandatory(codeEnv)
	config := parseEnvToSchedulerConfig()
	CheckScheduleAndSignal(config, code)

	return "", nil
}

func main() {
	lambda.Start(HandleRequest)
}

func parseEnvToSchedulerConfig() ScheduleClientConfig {
	username := parseEnvMandatory(userEnv)
	password := parseEnvMandatory(pwEnv)
	
	baseURL := os.Getenv(baseURLEnv)
	if baseURL == "" {
		baseURL = "https://www.stundenplan24.de/10124219/vplanle/vdaten/VplanLe"
	}
	date := os.Getenv(dateEnv)
	if date == "" {
		today := time.Now()
		date = fmt.Sprintf("%d%02d%02d", today.Year(), int(today.Month()), today.Day())
	}
	if !dateFormat.MatchString(date) {
		log.Fatalln("incorrect date format. Please format as YYYYMMDD")
	}
	return ScheduleClientConfig{User: username, Password: password, BaseURL: baseURL, Date: date}
}

func parseEnvMandatory(variableKEy string) string {
	variableValue := os.Getenv(variableKEy)
	if variableValue == "" {
		log.Fatalln("Environment variable: " + variableKEy + " is empty")
	}
	return variableValue
}

