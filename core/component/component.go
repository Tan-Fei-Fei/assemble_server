package component

import "assemble_server/core/message"

//组件接口
type Component interface {
	//当初始化的时候
	OnInit()
	//当组件被移除的时候
	OnRemove()
	//处理消息
	HandleMsg(msg message.Message)

}

type component struct {

}
