package router

import (
	"MyWeb/controller"
	"MyWeb/middleware/jwt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func SetupRouters() *gin.Engine {
	r := gin.Default()

	// 解决跨域问题
	r.Use(Cors())

	// 注册模板函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})

	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")
	r.GET("/", controller.IndexHandler)

	// 下载vaccine安卓安装包
	r.GET("/download/vaccine", controller.DownloadVaccineApk)

	apiGroup := r.Group("/api")
	{
		// VaccineCFDA数据，来自CFDA官网
		apiGroup.POST("/vaccine/cfda", controller.CreateAVaccineCFDA)
		apiGroup.GET("/vaccine/cfda", controller.GetVaccineCFDAList)
		apiGroup.DELETE("/vaccine/cfda", controller.DeleteVaccineCFDAByID)
		apiGroup.GET("/vaccine/cfda/example", controller.GetVaccineListExample)
		apiGroup.GET("/vaccine/cfda/example/:tid", controller.GetVaccineListExampleByTID)
		// 不良反应
		apiGroup.POST("/adverse", jwt.Auth, controller.CreateAdverseEvent)
		apiGroup.GET("/adverse", controller.GetAdverseEvent)
		apiGroup.DELETE("/adverse", jwt.Auth, controller.DeleteAdverseEvent)
		// OAE相关数据
		apiGroup.GET("/oae/label", controller.GetOaeTermsByLabel)
		apiGroup.GET("/oae/IRI", controller.GetOaeTermByIRI)
		apiGroup.GET("/oae/parent", controller.GetOaeParentTermsByIRI)
		// VaccineType 疫苗类型相关
		apiGroup.GET("/vaccine/type", controller.GetVaccineTypeList)
		// 用户操作
		apiGroup.POST("/user/login", controller.Login)
		apiGroup.POST("/user/register", controller.Register)
		apiGroup.GET("/user/status", controller.Status)
		apiGroup.DELETE("/user/logout", controller.Logout)
		apiGroup.GET("/user/list", controller.GetUserList)
		// 问卷系统
		// 问卷
		apiGroup.GET("/questionnaire", controller.GetQuestionnaireList)
		apiGroup.POST("/questionnaire", jwt.Auth, controller.CreateQuestionnaire)
		apiGroup.DELETE("/questionnaire/:id", jwt.Auth, controller.DeleteQuestionnaireByID)
		apiGroup.GET("/questionnaire/user/:uid", controller.GetQuestionnaireByUid)
		// 问题
		apiGroup.GET("/questionnaire/:id/questions", controller.GetQuestionListByQID)
		apiGroup.GET("/questionnaire/questions/type", controller.GetQuestionTypeList)
		apiGroup.POST("/questionnaire/:id/questions", jwt.Auth, controller.CreateQuestion)
		apiGroup.DELETE("/questionnaire/:id/questions/:qid", jwt.Auth, controller.DeleteQuestion)
		// 回答
		apiGroup.GET("/questionnaire/:id/response", jwt.Auth, controller.GetResponse)
		apiGroup.DELETE("/questionnaire/:id/response", jwt.Auth, controller.DeleteResponse)
		apiGroup.GET("/questionnaire/response/mine", jwt.Auth, controller.GetMineResponse)
		apiGroup.POST("/questionnaire/:id/response", jwt.Auth, controller.CreateResponse)
	}

	return r
}

// Cors 解决跨域问题
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
