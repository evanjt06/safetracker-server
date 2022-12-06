package main

import (
	"evantu/safetracker-server/internal"
	"fmt"
	common "github.com/aldelo/common"
	"github.com/aldelo/common/wrapper/aws/awsregion"
	"github.com/aldelo/common/wrapper/s3"
	"github.com/aldelo/common/wrapper/sns"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

//
// ROUTES
//

func Register(c *gin.Context, bindingInputPtr interface{}) {

	FirstName := c.PostForm("FirstName")
	LastName := c.PostForm("LastName")
	PhoneNumber := c.PostForm("PhoneNumber")
	EmailAddress := c.PostForm("Email")
	PointOfContact := c.PostForm("PointOfContact")
	PointOfContactPhoneNumber := c.PostForm("PointOfContactPhoneNumber")
	Password := c.PostForm("Password")
	Location := c.PostForm("Location")

	// null inputs == may be empty
	DiscordID := c.PostForm("DiscordID")
	TwitterID := c.PostForm("TwitterID")

	if FirstName == "" || LastName == "" || PhoneNumber == "" || EmailAddress == "" || PointOfContact == "" || PointOfContactPhoneNumber == "" || Password == "" || Location == "" {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, "invalid inputs")
		return
	}

	if !IsEmailValid(EmailAddress) {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, "invalid input (EMAIL)")
		return
	}

	// check if email address is already in database
	ux := internal.User{}
	ux.UseDBWriterPreferred()
	notFound, err := ux.GetByEmailAddress(EmailAddress)
	if !notFound {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, "EmailAddress already found")
		return
	}
	if err != nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, "47" + err.Error())
		return
	}

	if Password == "" || len(Password) < 8 {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, "invalid input (PASSWORD)")
		return
	}

	if !IsPhoneNumberValid(PhoneNumber) || !IsPhoneNumberValid(PointOfContactPhoneNumber) {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, "invalid input (PHONE NUMBER)")
		return
	}

	u := internal.User{
		FirstName:                 FirstName,
		LastName:                  LastName,
		PhoneNumber:               PhoneNumber,
		EmailAddress:              EmailAddress,
		PointOfContact:            PointOfContact,
		PointOfContactPhoneNumber: PointOfContactPhoneNumber,
		Password:                  Password,
		Location:                  Location,
	}
	u.UseDBWriterPreferred()
	err = u.Set()
	if err != nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, "74" + err.Error())
		return
	}
	dbot := internal.DiscordBot{
		DiscordID: common.ToNullString(DiscordID, true),
		UserID:    u.ID,
	}
	dbot.UseDBWriterPreferred()
	err = dbot.Set()
	if err != nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, "84" + err.Error())
		return
	}
	tbot := internal.TwitterBot{
		TwitterID: common.ToNullString(TwitterID, true),
		UserID:    u.ID,
	}
	err = tbot.Set()
	if err != nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, "93" + err.Error())
		return
	}

	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, "")
}

func Discord(c *gin.Context, bindingInputPtr interface{}) {

	claims := server.ExtractJwtClaims(c)
	uid := int64((claims["uid"]).(float64))

	Disd := internal.DiscordBot{}
	Disd.UseDBWriterPreferred()
	notFound, err := Disd.GetByUserID(uid)
	if notFound {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, "DiscordBot not found")
		return
	}
	if err != nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}

	DiscordBotList := internal.DiscordBotChildList{}
	DiscordBotList.UseDBWriterPreferred()
	err = DiscordBotList.GetByDiscordBotID(Disd.ID)
	if err != nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}

	if DiscordBotList.List == nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(200, gin.H{
			"DiscordData": []string{},
		})
		return
	}

	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, gin.H{
		"DiscordData": DiscordBotList.List,
	})
}

func Twitter(c *gin.Context, bindingInputPtr interface{}) {

	claims := server.ExtractJwtClaims(c)
	uid := int64((claims["uid"]).(float64))

	tw := internal.TwitterBot{}
	tw.UseDBWriterPreferred()
	notFound, err := tw.GetByUserID(uid)
	if notFound {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, "TwitterBot not found")
		return
	}
	if err != nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}

	TwitterBotList := internal.TwitterBotChildList{}
	TwitterBotList.UseDBWriterPreferred()
	err = TwitterBotList.GetByTwitterBotID(tw.ID)
	if err != nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}

	if TwitterBotList.List == nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(200, gin.H{
			"TwitterData": []string{},
		})
		return
	}

	err = TwitterBotList.GetCustom("WHERE TwitterBotID = " + strconv.Itoa(int(tw.ID)),"ORDER BY Datetime DESC", 0,0)

	if err != nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}

	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, gin.H{
		"TwitterData": TwitterBotList.List,
	})

}

