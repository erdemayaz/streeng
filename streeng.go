package streeng

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
	"regexp"
	"sync"
)

// Node is a struct of Streeng node
type Node struct {
	value       rune
	words       []int
	numberWords int
	characters  map[rune]*Node
}

// Streeng is a struct of Streeng
type Streeng struct {
	root         *Node
	reverseRoot  *Node
	words        []string
	nodeCount    int
	reverseCount int
	depth        int
	terms        map[string]int
	tokens       []int
	rate         float64
}

/*
StringFromFile function reads bytes from the file
and returns file's data as string
*/
func StringFromFile(fileName string) (string, error) {
	abs, err := filepath.Abs(fileName)
	if err != nil {
		return "", err
	}
	bytes, err2 := ioutil.ReadFile(abs)
	if err2 != nil {
		return "", err2
	}
	return string(bytes), nil
}

/*
StringFromURL function reads bytes from the URL content
and returns content's data as string
*/
func StringFromURL(url string) (string, bool) {
	res, err := http.Get(url)
	if err != nil {
		return "", false
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		data, err2 := ioutil.ReadAll(res.Body)
		if err2 != nil {
			panic("reading failed from body bytes")
		}
		return string(data), true
	}
	return "", false
}

/*
MakeStreeng makes a streeng struct with given string array
and it builds new tree. But reverse tree will not build
*/
func MakeStreeng(words []string) *Streeng {
	root := new(Node)
	root.characters = make(map[rune]*Node)
	root.words = nil
	s := new(Streeng)
	s.root = root
	s.nodeCount = 1
	s.depth = 0
	for k, v := range words {
		addString(s, k, v)
	}
	s.reverseCount = -1
	s.words = words
	s.rate = float64(len(words)) / float64(s.nodeCount)
	s.terms = nil
	s.tokens = nil
	s.reverseRoot = nil
	return s
}

// ReverseStreeng makes reverse tree and attach streeng
func (s *Streeng) ReverseStreeng() *Node {
	reverseRoot := new(Node)
	reverseRoot.characters = make(map[rune]*Node)
	reverseRoot.words = nil
	tempNode := reverseRoot
	count := 0
	for k, v := range s.words {
		tempNode = reverseRoot
		runic := []rune(v)
		lenOfValue := len(runic)
		for i := (lenOfValue - 1); i >= 0; i-- {
			isLast := i == 0
			if val, ok := tempNode.characters[runic[i]]; ok {
				tempNode = val
				if isLast {
					tempNode.words = append(tempNode.words, k)
				}
			} else {
				count++
				n := new(Node)
				n.characters = make(map[rune]*Node)
				n.value = runic[i]
				if isLast {
					n.words = append(n.words, k)
				}
				tempNode.characters[runic[i]] = n
				tempNode = tempNode.characters[runic[i]]
			}
		}
	}
	s.reverseRoot = reverseRoot
	s.reverseCount = count
	return reverseRoot
}

// Traverse function traverses nodes on given tree
func (s *Streeng) Traverse(sc func(*Node)) {
	if s != nil && s.root != nil {
		traverseChild(s.root, sc)
	}
}

// GoTraverse function traverses nodes on given tree with goroutines
func (s *Streeng) GoTraverse(sc func(*Node)) {
	if s != nil && s.root != nil {
		var wg sync.WaitGroup
		wg.Add(len(s.root.characters))
		for _, v := range s.root.characters {
			go goTraverseChild(&wg, v, sc)
		}
		wg.Wait()
	}
}

// Clean function cleans the tree
func (s *Streeng) Clean() {
	if s != nil && s.root != nil {
		cleanChild(s.root)
		cleanChild(s.reverseRoot)
		s.root.words = nil
		s.root.characters = nil
		s.words = nil
		s.nodeCount = 1
		s.reverseCount = 1
		s.terms = nil
		s.tokens = nil
	}
}

// Search function searches given word in the tree
func (s *Streeng) Search(word string) []int {
	runic := []rune(word)
	lenOfWord := len(runic)
	if s != nil && s.root != nil && lenOfWord > 0 {
		tempNode := s.root
		for i := 0; i < lenOfWord; i++ {
			if tempNode.characters[runic[i]] != nil {
				tempNode = tempNode.characters[runic[i]]
			} else {
				return nil
			}
		}
		return tempNode.words
	}
	return nil
}

// Match function matches words with given regular expression
func (s *Streeng) Match(regex string) ([]int, error) {
	re, err := regexp.Compile(regex)
	if err != nil {
		return nil, err
	}
	var mutex sync.Mutex
	results := []int{}
	s.GoTraverse(func(node *Node) {
		if re.MatchString(s.words[node.words[0]]) {
			for _, index := range node.words {
				mutex.Lock()
				results = append(results, index)
				mutex.Unlock()
			}
		}
	})
	return results, nil
}

// StartWith function searches words which start with given string
func (s *Streeng) StartWith(word string) []int {
	runic := []rune(word)
	lenOfWord := len(runic)
	if s != nil && s.root != nil && lenOfWord > 0 {
		tempNode := s.root
		for i := 0; i < lenOfWord; i++ {
			if tempNode.characters[runic[i]] != nil {
				tempNode = tempNode.characters[runic[i]]
			} else {
				return nil
			}
		}
		words := []int{}
		for _, v := range tempNode.words {
			words = append(words, v)
		}
		getSubstring(&words, tempNode)
		return words
	}
	return nil
}

