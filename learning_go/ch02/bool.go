package main

import "fmt"

func main(){
	var flag bool // no value assigned, set to false
	var isAwesome = true

	fmt.Println(flag || isAwesome)
}