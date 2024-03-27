package main

import "fmt"

func main() {
    fname := "Dalinar"
    lname := "Kholin"
    age := 45
    messageRate := 0.5
    isSubscribed := false
    message := "Sometimes hypocrite is nothing more han a man in the process of changing."

    userLog := fmt.Sprintf("Name: %s %s, Age: %d, Rate: %.1f, isSubscribed: %t, Message: %s\n", fname, lname, age, messageRate, isSubscribed, message)

    fmt.Println(userLog)
}