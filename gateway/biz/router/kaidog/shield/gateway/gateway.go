// Code generated by hertz generator. DO NOT EDIT.

package gateway

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	gateway "shield/gateway/biz/handler/kaidog/shield/gateway"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	root.POST("/login", append(_loginMw(), gateway.Login)...)
	root.POST("/logout", append(_logoutMw(), gateway.Logout)...)
	{
		_operator := root.Group("/operator", _operatorMw()...)
		{
			_admin := _operator.Group("/admin", _adminMw()...)
			_admin.POST("/create_account", append(_createaccountMw(), gateway.CreateAccount)...)
			_admin.POST("/create_user", append(_createuserMw(), gateway.CreateUser)...)
			_admin.GET("/query_account", append(_queryaccountMw(), gateway.QueryAccount)...)
			_admin.POST("/rest_password", append(_resetpasswordMw(), gateway.ResetPassword)...)
			_admin.POST("/switch_account_status", append(_switchaccountstatusMw(), gateway.SwitchAccountStatus)...)
			_admin.POST("/update_user", append(_updateuserinfoMw(), gateway.UpdateUserInfo)...)
		}
		{
			_user := _operator.Group("/user", _userMw()...)
			_user.GET("/query_login_record", append(_queryloginrecordMw(), gateway.QueryLoginRecord)...)
			_user.GET("/query_user_info", append(_queryuserinfoMw(), gateway.QueryUserInfo)...)
			_user.POST("/update_password", append(_updatepasswordMw(), gateway.UpdatePassword)...)
		}
	}
}
