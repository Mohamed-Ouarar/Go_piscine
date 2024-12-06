package main

func SortIntegerTable(table []int) {
	for i := 0; i < len(table); i++ {
		for j := i + 1;j < len(table); j++{
			if table[i] > table[j] {
				swap := table[j]
				table[j] = table[i]
				table[i] = swap
			}
		}
	}
}
