/*
@Time : 2022/11/12 21:23
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : trie
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package trie

// Node 子节点
type Node struct {
	Word     rune               // 当前节点存储的字符。byte只能表示英文字符，rune可以表示任意字符
	Child map[rune]*Node        // 子节点，用一个map存储
	Term     string		    // 标记关键词
}

// TrieTree Trie树
type TrieTree struct {
	root *Node
}

// add 把words[beginIndex:]插入到Trie树中
func (n *Node) add(words []rune, term string, beginIndex int) {
	if beginIndex >= len(words) { // words已经遍历完
		n.Term = term
		return
	}

	if n.Child == nil {
		n.Child = make(map[rune]*Node)
	}

	word := words[beginIndex] // 把这个word放到node的子节点中
	if child, exists := n.Child[word]; !exists {
		newNode := &Node{Word: word}
		n.Child[word] = newNode
		newNode.add(words, term, beginIndex+1) // 递归
	} else {
		child.add(words, term, beginIndex+1) // 递归
	}
}

// walk words[0]就是当前节点上存储的字符，按照words的指引顺着树往下走，最终返回words最后一个字符对应的节点
func (n *Node) walk(words []rune, beginIndex int) *Node {
	if beginIndex == len(words)-1 {
		return n
	}
	beginIndex += 1
	word := words[beginIndex]
	if child, exists := n.Child[word]; exists {
		return child.walk(words, beginIndex)
	} else {
		return nil
	}
}

// traverseTerms 遍历一个node下面所有的term。注意要传数组的指针，才能真正修改这个数组
func (n *Node) traverseTerms(terms *[]string) {
	if len(n.Term) > 0 {
		*terms = append(*terms, n.Term)
	}
	for _, child := range n.Child {
		child.traverseTerms(terms)
	}
}

// AddTerm 添加关键词
func (t *TrieTree) AddTerm(term string) {
	if len(term) <= 1 {
		return
	}
	words := []rune(term)

	if t.root == nil {
		t.root = new(Node)
	}

	t.root.add(words, term, 0)
}

// Retrieve 检索关键词
func (t *TrieTree) Retrieve(prefix string) []string {
	if t.root == nil || len(t.root.Child) == 0 {
		return nil
	}
	words := []rune(prefix)
	firstWord := words[0]
	if child, exists := t.root.Child[firstWord]; exists {
		end := child.walk(words, 0)
		if end == nil {
			return nil
		} else {
			terms := make([]string, 0, 100)
			end.traverseTerms(&terms)
			return terms
		}
	} else {
		return nil
	}
}