func Stats(c *gin.Context, bindingInputPtr interface{}) {

	claims := server.ExtractJwtClaims(c)
	uid := int64((claims["uid"]).(float64))

	liveFeedThreatCount := 0
	discordThreatCount := 0
	twitterThreatCount := 0

	LiveFeedObj := internal.LiveFeedList{}
	LiveFeedObj.UseDBWriterPreferred()
	err := LiveFeedObj.GetByUserID(uid)
	if err != nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}

	if LiveFeedObj.List != nil {
		for _,v := range *LiveFeedObj.List {

			x := internal.LiveFeedChildList{}
			x.UseDBWriterPreferred()
			err = x.GetByLiveFeedID(v.ID)
			if err != nil {
				c.Header("Access-Control-Allow-Credentials", "true")
				c.Header("Access-Control-Allow-Origin", "*")
				c.JSON(500, err.Error())
				return
			}
			liveFeedThreatCount += x.Count

		}
	}

	tw := internal.TwitterBot{}
	tw.UseDBWriterPreferred()
	notFound, err := tw.GetByUserID(uid)
	if notFound {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, "TwitterBot not found")
		return
	}
	if err != nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}

	TwitterBotList := internal.TwitterBotChildList{}
	TwitterBotList.UseDBWriterPreferred()
	err = TwitterBotList.GetByTwitterBotID(tw.ID)
	if err != nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}
	twitterThreatCount = TwitterBotList.Count

	Disd := internal.DiscordBot{}
	Disd.UseDBWriterPreferred()
	notFound, err = Disd.GetByUserID(uid)
	if notFound {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, "DiscordBot not found")
		return
	}
	if err != nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}

	DiscordBotList := internal.DiscordBotChildList{}
	DiscordBotList.UseDBWriterPreferred()
	err = DiscordBotList.GetByDiscordBotID(Disd.ID)
	if err != nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}

	discordThreatCount = DiscordBotList.Count

	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, gin.H{
		"totalThreats": liveFeedThreatCount + discordThreatCount + twitterThreatCount,
		"liveFeedThreatCount": liveFeedThreatCount,
		"discordThreatCount": discordThreatCount,
		"twitterThreatCount": twitterThreatCount,
	})
}

func User(c *gin.Context, bindingInputPtr interface{}) {
	claims := server.ExtractJwtClaims(c)
	uid := int64((claims["uid"]).(float64))

	u := internal.User{}
	u.UseDBWriterPreferred()
	notFound, err := u.GetByID(uid)
	if notFound {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, "User not found")
		return
	}
	if err != nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}

	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, gin.H{
		"FirstName": u.FirstName,
		"LastName": u.LastName,
		"Location": u.Location,
		"EmailAddress": u.EmailAddress,
		"PhoneNumber": u.PhoneNumber,
		"PointOfContact": u.PointOfContact,
		"PointOfContactPhoneNumber": u.PointOfContactPhoneNumber,
	})
}

func UpdateUser(c *gin.Context, bindingInputPtr interface{}) {

	claims := server.ExtractJwtClaims(c)
	uid := int64((claims["uid"]).(float64))

	FirstName := c.PostForm("FirstName")
	LastName := c.PostForm("LastName")
	PhoneNumber := c.PostForm("PhoneNumber")
	EmailAddress := c.PostForm("Email")
	PointOfContact := c.PostForm("PointOfContact")
	PointOfContactPhoneNumber := c.PostForm("PointOfContactPhoneNumber")
	Location := c.PostForm("Location")

	if FirstName == "" || LastName == "" || PhoneNumber == "" || EmailAddress == "" || PointOfContact == "" || PointOfContactPhoneNumber == "" || Location == "" {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, "invalid inputs")
		return
	}

	ux := internal.User{}
	ux.UseDBWriterPreferred()
	notFound, err := ux.GetByID(uid)
	if notFound {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, "User not found")
		return
	}
	if err != nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}
	ux2 := internal.User{}
	ux2.UseDBWriterPreferred()
	notFound, err = ux2.GetByEmailAddress(EmailAddress)
	if !notFound && ux.EmailAddress != EmailAddress {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, "Email already taken")
	}
	if err != nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}

	u := internal.User{}
	u.UseDBWriterPreferred()
	notFound, err = u.GetByID(uid)
	if notFound {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, "User not found")
		return
	}
	if err != nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}

	u.FirstName = FirstName
	u.LastName = LastName
	u.PhoneNumber = PhoneNumber
	u.EmailAddress = EmailAddress
	u.PointOfContact = PointOfContact
	u.PointOfContactPhoneNumber = PointOfContactPhoneNumber
	u.Location = Location

	err = u.Set()
	if err != nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}

	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, "")
}

