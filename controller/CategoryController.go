package controller

import (
	"github.com/gin-gonic/gin"
	"go-gin-gorm/common"
	"go-gin-gorm/model"
	"go-gin-gorm/response"
)

// SearchCategory /* controller/CategoryController.go */
// SearchCategory 查询分类
func SearchCategory(c *gin.Context) {
	db := common.GetDB()
	var categories []model.Category
	if err := db.Find(&categories).Error; err != nil {
		response.Fail(c, nil, "查找失败")
		return
	}
	response.Success(c, gin.H{"categories": categories}, "查找成功")
}

// SearchCategoryName 查询分类名
func SearchCategoryName(c *gin.Context) {
	db := common.GetDB()
	var category model.Category
	// 获取path中的分类id
	categoryId := c.Params.ByName("id")
	if err := db.Where("id = ?", categoryId).First(&category).Error; err != nil {
		response.Fail(c, nil, "分类不存在")
		return
	}
	response.Success(c, gin.H{"categoryName": category.CategoryName}, "查找成功")
}
