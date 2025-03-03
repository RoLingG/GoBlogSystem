package routers

import (
	"GoRoLingG/api"
	"GoRoLingG/middleware"
)

func (router RouterGroup) LargeScaleModelRouter() {
	largeScaleModelApi := api.ApiGroupApp.LargeScaleModelApi
	// 配置相关路由
	{
		router.GET("/large_scale_model/usable", middleware.JwtAdmin(), largeScaleModelApi.UsableModelListView)                    // 可用大模型列表获取接口
		router.GET("/large_scale_model/setting", middleware.JwtAuth(), largeScaleModelApi.ModelSettingView)                       // 大模型配置列表获取接口
		router.PUT("/large_scale_model/update", middleware.JwtAdmin(), largeScaleModelApi.ModelSettingUpdateView)                 //大模型配置更新接口
		router.GET("/large_scale_model/session_setting", middleware.JwtAdmin(), largeScaleModelApi.ModelSessionSettingView)       //大模型会话配置列表获取接口
		router.PUT("/large_scale_model/session_setting", middleware.JwtAdmin(), largeScaleModelApi.ModelSessionSettingUpdateView) //大模型会话配置更新接口
		router.PUT("/large_scale_model/auto_reply", middleware.JwtAdmin(), largeScaleModelApi.AutoReplyUpdateView)                //大模型自动回复更新接口
		router.GET("/large_scale_model/auto_reply", middleware.JwtAdmin(), largeScaleModelApi.AutoReplyListView)                  //大模型自动回复列表获取接口
		router.DELETE("/large_scale_model/auto_reply", middleware.JwtAdmin(), largeScaleModelApi.AutoReplyDeleteView)             //大模型自动回复删除接口
	}

	// 用户积分相关路由
	{
		router.GET("/large_scale_model/scope", middleware.JwtAuth(), largeScaleModelApi.UserScopeEnableView) //用户积分是否可领取接口
		router.POST("/large_scale_model/scope", middleware.JwtAuth(), largeScaleModelApi.UserGetScopeView)   //用户领取积分接口
	}

	//大模型角色相关路由
	{
		router.PUT("/large_scale_model/tag", middleware.JwtAdmin(), largeScaleModelApi.LargeScaleModelTagUpdateView)             //大模型角色标签更新接口
		router.GET("/large_scale_model/tag", middleware.JwtAdmin(), largeScaleModelApi.LargeScaleModelTagListView)               //大模型角色标签列表获取接口
		router.GET("/large_scale_model/tag/options", middleware.JwtAdmin(), largeScaleModelApi.LargeScaleModelTagOptionListView) //大模型角色标签选项获取接口
		router.DELETE("/large_scale_model/tag", middleware.JwtAdmin(), largeScaleModelApi.LargeScaleModelTagDeleteView)          //大模型角色标签删除接口
		router.PUT("/large_scale_model/role", middleware.JwtAdmin(), largeScaleModelApi.ModelRoleUpdateView)                     //大模型角色更新接口
		router.PUT("/large_scale_model/role_part", middleware.JwtAdmin(), largeScaleModelApi.ModelRolePartUpdateView)            //大模型劫色部分属性更新接口
		router.GET("/large_scale_model/roleList", middleware.JwtAdmin(), largeScaleModelApi.ModelRoleListView)                   //大模型角色列表获取接口
		router.DELETE("/large_scale_model/role", middleware.JwtAdmin(), largeScaleModelApi.ModelRoleDeleteView)                  //大模型角色删除接口
		router.GET("/large_scale_model/role_tag", middleware.JwtAuth(), largeScaleModelApi.LargeScaleModelRoleTagListView)       //大模型角色广场接口
		router.GET("/large_scale_model/role_info/:id", middleware.JwtAuth(), largeScaleModelApi.ModelRoleInfoView)               //大模型角色信息详情接口
		router.GET("/large_scale_model/role_icon/options", middleware.JwtAuth(), largeScaleModelApi.ModelRoleIconsView)          //大模型角色头像获取接口
	}

	//大模型会话相关路由
	{
		router.POST("/large_scale_model/session", middleware.JwtAuth(), largeScaleModelApi.ModelSessionCreateView)                   //大模型会话创建接口
		router.PUT("/large_scale_model/session", middleware.JwtAuth(), largeScaleModelApi.ModelSessionUpdateView)                    //大模型会话更新接口
		router.DELETE("/large_scale_model/session/:id", middleware.JwtAuth(), largeScaleModelApi.ModelSessionDeleteView)             //大模型会话删除接口
		router.DELETE("/large_scale_model/session_list", middleware.JwtAdmin(), largeScaleModelApi.ModelSessionListDeleteView)       //大模型会话批量删除接口
		router.GET("/large_scale_model/historical_session", middleware.JwtAuth(), largeScaleModelApi.ModelHistoricalSessionListView) //大模型历史会话记录获取接口
		router.GET("/large_scale_model/session", middleware.JwtAuth(), largeScaleModelApi.ModelSessionListView)                      //大模型会话列表获取接口
		router.GET("/large_scale_model/role_session", middleware.JwtAuth(), largeScaleModelApi.ModelRoleSessionsView)                //大模型角色会话列表获取接口
	}

	//大模型对话相关路由
	{
		router.GET("/large_scale_model/chat_sse", largeScaleModelApi.ModelChatCreateView)                              //大模型对话创建接口	Get请求是因为要从大模型api那里拿会话数据
		router.GET("/large_scale_model/chat_list", middleware.JwtAuth(), largeScaleModelApi.ModelChatListView)         //大模型对话列表获取接口
		router.DELETE("/large_scale_model/chat/:id", middleware.JwtAuth(), largeScaleModelApi.ModelUserChatDeleteView) //大模型普通用户对话删除接口
		router.DELETE("/large_scale_model/chat", middleware.JwtAdmin(), largeScaleModelApi.ModelAdminChatDeleteView)   //大模型管理员对话删除接口
	}
}
