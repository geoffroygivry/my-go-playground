package main

import (
	"fmt"
	"os"
	"time"

	myprint "./myPrint"
)

var args = os.Args

func main() {
	myprint.Printit("hey hey")
	myMessage, myTime := myprint.MyTime()
	fmt.Println(myMessage, myTime)
	fmt.Println(time.Now())
	fmt.Println("Hello World!")
	myUniqueSlice := uniqueArgs(args[1:])
	fmt.Println(myUniqueSlice)
}

func uniqueArgs(args []string) []string {
	var uniqueStrings []string
	for _, v := range args {
		for _, i := range uniqueStrings {
			if v == i {
				continue
			} else {
				uniqueStrings = append(uniqueStrings, v)
			}
		}
	}
	return uniqueStrings
}
