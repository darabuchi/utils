package utils

import (
	"sync"
)

type TrieNode struct {
	Val    rune
	Depth  int
	Count  int // 统计分支数量
	Child  map[rune]*TrieNode
	IsWord bool // 标记是否为一个完整的字符串
}

// 创建节点
func NewTrieNode() *TrieNode {
	return &TrieNode{Child: make(map[rune]*TrieNode)}
}

type TrieTree struct {
	Root *TrieNode
	lock sync.RWMutex
}

func NewTrieTree() *TrieTree {
	return &TrieTree{Root: NewTrieNode()}
}

func (p *TrieTree) Add(words ...string) {
	for _, word := range words {
		p.add(word)
	}
}

// 添加节点
func (p *TrieTree) add(word string) {
	if len(word) == 0 {
		return
	}

	p.lock.Lock()
	defer p.lock.Unlock()

	bt := []rune(word)
	node := p.Root
	for _, val := range bt {
		child, ok := node.Child[val]
		if !ok {
			child = NewTrieNode()
			child.Val = val
			node.Child[val] = child
			node.Count += 1
			child.Depth = node.Depth + 1
		}
		node = child
	}
	node.IsWord = true
}

func (p *TrieTree) Find(word string) bool {
	bt := []rune(word)
	node := p.Root

	p.lock.RLock()
	defer p.lock.RUnlock()

	for _, val := range bt {
		child, ok := node.Child[val]
		if !ok {
			return false
		}
		node = child
	}
	return node.IsWord
}

func (p *TrieTree) Del(words ...string) {
	for _, word := range words {
		p.add(word)
	}
}

// 节点的删除分为如下种情况
// 1. 前缀的删除：判断Count是否大于0， 标记IsWord 为false。
// 3. 字符串的删除：
//     a. 如果是无分支，则整个删除。
//     b. 如果有分支，仅删除不是共有前缀的部分。
func (p *TrieTree) del(word string) {
	bt := []rune(word)
	if len(word) == 0 {
		return
	}

	p.lock.Lock()
	defer p.lock.Unlock()

	node := p.Root
	var lastBranch *TrieNode
	var delVal rune

	for index, val := range bt {
		child, ok := node.Child[val]
		if ok {
			if child.Count > 1 {
				lastBranch = child
				delVal = bt[index+1]
			}
		}
		node = child
	}

	if node.Count > 0 {
		// 删除前缀
		node.IsWord = false
	} else {
		if lastBranch == nil {
			// 删除整个字符串
			lastBranch = p.Root
			delVal = bt[0]
		}
		delete(lastBranch.Child, delVal)
		lastBranch.Count -= 1
	}
}

func (p *TrieTree) Associate(word string, max int) []string {
	if max <= 0 {
		return nil
	}

	p.lock.RLock()
	defer p.lock.RUnlock()

	bt := []rune(word)
	node := p.Root

	for _, val := range bt {
		child, ok := node.Child[val]
		if !ok {
			return nil
		}

		node = child
	}

	var words []string
	var cnt int
	var findSub func(node *TrieNode, base string)
	findSub = func(node *TrieNode, base string) {
		if cnt >= max {
			return
		}
		base += string(node.Val)
		if node.IsWord {
			words = append(words, base)
			cnt++
		}

		for _, trieNode := range node.Child {
			findSub(trieNode, base)
		}
	}

	if len(bt) == 0 {
		findSub(node, "")
	} else {
		findSub(node, string(bt[:len(bt)-1]))
	}

	return words
}
