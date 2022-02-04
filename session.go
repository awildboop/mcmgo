package mcmgo

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/imroc/req/v2"
)

type Session struct {
	apiToken    string
	sharedToken bool
	Member      *Member
	RateLimited bool
	RetryAfter  uint
	HttpClient  *req.Client
}

func NewSession(token string, sharedToken bool) (*Session, error) {
	session = &Session{}

	session.apiToken = token
	session.sharedToken = sharedToken

	session.RateLimited = false
	session.RetryAfter = 0

	var tokenType string
	if sharedToken {
		tokenType = "Shared "
	} else {
		tokenType = "Private "
	}

	session.HttpClient = req.C().
		SetUserAgent("mcmgo v0.0.1 - github.com/awildboop/mcmgo").
		SetCommonHeader("Authorization", tokenType+token).
		OnBeforeRequest(reqCheck).
		OnAfterResponse(resCheck).
		SetBaseURL("https://api.mc-market.org/v1/")

	return session, nil
}

func (ses *Session) Connect() (err error) {

	return
}

func reqCheck(*req.Client, *req.Request) error {
	if session.RateLimited && session.RetryAfter >= uint(time.Now().UnixMilli()) {
		return errors.New("api rate limit is active")
	} else {
		return nil
	}
}

func resCheck(req *req.Client, res *req.Response) error {
	if res.Header.Get("Retry-After") != "" {
		log.Println("Rate-limit triggered...")
		delay, err := strconv.ParseInt(res.Header.Get("Retry-After"), 10, 64)
		if err != nil {
			log.Fatalln("Failed parsing Retry-After delay, exiting to prevent accidental abuse...")
			return errors.New("failed parsing retry-after delay") // this is pointless, but good for posterity
		}

		session.RetryAfter = uint(time.Now().UnixMilli() + delay)
		session.RateLimited = true

		go func() {
			time.Sleep(time.Duration(delay) * time.Millisecond)
			session.RateLimited = false
		}()
	}

	var ResData Response
	err := res.UnmarshalJson(&ResData)
	if err != nil {
		return err
	}

	if ResData.Result == "error" && ResData.Error != nil {
		var ErrorData ErrorResponse
		err := json.Unmarshal(ResData.Error, &ErrorData)
		if err != nil {
			return err
		}

		return errors.New(strings.TrimSuffix(strings.ToLower(ErrorData.Message), "."))
	} else if ResData.Result != "success" && ResData.Data != nil {
		return errors.New("encountered unexpected result & data type")
	} else {
		return nil
	}
}
