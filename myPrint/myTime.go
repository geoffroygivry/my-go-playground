package myprint

import "time"

// MyTime returns an value of time in int64 type. Trulufu.
func MyTime() (myMessage string, myTime int64) {
	myTime = time.Now().UnixNano()
	myMessage = ("the time at this point in nano second is....     ")
	return
}
