package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTrans "github.com/go-playground/validator/v10/translations/en"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"go-study/component/protobuf"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	InitTrans("zh")

	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery(), MyLogger(), MyValidator())

	//default 带logger和recovery
	//r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		//panic("1")
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	magicMirrorRouter := r.Group("/magicMirror")

	magicMirrorRouter.GET("/report/v1/:id", magicMirrorReportV1)
	magicMirrorRouter.POST("/report/v1/:id", magicMirrorReportV1)
	magicMirrorRouter.GET("/report/v2/:id", magicMirrorReportV2)
	magicMirrorRouter.POST("/report/v2/:id", magicMirrorReportV2)
	magicMirrorRouter.GET("/report/v3/:id", magicMirrorReportV3)

	go func() { r.Run(":9002") }()

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	fmt.Println("服务器关闭")
	fmt.Println("注销服务")
}

var trans ut.Translator

//验证器翻译
func InitTrans(locale string) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhNew := zh.New()
		enNew := en.New()
		uni := ut.New(enNew, zhNew, enNew)
		trans, ok = uni.GetTranslator(locale)

		if ok {
			switch locale {
			case "zh":
				zhTrans.RegisterDefaultTranslations(v, trans)
			case "en":
				fallthrough
			default:
				enTrans.RegisterDefaultTranslations(v, trans)
			}
		}
	}
}

//token验证器
func MyValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		for k, v := range c.Request.Header {
			if k == "X-Token" {
				token = v[0]
			}
		}

		if token != "test" {
			c.JSON(http.StatusUnauthorized, gin.H{"errMsg": "未登录"})
			c.Abort()
		} else {
			c.Next()
		}
	}
}

//日志中间件
func MyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		fmt.Println("耗时：", time.Since(t))
		fmt.Println("状态：", c.Writer.Status())
	}
}

func magicMirrorReportV1(c *gin.Context) {
	//Get请求 获取url query 参数
	a := c.DefaultQuery("a", "a")
	b := c.Query("b")
	//Post请求 获取表单参数
	cc := c.DefaultPostForm("c", "c")
	d := c.PostForm("d")
	c.JSON(http.StatusOK, gin.H{"id": c.Param("id"), "a": a, "b": b, "c": cc, "d": d})
}

type MagicMirrorReportV2Request struct {
	ID int32 `uri:"id" binding:"required"`
}

type MagicMirrorReportV2Response struct {
	Id      int32  `json:"id"`
	Version string `json:"version"`
}

func magicMirrorReportV2(c *gin.Context) {
	var request MagicMirrorReportV2Request
	//c.ShouldBind()  表单验证
	if err := c.ShouldBindUri(&request); err != nil {
		validationError, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationError.Translate(trans)})
		}
		return
	}

	response := &MagicMirrorReportV2Response{
		Id:      request.ID,
		Version: "v2",
	}

	c.JSON(http.StatusOK, response)
}

func magicMirrorReportV3(c *gin.Context) {

	response := &protobuf.MagicMirrorResponse{
		Id:      c.Param("id"),
		Version: "v3",
	}

	fmt.Println(response)

	c.ProtoBuf(http.StatusOK, response)
}
