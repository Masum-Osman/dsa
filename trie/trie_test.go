package trie

import "fmt"

func TestTrie() {
	goTrie := InitTrie()

	toAdd := []string{
		"aragorn",
		"aragon",
		"argon",
		"eragon",
		"oregon",
		"oregano",
		"oreo",
	}

	for i := 0; i < len(toAdd); i++ {
		goTrie.InsertWord(toAdd[i])
	}

	fmt.Println(goTrie.Search("oreo"))
}
