package main

import "fmt"

func main() {
    fname := "Dalinar"
    lname := "Kholin"
    age := 45
    messageRate := 0.5
    isSubscribed := false
    message := "Sometimes hypocrite is nothing more han a man in the process of changing."

    userLog := fmt.Sprintf("Name: ? ?, Age: ?, Rate: ?, isSubscribed: ?, Message: ?")

    fmt.Println(userLog)
}