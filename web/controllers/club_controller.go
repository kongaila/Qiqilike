package controllers

import (
	"QiqiLike/datamodels/domain"
	"QiqiLike/datamodels/vo"
	"QiqiLike/service"
	"QiqiLike/tools"
	"github.com/kataras/iris/v12"
	"time"
)

type ClubController struct {
	AttrLoginService    service.LoginService
	AttrClubService     service.ClubService
	AttrUserClubService service.UserClubService
	AttrArticleService  service.ArticleService
	Ctx                 iris.Context
}

// 创建贴吧
func (c *ClubController) PostCreate() (result *vo.RespVO) {
	club := domain.TbClub{}
	if err := c.Ctx.ReadJSON(&club); err != nil {
		result = vo.Req204RespVO(0, "数据有误", nil)
		return
	}
	if club.MasterUuid, club.MasterName, _ = tools.ParseHeaderToken(c.Ctx); club.MasterUuid == "" || club.MasterName == "" {
		result = vo.Req200RespVO(1, "登录信息有误", nil)
		return
	}
	if ok := c.AttrClubService.Create(&club); !ok {
		result = vo.Req200RespVO(1, "添加失败", nil)
		return
	}
	userClub := domain.TbUserClub{
		Uuid:     tools.GenerateUUID(),
		ClubUuid: club.Uuid,
		UserUuid: club.MasterUuid,
		Identity: 1, // 吧主
		CreateAt: time.Now(),
	}
	c.AttrUserClubService.Create(&userClub)
	result = vo.Req200RespVO(1, "添加成功", club.Uuid)
	return
}

// 你这坏孩子 不要不说话 没有眼泪要擦 就别揉眼了
// 你这坏孩子 没人怪你啊 爱本是自由的 我该承受这 变化

// 获得贴吧列表
func (c *ClubController) GetMany() (result *vo.RespVO) {
	c.Ctx.URLParam("uuid")
	params := c.Ctx.URLParams()
	data, count, err := c.AttrClubService.GetClubMany(params)
	if err != nil {
		result = vo.Req500RespVO(0, "查询失败", nil)
		return
	}
	result = vo.Req200RespVO(count, "查询成功", data)
	return
}

// 获得一个贴吧详情
func (c *ClubController) GetBy() (result *vo.RespVO) {
	uuid := c.Ctx.URLParam("uuid")
	var club domain.TbClub
	var err error
	club, err = c.AttrClubService.GetClubDetail(uuid)
	if err != nil {
		result = vo.Req500RespVO(0, "查询失败", nil)
		return
	}
	result = vo.Req200RespVO(1, "查询成功", club)
	return
}

// 获得一个贴吧的全部帖子(分页)
func (c *ClubController) GetArticleMany() (result *vo.RespVO) {
	params := c.Ctx.URLParams()
	articles, count, err := c.AttrArticleService.GetArticleMany(params)
	if err != nil {
		result = vo.Req500RespVO(0, "查询失败", nil)
		return
	}
	result = vo.Req200RespVO(count, "查询成功", articles)
	return
}
