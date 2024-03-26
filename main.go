package main

import "fmt"

func main() {
    var username string = "presidentBilly"
    var password string = "12345"
    // var password int = 12345

    fmt.Println("Authorization: Basic", username+":"+password)
}