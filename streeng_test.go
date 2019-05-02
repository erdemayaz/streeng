package streeng

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
	"time"
)

/*
Test functions wrote for PRIDE AND PREJUDICE book text.
*/

func TestSearch(t *testing.T) {
	fileName := "pp.txt"
	tests := []string{
		`discretion`,
		`exclamation`,
		`want`,
		`fortune`,
		`objection`,
		`name`,
		`daughters`,
		`pleasure`,
		`Mrs.`,
		`asd`,
		`importance`,
		`neighbour`,
	}
	text, err := StringFromFile(fileName)
	if err != nil {
		t.Errorf("Test Fail: StringFromFile Error: %s\n", err.Error())
	}
	words := strings.Fields(text)
	streeng := MakeStreeng(words)
	for _, test := range tests {
		i := 0
		for _, word := range streeng.words {
			if strings.Compare(test, word) == 0 {
				i++
			}
		}
		results := streeng.Search(test)
		if len(results) == i {
			t.Logf("Test Successful: word: %s", test)
		} else {
			t.Errorf("Test Fail:\t word: %s \t expected: %d \t result: %d",
				test, i, len(results))
		}
	}
}

func TestStartWith(t *testing.T) {
	fileName := "pp.txt"
	tests := []string{
		`disc`,
		`ex`,
		`w`,
		`fort`,
		`obj`,
		`nam`,
		`daughte`,
		`pleas`,
		`Mr`,
		`asd`,
		`import`,
		`neigh`,
	}
	text, err := StringFromFile(fileName)
	if err != nil {
		t.Errorf("Test Fail: StringFromFile Error: %s\n", err.Error())
	}
	words := strings.Fields(text)
	streeng := MakeStreeng(words)
	for _, test := range tests {
		i := 0
		for _, word := range streeng.words {
			if strings.HasPrefix(word, test) {
				i++
			}
		}
		results := streeng.StartWith(test)
		if len(results) == i {
			t.Logf("Test Successful: word: %s", test)
		} else {
			t.Errorf("Test Fail:\t word: %s \t expected: %d \t result: %d",
				test, i, len(results))
		}
	}
}

func TestEndtWith(t *testing.T) {
	fileName := "pp.txt"
	tests := []string{
		`etion`,
		`mation`,
		`nt`,
		`ne`,
		`ction`,
		`ame`,
		`ters`,
		`ure`,
		`rs.`,
		`asd`,
		`ance`,
		`bour`,
	}
	text, err := StringFromFile(fileName)
	if err != nil {
		t.Errorf("Test Fail: StringFromFile Error: %s\n", err.Error())
	}
	words := strings.Fields(text)
	streeng := MakeStreeng(words)
	streeng.ReverseStreeng()
	for _, test := range tests {
		i := 0
		for _, word := range streeng.words {
			if strings.HasSuffix(word, test) {
				i++
			}
		}
		results := streeng.EndWith(test)
		if len(results) == i {
			t.Logf("Test Successful: word: %s", test)
		} else {
			t.Errorf("Test Fail:\t word: %s \t expected: %d \t result: %d",
				test, i, len(results))
		}
	}
}

func TestTermDocument(t *testing.T) {
	fileName := "pp.txt"
	text, err := StringFromFile(fileName)
	if err != nil {
		t.Errorf("Test Fail: StringFromFile Error: %s\n", err.Error())
	}
	words := strings.Fields(text)
	streeng := MakeStreeng(words)
	streeng.Terms()
	for k, v := range streeng.tokens {
		if v == -1 {
			t.Errorf("Test Fail:\t %d of token is empty", k)
			break
		}
	}
}

func TestContains(t *testing.T) {
	fileName := "pp.txt"
	tests := []string{
		`12123123`,
		`asdasf`,
		`_?=)%'^++-*/`,
		`cvbdf`,
		`objection`,
		`wersfd`,
		`daughters`,
		`dfcbv`,
		``,
		`asd`,
		`importance`,
		`ggsdf`,
	}
	text, err := StringFromFile(fileName)
	if err != nil {
		t.Errorf("Test Fail: StringFromFile Error: %s\n", err.Error())
	}
	words := strings.Fields(text)
	streeng := MakeStreeng(words)
	for _, test := range tests {
		i := false
		for _, word := range streeng.words {
			if strings.Compare(test, word) == 0 {
				i = true
				break
			}
		}
		result := streeng.Contains(test)
		if result == i {
			t.Logf("Test Successful: word: %s", test)
		} else {
			t.Errorf("Test Fail:\t word: %s \t expected: %t \t result: %t",
				test, i, result)
		}
	}
}

