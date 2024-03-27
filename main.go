package main

import "fmt"

func main() {
    var insufficientFundMesssage string = "Purchase failed. Insuffficient funds."
    var purchaseSuccessMessage string = "Purchase succesful."
    var accountBalance float64 = 100.0
    var bulkMessageCost float64 = 75.0 
    var isPremiumUser bool = true
    var discountRate float64 = 0.10
    var finalCost float64


//set finalCost to bulkMessageCost
finalCost = bulkMessageCost

if isPremiumUser == true {
    finalCost -= bulkMessageCost * discountRate
}


if accountBalance > finalCost {
    fmt.Println(purchaseSuccessMessage)
} else {
    fmt.Println(insufficientFundMesssage)
}
accountBalance = accountBalance - finalCost
//if user is premium, discount should apply to final cost
//if there is more money in account balance print purchase succcess message
//if there is less money in accound balance print the puchase failed message


    fmt.Println("Account balance:", accountBalance)
}