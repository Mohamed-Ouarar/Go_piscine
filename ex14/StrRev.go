package main

func StrRev(s string) string {
	str := []rune(s)
	end := len(str) - 1
	start := 0
	for  end > start {
		swap := str[end]
		str[end] ,str[start] = str[start] , swap
		end--
		start++
	}
	return string(str)
}
