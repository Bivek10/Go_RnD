package main

import "fmt"

func main() {

	//map making approach
	userdetail := make(map[string]any)
	userdetail["username"] = "Bivek Karki"
	userdetail["contact"] = 9800965652
	userdetail["location"] = "Dilibazar, 5"
	fmt.Println(userdetail["contact"])

	//
	errorDetail := ErrorDetail{
		errorCode:    200,
		errorMessage: "Something went wrong",
	}

	fmt.Println(errorDetail)

}

type ErrorDetail struct {
	errorCode    int
	errorMessage string
}