func CreateDiscordRecord(c *gin.Context, bindingInputPtr interface{}) {

	claims := server.ExtractJwtClaims(c)
	uid := int64((claims["uid"]).(float64))
	uid_string := strconv.Itoa(int(uid))

	TextContent := c.PostForm("TextContent")
	AuthorDiscordID := c.PostForm("AuthorDiscordID")
	AuthorDiscordTag := c.PostForm("AuthorDiscordTag")
	ImageFile, err := c.FormFile("ImageFile")
	if err != nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}

	if TextContent == "" || AuthorDiscordID == "" || AuthorDiscordTag == "" {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, "invalid inputs")
		return
	}

	// get the discordbot id by using the UID
	dbot := internal.DiscordBot{}
	dbot.UseDBWriterPreferred()
	notFound, err := dbot.GetByUserID(uid)
	if notFound {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, "Discord record not found")
		return
	}
	if err != nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}
	DiscordBotID := dbot.ID


	// upload file to S3, given the user id -> discord bot id as subfolders
	s := s3.S3{
		AwsRegion:   awsregion.AWS_us_west_2_oregon,
		HttpOptions: nil,
		BucketName:  "calc.masa.space",
	}
	err = s.Connect()
	if err != nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}
	defer s.Disconnect()

	// if the image extension is not png or jpg or jpeg, error
	validExtensions := []string{".png", ".jpg", ".jpeg"}
	if !Contains(validExtensions, filepath.Ext(ImageFile.Filename)) {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, "image file type not supported")
		return
	}

	newFileName := uid_string + "/discord/ " + ImageFile.Filename

	fileContent, err := ImageFile.Open()
	if err != nil {
		fmt.Println(err.Error())
	}
	byteContainer, err := ioutil.ReadAll(fileContent)

	location, err := s.Upload(nil, byteContainer, newFileName)

	if err != nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}

	// create discord bot child record
	dbotChild := internal.DiscordBotChild{
		DiscordBotID:     DiscordBotID,
		ImageURL:         location,
		TextContent:      TextContent,
		Datetime:         time.Now(),
		AuthorDiscordID:  AuthorDiscordID,
		AuthorDiscordTag: AuthorDiscordTag,
	}
	dbotChild.UseDBWriterPreferred()
	err = dbotChild.Set()
	if err != nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}

	// send SMS text message to user
	t := &sns.SNS{
		AwsRegion: awsregion.AWS_us_west_2_oregon,
		SMSTransactional: true,
	}

	if err := t.Connect(); err != nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}
	defer t.Disconnect()

	ux := internal.User{}
	ux.UseDBWriterPreferred()
	notFound, err = ux.GetByID(uid)
	if notFound {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, "User record not found")
		return
	}
	if err != nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}

	message := "[SAFETRACKER] Critical threat detected by Discord Analyzer"

	_, err = t.SendSMS("+1"+ ux.PhoneNumber, message, 3*time.Second)
	if err != nil{
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}

	_, err = t.SendSMS("+1"+ ux.PointOfContactPhoneNumber, message, 3*time.Second)
	if err != nil{
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}

	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, "")
}

