package main

func main() {
	r, err := InitApplication()
	if err != nil {
		panic(err)
	}
	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
