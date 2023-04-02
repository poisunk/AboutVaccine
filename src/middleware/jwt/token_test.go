package jwt

import "testing"

func TestGenerate(t *testing.T) {
	token, err := GenerateToken(4, "sadas")
	if err != nil {
		panic(err)
	}
	println(token)
	c, err := ParseToken(token)
	if err != nil {
		panic(err)
	}
	println(c.Audience)
}
