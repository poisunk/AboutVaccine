package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vax/internal/controller"
)

type ApiRouter struct {
	UserController         *controller.UserController
	VaersController        *controller.VaersController
	VaccineController      *controller.VaccineController
	OAETermController      *controller.OAETermController
	AdverseEventController *controller.AdverseReportController
}

func NewApiRouter(
	userController *controller.UserController,
	vaersController *controller.VaersController,
	vaccineController *controller.VaccineController,
	oaeTermController *controller.OAETermController,
	adverseEventController *controller.AdverseReportController,
) *ApiRouter {
	return &ApiRouter{
		UserController:         userController,
		VaersController:        vaersController,
		VaccineController:      vaccineController,
		OAETermController:      oaeTermController,
		AdverseEventController: adverseEventController,
	}
}

func (a *ApiRouter) Run(address string) error {
	r := a.InitEngine()
	return r.Run(address)
}

func (a *ApiRouter) InitEngine() *gin.Engine {
	r := gin.Default()
	// 解决跨域问题
	r.Use(a.Cors())
	a.SetupRouters(r)
	return r
}

func (a *ApiRouter) SetupRouters(r *gin.Engine) {
	apiGroup := r.Group("/api")
	{
		// VaccineCFDA数据，来自CFDA官网
		apiGroup.GET("/vaccine/cfda", a.VaccineController.GetVaccine)
		apiGroup.GET("/vaccine/type", a.VaccineController.GetVaccineType)
		// 不良反应
		apiGroup.POST("/adverse", a.AdverseEventController.CreateAdverseEvent)
		apiGroup.GET("/adverse", a.AdverseEventController.GetAdverseEvent)
		apiGroup.DELETE("/adverse", a.AdverseEventController.DeleteAdverseEvent)
		apiGroup.GET("/adverse/result", a.AdverseEventController.GetAdverseResult)
		// OAE相关数据
		apiGroup.GET("/oae/label", a.OAETermController.GetOaeTermsByLabel)
		apiGroup.GET("/oae/IRI", a.OAETermController.GetOaeTermByIRI)
		apiGroup.GET("/oae/parent", a.OAETermController.GetOaeTermParents)
		apiGroup.GET("/oae/:id", a.OAETermController.GetOaeTermByID)
		// VaccineType 疫苗类型相关
		// 用户操作
		apiGroup.POST("/user/login", a.UserController.Login)
		apiGroup.POST("/user/register", a.UserController.Register)
		apiGroup.DELETE("/user/logout", a.UserController.Logout)
		apiGroup.GET("/user/search", a.UserController.SearchUser)
		// Vaers不良反应检索
		apiGroup.GET("/vaers", a.VaersController.SearchVaersResult)
		apiGroup.GET("/vaers/symptom", a.VaersController.GetVaersSymptomList)
		apiGroup.GET("/vaers/vaccine", a.VaersController.GetVaersVaccineList)
	}
}

// Cors 解决跨域问题
func (a *ApiRouter) Cors() gin.HandlerFunc {
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
