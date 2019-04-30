package myoop

import "fmt"

type Show struct {
	ShowName string
	CodeName string
}

func (s Show) Showcreate() {
	fmt.Printf("show %s has been created!\n\n", s.ShowName)
}