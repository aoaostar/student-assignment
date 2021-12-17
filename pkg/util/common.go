package util

import (
	"github.com/kataras/iris/v12"
	"golang.org/x/crypto/bcrypt"
	"os"
)

func FileExist(filepath string) bool {
	_, err := os.Stat(filepath)
	if err != nil {
		return !os.IsNotExist(err)
	}
	return true

}

// HashAndSalt 加密密码
func HashAndSalt(pwd string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	return string(hash)
}

// ValidatePasswords 验证密码
func ValidatePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))
	if err != nil {
		return false
	}
	return true
}


func Msg(status string, message string, data ...interface{}) iris.Map {

	var data2 interface{}
	if len(data) == 0 {
		data2 = iris.Map{}
	} else {
		if data[0] == nil {
			data2 = iris.Map{}
		} else {
			data2 = data[0]
		}
	}
	return iris.Map{
		"status":  status,
		"message": message,
		"data":    data2,
	}

}
func JSON(c iris.Context, status string, message string, data ...interface{}) {

	var data2 interface{}
	if len(data) == 0 {
		data2 = iris.Map{}
	} else {
		if data[0] == nil {
			data2 = iris.Map{}
		} else {
			data2 = data[0]
		}
	}
	c.JSON(Msg(status, message, data2))
}