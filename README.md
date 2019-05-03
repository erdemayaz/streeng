# Streeng
Streeng is an indexing tool for string arrays. If you make `MakeStreeng` with string array, it builds new rune tree.

## How does it work?

Example: This is a text to test

String array: `This`, `is`, `a`, `text`, `to`, `test`

String tree:

![Streeng Tree](https://github.com/erdemayaz/streeng/blob/master/assets/tree.png)

## Functions
|Name| Description | Parameter(s) | Return |
|--|--|--|--|
| `MakeStreeng` | It makes a streeng struct with given string array | []string | *streeng.Streeng
| `Search` | This function searches given word in the tree | string| []int | 
| `Match` | It matches words with given regular expression | string | []int |
| `StartWith` | It searches words which start with given string | string | []int | 
| `EndWith` | It searches words which end with given string | string | []int | 
| `Contains` | It returns whether or not the word exists | string | bool |
| `Terms` | It calculates term of tree with frequency as map | | map[string]int | 
| `FindFreqTerms` | It reports frequent of terms bigger than min value | int | map[string]int | 
| `Traverse` | Traverse function traverses nodes on given tree | func(*streeng.Node)|  |
| `GoTraverse` | It traverses nodes on given tree with goroutines | func(*streeng.Node) |  |
| `Clean` | Clean function cleans the tree | |  |
| `ReverseStreeng` | It makes reverse tree and attach streeng | | *streeng.Node |
| `StringFromFile` | It reads bytes from the file | string | string, error |
| `StringFromURL` | It reads bytes from the URL content | string | string, bool |
| `Depth` | It returns depth of streeng | | int |
| `Words` | It returns element of words | int | string |
| `NodeCount` | It returns count of streeng's tree | | int |
| `ReverseNodeCount` | It returns count of streeng's reverse tree | | int |
| `Rate` | It returns rate streeng | | float64 |
| `TermList` | It returns list of terms | | map[string]int |
| `TokenList` | It returns list of tokens | | []int |
| `Value` | It returns rune value of node | | rune |
| `Words` | It returns word of index | int | int |
| `Character` | It returns node's rune child | rune | *streeng.Node |
