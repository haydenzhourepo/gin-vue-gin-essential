package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"oceanlearn.teach/ginessential/model"
	"oceanlearn.teach/ginessential/repository"
	"oceanlearn.teach/ginessential/response"
	"oceanlearn.teach/ginessential/vo"
	"strconv"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	Repository repository.ICategoryRepository
}

func NewCategoryController() ICategoryController {
	categoryController := CategoryController{Repository:repository.NewCategoryRepository()}
	categoryController.Repository.(repository.CategoryRepository).DB.AutoMigrate(model.Post{})
	return categoryController
}

func (c CategoryController) Create(ctx *gin.Context) {
	// 获取参数
	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		log.Println(err.Error())
		response.Fail(ctx, "数据验证错误", nil)
		return
	}
	category, err := c.Repository.Create(requestCategory.Name)
	if err != nil  {
		response.Fail(ctx, "数据验证错误", nil)
		return
	}

	response.Success(ctx, gin.H{"category": category}, "创建成功")
}

func (c CategoryController) Update(ctx *gin.Context) {
	// 获取body 参数
	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		log.Println(err.Error())
		response.Fail(ctx, "数据验证错误", nil)
		return
	}

	// 获取 path 参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	var updateCategory *model.Category
	category, err := c.Repository.SelectById(categoryId)
	if err != nil {
		panic(err)
	}
	updateCategory, err = c.Repository.Update(*category, requestCategory.Name)
	if err != nil {
		panic(err)
	}
	response.Success(ctx, gin.H{"category": updateCategory}, "更新成功")
}

func (c CategoryController) Show(ctx *gin.Context) {
	// 获取 path 参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	category, err := c.Repository.SelectById(categoryId)

	if err != nil {
		panic(err)
	}

	response.Success(ctx, gin.H{"category": category}, "")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	// 获取 path 参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	c.Repository.DeleteById(categoryId)

	response.Success(ctx, nil, "删除成功")
}



