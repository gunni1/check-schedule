package main

import "schedulecheck/fn"


func main() {
	msg := fn.PubSubMessage{Data: []byte("bla")}
	fn.CheckScheduleAndSignal(nil,msg)

}


