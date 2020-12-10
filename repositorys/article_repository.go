package repositorys

import (
	"QiqiLike/datamodels/domain"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
	"sync"
)

type ArticleRepository interface {
	GetArticleManyRepo(map[string]string) ([]domain.TbArticle, int, error)
	CreateRepo(article *domain.TbArticle) bool
	SelectArticleDetailRepo(uuid string) (domain.TbArticle, error)
}

func NewArticleRepository(source *gorm.DB) ArticleRepository {
	return &articleRepository{source: source}
}

type articleRepository struct {
	source *gorm.DB
	mux    sync.RWMutex
}

func (a *articleRepository) SelectArticleDetailRepo(uuid string) (article domain.TbArticle, err error) {
	// 一组操作事务， 先将热度+1， 然后查询数据
	tx := a.source.Begin()
	if err = tx.Model(domain.TbArticle{}).Update("open_num", gorm.Expr("open_num + 1")).Where("uuid = ? ", uuid).Error; err != nil {
		tx.Rollback()
		return
	}
	if err = tx.Model(domain.TbArticle{}).First(&article).Error; err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

func (a *articleRepository) CreateRepo(article *domain.TbArticle) bool {
	if db := a.source.Create(&article); db.Error != nil {
		return false
	}
	return true
}

func (a *articleRepository) GetArticleManyRepo(params map[string]string) (articles []domain.TbArticle, count int, err error) {
	uuid := params["uuid"]
	db := a.source
	if !strings.EqualFold(uuid, "") {
		db = db.Where("club_uuid = ? ", uuid)
	}
	// 获取总条数
	db.Table("tb_article").Count(&count)
	page, _ := strconv.Atoi(params["page"])
	limit, _ := strconv.Atoi(params["limit"])
	if page != 0 && limit != 0 {
		db = db.Limit(limit).Offset((page - 1) * limit)
	}
	// TODO 添加可能的随机
	db.Order("open_num desc ").Find(&articles)
	return
}