func CreateTwitterRecord(c *gin.Context, bindingInputPtr interface{}) {

	claims := server.ExtractJwtClaims(c)
	uid := int64((claims["uid"]).(float64))

	//TextContent := c.PostForm("TextContent")
	//AuthorTID := c.PostForm("AuthorTwitterID")
	//AuthorTTag := c.PostForm("AuthorTwitterTag")
	//ImageFile := c.PostForm("ImageFile")
	//
	//fmt.Println(TextContent, AuthorTID, AuthorTTag, ImageFile)

	var tj TwitterJSONData
	err := c.BindJSON(&tj)
	if err != nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}

	if tj.TextContent == "" || tj.AuthorTwitterTag == "" || tj.AuthorTwitterID == "" || tj.ImageFile == "" {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, "invalid inputs")
		return
	}

	// get the discordbot id by using the UID
	tbot := internal.TwitterBot{}
	tbot.UseDBWriterPreferred()
	notFound, err := tbot.GetByUserID(uid)
	if notFound {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, "Twitter record not found")
		return
	}
	if err != nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}
	TwitterBotID := tbot.ID

	//// upload file to S3, given the user id -> discord bot id as subfolders
	//s := s3.S3{
	//	AwsRegion:   awsregion.AWS_us_west_2_oregon,
	//	HttpOptions: nil,
	//	BucketName:  "calc.masa.space",
	//}
	//err = s.Connect()
	//if err != nil {
	//	c.JSON(500, err.Error())
	//	return
	//}
	//defer s.Disconnect()
	//
	//// if the image extension is not png or jpg or jpeg, error
	//validExtensions := []string{"png", "jpg", "jpeg"}
	//if !Contains(validExtensions, filepath.Ext(ImageFile.Filename)) {
	//	c.JSON(500, "image file type not supported")
	//	return
	//}
	//
	//newFileName := uid_string + "/twitter/ " + ImageFile.Filename
	//
	//fileContent, err := ImageFile.Open()
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//byteContainer, err := ioutil.ReadAll(fileContent)
	//
	//location, err := s.Upload(nil, byteContainer, newFileName)
	//
	//if err != nil {
	//	c.JSON(500, err.Error())
	//	return
	//}

	// create discord bot child record
	tbotChild := internal.TwitterBotChild{
		TwitterBotID:     TwitterBotID,
		ImageURL:         tj.ImageFile,
		TextContent:      tj.TextContent,
		Datetime:         time.Now(),
		AuthorTwitterID:  tj.AuthorTwitterID,
		AuthorTwitterTag: tj.AuthorTwitterTag,
	}
	tbotChild.UseDBWriterPreferred()
	err = tbotChild.Set()
	if err != nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}

	// send SMS text message to user
	t := &sns.SNS{
		AwsRegion: awsregion.AWS_us_west_2_oregon,
		SMSTransactional: true,
	}

	if err := t.Connect(); err != nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}
	defer t.Disconnect()

	ux := internal.User{}
	ux.UseDBWriterPreferred()
	notFound, err = ux.GetByID(uid)
	if notFound {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, "User record not found")
		return
	}
	if err != nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}

	fin := ""
	x := tbotChild.Datetime.Format("2006-01-02 15:04:05")
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


	message := "[SafeTracker] Twitter Feed Analyzer detected critical threat by @" + tbotChild.AuthorTwitterID + " on " + fin

	_, err = t.SendSMS("+1"+ ux.PhoneNumber, message, 3*time.Second)
	if err != nil{
		fmt.Println(err.Error())
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}

	_, err = t.SendSMS("+1"+ ux.PointOfContactPhoneNumber, message, 3*time.Second)
	if err != nil{
		fmt.Println(err.Error())
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}

	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, "")

}

func CreateLiveFeed(c *gin.Context, bindingInputPtr interface{}) {

	// create live feed record
	claims := server.ExtractJwtClaims(c)
	uid := int64((claims["uid"]).(float64))
	YouTubeLiveLink := c.PostForm("YouTubeLiveLink")
	Location := c.PostForm("Location")

	if YouTubeLiveLink == "" {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, "invalid inputs")
		return
	}

	ux := internal.User{}
	ux.UseDBWriterPreferred()
	notFound, err := ux.GetByID(uid)
	if notFound {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, "User record not found")
		return
	}
	if err != nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}

	if len(Location) == 0 {
		// get the users location
		lf := internal.LiveFeed{
			UserID:          uid,
			YouTubeLiveLink: YouTubeLiveLink,
			Location:  ux.Location,
		}
		lf.UseDBWriterPreferred()
		err = lf.Set()
		if err != nil{
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Origin", "*")
			c.JSON(500, err.Error())
			return
		}

		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(200, gin.H{
			"LiveFeedID": lf.ID,
		})
	} else {
		lf := internal.LiveFeed{
			UserID:          uid,
			YouTubeLiveLink: YouTubeLiveLink,
			Location: Location,
		}
		lf.UseDBWriterPreferred()
		err = lf.Set()
		if err != nil{
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Origin", "*")
			c.JSON(500, err.Error())
			return
		}

		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(200, gin.H{
			"LiveFeedID": lf.ID,
		})
	}

}

