package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r, "<- panic1-recovered")
			panic("panic2")
		}
	}()

	panic("panic1")

}
