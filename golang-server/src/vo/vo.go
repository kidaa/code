package vo

//用户的全部数据，服务器存取专用
type UserData struct {
	Userid     int    // 用户id
	Nickname   string // 用户昵称
	Remark     string // 好友的备注
	Sex        int    // 用户性别,男1 女2 非男非女3
	Headpic    int    // 用户头像,其实就是用户创建时的Unix时间，例如：http://image.wa.com/headpic/2015/05/26/01/60434/80_60434.jpg。80表示头像大小
	WorldID    int    //世界ID
	Provinceid int    // 省份ID
	Cityid     int    // 市ID
	Sign       string // 用户签名
	Isfriend   int    // 是否是好友
	Email      string // 绑定的邮箱地址
	Phone      string //	绑定的手机号码
	Auth       string // 密码验证码
	Pwd        string // MD5密码
	Qq_uid     string

	Birthday       int64   // 用户生日日期
	CreaterIP      int64   // 	创建账户是的IP地址
	Currency       int     // 货币
	Grade          int     // 用户等级
	Exp            float32 // 经验
	Score          int     // 积分
	TerminalType   int     // 终端类型 1:PC，2:手机
	Star           int     // 星座 1:白羊 2:金牛 3:双子 4:巨蟹 5:狮子 6:处女 7:天秤 8:天蝎 9:射手 10:摩羯 11:水瓶 12:双鱼
	UserType       int     // 普通用户1 付费用户2  商家3
	ValidateStatus int     // 验证状态；0：手机、邮箱均为验证，1：手机已验证，2：邮箱已验证，3：手机邮箱均已通过验证
	Status         int     // 正常1  锁定2  黑名单3

	Session string // 登陆session
}

//聊天消息数据
type MsgData struct {
	Userid int    `json:"userid"`
	Msg    string `json:"msg"`
	//	Status   int    `json:"status"`
	Sendtime int64 `json:"sendtime"`
}

// 好友关系数据
type Relationship struct {
	GroupID    int
	FriendID   int
	Remarks    string
	Status     int // 友好关系; 默认值1:正常, 黑名单:2
	CreateTime int
}

//用户的道具
type WidgetData struct {
	Id         int   `json:"id"`         //物品的唯一id，如：道具编号，奖品id
	Type       int   `json:"type"`       // 物品类别； 1：道具， 2：奖品
	Count      int   `json:"count"`      // 物品总数量
	Position   int   `json:"position"`   //仓库货位编号
	CreateTime int64 `json:"createTime"` //入库时间
}

//缘分纸条的数据结构
type KarmaData struct {
	ID         int
	SendID     int    //	发送人
	WidgetID   int    // 物品的唯一id
	ReadTimes  int    // 纸条被读取的次数
	Content    string //纸条携带的消息
	Provinceid int    // 省份ID
	Cityid     int    // 市ID
	Sex        int
	Birthday   int64
	Status     int   //0:正常状态，1待清楚状态
	CreateTime int64 //入库时间
}

// 缘分关系
type RelationKarma struct {
	Userid     int
	Friendid   int
	Status     int
	CreateTime int64
}

// 对白及属性结构
type DialogData struct {
	Dialogid     int    //   对白的id
	Abspath      string //对白的位置
	Userid       int    // 发布对白的用户id
	Dialogtype   int    // 对白类型；默认值1:用户对白， 2:原始对白
	Content      string
	Worldid      int
	Buildid      int
	Status       int // 默认值0:正常, 1:屏蔽
	Isanno       int // 是否匿名; 默认值0:实名，1:匿名
	CreateTime   int64
	TerminalType int
	Upworth      int
	UpFlag       int //置顶标示位; 默认值0,1:置顶
	UpRemainTime int //置顶剩余时间, 默认值0
	UpTotalTime  int // 置顶总共已经停留时间, 默认值0

	UpEndTime int // 当处于置顶时的结束时间; 默认值0
	FollowNum int // 跟帖数, 默认值0
	HotNum    int //热度数, 默认值0
}


//新闻资讯数据结构
type NewsData struct {
	ID         int    //
	Category   string //
	Title      string //
	Source     string //
	Url        string //
	CreateTime int64  //
}

//大楼云、部分数据结构
type BuildPartCloud struct {
	AccountNum int
	Buildid    int    //
	Part       string //
	Cloud      string //
}

//用户所属世界
type UserWorld struct {
	AccountNum int
	WorldID    int //
	Currently  int // 是否当前世界 0:否 1:是
}
