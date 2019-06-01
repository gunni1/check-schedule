package main

import "schedulecheck/pkg"

func main() {
	pkg.CheckScheduleAndSignal(nil, pkg.PubSubMessage{Data: []byte("")})
}