func TestTraverse(t *testing.T) {
	text, err := StringFromFile("pp.txt")
	if err != nil {
		t.Errorf("Test Fail: StringFromFile Error: %s\n", err.Error())
	}
	words := strings.Fields(text)
	streeng := MakeStreeng(words)
	streeng.Terms()
	i := 0
	streeng.Traverse(func(n *Node) {
		if len(n.words) > i {
			i = len(n.words)
		}
	})
	j := 0
	for _, v := range streeng.terms {
		if v > j {
			j = v
		}
	}
	if i != j {
		t.Errorf("Test Fail:\t traverse: %d \t expected: %d", i, j)
	} else {
		t.Logf("Test Successful:\t traverse: %d \t expected: %d", i, j)
	}
}

func TestMatch(t *testing.T) {
	fileName := "pp.txt"
	tests := []string{
		`http(s)?://.*`,
		`c..t`,
		`want(s)?`,
		`.*c`,
		`.*ion`,
		``,
		`daughters`,
		`nd$`,
		`Mr(s)?\.`,
		`asdgh`,
		`^(The)`,
		`[0-9]+`,
	}
	text, err := StringFromFile(fileName)
	if err != nil {
		t.Errorf("Test Fail: StringFromFile Error: %s\n", err.Error())
	}
	words := strings.Fields(text)
	streeng := MakeStreeng(words)
	for _, test := range tests {
		re, err := regexp.Compile(test)
		if err != nil {
			t.Errorf("Test Fail:\t regex: %s\t could not compile", test)
		}
		i := 0
		for _, word := range streeng.words {
			if re.MatchString(word) {
				i++
			}
		}
		results, err2 := streeng.Match(test)
		if err2 != nil {
			t.Errorf("Test Fail:\t regex: %s\t error in MatchRegex", test)
		}
		if len(results) == i {
			t.Logf("Test Successful: regex: %s", test)
		} else {
			t.Errorf("Test Fail:\t regex: %s \t expected: %d \t result: %d",
				test, i, len(results))
		}
	}
}

func TestPerformanceMatch(t *testing.T) {
	fileName := "pp.txt"
	tests := []string{
		`http(s)?://.*`,
		`c..t`,
		`want(s)?`,
		`.*c`,
		`.*ion`,
		``,
		`daughters`,
		`nd$`,
		`Mr(s)?\.`,
		`asdgh`,
		`^(The)`,
		`[0-9]+`,
	}
	text, err := StringFromFile(fileName)
	if err != nil {
		t.Errorf("Test Fail: StringFromFile Error: %s\n", err.Error())
	}
	words := strings.Fields(text)
	streeng := MakeStreeng(words)
	var start, end time.Time

	var average int64
	durations := []int64{}
	average = 0
	for _, test := range tests {
		start = time.Now()
		streeng.Match(test)
		end = time.Now()
		durations = append(durations,
			end.Sub(start).Nanoseconds()/1000000)
	}
	for _, v := range durations {
		average += v
	}
	average /= int64(len(durations))
	fmt.Printf("Regex Average of Streeng:\t %dms\n", average)
}

func TestPerformanceSearch(t *testing.T) {
	fileName := "words.txt"
	tests := []string{
		`discretion`,
		`exclamation`,
		`want`,
		`fortune`,
		`objection`,
		`name`,
		`daughters`,
		`pleasure`,
		`Mrs.`,
		`asd`,
		`importance`,
		`neighbour`,
	}
	text, err := StringFromFile(fileName)
	if err != nil {
		t.Errorf("Test Fail: StringFromFile Error: %s\n", err.Error())
	}
	words := strings.Fields(text)
	streeng := MakeStreeng(words)

	var start, end time.Time

	start = time.Now()
	for _, test := range tests {
		streeng.Search(test)
	}
	end = time.Now()
	duration := end.Sub(start).Nanoseconds() / 1000000
	fmt.Printf("Find Duration of Streeng:\t %dms\n", duration)
}
