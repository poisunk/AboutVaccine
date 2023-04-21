package main

import (
	"about-vaccine/dao"
	"about-vaccine/router"
)

func main() {
	// 链接数据库
	dao.InitMysql()
	// 关闭数据库
	defer dao.DB.Close()
	// 设置路由
	var r = router.SetupRouters()
	err := r.Run(":8080")
	if err != nil {
		println(err)
		return
	}
}