func FetchLiveFeeds(c *gin.Context, bindingInputPtr interface{}) {

	claims := server.ExtractJwtClaims(c)
	uid := int64((claims["uid"]).(float64))

	lf := internal.LiveFeedList{}
	lf.UseDBWriterPreferred()
	err := lf.GetByUserID(uid)
	if err != nil{
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}

	if lf.List == nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(200, "")
		return
	}

	arr := make([]LiveFeedRecord, lf.Count)
	for _,v := range *lf.List {

		lfA := LiveFeedRecord_A{
			ID:              v.ID,
			YouTubeLiveLink: v.YouTubeLiveLink,
			Location:        v.Location,
		}

		// get all live feed children
		lfc := internal.LiveFeedChildList{}
		lfc.UseDBWriterPreferred()
		err = lfc.GetCustom("WHERE LiveFeedID = " + strconv.Itoa(int(v.ID)), "LiveFeedID DESC",0,0)
		if err != nil{
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Origin", "*")
			c.JSON(500, err.Error())
			return
		}

		if lfc.Count == 0 {
			continue
		}

		temp := new([]LiveFeedRecord_B)
		for _,j := range *lfc.List {
			lfB := LiveFeedRecord_B{
				Datetime: j.Datetime,
				ImageURL: j.ImageURL,
			}
			*temp = append(*temp, lfB)
		}

		lfR := LiveFeedRecord{
			Record:     lfA,
			RecordData: *temp,
		}

		arr = append(arr, lfR)
	}

	// reverse array to filter from latest to earliest
	fn := new([]LiveFeedRecord)
	for i := len(arr) - 1; i >= 0; i-- {
		*fn = append(*fn, arr[i])
	}

	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, gin.H{
		"res": fn,
	})
}

// change into json
func CreateLiveFeedThreat(c *gin.Context, bindingInputPtr interface{}) {

	claims := server.ExtractJwtClaims(c)
	uid := int64((claims["uid"]).(float64))

	var lfj LiveFeedJSONData
	err := c.BindJSON(&lfj)
	if err != nil {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}

	LiveFeedID,err := strconv.Atoi(lfj.LiveFeedID)
	if err != nil {
		fmt.Println(err.Error())
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}

	lfc := internal.LiveFeedChild{
		LiveFeedID: int64(LiveFeedID),
		ImageURL:   lfj.ImageFile,
		Datetime:   time.Now(),
	}

	lff := internal.LiveFeed{}
	notFound, err := lff.GetByID(int64(LiveFeedID))
	if notFound {
		fmt.Println("not found")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, "User record not found")
		return
	}
	if err != nil {
		fmt.Println(err.Error())
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}

	lfc.UseDBWriterPreferred()
	err = lfc.Set()
	if err != nil {
		fmt.Println(err.Error())
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}

	// send SMS text message to user
	t := &sns.SNS{
		AwsRegion: awsregion.AWS_us_west_2_oregon,
		SMSTransactional: true,
	}

	if err := t.Connect(); err != nil {
		fmt.Println(err.Error())
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}
	defer t.Disconnect()

	ux := internal.User{}
	ux.UseDBWriterPreferred()
	notFound, err = ux.GetByID(uid)
	if notFound {
		fmt.Println("not found")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, "User record not found")
		return
	}
	if err != nil {
		fmt.Println(err.Error())
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}

	fin := ""
	x := lfc.Datetime.Format("2006-01-02 15:04:05")
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

	message := "[SafeTracker] Live Feed Analyzer detected critical threat at " + lff.Location + " on " + fin

	fmt.Println(message)

	_, err = t.SendSMS("+1"+ ux.PhoneNumber, message, 3*time.Second)
	if err != nil{
		fmt.Println(err.Error())
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}
	_, err = t.SendSMS("+1"+ ux.PointOfContactPhoneNumber, message, 3*time.Second)
	if err != nil{
		fmt.Println(err.Error())
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(500, err.Error())
		return
	}

	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, "")

}