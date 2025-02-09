package StringFunc

func validPalindrome(s string) bool {
	for l, h := 0, len(s)-1; l < h; l, h = l+1, h-1 {
		if s[l] != s[h] {
			return isPalindrome(s[l:h]) || isPalindrome(s[l+1:h+1])
		}
	}
	return true
}

func isPalindrome(s string) bool {
	for l, h := 0, len(s)-1; l < h; l, h = l+1, h-1 {
		if s[l] != s[h] {
			return false
		}
	}
	return true
}
