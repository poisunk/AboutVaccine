package router

import (
	"about-vaccine/internal/controller"
	"about-vaccine/internal/middleware/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type APIRouter struct {
	UserController         *controller.UserController
	VaersController        *controller.VaersController
	VaccineController      *controller.VaccineController
	OAETermController      *controller.OAETermController
	AdverseEventController *controller.AdverseEventController
}

func NewAPIRouter(
	userController *controller.UserController,
	vaersController *controller.VaersController,
	vaccineController *controller.VaccineController,
	oaeTermController *controller.OAETermController,
	adverseEventController *controller.AdverseEventController,
) *APIRouter {
	return &APIRouter{
		UserController:         userController,
		VaersController:        vaersController,
		VaccineController:      vaccineController,
		OAETermController:      oaeTermController,
		AdverseEventController: adverseEventController,
	}
}

func (a *APIRouter) Run(address string) error {
	r := a.InitEngine()
	return r.Run(address)
}

func (a *APIRouter) InitEngine() *gin.Engine {
	r := gin.Default()
	// 解决跨域问题
	r.Use(a.Cors())
	a.SetupRouters(r)
	return r
}

func (a *APIRouter) SetupRouters(r *gin.Engine) {
	apiGroup := r.Group("/api")
	{
		// VaccineCFDA数据，来自CFDA官网
		apiGroup.GET("/vaccine/cfda", a.VaccineController.GetVaccineList)
		apiGroup.GET("/vaccine/type", a.VaccineController.GetVaccineTypeList)
		// 不良反应
		apiGroup.POST("/adverse", jwt.AuthWithoutLogin, a.AdverseEventController.CreateAdverseEvent)
		apiGroup.GET("/adverse", a.AdverseEventController.GetAdverseEvent)
		apiGroup.DELETE("/adverse", jwt.Auth, a.AdverseEventController.DeleteAdverseEvent)
		// OAE相关数据
		apiGroup.GET("/oae/label", a.OAETermController.GetOaeTermsByLabel)
		apiGroup.GET("/oae/IRI", a.OAETermController.GetOaeTermByIRI)
		apiGroup.GET("/oae/parent", a.OAETermController.GetOaeTermParents)
		// VaccineType 疫苗类型相关
		// 用户操作
		apiGroup.POST("/user/login", a.UserController.Login)
		apiGroup.POST("/user/register", a.UserController.Register)
		apiGroup.DELETE("/user/logout", a.UserController.Logout)
		apiGroup.GET("/user/search", a.UserController.SearchUser)
		//// 问卷系统
		//// 问卷
		//apiGroup.GET("/questionnaire", controller.GetQuestionnaireList)
		//apiGroup.POST("/questionnaire", jwt.Auth, controller.CreateQuestionnaire)
		//apiGroup.DELETE("/questionnaire/:id", jwt.Auth, controller.DeleteQuestionnaireById)
		//apiGroup.GET("/questionnaire/user/:uid", controller.GetQuestionnaireByUid)
		//// 问题
		//apiGroup.GET("/questionnaire/:id/questions", controller.GetQuestionListByQId)
		//apiGroup.GET("/questionnaire/questions/type", controller.GetQuestionTypeList)
		//apiGroup.POST("/questionnaire/:id/questions", jwt.Auth, controller.CreateQuestion)
		//apiGroup.DELETE("/questionnaire/:id/questions/:qid", jwt.Auth, controller.DeleteQuestion)
		//// 回答
		//apiGroup.GET("/questionnaire/:id/response", jwt.Auth, controller.GetResponse)
		//apiGroup.DELETE("/questionnaire/:id/response", jwt.Auth, controller.DeleteResponse)
		//apiGroup.GET("/questionnaire/response/mine", jwt.Auth, controller.GetMineResponse)
		//apiGroup.POST("/questionnaire/:id/response", jwt.Auth, controller.CreateResponse)
		// Vaers不良反应检索
		apiGroup.GET("/vaers", a.VaersController.SearchVaersResult)
		apiGroup.GET("/vaers/symptom", a.VaersController.GetVaersSymptomList)
		apiGroup.GET("/vaers/vaccine", a.VaersController.GetVaersVaccineList)
	}
}

// Cors 解决跨域问题
func (a *APIRouter) Cors() gin.HandlerFunc {
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
