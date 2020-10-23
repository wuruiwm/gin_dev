package api

import (
	"gin_dev/app/model"
	"gin_dev/app/response"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func ArticleList(c *gin.Context){
	//获取参数
	keyword := c.Query("keyword")
	categoryId := com.StrTo(c.Query("category_id")).MustInt()
	page := com.StrTo(c.DefaultQuery("page","1")).MustInt()
	limit := com.StrTo(c.DefaultQuery("limit","15")).MustInt()

	//where条件
	where := make(map[string]interface{})
	if categoryId != 0 {
		where["category_id"] = categoryId
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

func ArticleCreate(c *gin.Context){
	defer func() {
		if err := recover();err != nil{
			if msg,ok := err.(string);ok{
				response.Error(c,msg)
			}else{
				response.Error(c,"创建失败")
			}
		}
	}()
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

func ArticleUpdate(c *gin.Context){
	defer func() {
		if err := recover();err != nil{
			if msg,ok := err.(string);ok{
				response.Error(c,msg)
			}else{
				response.Error(c,"创建失败")
			}
		}
	}()
	//获取参数
	id := com.StrTo(c.PostForm("id")).MustInt()
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

func ArticleDelete(c *gin.Context){
	//获取参数
	id := com.StrTo(c.PostForm("id")).MustInt()

	//参数校验
	if id == 0{
		response.Error(c,"id不能为空")
		return
	}

	err := model.ArticleDelete(id)
	if err != nil {
		response.Error(c,"删除失败")
		return
	}
	response.Success(c,"删除成功",nil)
}