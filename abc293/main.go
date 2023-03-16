package main

import "fmt"

func main() {
	var s string

	fmt.Scan(&s)

	for i := 0; i < len(s)/2; i++ {
    fmt.Print(s[2*i+1:2*i+2])
    fmt.Print(s[2*i:2*i+1])
	}
}
