package node

import (
	_ "assemble_server/core/component"
	"sync"

	"assemble_server/core/message"
	"assemble_server/core/component"
	"fmt"
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

	//通过标签删除所有符合条件的孩子节点
	RemoveChildNodesByTag(name string)

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
	//根据标签获得一组孩子节点
	GetChildNodesByTag(tag string)[]Node

	//设置标签
	SetTag(tag string)
	//获得标签
	GetTag()string

	//设置父节点
	SetParent(parent Node)

	//获得父节点
	GetParent()Node

	//通过名字给指定孩子节点发送消息
	SendMsg2Child(name string, msg message.Message)

	//通过名字给指定孩子节点发送消息
	SendMsg2Parent(msg message.Message)


	//通过名字给指定孩子节点发送消息
	SendMsg2GroupByTag(tag string, msg message.Message)

	//给孩子所有孩子节点发送消息
	Broadcast(msg message.Message)

	//获得消息通道
	GetMsgChannel() chan<-message.Message

	//添加组件
	AddComponent(key string,comp component.Component) component.Component

	//删除组件
	RemoveComponent(key string) component.Component

	//获取组件
	GetComponent(key string)component.Component

}

//节点类
type node struct {
	//名字
	name string
	//标签，可用于分组
	tag string
	//子节点列表
	childList []Node
	//父节点
	parent Node

	//父节点读写锁
	parentGuard sync.RWMutex
	//名字读写读写锁
	nameGuard sync.RWMutex
	//标记读写锁
	tagGuard sync.RWMutex
	//孩子切片读写锁
	childListGuard sync.RWMutex
	//消息通道
	msgChannel chan<- message.Message
	//组件表
	componentMap map[string]component.Component
	//组件表读写锁
	componentMapGuard sync.RWMutex


}

//创建一个节点方法
func  NewNode() Node {
	return &node{childList: make([]Node,0),
		         msgChannel: make(chan <-message.Message),
				 componentMap:make(map[string]component.Component)}
}

func (self * node)GetComponent(key string)component.Component  {
	self.componentMapGuard.RLock()
	defer self.componentMapGuard.RUnlock()
	return self.componentMap[key]
}

func (self *node)AddComponent(key string,comp component.Component) component.Component {
	self.componentMapGuard.Lock()
	defer self.componentMapGuard.Unlock()

	comp,exits:=self.componentMap[key]
	if exits{
		panic(fmt.Sprintf( "该key: %s 已经存在",key))
	}else {

		self.componentMap[key]=comp
	}

	return comp
}

func (self *node)RemoveComponent(key string) component.Component{
	self.componentMapGuard.Lock()
	defer self.componentMapGuard.Unlock()
	comp,exits:=self.componentMap[key]
	if exits{
		delete(self.componentMap,key)
	}else {

		panic(fmt.Sprintf( "该key: %s 不存在",key))
	}

	return comp
}



func (self *node) GetMsgChannel()chan<-message.Message  {

	return self.msgChannel

}

func (self *node)Broadcast(msg message.Message){
	childList:=self.GetChildNodes()
	for _,child:=range childList{
		child.GetMsgChannel()<-msg
	}
}

func (self *node) SendMsg2Child(name string, msg message.Message)  {
	targetNode:=self.GetChildNodeByName(name)
	targetNode.GetMsgChannel() <-msg

}



func (self *node) SendMsg2Parent(msg message.Message){
	if self.parent!=nil{
		self.parent.GetMsgChannel()<-msg
	}else {
		panic("该节点没有父节点，复发发送消息")
	}

}

func (self *node) SendMsg2GroupByTag(tag string, msg message.Message)  {
	childList:=self.GetChildNodesByTag(tag)
	for _,child:=range childList{
		child.GetMsgChannel()<-msg
	}

}


func (self *node)SetParent(parent Node)  {
	self.parentGuard.Lock()
	defer self.parentGuard.Unlock()

	self.parent=parent

}

func (self *node)GetParent() Node {
	self.parentGuard.RLock()
	defer self.parentGuard.RUnlock()
	return self.parent

}


func (self *node)SetTag(tag string)  {
	self.tagGuard.Lock()
	defer self.tagGuard.Unlock()
	self.tag = tag
}

func (self *node)GetTag() string  {
	self.tagGuard.RLock()
	defer self.tagGuard.RUnlock()

	return  self.tag
}

//设置节点名
func (self *node) SetName(name string)  {
	self.nameGuard.Lock()
	defer self.nameGuard.Unlock()

	self.name=name

}

//通过名称删除第一个孩子节点
func (self *node)RemoveChildNodeByName(name string) (removeNode Node)  {
	self.childListGuard.Lock()
	defer self.childListGuard.Unlock()

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
	self.childListGuard.Lock()
	defer self.childListGuard.Unlock()

	newList:=[]Node{}
	for i :=0; i <len(self.childList); i++{
		if self.childList[i].GetName()!=name{
			newList=append(newList,self.childList[i])
		}

	}

	self.childList=newList

}


//通过标签删除所有符合条件的孩子节点
func (self *node)RemoveChildNodesByTag(tag string)  {
	self.childListGuard.Lock()
	defer self.childListGuard.Unlock()

	newList:=[]Node{}
	for i :=0; i <len(self.childList); i++{
		if self.childList[i].GetTag()!=tag{
			newList=append(newList,self.childList[i])
		}

	}

	self.childList=newList

}


//获得节点名
func (self *node) GetName() string  {
	self.nameGuard.RLock()
	defer self.nameGuard.RUnlock()

	return self.name
}

//添加孩子节点
func (self *node) AddChildNode(child Node) Node {
	self.childListGuard.Lock()
	defer self.childListGuard.Unlock()

	self.childList=append(self.childList,child)
	return child

}

//获得所有孩子节点
func (self *node)GetChildNodes() []Node {
	self.childListGuard.RLock()
	defer self.childListGuard.RUnlock()

	len:=len(self.childList)
	return self.childList[0:len:len]
}
//根据名字返回第一个出现的孩子节点
func (self *node)GetChildNodeByName(name string) Node  {
	self.childListGuard.RLock()
	defer self.childListGuard.RUnlock()

	for _,child:= range self.childList{
		if child.GetName()==name{
			return child
		}
	}
	return nil
}

//根据名称获得相同名称的所有孩子节点
func (self *node)GetChildNodesByName(name string) []Node  {

	self.childListGuard.RLock()
	defer self.childListGuard.RUnlock()

	childList:=[]Node{}
	for _,child:= range self.childList{
		if child.GetName()==name{
			childList=append(childList,child)
		}
	}

	return childList
}

//根据标签获得一组孩子节点
func (self *node)GetChildNodesByTag(tag string) []Node  {
	self.childListGuard.RLock()
	defer self.childListGuard.RUnlock()

	childList:=[]Node{}
	for _,child:= range self.childList{
		if child.GetTag()==tag{
			childList=append(childList,child)
		}
	}

	return childList
}


func (self *node) ClearChildNodes() {
	self.childListGuard.Lock()
	defer self.childListGuard.Unlock()

	self.childList=[]Node{}

}
