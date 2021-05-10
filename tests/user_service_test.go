package tests

import (
	"backend/dao"
	"backend/model"
	"backend/service"
	"backend/utils"
	"testing"

	"github.com/prashantv/gostub"
)

var userAuth_not_exist model.UserAuth = model.UserAuth{
	Uid: 0,
}

var user_right model.User = model.User{
	Uid:      1,
	Username: "mok",
	Email:    "mokkkkk@sjtu.edu.cn",
}

var user_not_exist model.User = model.User{
	Uid: 0,
}

var userAuth_right model.UserAuth = model.UserAuth{
	Uid:      1,
	Username: "mok",
	Password: "123456",
}

var loginParam_right utils.LoginParams = utils.LoginParams{
	Username: "mok",
	Password: "123456",
}

var loginParam_password_wrong utils.LoginParams = utils.LoginParams{
	Username: "mok",
	Password: "12345",
}

var registerParam utils.RegisterParams = utils.RegisterParams{
	Username: "mok",
	Password: "123456",
	Email:    "mokkkkk@sjtu.edu.cn",
}

var getUserParam utils.GetUserParams = utils.GetUserParams{
	Token: "token123",
}

var modifyUserParam_same_username utils.ModifyUserParams = utils.ModifyUserParams{
	Token:    "token123",
	Username: "mok",
	Email:    "mokkkkk@sjtu.edu.cn",
}

var modifyUserParam_diff_username utils.ModifyUserParams = utils.ModifyUserParams{
	Token:    "token123",
	Username: "mokk",
	Email:    "mokkkkk@sjtu.edu.cn",
}

var modifyUserAuthParam utils.ModifyUserAuthParams = utils.ModifyUserAuthParams{
	Token:    "token123",
	Password: "123456",
}

func TestLoginService(t *testing.T) {
	//9-10-11-12-21
	defer gostub.StubFunc(&dao.GetUserAuthByUsername, userAuth_not_exist).Reset()
	service.Login(loginParam_right)

	//9-10-11-13-14-21
	defer gostub.StubFunc(&dao.GetUserAuthByUsername, userAuth_right).Reset()
	service.Login(loginParam_password_wrong)

	//9-10-11-13-15-16-17-18-19-21
	defer gostub.StubFunc(&dao.GetUserAuthByUsername, userAuth_right).
		StubFunc(&service.NewToken, "token123").
		StubFunc(&dao.GetUserByUid, model.User{}).Reset()
	service.Login(loginParam_right)
}

func TestRegisterService(t *testing.T) {
	//24-25-26-27-28-29-30-31-32-33-34-35-36-40
	defer gostub.StubFunc(&dao.GetUserAuthByUsername, userAuth_not_exist).
		StubFunc(&dao.CreateUserAuth, uint(1)).
		StubFunc(&dao.CreateUser, uint(1)).
		StubFunc(&service.NewToken, "token123").Reset()
	service.Register(registerParam)

	//24-25-26-38-40
	defer gostub.StubFunc(&dao.GetUserAuthByUsername, userAuth_right).Reset()
	service.Register(registerParam)
}

func TestGetUserService(t *testing.T) {
	//43-44-45-46-51
	defer gostub.StubFunc(&service.CheckToken, uint(0)).Reset()
	service.GetUser(getUserParam)
	//43-44-45-48-51
	defer gostub.StubFunc(&service.CheckToken, uint(1)).
		StubFunc(&dao.GetUserByUid, user_right).Reset()
	service.GetUser(getUserParam)
}

func TestModifyUserService(t *testing.T) {
	//54-55-56-74-77
	defer gostub.StubFunc(&service.CheckToken, uint(0)).Reset()
	service.ModifyUser(modifyUserParam_same_username)

	//54-55-56-57-58-59-77
	defer gostub.StubFunc(&service.CheckToken, uint(1)).
		StubFunc(&dao.GetUserByUid, user_right).
		StubFunc(&dao.GetUserByUsername, user_right).Reset()
	service.ModifyUser(modifyUserParam_diff_username)

	//54-55-56-57-58-61-62-63-64-65-66-67-68-69-70-71-77
	defer gostub.StubFunc(&service.CheckToken, uint(1)).
		StubFunc(&dao.GetUserByUid, user_right).
		StubFunc(&dao.GetUserByUsername, user_not_exist).
		StubFunc(&dao.GetUserAuthByUid, userAuth_right).
		StubFunc(&dao.SetUserAuth).
		StubFunc(&dao.SetUser).Reset()
	service.ModifyUser(modifyUserParam_diff_username)

	//54-55-56-57-58-61-66-67-68-69-70-71-77
	defer gostub.StubFunc(&service.CheckToken, uint(1)).
		StubFunc(&dao.GetUserByUid, user_right).
		StubFunc(&dao.GetUserAuthByUid, userAuth_right).
		StubFunc(&dao.SetUser).Reset()
	service.ModifyUser(modifyUserParam_same_username)
}

func TestModifyUserAuthService(t *testing.T) {
	//80-81-82-88-91
	defer gostub.StubFunc(&service.CheckToken, uint(0)).Reset()
	service.ModifyUserAuth(modifyUserAuthParam)
	//80-81-82-83-84-85-86-91
	defer gostub.StubFunc(&service.CheckToken, uint(1)).
		StubFunc(&dao.GetUserAuthByUid, userAuth_right).
		StubFunc(&dao.SetUserAuth).Reset()
	service.ModifyUserAuth(modifyUserAuthParam)
}

func TestRemoveUserAndUserAuthService(t *testing.T) {
	//94-95-96-97
	defer gostub.StubFunc(&dao.GetUserAuthByUsername, userAuth_right).
		StubFunc(&dao.RemoveUser).
		StubFunc(&dao.RemoveUserAuth).Reset()
	service.RemoveUserAndUserAuth("mok")
}
