package main

import "about-vaccine/src/base/dao"

func main() {
	r, err := InitApplication(dao.DSN)
	if err != nil {
		panic(err)
	}
	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
