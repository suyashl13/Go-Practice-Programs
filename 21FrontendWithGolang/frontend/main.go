package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// PerformGetRequest()
	// PerformPostRequest()
	// PerformFormDataRequest()
	DecodeJson()
}

func PerformGetRequest() {
	var myUrl = "http://localhost:8000/get"

	response, error := http.Get(myUrl)
	if error != nil {
		panic("Something went wrong..")
	}
	defer response.Body.Close()

	fmt.Println("Status Code :", response.StatusCode)
	fmt.Println("Content Length :", response.ContentLength)

	var responseBuilder strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseBuilder.Write(content)

	fmt.Println("Byte Count is : ", byteCount)
	fmt.Println(responseBuilder.String())

}

func PerformPostRequest() {
	const myUrl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
		{
			"course_name": "lets go with golang",
			"price":0,
			"platform" : "learncodeonline.com"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	data_contnet, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Response : ", string(data_contnet))
}

func PerformFormDataRequest() {
	const myUrl = "http://localhost:8000/postform"

	// create a formdata
	data := url.Values{}
	data.Add("firstname", "hitesh")
	data.Add("lastname", "hitesh@go.dev")

	res, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	content, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(content))

}

// type course struct {
// 	Name     string `json:"coursename"`
// 	Price    int
// 	Platform string   `json:"website"`
// 	Password string   `json:"-"`
// 	Tags     []string `json:"tags,omitempty"`
// }

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website" : "learncodeonline.in",
		"tags": ["web-dev", "js"]
	}
	`)

	// var lcocourse course
	// isValid := json.Valid(jsonDataFromWeb)

	// if isValid {
	// 	fmt.Println("JSON is valid")
	// 	json.Unmarshal(jsonDataFromWeb, &lcocourse)
	// 	fmt.Printf("%#v\n", lcocourse)
	// } else {
	// 	fmt.Println("Json was not valid.")
	// }

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	// fmt.Println(myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("KEY %v, VAL %v (%T)\n", k, v, v)
	}

}
