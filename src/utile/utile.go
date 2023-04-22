package utile

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

// HandleSearchWord 处理搜索关键字
func HandleSearchWord(word string) (s string) {
	for _, x := range word {
		s += "%" + string(x)
	}
	s += "%"
	return s
}

// EnCoder 字符串加密
func EnCoder(password string) string {
	h := hmac.New(sha256.New, []byte(password))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

func StructConv(origin any, target any) error {
	s, err := json.Marshal(origin)
	if err != nil {
		return err
	}
	err = json.Unmarshal(s, target)
	if err != nil {
		return err
	}
	return nil
}
