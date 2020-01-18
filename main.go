package main

import "fmt"

// Parse range qualifiers first
// RangeQuantifier ::= "{" Number ( "," Number? )? "}"
/*
	a{n} – matches “a” exactly n times
	a{n,} – matches “a” at least n times
	a{n,m} – matches “a” from n to m times
*/
func parseRangeQualifier(s string) {

	// parse the initial opening bracket
	testing := parseString("{", s)
	fmt.Println(testing)
}

func parseString(prefix string, full string) string {
	if full[0:len(prefix)] == prefix {

		fmt.Println("matched")

		toReturn := full[len(prefix):]
		return toReturn
	}

	return "nothing"

}

// check equality of a single character
func matchOne(pattern string, text string) bool {

	// all text matches empty pattern
	if pattern == "" {
		return true
	}
	if text == "" {
		return false
	}
	if pattern == "." {
		return true
	}
	return pattern == text

}

func match(pattern string, text string) bool {

	if pattern == "" {
		return true
	} else {

		return matchOne(string(pattern[0]), string(text[0])) && match(pattern[1:], text[1:])

	}

}
func main() {
	fmt.Println("vim-go")

	//regex_expression := "{1,2}"
	fmt.Println(match("abcd", "abcd"))
	//parseRangeQualifier(regex_expression)
}
