package insertData

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	)

type Result struct {
	Data []struct{
		Type string `json:"type"`
		Attributes struct{
			TeamOne 	string 		`json:"teamOne"`
			TeamTwo 	string 		`json:"teamTwo"`
			TeamOneGoals string 	`json:"teamOneGoals"`
			TeamTwoGoals string 	`json:"teamTwoGoals"`
		} 	`json:"attributes"`
	} `json:"data"`
}

// now will the error json that endpoint will return if something is wrong
// with the request's json

type Error struct {
	Status string `json:"status"`
	Source struct{
		Pointer 	string 	`json:"pointer"`
	} `json:"source"`
	Title string	`json:"title"`
	Detail string 	`json:"detail"`
}

type SuccessMessage struct {
	Meta struct{
		Success struct{
			Title string `json:"title"`
			Message string `json:"message"`
			Status string `json:"status"`
		} `json:"success"`
	} `json:"meta"`
}

func InsertData(r *http.Request) ([]byte, error) {
	var result Result
	b, _ := ioutil.ReadAll(r.Body) // reads the body of the request to /insert
	err := json.Unmarshal(b, &result)
	if err != nil{
		fmt.Println(err)
	}
	// checks if the request body is correct
	isSuccess, message := isJSONCorrect(result) // will create isJSONCorrect
	if isSuccess{
		return getSuccessJsonBody("Game result was successfully inserted"), nil
	}else {
		return getErrorFromResult("/insert", "Invalid Attribute", message), nil
	}

}

func getErrorFromResult(pointer string, title string, detail string) []byte {
	var errorMessage Error
	errorMessage.Status = "422"
	errorMessage.Source.Pointer = pointer
	errorMessage.Title = title
	errorMessage.Detail = detail
	jsonBody, _ := json.Marshal(errorMessage)
	return jsonBody
}

func getSuccessJsonBody(message string) []byte {
	var successMessage SuccessMessage
	successMessage.Meta.Success.Title = "success"
	successMessage.Meta.Success.Status = "200"
	jsonBody, _ := json.Marshal(successMessage)
	return jsonBody
}

func isJSONCorrect(r Result) (bool, string) {
	if (len(r.Data[0].Attributes.TeamOne)==0 || len(r.Data[0].Attributes.TeamTwo)==0) {
		fmt.Println("TeamOne and TeamTwo need to have a value")
		return false, "TeamOneGoals and TeamTwoGoals need to have a value"
	}
	return true, "success"
}