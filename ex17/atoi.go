package main

func Atoi(s string) int {
	result := 0
	condition := 1
	sign := 1
	for _, char := range s{
		if condition == 1 && (char == '+' || char == '-'){
			if char == '-' {
				sign = -1
			}
			condition = 0
			continue
		}
        if char < '0' || char > '9' {
            return 0
		}
		result = result * 10 + int(char) - 48
	}
	return result * sign
}
