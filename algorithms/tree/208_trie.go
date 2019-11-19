package tree

type m map[string]string

type Trie struct {
	Node map[string]string
	Next *Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{Node: make(map[string]string)}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, s := range word {
		this.Node = map[string]string{s: s}
		if this == nil {
			this.Next = Constructor()
		}
		this = this.Next
	}
	this.Node = map[string]string{"#": "#"}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, s := range word {
		if _, found := this.Node[s]; found {
			if this.Next != nil {
				this = this.Next
				continue
			}
			return false
		}
		return false
	}
	return true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, s := range word {
		if _, found := this.Node[s]; found {
			if this.Next != nil {
				this = this.Next
				continue
			}
			return false
		}
		return false
	}
	return true
}
