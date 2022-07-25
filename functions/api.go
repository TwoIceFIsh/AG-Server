package functions

import (
	"github.com/gin-gonic/gin"
)

type ResultJson struct {
	Status string `json:"empname"`
	Body   string `json:"empbody"`
}

func Ping(c *gin.Context) {

	var err error
	var raw *user_raw

	_, err = connectDB("file:memory.db")

	if err != nil {
		panic("데이터베이스가 연결되지 않았습니다.")
	}

	c.JSON(200, gin.H{"responseData": raw})

	return

}