// EndWith function searches words which end with given string
func (s *Streeng) EndWith(word string) []int {
	runic := []rune(word)
	lenOfWord := len(runic)
	if s != nil && s.reverseRoot != nil && lenOfWord > 0 {
		tempNode := s.reverseRoot
		for i := (lenOfWord - 1); i >= 0; i-- {
			if tempNode.characters[runic[i]] != nil {
				tempNode = tempNode.characters[runic[i]]
			} else {
				return nil
			}
		}
		words := []int{}
		for _, v := range tempNode.words {
			words = append(words, v)
		}
		getSubstring(&words, tempNode)
		return words
	}
	return nil
}

// Terms function calculates term of tree with frequency as map
func (s *Streeng) Terms() map[string]int {
	if s != nil {
		s.terms = make(map[string]int)
		s.tokens = make([]int, len(s.words))
		for j := 0; j < len(s.words); j++ {
			s.tokens[j] = -1
		}
		i := 1
		collectTerm(s, s.root, &i)
	}
	return s.terms
}

// FindFreqTerms reports frequent of terms bigger than minimum value
func (s *Streeng) FindFreqTerms(min int) map[string]int {
	if s != nil && s.terms != nil && min >= 0 {
		if min > 0 {
			terms := make(map[string]int)
			for k, v := range s.terms {
				if v >= min {
					terms[k] = v
				}
			}
			return terms
		}
		return s.terms
	}
	return nil
}

// Contains returns whether or not the word exists
func (s *Streeng) Contains(word string) bool {
	runic := []rune(word)
	lenOfWord := len(runic)
	if s != nil && s.root != nil && lenOfWord > 0 {
		tempNode := s.root
		for i := 0; i < lenOfWord; i++ {
			if tempNode.characters[runic[i]] != nil {
				tempNode = tempNode.characters[runic[i]]
			} else {
				return false
			}
		}
		return len(tempNode.words) > 0
	}
	return false
}

// Depth returns depth of streeng
func (s *Streeng) Depth() int {
	return s.depth
}

// Words returns element of words
func (s *Streeng) Words(index int) string {
	if index >= 0 {
		return s.words[index]
	}
	return ""
}

// NodeCount returns count of streeng's tree
func (s *Streeng) NodeCount() int {
	return s.nodeCount
}

// ReverseNodeCount returns count of streeng's reverse tree
func (s *Streeng) ReverseNodeCount() int {
	return s.reverseCount
}

// Rate returns rate streeng
func (s *Streeng) Rate() float64 {
	return s.rate
}

// TermList returns list of terms
func (s *Streeng) TermList() map[string]int {
	return s.terms
}

// TokenList returns list of tokens
func (s *Streeng) TokenList() []int {
	return s.tokens
}

// Value returns rune value of node
func (n *Node) Value() rune {
	return n.value
}

/*
Words returns word of index,
if index is smaller than 0 or
bigger than number of words, it returns -1
*/
func (n *Node) Words(index int) int {
	if index >= 0 && index < n.numberWords {
		return n.words[index]
	}
	return -1
}

/*
Character returns node's rune child.
If there is not, it returns nil
*/
func (n *Node) Character(char rune) *Node {
	if val, ok := n.characters[char]; ok {
		return val
	}
	return nil
}

// NumberWords returns number of node's words
func (n *Node) NumberWords() int {
	return n.numberWords
}

func getSubstring(words *[]int, node *Node) {
	if node != nil {
		for _, v := range node.words {
			*words = append(*words, v)
		}
		for _, v := range node.characters {
			getSubstring(words, v)
		}
	}
}

func addString(s *Streeng, index int, value string) {
	tempNode := s.root
	runic := []rune(value)
	lenOfValue := len(runic)
	if lenOfValue > s.depth {
		s.depth = lenOfValue
	}
	for i := 0; i < lenOfValue; i++ {
		isLast := i+1 == lenOfValue
		if val, ok := tempNode.characters[runic[i]]; ok {
			tempNode = val
			if isLast {
				tempNode.words = append(tempNode.words, index)
				tempNode.numberWords++
			}
		} else {
			n := new(Node)
			n.characters = make(map[rune]*Node)
			n.value = runic[i]
			n.numberWords = 0
			s.nodeCount++
			if isLast {
				n.words = append(n.words, index)
				n.numberWords++
			}
			tempNode.characters[runic[i]] = n
			tempNode = tempNode.characters[runic[i]]
		}
	}
}

func collectTerm(s *Streeng, node *Node, i *int) {
	if node != nil {
		if len(node.words) > 0 {
			s.terms[s.words[node.words[0]]] = len(node.words)
			for _, v := range node.words {
				s.tokens[v] = *i
			}
			*i++
		}
		for _, v := range node.characters {
			collectTerm(s, v, i)
		}
	}
}

func traverseChild(node *Node, sc func(*Node)) {
	if node != nil {
		if len(node.words) > 0 {
			sc(node)
		}
		for _, v := range node.characters {
			traverseChild(v, sc)
		}
	}
}

func goTraverseChild(wg *sync.WaitGroup, node *Node, sc func(*Node)) {
	if node != nil {
		if len(node.words) > 0 {
			sc(node)
		}
		for _, v := range node.characters {
			traverseChild(v, sc)
		}
	}
	wg.Done()
}

func cleanChild(node *Node) {
	if node != nil {
		for _, v := range node.characters {
			cleanChild(v)
		}
		node.words = nil
		node.characters = nil
	}
}
