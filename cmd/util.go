package main

import (
	"evantu/safetracker-server/internal"
	"fmt"
	"regexp"
	"time"
)

type Credentials struct {
	EmailAddress    string `form:"EmailAddress" json:"EmailAddress"`
	Password string `form:"Password" json:"Password"`

	// for the claims
	UserID   int64  `form:"UserID" json:"UserID"`
}

type DashboardData struct {
	YouTubeLiveLink string `form:"YouTubeLiveLink" json:"YouTubeLiveLink"`
	LiveFeedChildID int64 `form:"LiveFeedChildID" json:"LiveFeedChildID"`

	LiveFeedID int64    `form:"LiveFeedID" json:"LiveFeedID"`
	ImageURL   string   `form:"ImageURL" json:"ImageURL"`
	Datetime   time.Time `form:"Datetime" json:"Datetime"`
}

type LiveFeedRecord struct {

	Record LiveFeedRecord_A `form:"Record" json:"Record"`

	RecordData []LiveFeedRecord_B `form:"RecordData" json:"RecordData"`
}

type LiveFeedRecord_A struct {
	ID int64  `form:"UserID" json:"UserID"`
	YouTubeLiveLink string  `form:"YouTubeLiveLink" json:"YouTubeLiveLink"`
	Location string  `form:"Location" json:"Location"`
}

type LiveFeedRecord_B struct {
	Datetime time.Time `form:"Datetime" json:"Datetime"`
	ImageURL string  `form:"ImageURL" json:"ImageURL"`
}

type TwitterJSONData struct {
	TextContent string `json:"TextContent"`
	AuthorTwitterID string `json:"AuthorTwitterID"`
	AuthorTwitterTag string `json:"AuthorTwitterTag"`
	ImageFile string `json:"ImageFile"`
}

type LiveFeedJSONData struct {
	ImageFile string `json:"ImageFile"`
	LiveFeedID string `json:"LiveFeedID"`
}

func ValidateCredentials(cred Credentials) (int64, error) {
	if cred.EmailAddress == "" {
		return 0,fmt.Errorf("Email invalid")
	}
	if cred.Password == "" {
		return 0,fmt.Errorf("Password invalid")
	}

	user := internal.User{}
	user.UseDBWriterPreferred()
	notFound, err := user.GetByPassword(cred.Password, cred.EmailAddress)
	if notFound {
		return 0,fmt.Errorf("User is nil")
	}
	if err != nil {
		return 0,err
	}
	return user.ID, nil

}


func IsEmailValid(e string) bool {

	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}

func IsPhoneNumberValid(p string) bool {
	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)

	return re.MatchString(p)
}

func Contains(s []string, str string) bool {
	for _,v := range s {
		if v == str {
			return true
		}
	}

	return false
}