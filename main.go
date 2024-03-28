package main

func monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth int) int {
	var lastMonthBill int
	var thisMonthBill int
	getBillForMonth(lastMonthBill, costPerSend, numLastMonth)
	getBillForMonth(thisMonthBill, costPerSend, numThisMonth)
	return thisMonthBill - lastMonthBill
}

func getBillForMonth(bill, costPerSend, messagesSent, int) {
	bill = costPerSend * messagesSent
}


// INSTRUCTIONS
// Fix the bugs in the monthlyBillIncrease and getBillForMonth functions.

// monthlyBillIncrease: Returns the increase in the bill from the previous to the current month. If the bill decreased, it should return a negative number.
// getBillForMonth: Returns the bill for the given month.
// It looks like whoever wrote the getBillForMonth function thought that they could pass in the bill parameter, update it inside the function, and that update would apply in the parent function (monthlyBillIncrease). They were wrong.

// Change the getBillForMonth function to explicitly return the bill for the given month, and be sure to capture that return value properly in the monthlyBillIncrease function.

// The function signature for getBillForMonth should only take 2 parameters once you're done.