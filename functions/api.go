package functions

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

type ResultJson struct {
	Hash     string `json:"client_hash"`
	HostName string `json:"host_name"`
	HostIP   string `json:"host_ip"`
}

// Client hash 정보를 요청한다.
func Check(c *gin.Context) {

	// 1. body의 json 정보를 획득한다.
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		// Handle error
	}
	log.Println(jsonData)

	// 2. DB Search (OK hash)

	// 2-1. 신규정보 등록 (OK 1)
	// Client 리스트에 등록한다.

	// 2-2. DB Search (OK hash)
	// Client hash를 반환해준다 해당 값을 기준으로 서치 가능하도록 한다.

	c.JSON(200, ResultJson{"1", "2", "3"})

	return

}
