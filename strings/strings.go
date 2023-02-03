package main

import (
	"fmt"
	"strings" 
)


func Stringy(mainString, subString string) bool {
	return strings.Contains(mainString,subString)
}

func main(){
	fmt.Println(Stringy("messi", "mess"))
}