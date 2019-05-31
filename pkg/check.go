package pkg

import (
	"encoding/xml"
	"log"
	"strings"
)

func CheckScheduleAndSignal(config ScheduleClientConfig, code string) error {
	xmlResponse, _ := config.RequestXML()
	var document Schedule
	xml.Unmarshal(xmlResponse, &document)
	log.Println("Prüfe Plan für: " + document.Head.Titel)
	log.Println("Suche nach Kürzel: " + code)
	if strings.Contains(document.Head.Info.ChangesTeacher, code+";") {
		log.Println("Änderungen gefunden!")
	} else {
		log.Println("Keine relevanten Änderungen.")
	}
	return nil
}
