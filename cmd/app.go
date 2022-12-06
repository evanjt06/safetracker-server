package main

import (
	"evantu/safetracker-server/internal"
	ginw "github.com/aldelo/common/wrapper/gin"
	"github.com/aldelo/common/wrapper/gin/ginhttpmethod"
	"github.com/aldelo/connector/webserver"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

var server *webserver.WebServer

func main() {

	// db init
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}

	port, _ := strconv.Atoi(os.Getenv("PORT"))

	// db init
	err = internal.SetWriterDBInfo(os.Getenv("HOST"), port, os.Getenv("DBNAME"), os.Getenv("USERNAME"), os.Getenv("PASSWORD"))
	if err != nil {
		log.Fatal(err)
	}

	err = internal.ConnectToWriterDB()

	if err != nil {
		log.Fatal(err)
	}

	// ping test
	isReady := internal.IsWriterDBReady()

	// ping 200
	if !isReady {
		log.Println("MYSQL DB NOT READY")
	}

	// last func
	defer internal.DisconnectFromWriterDB()

	server = webserver.NewWebServer("SafeTracker", "config", "")

	server.LoginRequestDataPtr = &Credentials{}

	server.LoginResponseHandler = func(c *gin.Context, statusCode int, token string, expires time.Time) {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")

		c.JSON(statusCode, gin.H{
			"token": token,
			"exp":   expires,
		})
	}
	server.AuthenticateHandler = func(loginRequestDataPtr interface{}) (loggedInCredentialPtr interface{}) {
		if lg, ok := loginRequestDataPtr.(*Credentials); !ok {
			return nil
		} else {
			defer func() {
				lg.Password = ""
				lg.EmailAddress = ""
				lg.UserID = 0
				lg = nil
			}()

			// authenticate user
			uid, err := ValidateCredentials(*lg)
			if err != nil {
				log.Println(err.Error())
				return nil
			}

			return &Credentials{
				EmailAddress:    lg.EmailAddress,
				Password: lg.Password,
				UserID:   uid,
			}

		}
	}

	server.AddClaimsHandler = func(loggedInCredentialPtr interface{}) (identityKeyValue string, claims map[string]interface{}) {

		ptr, ok := loggedInCredentialPtr.(*Credentials)

		if !ok {
			return "", nil
		}

		if loggedInCredentialPtr != nil {
			return "app", map[string]interface{}{
				"uid": ptr.UserID,
			}
		}

		return "", nil
	}

	server.AuthorizerHandler = func(loggedInCredentialPtr interface{}, c *gin.Context) bool {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		return true
	}

	server.Routes = map[string]*ginw.RouteDefinition{
		"base": {
			Routes: []*ginw.Route{
				{
					RelativePath: "/register",
					Method: ginhttpmethod.POST,
					Handler: Register,
				},
			},
		},
		"auth": {
			Routes: []*ginw.Route{
				{
					RelativePath: "/discord",
					Method: ginhttpmethod.GET,
					Handler: Discord,
				},
				{
					RelativePath: "/twitter",
					Method: ginhttpmethod.GET,
					Handler: Twitter,
				},
				{
					RelativePath: "/stats",
					Method: ginhttpmethod.GET,
					Handler: Stats,
				},
				{
					RelativePath: "/user",
					Method: ginhttpmethod.GET,
					Handler: User,
				},
				{
					RelativePath: "/user",
					Method: ginhttpmethod.PUT,
					Handler: UpdateUser,
				},
				{
					RelativePath: "/discord",
					Method: ginhttpmethod.POST,
					Handler: CreateDiscordRecord,
					// retired, will not be used by API
				},
				{
					RelativePath: "/twitter",
					Method: ginhttpmethod.POST,
					Handler: CreateTwitterRecord,
				},
				{
					RelativePath: "/livefeed",
					Method: ginhttpmethod.POST,
					Handler: CreateLiveFeed,
				},
				{
					RelativePath: "/livefeed",
					Method: ginhttpmethod.GET,
					Handler: FetchLiveFeeds,
				},
				{
					RelativePath: "/livethreat",
					Method: ginhttpmethod.POST,
					Handler: CreateLiveFeedThreat,
				},
			},
		},
	}

	err = server.Serve()

	if err != nil {
		log.Fatal(err)
	}

}