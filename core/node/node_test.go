package node

import (
	"testing"
	"fmt"
)



func TestNode_AddChildNode(t *testing.T) {
	t.Log("******添加孩子节点测试开始*******")
	inode:= NewNode()
	inode.SetName("tom")

	inode2:= NewNode()
	//inode2.SetName("jack")
	fmt.Println(inode2.GetName())
	inode.AddChildNode(inode2)
	inode.AddChildNode(inode2)

	t.Log(inode.GetChildNodes())


}

func TestNode_RemoveChildNodeByName(t *testing.T) {
	t.Log("******删除孩子节点测试*****")
	aNode:=NewNode()
	bNode:=NewNode()
	bNode.SetName("tom")
	cNode:=NewNode()
	cNode.SetName("jack")
	dNode:=NewNode()
	dNode.SetName("tom")
	aNode.AddChildNode(bNode)
	aNode.AddChildNode(cNode)
	deleteNode :=aNode.RemoveChildNodeByName("jack")
	for i:=0;i<len(aNode.GetChildNodes()) ;i++  {
		t.Log(aNode.GetChildNodes()[i].GetName())
	}
	t.Log("删除的元素",deleteNode.GetName())
}

func TestNode_GetChildNodesByName(t *testing.T) {
	t.Log("******根据名字获得多个节点******")
	aNode:=NewNode()
	bNode:=NewNode()
	bNode.SetName("tom")
	cNode:=NewNode()
	cNode.SetName("tom")
	aNode.AddChildNode(bNode)
	aNode.AddChildNode(cNode)
	list:=aNode.GetChildNodesByName("tom")
	for i:=0;i<len(list) ;i++  {
		t.Log(list[i].GetName())
	}
}

func TestNode_RemoveChildNodesByName(t *testing.T) {
	t.Log("******删除符合名字的所有孩子节点测试*****")
	aNode:=NewNode()


	bNode:=NewNode()
	bNode.SetName("tom")
	cNode:=NewNode()
	cNode.SetName("jack")
	dNode:=NewNode()
	dNode.SetName("tom")
	aNode.AddChildNode(bNode)
	aNode.AddChildNode(cNode)
	aNode.AddChildNode(dNode)
	aNode.RemoveChildNodesByName("tom")
	for i:=0;i<len(aNode.GetChildNodes()) ;i++  {
		t.Log(aNode.GetChildNodes()[i].GetName())
	}

}

