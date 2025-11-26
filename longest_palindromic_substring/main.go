package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, maxLength := 0, 1

	expandAroundCenter := func(l, r int) {
		// Expand while characters match and indices are valid
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}

		// Calculate length of the found palindrome
		// Note: indices l and r are now one step 'too far', so length is (r-1) - (l+1) + 1 = r - l - 1
		currentLength := r - l - 1

		if currentLength > maxLength {
			maxLength = currentLength
			start = l + 1
		}
	}

	for i := 0; i < len(s); i++ {
		// Check for odd-length palindromes (e.g., "aba")
		expandAroundCenter(i, i)
		// Check for even-length palindromes (e.g., "abba")
		expandAroundCenter(i, i+1)
	}

	return s[start : start+maxLength]
}

func main() {
	// Test cases from the problem description + edge cases
	tests := []string{
		"babad",   // Expected: "bab" or "aba"
		"cbbd",    // Expected: "bb"
		"a",       // Expected: "a"
		"ac",      // Expected: "a" (or "c")
		"racecar", // Expected: "racecar"
	}

	fmt.Println("Longest Palindromic Substring Tests:")
	fmt.Println("------------------------------------")

	for _, t := range tests {
		result := longestPalindrome(t)
		fmt.Printf("Input:  %s\nOutput: %s\n\n", t, result)
	}
}
