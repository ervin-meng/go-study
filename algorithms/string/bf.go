package main

func bf() int {
	haystack := "dasdasdba"
	needle := "aa"

	for i := 0; i < len(haystack); {
		j := 0
		for i < len(haystack) && j < len(needle) && haystack[i] == needle[j] {
			i++
			j++
		}
		if j == len(needle) {
			return i - j
		}
		i = i - j + 1
	}

	return -1
}
