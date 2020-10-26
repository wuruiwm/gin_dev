package admin

import (
	"fmt"
	"gin_dev/app/model"
	"gin_dev/app/response"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

//文章列表
func ArticleList(c *gin.Context){
	//获取参数
	keyword := c.Query("keyword")
	categoryId := com.StrTo(c.Query("category_id")).MustInt()
	page := com.StrTo(c.DefaultQuery("page","1")).MustInt()
	limit := com.StrTo(c.DefaultQuery("limit","15")).MustInt()

	//where条件
	where := make(map[string]interface{})
	if categoryId != 0 {
		where["a.category_id"] = categoryId
	}

	//查询数据
	article,err := model.ArticleList(page,limit,where,keyword)
	if err != nil{
		response.Error(c,"获取失败")
		return
	}
	response.Success(c,"获取成功",article)
	return
}

//文章创建
func ArticleCreate(c *gin.Context){
	defer response.RecoverError(c,"创建失败")
	//获取参数
	title := c.PostForm("title")
	content := c.PostForm("content")

	//参数校验
	if title == ""{
		panic("请输入标题")
	}
	if content == ""{
		panic("请输入内容")
	}

	err := model.ArticleCreate(title,content)
	if err != nil{
		panic("创建失败")
	}
	response.Success(c,"创建成功",nil)
}

//文章修改
func ArticleUpdate(c *gin.Context){
	defer response.RecoverError(c,"修改失败")
	//获取参数
	id := com.StrTo(c.Param("id")).MustInt()
	title := c.PostForm("title")
	content := c.PostForm("content")

	//参数校验
	if id == 0{
		panic("id不能为空")
	}
	if title == ""{
		panic("请输入标题")
	}
	if content == ""{
		panic("请输入内容")
	}

	err := model.ArticleUpdate(id,title,content)
	if err != nil{
		panic("修改失败")
	}
	response.Success(c,"修改成功",nil)
}

//文章删除
func ArticleDelete(c *gin.Context){
	defer response.RecoverError(c,"删除失败")
	//获取参数
	id := com.StrTo(c.Param("id")).MustInt()

	//参数校验
	if id == 0{
		panic("id不能为空")
	}

	err := model.ArticleDelete(id)
	if err != nil {
		panic("删除失败")
	}
	response.Success(c,"删除成功",nil)
}

//文章详情
func ArticleDetail(c *gin.Context){
	defer response.RecoverError(c,"数据不存在,请重试")
	//获取参数
	id := com.StrTo(c.Param("id")).MustInt()
	fmt.Println(id)

	//参数校验
	if id == 0{
		panic("id不能为空")
	}

	article,err := model.ArticleDetail(id)
	if err != nil {
		panic("数据不存在")
	}
	response.Success(c,"获取成功",article)
}

//文章分类 用于文章列表搜索
func ArticleCategory(c *gin.Context){
	defer response.RecoverError(c,"获取失败")

	articleCategory,err := model.ArticleCategoryAll()
	if err != nil {
		panic("获取失败")
	}

	response.Success(c,"获取成功",articleCategory)
}