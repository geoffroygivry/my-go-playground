package main

import (
	"fmt"
	"os"
	"time"

	myprint "./myPrint"
	sliceexplanation "./slicesTuto"
	myoop "./interfacesOop"
)

var args = os.Args
var newShow = myoop.Show{
	"Fufu",
	"Fu",
}

func main() {
	myprint.Printit("hey hey")
	fmt.Println(args[1:])
	myMessage, myTime := myprint.MyTime()
	fmt.Println(myMessage, myTime)
	fmt.Println(time.Now())
	fmt.Println("Hello World!")
	myUniqueSlice := uniqueArgs(args[1:])
	fmt.Println(myUniqueSlice)
	mySuperSlice := sliceexplanation.PrintMySlice()
	fmt.Println(mySuperSlice)

	newShow.Showcreate()
}

func uniqueArgs(args []string) []string {
	var uniqueStrings []string
	uniqueStrings = append(uniqueStrings, args[0])
	for _, v := range args {
		if len(uniqueStrings) > 0 {
			for _, i := range uniqueStrings {
				fmt.Println("i, v :", i, v)
				if v == i {
					continue
				} else {
					uniqueStrings = append(uniqueStrings, v)
				}
			}
		}
	}
	return uniqueStrings
}
