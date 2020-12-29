package gee

import "strings"

//参数匹配符":"  如/p/:lang/doc 可以匹配/p/c/doc和/p/go/doc
//通配符"*"  如/static/*filepath  可以匹配/static/fav.ico和/static/js/jQuery.js

//前缀树节点定义
type node struct {
	pattern  string  //待匹配路由   如/p/:lang
	part     string  //路由中的一部分如:lang
	children []*node //子节点
	isWild   bool    //是否完整匹配,这是和普通树的区别所在
	//part 含有：或*时为true
}

//第一个匹配成功的节点,用于插入节点
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild { //判断这个新插入的part应该插入的位置
			return child
		}
	}
	return nil
}

//查找所有匹配成功的点，在Search函数中使用
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild { //找到包含的部分或完全匹配
			nodes = append(nodes, child)
		}
	}
	return nodes
}

//插入节点
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height { //如/p/:lang/doc只有在第三层节点
		//即doc节点，pattern才会设置为/p/:lang/doc
		//其实就是因为如果没到最后parts的长度会和height共同增长，
		//但是parts长度从1开始而height从0开始，到最后一个node时
		//parts不会再涨了因为
		n.pattern = pattern //[注意]pattern从头到尾都没有变过，就是最开始传入的pattern

		return
	}
	//安排递归参数
	part := parts[height]
	child := n.matchChild(part) //通过matchChild告诉你下一步该往哪走
	if child == nil {           //说明找不到child能满足要求了
		child = &node{ //新建一个节点
			part: part, //先初始化part=parts[height]
			//比如之前是/p/:lang/doc-->parts[2]=/doc=part
			isWild: part[0] == ':' || part[0] == '*',
			//如果part的首字母是通配符或者参数匹配符标记为完整匹配
		}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1) //递归，对part的进一步划分
		if result != nil {
			return result //返回节点
		}
	}
	return nil
}
