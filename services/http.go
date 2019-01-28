package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// HTTPRequest create
func HTTPRequest() {

}

// GetToken .-
func GetToken() string {
	bodyToJSON := map[string]string{"email": os.Getenv("USER_EMAIL"), "password": os.Getenv("USER_PASSWORD")}
	jsonBody, _ := json.Marshal(bodyToJSON)
	url := os.Getenv("BYRD_API")
	req, _ := http.NewRequest("POST", url+"/auth", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		panic(err)
	}
	fmt.Println("response Body:", string(body))

	return string(body)
}
