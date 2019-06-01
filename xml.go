package pkg

import "encoding/xml"

type Schedule struct {
	XMLName xml.Name `xml:"vp"`
	Head    Head     `xml:"kopf"`
}

type Head struct {
	Titel      string `xml:"titel"`
	UploadDate string `xml:"datum"`
	Info       Info   `xml:"kopfinfo"`
}

//Kopfinfo im XML
type Info struct {
	ChangesTeacher string `xml:"aenderungl"`
}
