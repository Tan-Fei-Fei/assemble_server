package node

import (
	"testing"
	"fmt"
)

var inode Node

func TestNode_AddChildNode(t *testing.T) {

	inode= NewNode()
	inode.SetName("tom")

	inode2:= NewNode()
	//inode2.SetName("jack")
	fmt.Println(inode2.GetName())
	inode.AddChildNode(inode2)
	inode.AddChildNode(inode2)

	fmt.Println(inode.GetChildNodes())


}

func TestNode_GetChildNodesByName(t *testing.T) {
	aNode:=NewNode()
	bNode:=NewNode()
	bNode.SetName("tom")
	cNode:=NewNode()
	cNode.SetName("tom")
	aNode.AddChildNode(bNode)
	aNode.AddChildNode(cNode)
	list:=aNode.GetChildNodesByName("tom")
	for i:=0;i<len(list) ;i++  {
		fmt.Println(list[i].GetName())
	}
}
