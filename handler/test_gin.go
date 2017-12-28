package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//处理最基本的get请求
func TestGinGet(context *gin.Context){
	// 回复一个200OK,在client的http-get的resp的body中获取数据
	context.String(http.StatusOK, "test GET ok")

}


//处理post请求
func TestGinPost(context *gin.Context){
	// 回复一个200 OK, 在client的http-post的resp的body中获取数据
	context.String(http.StatusOK, "test2 OK")
}

func main(){
	// 注册一个默认的路由器
	router := gin.Default()
	router.GET("/testGet", TestGinGet)
	router.POST("/testPost", TestGinPost)
	router.Run("8888")
}