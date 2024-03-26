package main

import "fmt"

func main() {
    messagesFromHelga := []string{
        "Have ya seen me biceps lately",
        "I can throw a log thrice the distance of half of what you can",
        "I'm sorry I'm like this please respond",
        "or not I'll just kick you ass in the town square next time I see you!",
    }
numMessages := float64(len(messagesFromHelga))
costPerMessage := .02

// totalCost := costPerMessage + numMessages

totalCost := costPerMessage * numMessages

fmt.Printf("Helga spent %.2f on text messages today/n", totalCost)
}