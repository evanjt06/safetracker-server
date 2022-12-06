package test

import (
	"fmt"
	"github.com/aldelo/common/wrapper/aws/awsregion"
	"github.com/aldelo/common/wrapper/sns"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"
	"github.com/gorilla/websocket"
)

func TestA(tj *testing.T) {
	t := &sns.SNS{
		AwsRegion: awsregion.AWS_us_west_2_oregon,
		SMSTransactional: true,
	}

	if err := t.Connect(); err != nil {
		log.Fatal(err.Error())
		return
	}
	defer t.Disconnect()
	_, err := t.SendSMS("+1"+ "9257504638", "test", 3*time.Second)
	if err != nil{
		log.Fatal(err.Error())
		return
	}
}

func TestASD(tdd *testing.T) {
	t := time.Now()

	fin := ""
	x := t.Format("2006-01-02 15:04:05")
	arr := strings.Split(x, " ")

	l := strings.Split(arr[1], ":")
	l_int, _ := strconv.Atoi(l[0])

	if l_int < 12 {
		if l_int == 0 {
			fin = "12:" + l[1] + ":" + l[2] + " AM"
		} else {
			fin = arr[1] + " AM"
		}
	}
	if l_int > 12 {
		fin = strconv.Itoa(l_int - 12) + ":" + l[1] + ":" + l[2] + " PM"
	}
	if l_int == 12 {
		fin = arr[1] + " PM"
	}
	fin = arr[0] + " " + fin

	fmt.Println(fin)
}

func TestWebSocket(tj *testing.T) {


	var upgrader = websocket.Upgrader{
		//check origin will check the cross region source (note : please not using in production)
		CheckOrigin: func(r *http.Request) bool {
			//Here we just allow the chrome extension client accessable (you should check this verify accourding your client source)
			return true
		},
	}


		r := gin.Default()
		r.GET("/", func(c *gin.Context) {
			//upgrade get request to websocket protocol
			ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer ws.Close()


			//Response message to client
			err = ws.WriteMessage(1, []byte("REFRESH"))
			if err != nil {
				fmt.Println(err)
			}

		})
		r.Run() // listen and serve on 0.0.0.0:8080


}