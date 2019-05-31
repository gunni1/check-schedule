package fn 

import (
	"context"
	"log"
	"fmt"
	"os"
	"regexp"
	"time"
	"schedulecheck/pkg"
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

// PubSubMessage is the payload of a Pub/Sub event. Please refer to the docs for
// additional information regarding Pub/Sub events.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

// Triggered by a Pub/Sub message.
func CheckScheduleAndSignal(ctx context.Context, m PubSubMessage) error {
	code := parseEnvMandatory(codeEnv)
	config := parseEnvToSchedulerConfig()
	return pkg.CheckScheduleAndSignal(config, code)
}

func parseEnvToSchedulerConfig() pkg.ScheduleClientConfig {
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
	return pkg.ScheduleClientConfig{User: username, Password: password, BaseURL: baseURL, Date: date}
}

func parseEnvMandatory(variableKEy string) string {
	variableValue := os.Getenv(variableKEy)
	if variableValue == "" {
		log.Fatalln("Environment variable: " + variableKEy + " is empty")
	}
	return variableValue
}

