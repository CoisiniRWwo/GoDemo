package admin

import (
	"fmt"
	"gin08/models"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (con UserController) Index(c *gin.Context) {

	//查询数据库
	//accountList := []models.Account{}
	//models.DB.Find(&accountList)
	//
	//c.JSON(http.StatusOK, gin.H{
	//	"result": accountList,
	//})

	//查询balance大于500的用户
	accountList := []models.Account{}
	models.DB.Where("balance=300").Find(&accountList)
	c.JSON(200, gin.H{
		"result": accountList,
	})

}

func (con UserController) Add(c *gin.Context) {

	account := models.Account{
		Username: "GORM",
		Balance:  5000,
	}

	models.DB.Create(&account)
	fmt.Println(account)
	c.String(200, "增加数据成功")
}

func (con UserController) Edit(c *gin.Context) {
	//保存所有字段

	//查询id等于6的数据
	//account := models.Account{Id: 8}
	//models.DB.Find(&account)
	////更新数据
	//account.Username = "gorm"
	//account.Balance = 10000
	//models.DB.Save(&account)

	//更新单个列
	account := models.Account{}
	models.DB.Model(&account).Where("id = ?", 8).Update("username", "GORM")

	//user := models.User{}
	//models.DB.Where("id = ?", 6).Find(&user)
	//user.Username = "哈哈"
	//user.Email = "aaa@qqq.com"
	//user.AddTime = int(models.GetUnix())
	//models.DB.Save(&user)
	//
	c.String(200, "修改用户成功")
}
func (con UserController) Delete(c *gin.Context) {

	account := models.Account{Id: 1}
	models.DB.Delete(&account)

	//删除数据
	//user := models.User{}
	//models.DB.Where("username = ?", "gorm").Delete(&user)
	//
	c.String(200, "删除用户")
}
