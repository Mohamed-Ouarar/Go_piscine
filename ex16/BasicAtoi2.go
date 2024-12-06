package main

func BasicAtoi2(s string) int {
	result := 0
	for _, char := range s{
        if (char < '0' || char > '9') {
            return 0
        }
		result = result * 10 + int(char) - 48
	}
	return result
}
