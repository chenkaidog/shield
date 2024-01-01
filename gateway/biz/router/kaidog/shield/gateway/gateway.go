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
			{
				_account := _admin.Group("/account", _accountMw()...)
				_account.POST("/create", append(_createaccountMw(), gateway.CreateAccount)...)
				_account.GET("/query", append(_queryaccountMw(), gateway.QueryAccount)...)
			}
			{
				_account_status := _admin.Group("/account_status", _account_statusMw()...)
				_account_status.POST("/change", append(_switchaccountstatusMw(), gateway.SwitchAccountStatus)...)
			}
			{
				_login_record := _admin.Group("/login_record", _login_recordMw()...)
				_login_record.GET("/query", append(_queryloginrecordMw(), gateway.QueryLoginRecord)...)
			}
			{
				_password := _admin.Group("/password", _passwordMw()...)
				_password.POST("/reset", append(_resetpasswordMw(), gateway.ResetPassword)...)
			}
			{
				_user_info := _admin.Group("/user_info", _user_infoMw()...)
				_user_info.POST("/create", append(_createuserMw(), gateway.CreateUser)...)
				_user_info.GET("/query", append(_queryuserinfoMw(), gateway.QueryUserInfo)...)
				_user_info.POST("/update", append(_updateuserinfoMw(), gateway.UpdateUserInfo)...)
			}
		}
		{
			_user := _operator.Group("/user", _userMw()...)
			{
				_login_record0 := _user.Group("/login_record", _login_record0Mw()...)
				_login_record0.GET("/query", append(_queryselfloginrecordMw(), gateway.QuerySelfLoginRecord)...)
			}
			{
				_password0 := _user.Group("/password", _password0Mw()...)
				_password0.POST("/update", append(_updatepasswordMw(), gateway.UpdatePassword)...)
			}
			{
				_user_info0 := _user.Group("/user_info", _user_info0Mw()...)
				_user_info0.GET("/query", append(_queryselfuserinfoMw(), gateway.QuerySelfUserInfo)...)
			}
		}
	}
}
