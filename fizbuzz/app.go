package fizbuzz

// fizz_buzz to validate whether or not
func fizzBbuzz(amount int) string {
	result := ""
	if amount % 3 == 0 {
		result += "Fizz"
	}
	if amount % 5 == 0 {
        result += "Buzz"
	}
	return result
}
