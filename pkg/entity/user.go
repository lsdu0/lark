package entity

type User struct {
	GormEntityTs
	Uid         int64  `gorm:"column:uid;primary_key" json:"uid"`                          // 用户ID 系统生成
	LarkId      string `gorm:"column:lark_id;NOT NULL" json:"lark_id"`                     // 账户ID 用户设置
	Password    string `gorm:"column:password;NOT NULL" json:"password"`                   // 密码
	Udid        string `gorm:"column:udid;NOT NULL" json:"udid"`                           // 注册设备唯一标识
	Status      int    `gorm:"column:status;default:0;NOT NULL" json:"status"`             // 用户状态
	Nickname    string `gorm:"column:nickname;NOT NULL" json:"nickname"`                   // 昵称
	Firstname   string `gorm:"column:firstname;NOT NULL" json:"firstname"`                 // firstname
	Lastname    string `gorm:"column:lastname;NOT NULL" json:"lastname"`                   // lastname
	Gender      int    `gorm:"column:gender;default:0;NOT NULL" json:"gender"`             // 性别
	BirthTs     int64  `gorm:"column:birth_ts;default:0;NOT NULL" json:"birth_ts"`         // 生日
	Email       string `gorm:"column:email;NOT NULL" json:"email"`                         // Email
	Mobile      string `gorm:"column:mobile;NOT NULL" json:"mobile"`                       // 手机号
	RegPlatform int    `gorm:"column:reg_platform;default:0;NOT NULL" json:"reg_platform"` // 注册平台
	Platform    int    `gorm:"column:platform;default:0;NOT NULL" json:"platform"`         // 最后登录的平台
	ServerId    int    `gorm:"column:server_id;default:0;NOT NULL" json:"server_id"`       // 分配的ws服务器
	CityId      int    `gorm:"column:city_id;default:0;NOT NULL" json:"city_id"`           // 城市ID
}
