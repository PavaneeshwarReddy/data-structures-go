package greedy

/*
Lemonade
At a lemonade stand, each lemonade costs $5. Customers are standing in a queue to buy from you and order one at a time (in the order specified by bills). Each customer will only buy one lemonade and pay with either a $5, $10, or $20 bill. You must provide the correct change to each customer so that the net transaction is that the customer pays $5.
Note that you do not have any change in hand at first.
Given an integer array bills where bills[i] is the bill the ith customer pays, return true if you can provide every customer with the correct change, or false otherwise.

Explanation:
- If we have five just increment 5
- If we got 10 then check if five exisits if not return false
- If we got 20 then check if 10s if not check with 5s if not return false
*/

func lemonadeChange(bills []int) bool {
	five, ten := 0, 0

	for _, val := range bills {
		if val == 5 {
			five++
		}
		if val == 10 {
			if five > 0 {
				five--
				ten++
			} else {
				return false
			}
		}
		if val == 20 {
			val -= 5
			if ten > 0 {
				val -= 10
				ten--
			}
			if (val / 5) > five {
				return false
			}
			five -= (val / 5)
		}
	}

	return true
}
