package fz

import "strconv"

func Fizzbuss(n int) string {
	var answer string

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			answer = answer + "FizzBuzz "
		} else if i%3 == 0 {
			answer = answer + "Fizz "
		} else if i%5 == 0 {
			answer = answer + "Buzz "
		} else {
			answer = answer + strconv.Itoa(i) + " "
		}
	}
	return answer
}
