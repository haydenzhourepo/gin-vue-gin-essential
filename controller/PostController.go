package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"oceanlearn.teach/ginessential/common"
	"oceanlearn.teach/ginessential/model"
	"oceanlearn.teach/ginessential/response"
)

type IPostController interface {
	RestController
}

type PostController struct {
	DB *gorm.DB
}

func (p PostController) Create(ctx *gin.Context) {
	var post model.Post
	ctx.Bind(&post)

	// 数据验证


	// 创建
	if err := p.DB.Create(&post).Error; err != nil {
		panic(err)
		return
	}

	response.Success(ctx, nil, "创建成功")
}

func (p PostController) Update(ctx *gin.Context) {
	// 获取body 参数
	var requestCategory model.Category
	ctx.Bind(&requestCategory)

	// 获取 path 参数
	//categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	if requestCategory.Name == "" {
		response.Fail(ctx, "请输入名称", nil)
		return
	}

	var updateCategory model.Category
	//c.DB.First(&updateCategory, categoryId)
	//c.DB.Model(&updateCategory).Update("name", requestCategory.Name)

	response.Success(ctx, gin.H{"category": updateCategory}, "更新成功")
}

func (p PostController) Show(ctx *gin.Context) {
	panic("implement me")
}

func (p PostController) Delete(ctx *gin.Context) {
	panic("implement me")
}

func NewPostController() IPostController {
	db := common.GetDB()
	db.AutoMigrate(model.Post{})
	return PostController{DB:db}
}
