package fetchzinc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mail-indexer/constants"
	"mail-indexer/models"
	"net/http"
	"time"
)

// Creates the Index for email-indexer (THIS ALSO COULD BE DONE MANUALLY :/)
func FetchCreateZincIndex() {
	username := constants.USER_NAME
	password := constants.PASSWORD

	body := constants.EMAIL_INDEXER_INDEX

	jsonBody, parserError := json.Marshal(body)
	if parserError != nil {
		fmt.Println("Error when parsing JSON:", parserError)
		return
	}

	url := constants.SERVER + constants.ENDPOINT_INDEX
	request, requestError := http.NewRequest(constants.METHOD_POST, url, bytes.NewBuffer(jsonBody))
	if requestError != nil {
		fmt.Println("Error creating request:", requestError)
		return
	}

	request.SetBasicAuth(username, password)
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, responseError := client.Do(request)
	if responseError != nil {
		fmt.Println("Error sending request:", responseError)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Error:", response.Status)
		return
	}
}

// receives an array of emails and then send them to zincSearch
func FetchZing(allEmails []models.Email) {
	username := constants.USER_NAME
	password := constants.PASSWORD

	var body models.ZincRequest
	body.Index = constants.INDEXER_NAME
	body.Records = allEmails

	jsonBody, parserError := json.Marshal(body)
	if parserError != nil {
		fmt.Println("Error encoding JSON:", parserError)
		return
	}

	url := constants.SERVER + constants.ENDPOINT
	request, requestError := http.NewRequest(constants.METHOD_POST, url, bytes.NewBuffer(jsonBody))
	if requestError != nil {
		fmt.Println("Error creating request:", requestError)
		return
	}

	request.SetBasicAuth(username, password)
	request.Header.Set("Content-Type", "application/json")

	start := time.Now()

	client := &http.Client{}
	response, responseError := client.Do(request)
	if responseError != nil {
		fmt.Println("Error sending request:", responseError)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Error:", response.Status)
		return
	}

	fmt.Println("start fetch: ", start)
	fmt.Println("end fetch: ", time.Since(start))
}
