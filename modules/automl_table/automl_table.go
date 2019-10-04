package automl_table

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	auth "../auth"
	"github.com/imroc/req"
)

type Body struct {
	Payload Payload `json:"payload"`
}
type Payload struct {
	Row Row `json:"row"`
}
type Row struct {
	Values        []string `json:"values"`
	ColumnSpecIds []string `json:"columnSpecIds"`
}

func OnlinePredict() error {
	token, _ := auth.ServiceAccount("./authentication.json", "https://www.googleapis.com/auth/cloud-platform")

	header := req.Header{
		"Accept":        "application/json",
		"Content-Type":  "application/json; charset=utf-8",
		"Authorization": "Bearer " + token.AccessToken,
	}

	body := Body{
		Payload: Payload{
			Row: Row{
				Values: []string{
					"39", "admin.", "married", "secondary", "no", "70", "yes", "no", "cellular", "31", "jul", "13", "11", "-1", "0", "unknown",
				},
				ColumnSpecIds: []string{
					"461385865340387328", "5073071883767775232", "1614307369947234304", "2767228874554081280", "6225993388374622208", "7378914892981469184", "3920150379160928256", "5649532636071198720", "8531836397588316160", "1037846617643810816", "3343689626857504768", "7955375645284892672", "2190768122250657792", "9108297149891739648", "6802454140678045696", "4496611131464351744",
				},
			},
		},
	}
	json_string, _ := json.Marshal(body)

	param := req.BodyJSON(json_string)
	// only url is required, others are optional.
	//
	r, err := req.Post(
		fmt.Sprintf("https://automl.googleapis.com/v1beta1/projects/%s/locations/us-central1/models/%s:predict", os.Getenv("PROJECT_NUMBER"), os.Getenv("MODEL_NAME")),
		header,
		param,
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", r) // print info (try it, you may surprise)

	return nil
}
