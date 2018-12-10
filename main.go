package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
	"bytes"
	"flag"
	"github.com/alexcesaro/log/stdlog"
)



type Response struct {
	Name    string    `json:"name"`
}


type AuthData struct {
	AccessToken string `json:"access_token"`
	EventID     int    `json:"event_id"`
	EventUserID int    `json:"event_user_id"`
}

	
type VoteData struct {
	EventQuestionID        int `json:"event_question_id"`
	EventQuestionScore     int `json:"event_question_score"`
	EventQuestionUserScore int `json:"event_question_user_score"`
}

type EventData []struct {
	URL struct {
		App             string `json:"app"`
		Admin           string `json:"admin"`
		Wall            string `json:"wall"`
		WallDirect      string `json:"wall_direct"`
		AdminEmbeddable struct {
			Questions string `json:"questions"`
			Twitter   string `json:"twitter"`
		} `json:"admin_embeddable"`
	} `json:"url"`
	OutOfDate int `json:"out_of_date"`
	Wall      struct {
		TransparentLogoBox     bool `json:"transparent_logo_box"`
		TransparentPartnersBox bool `json:"transparent_partners_box"`
	} `json:"wall"`
	Attrs struct {
		Signin struct {
		} `json:"signin"`
		Captcha struct {
			Enabled bool `json:"enabled"`
		} `json:"captcha"`
		Whitelabel          bool `json:"whitelabel"`
		EnableWelcomeScreen bool `json:"enable_welcome_screen"`
		Questions           struct {
			DisableAnonymous bool `json:"disable_anonymous"`
		} `json:"questions"`
		SsoRequiresConsent bool `json:"sso_requires_consent"`
		EnableIdeas        bool `json:"enable_ideas"`
	} `json:"attrs"`
	Code      string `json:"code"`
	Localized struct {
		DateFrom time.Time `json:"date_from"`
		DateTo   time.Time `json:"date_to"`
	} `json:"localized"`
	PartnersFiles      []interface{} `json:"partners_files"`
	EventID            int           `json:"event_id"`
	EventGroupID       int           `json:"event_group_id"`
	UUID               string        `json:"uuid"`
	Hash               string        `json:"hash"`
	IsPublic           bool          `json:"is_public"`
	EnableQuestions    bool          `json:"enable_questions"`
	EnablePolls        bool          `json:"enable_polls"`
	Name               string        `json:"name"`
	Img                string        `json:"img"`
	Location           string        `json:"location"`
	DateFrom           time.Time     `json:"date_from"`
	DateTo             time.Time     `json:"date_to"`
	PlanID             int           `json:"plan_id"`
	Timezone           string        `json:"timezone"`
	Locale             string        `json:"locale"`
	EnableAutocomplete bool          `json:"enable_autocomplete"`
	DateCreated        time.Time     `json:"date_created"`
	Account            struct {
		UUID string `json:"uuid"`
		Name string `json:"name"`
	} `json:"account"`
	Owner struct {
		Features []interface{} `json:"features"`
		Name     string        `json:"name"`
		UUID     string        `json:"uuid"`
	} `json:"owner"`
}


func GetEventData(event_code string) EventData {
    eventUrl := fmt.Sprintf("https://app2.sli.do/api/v0.5/events?code=%s", event_code)
    // logger.Debug("eventUrl:", eventUrl)
	response, err := http.Get(eventUrl)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err)
	}
	var ev EventData
	json.Unmarshal(responseData, &ev)
    return ev
}


func AuthEvent(event_uuid string) AuthData {
    authUrl := fmt.Sprintf("https://app2.sli.do/api/v0.5/events/%s/auth", event_uuid)
	form := url.Values{}
	body := bytes.NewBufferString(form.Encode())
	rsp, err := http.Post(authUrl, "", body)
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()
	body_byte, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}
	var auth AuthData
    json.Unmarshal(body_byte, &auth)    
    // logger.Debug("AccessToken:", auth.AccessToken)
    return auth
}


func VoteUp(auth AuthData, question_id string) VoteData {
    var jsonStr = []byte(`{"score":"1"}`)
    client := &http.Client{}
    voteUrl := fmt.Sprintf("https://app2.sli.do/api/v0.5/events/%d/questions/%s/like", auth.EventID, question_id) 
    // logger.Debug("voteUrl:", voteUrl)
    req, err := http.NewRequest("POST", voteUrl, bytes.NewBuffer(jsonStr))
    req.Header.Add("Authorization", fmt.Sprintf("Bearer %s",string(auth.AccessToken)))
    req.Header.Set("Content-Type", "application/json")
    rsp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
    defer rsp.Body.Close()
    // logger.Debug("VoteUp response Status:", rsp.Status)
	body_byte, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}
	var vote VoteData
    json.Unmarshal(body_byte, &vote)    
    return vote
}

func printUsageErrorAndExit(format string, values ...interface{}) {
	fmt.Fprintf(os.Stderr, "ERROR: %s\n", fmt.Sprintf(format, values...))
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "Available command line options:")
	flag.PrintDefaults()
	os.Exit(64)
}


var (
    eventPtr = flag.String("event", "xxxxx", "event code")
    voteCountPtr = flag.Int("votes", 42, "votes count")
    questionPtr = flag.String("question", "zzzzz", "question id")
)


func main() {
    
    flag.Parse()
    logger := stdlog.GetFromFlags()
    if *eventPtr == "" {
		printUsageErrorAndExit("-event is required")
    }
    
    if *questionPtr == "" {
		printUsageErrorAndExit("-question is required")
    }

    if *voteCountPtr  < 1 {
		printUsageErrorAndExit("-votes must me greater than 1")
	}    

    logger.Info("Starting sli.do voter")
    logger.Debug("voting event: ", *eventPtr)
    logger.Debug("voting count: ", *voteCountPtr)
    logger.Debug("question to vote: ", *questionPtr)

    event := GetEventData(*eventPtr)
    var vote VoteData
    for i := 0; i < *voteCountPtr; i++ {
        auth := AuthEvent(event[0].UUID)
        vote = VoteUp(auth, *questionPtr)
	}
    logger.Info("Final Vote: ", vote.EventQuestionScore)
}