package pe

func NewTrie() *Trie {
	t := new(Trie)
	t.trieMap = make(map[string]*Trie)
	return t
}

type Trie struct {
	endsWord bool
	trieMap  map[string]*Trie
}

func (t *Trie) AddWord(s string) {
	if len(s) == 0 {
		t.endsWord = true
		return
	}
	ch := string(s[0])
	if _, ok := t.trieMap[ch]; !ok {
		newTrie := NewTrie()
		t.trieMap[ch] = newTrie
	}
	t.trieMap[ch].AddWord(s[1:])
	return
}

func (t *Trie) IsWord(s string) bool {
	if len(s) == 0 {
		if t.endsWord {
			return true
		} else {
			return false
		}
	}
	ch := string(s[0])
	if _, ok := t.trieMap[ch]; !ok {
		return false
	}
	return t.trieMap[ch].IsWord(s[1:])
}

func (t *Trie) getWordsHelper(s string) []string {
	ret := []string{}
	if t.endsWord {
		ret = append(ret, s)
	}
	for ch, ptr := range t.trieMap {
		ret = append(ret, ptr.getWordsHelper(s+string(ch))...)
	}
	return ret
}

func (t *Trie) GetAllWords() []string {
	ret := []string{}
	for ch, ptr := range t.trieMap {
		ret = append(ret, ptr.getWordsHelper(string(ch))...)
	}
	return ret
}
