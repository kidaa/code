package vo


// 广播消息结构
type Broadcast struct {
	Msg     []byte		// 消息体
	Channel int			// -1:全服广播，<10000:世界广播，>=10000:个人广播
	Kick bool			// 是否是T人，只针对个人广播才生效,Msg数据不会被发送
}
