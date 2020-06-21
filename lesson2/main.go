package main

import (
	"fmt"
)

func main() {
	var users [3]string = [3]string{"Dima", "Max", "Nick"}

	transaction := make(map[int]map[string]string)
	transaction[1] = map[string]string{
		"User":   "Dima",
		"Status": "Paid",
		"Sum":    "100RUR",
	}

	transaction[2] = map[string]string{
		"User":   "Max",
		"Status": "Refund",
		"Sum":    "200RUR",
	}

	transaction[3] = map[string]string{
		"User":   "Nick",
		"Status": "Paid",
		"Sum":    "300RUR",
	}

	transaction[4] = map[string]string{
		"User":   "Dima",
		"Status": "Paid",
		"Sum":    "100RUR",
	}

	fmt.Println(users)
	fmt.Printf("Unsorted trxs: %v \n", transaction)

	var sortedTransaction []map[string]string

	for _, value := range users {
		for i := 0; i <= len(transaction); i++ {
			if trx, ok := transaction[i]; ok {
				if trx["User"] == value {
					sortedTransaction = append(sortedTransaction, transaction[i])
				}
			}
		}
	}

	for _, value := range sortedTransaction {
		fmt.Printf("User: %s, Status, %s, Sum: %s \n", value["User"], value["Status"], value["Sum"])
	}

}
