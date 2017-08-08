package node

import (
	_ "assemble_server/core/component"

)

//节点接口
type Node interface {
	//设置节点名称
	SetName(name string)
	//获得节点名
	GetName() string

	//添加孩子节点
	AddChildNode(child Node) Node

	//通过名称删除第一个符合条件的孩子节点
	RemoveChildNodeByName(name string) Node
	//通过名称删除所有符合条件的孩子节点
	RemoveChildNodesByName(name string)
	//删除所有孩子节点
	ClearChildNodes()

	//获得所有孩子节点
	GetChildNodes() []Node
	//根据名称获得孩子节点
	GetChildNodeByName(name string) Node

	//根据名称获得相同名称的孩子节点
	GetChildNodesByName(name string) []Node
	////添加组件
	//AddComponent(comp *component.Component) *component.Component
	////获得组件
	//GetComponent() *component.Component


}

//节点类
type node struct {
	name string
	childList []Node

}

//创建一个节点方法
func  NewNode() Node {
	return &node{childList: make([]Node,0)}
}

//设置节点名
func (self *node) SetName(name string)  {
	self.name=name

}

//通过名称删除第一个孩子节点
func (self *node)RemoveChildNodeByName(name string) (removeNode Node)  {
	index:=0
	for ;index<len(self.childList);index++{
		if self.childList[index].GetName()==name{
			break
		}
	}
	removeNode=self.childList[index]
	self.childList= append(self.childList[0:index],self.childList[index+1:]...)

	return
}



//通过名称删除所有符合条件的孩子节点
func (self *node)RemoveChildNodesByName(name string)  {

	newList:=[]Node{}
	for i :=0; i <len(self.childList); i++{
		if self.childList[i].GetName()!=name{
			newList=append(newList,self.childList[i])
		}

	}

	self.childList=newList

}
//获得节点名
func (self *node) GetName() string  {
	return self.name
}

//添加孩子节点
func (self *node) AddChildNode(child Node) Node {

	self.childList=append(self.childList,child)
	return child

}

//获得所有孩子节点
func (self *node)GetChildNodes() []Node {
	len:=len(self.childList)
	return self.childList[0:len:len]
}
//根据名字返回第一个出现的孩子节点
func (self *node)GetChildNodeByName(name string) Node  {
	for _,child:= range self.childList{
		if child.GetName()==name{
			return child
		}
	}
	return nil
}

//根据名称获得相同名称的所有孩子节点
func (self *node)GetChildNodesByName(name string) []Node  {
	childList:=[]Node{}
	for _,child:= range self.childList{
		if child.GetName()==name{
			childList=append(childList,child)
		}
	}

	return childList
}


func (self *node) ClearChildNodes() {

	self.childList=[]Node{}

}
