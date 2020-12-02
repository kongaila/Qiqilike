package controllers

import (
	"QiqiLike/datamodels/domain"
	"QiqiLike/datamodels/vo"
	"github.com/kataras/iris/v12"
)

// 注册接口
func PostRegister(ctx iris.Context) {
	// 定义返回结果模型
	var result *vo.RespVO
	user := domain.TbUser{}
	_ = ctx.ReadJSON(&user)
	defer ctx.JSON(result)
	if ok := user.CheckUserNameAndPass(); !ok {
		result = vo.Req204RespVO(0, "账号不符合规则", nil)
		return
	}

	result = vo.Req200RespVO(1, "账号创建成功", nil)

}
