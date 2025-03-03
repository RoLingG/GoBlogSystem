package models

import "time"

const (
	AdminRole   = 1
	UserRole    = 2
	TouristRole = 3
)

type Model struct {
	ID       uint      `gorm:"primarykey" json:"id"`
	CreateAt time.Time `gorm:"default:current_timestamp(3)" json:"create_at"`
	UpdateAt time.Time `gorm:"default:current_timestamp(3)" json:"-"`

	//Error 1067 (42000): Invalid default value for 'create_at'
	//[0.000ms] [rows:0] ALTER TABLE `image_models` MODIFY COLUMN `create_at` datetime(3) NULL DEFAULT current_timestamp
	//这里我一开始会出现上面的报错是因为一开始设置的下面两个对象
	//CreateAt time.Time `json:"create_at"`
	//UpdateAt time.Time `json:"-"`
	//可以看到这两个对象并没有设置初始值，我看不顺眼给了个初始值（gorm:"default:current_timestamp"
	//然后衰仔的是time.Time在gorm模型里对应的mysql的类型是datetime(3)，妈的它默认设置了精度，所以加个精度就好了，虽然很蠢，但我还是得提醒自己
}

// APIPOST的参数，不设置默认值都是0
// 因为Page会被高频使用(是分页就会用到)，所以可以封装起来
type PageInfo struct {
	Page  int    `form:"page"`  //页数
	Key   string `form:"key"`   //模糊匹配的关键字
	Limit int    `form:"limit"` //每页限制显示量
	Sort  string `form:"sort"`  //排序
	Role  []int  `form:"role"`  //筛选角色权限
}

type Options[T any] struct {
	Label string `json:"label"`
	Value T      `json:"value"`
}

type IDRequest struct {
	ID string `json:"id" form:"id" uri:"id"`
}

type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}

type ESIDRequest struct {
	ID string `json:"id" form:"id" uri:"id"`
}

type ESIDListRequest struct {
	IDList []string `json:"id_list" binding:"required" msg:"请输入要取消的文章"`
}
