package gee

import (
	"fmt"
	"strings"
)

// Trie 树节点
type node struct {
	pattern  string  // 待匹配的路由
	part     string  // 当前节点对应部分
	children []*node // 子节点
	isWild   bool    // 通配标志
}

func (n *node) String() string {
	return fmt.Sprintf("node{pattern= %s, part= %s, isWild=%t", n.pattern, n.part, n.isWild)
}

// 寻找第一个匹配的节点
func (n *node) matchChild(part string) *node {
	// 查找子节点
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 寻找所有匹配的节点
func (n *node) matchChildren(part string) *[]*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return &nodes
}

// 插入新节点
func (n *node) insert(pattern string, parts []string, height int) {
	// 递归出口
	if len(parts) == height {
		// 记录匹配模式
		n.pattern = pattern
		return
	}

	part := parts[height]
	// 查找匹配节点
	child := n.matchChild(part)
	if child == nil {
		// 未找到则新增节点
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}

	// 递归新增节点
	child.insert(pattern, parts, height+1)
}

// 查找匹配节点
func (n *node) search(parts []string, height int) *node {
	// 递归出口
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		// 匹配模式为空，未找到
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	// 查找所有匹配的节点
	children := n.matchChildren(part)

	for _, child := range *children {
		// 递归搜索每个子节点
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}

// 遍历前缀树，获取所有的路由模式节点
func (n *node) travel(list *[]*node) {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		child.travel(list)
	}
}
