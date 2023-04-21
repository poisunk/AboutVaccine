package router

import (
	"about-vaccine/src/controller"
	"about-vaccine/src/middleware/jwt"
	"about-vaccine/src/service"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"strconv"
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
	r.GET("/vaers", func(c *gin.Context) {
		c.HTML(http.StatusOK, "search.html", gin.H{})
	})
	r.GET("/vaers/result", func(c *gin.Context) {
		vaccineId, _ := c.GetQuery("vaccineId")
		symptomId, _ := c.GetQuery("symptomId")

		// 准备服务
		if vaccineId == "" && symptomId == "" {
			c.HTML(http.StatusBadRequest, "vaers.html", gin.H{})
			return
		}

		vaersService := service.InitVaersService()
		if len(vaccineId) == 0 {
			id, _ := strconv.ParseInt(symptomId, 10, 64)
			vaers, _, _ := vaersService.GetVaersResultsBySymptomId(id, 1, 20)
			c.HTML(http.StatusOK, "vaers.html", gin.H{
				"vaers": vaers,
			})
			return
		} else if len(symptomId) == 0 {
			id, _ := strconv.ParseInt(vaccineId, 10, 64)
			vaers, _, _ := vaersService.GetVaersResultsByVaccineId(id, 1, 20)
			c.HTML(http.StatusOK, "vaers.html", gin.H{
				"vaers": vaers,
			})
			return
		} else {
			id, _ := strconv.ParseInt(vaccineId, 10, 64)
			id2, _ := strconv.ParseInt(symptomId, 10, 64)
			vaers, _ := vaersService.GetVaersResults(id, id2)
			c.HTML(http.StatusOK, "vaers.html", gin.H{
				"vaers": []*service.VaersResult{vaers},
			})
		}
	})

	// 下载vaccine安卓安装包
	r.GET("/download/vaccine", controller.DownloadVaccineApk)

	apiGroup := r.Group("/api")
	{
		// VaccineCFDA数据，来自CFDA官网
		apiGroup.POST("/vaccine/cfda", controller.CreateAVaccineCFDA)
		apiGroup.GET("/vaccine/cfda", controller.GetVaccineCFDAList)
		apiGroup.DELETE("/vaccine/cfda", controller.DeleteVaccineCFDAById)
		apiGroup.GET("/vaccine/cfda/example", controller.GetVaccineListExample)
		apiGroup.GET("/vaccine/cfda/example/:tid", controller.GetVaccineListExampleByTId)
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
		apiGroup.DELETE("/questionnaire/:id", jwt.Auth, controller.DeleteQuestionnaireById)
		apiGroup.GET("/questionnaire/user/:uid", controller.GetQuestionnaireByUid)
		// 问题
		apiGroup.GET("/questionnaire/:id/questions", controller.GetQuestionListByQId)
		apiGroup.GET("/questionnaire/questions/type", controller.GetQuestionTypeList)
		apiGroup.POST("/questionnaire/:id/questions", jwt.Auth, controller.CreateQuestion)
		apiGroup.DELETE("/questionnaire/:id/questions/:qid", jwt.Auth, controller.DeleteQuestion)
		// 回答
		apiGroup.GET("/questionnaire/:id/response", jwt.Auth, controller.GetResponse)
		apiGroup.DELETE("/questionnaire/:id/response", jwt.Auth, controller.DeleteResponse)
		apiGroup.GET("/questionnaire/response/mine", jwt.Auth, controller.GetMineResponse)
		apiGroup.POST("/questionnaire/:id/response", jwt.Auth, controller.CreateResponse)
		// Vaers不良反应检索
		apiGroup.GET("/vaers", controller.SearchVaers)
		apiGroup.GET("/vaers/:vaersId", controller.GetVaers)
		apiGroup.GET("/vaers/symptom", controller.GetVaersSymptomList)
		apiGroup.GET("/vaers/vaccine", controller.GetVaersVaccineList)
	}

	return r
}

// Cors 解决跨域问题
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
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
