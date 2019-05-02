package main

import (
	"fmt"
	"strings"

	"github.com/erdemayaz/streeng"
)

func main() {
	exampleSearch()
	exampleMatch()
	exampleStartWith()
	exampleEndWith()
	exampleContains()
	exampleTerms()
	exampleFindFreqTerms()
	exampleTraverse()
}

// Example for searching
func exampleSearch() {
	// Read text from file
	text, err := streeng.StringFromFile("pp.txt")
	if err != nil {
		fmt.Println("String from file error:", err.Error())
		return
	}

	// split text with whitespace seperator
	words := strings.Fields(text)

	// Make a new Streeng
	s := streeng.MakeStreeng(words)

	// Search exact string in streeng
	listOfIndex := s.Search("begin")

	fmt.Println(listOfIndex)
}

// Example for matching
func exampleMatch() {
	// Read text from file
	text, err := streeng.StringFromFile("pp.txt")
	if err != nil {
		fmt.Println("String from file error:", err.Error())
		return
	}

	// split text with whitespace seperator
	words := strings.Fields(text)

	// Make a new Streeng
	s := streeng.MakeStreeng(words)

	// Match string using regular expression in streeng
	listOfIndex, err2 := s.Match("http(s)?://.*")
	if err2 != nil {
		fmt.Println("Regular expression error: ", err2.Error())
	}
	fmt.Println(listOfIndex)
}

// Example for start with
func exampleStartWith() {
	// Read text from file
	text, err := streeng.StringFromFile("pp.txt")
	if err != nil {
		fmt.Println("String from file error:", err.Error())
		return
	}

	// split text with whitespace seperator
	words := strings.Fields(text)

	// Make a new Streeng
	s := streeng.MakeStreeng(words)

	// Match string which start with given string in streeng
	listOfIndex := s.StartWith("sta")

	fmt.Println(listOfIndex)
}

// Example for start with
func exampleEndWith() {
	// Read text from file
	text, err := streeng.StringFromFile("pp.txt")
	if err != nil {
		fmt.Println("String from file error:", err.Error())
		return
	}

	// split text with whitespace seperator
	words := strings.Fields(text)

	// Make a new Streeng
	s := streeng.MakeStreeng(words)

	/*
		Build a reverse tree in streeng. If you don't build
		before calling EndWith function, result will not be empty
	*/
	s.ReverseStreeng()

	// Match string which end with given string in streeng
	listOfIndex := s.EndWith("ous")

	fmt.Println(listOfIndex)
}

// Example for contains
func exampleContains() {
	// Read text from file
	text, err := streeng.StringFromFile("pp.txt")
	if err != nil {
		fmt.Println("String from file error:", err.Error())
		return
	}

	// split text with whitespace seperator
	words := strings.Fields(text)

	// Make a new Streeng
	s := streeng.MakeStreeng(words)

	// Check for given exact string
	isContain := s.Contains("secret")

	fmt.Println("secret: ", isContain)
}

// Example for terms
func exampleTerms() {
	// Read text from file
	text, err := streeng.StringFromFile("pp.txt")
	if err != nil {
		fmt.Println("String from file error:", err.Error())
		return
	}

	// split text with whitespace seperator
	words := strings.Fields(text)

	// Make a new Streeng
	s := streeng.MakeStreeng(words)

	// Evaluate all terms and tokenize all words
	terms := s.Terms()

	fmt.Println(terms)
}

// Example for find freq terms
func exampleFindFreqTerms() {
	// Read text from file
	text, err := streeng.StringFromFile("pp.txt")
	if err != nil {
		fmt.Println("String from file error:", err.Error())
		return
	}

	// split text with whitespace seperator
	words := strings.Fields(text)

	// Make a new Streeng
	s := streeng.MakeStreeng(words)

	/*
		Evaluate all terms and tokenize all words.
		If you don't run terms before FindFreqTerms,
		it will return nil
	*/
	s.Terms()

	// Find a terms with minimum frequency as string int map
	freq := s.FindFreqTerms(3000)

	fmt.Println(freq)
}

// Example for traverse
func exampleTraverse() {
	// Read text from file
	text, err := streeng.StringFromFile("pp.txt")
	if err != nil {
		fmt.Println("String from file error:", err.Error())
		return
	}

	// split text with whitespace seperator
	words := strings.Fields(text)

	// Make a new Streeng
	s := streeng.MakeStreeng(words)

	s.Traverse(func(node *streeng.Node) {
		fmt.Println(node.NumberWords())
	})
}
