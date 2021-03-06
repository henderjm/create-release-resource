package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/henderjm/create-release-resource/check"
	"github.com/henderjm/create-release-resource/github"
)

func main() {
	request := check.CheckRequest{}
	err := json.NewDecoder(os.Stdin).Decode(&request)
	if err != nil {
		fmt.Println("Error has occurred unmarshalling check request")
		os.Exit(1)
	}
	checkCommand := check.NewCheckCommand(
		github.NewClient(
			"fakeuser",
			"fakepassword",
			"repository",
			"branch",
		))
	response, err := checkCommand.Execute(request)
	if err != nil {
		fmt.Println("Could not execute check")
		os.Exit(1)
	}
	err = json.NewEncoder(os.Stdout).Encode(response)
	// responseToReturn := check.CheckResponse{}
	// if request.Version.Number == response[0].Number {
	// 	err = json.NewEncoder(os.Stdout).Encode(check.CheckResponse{})
	// } else {
	// 	responseToReturn = append(responseToReturn, concourse.Version{Number: response[0].Number})
	// 	err = json.NewEncoder(os.Stdout).Encode(responseToReturn)
	// }
	if err != nil {
		fmt.Println("Can not return version to concourse")
		fmt.Println(err)
		os.Exit(1)
	}
}
