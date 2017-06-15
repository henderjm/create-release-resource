package main

import (
	"encoding/json"
	"os"

	"github.com/henderjm/create-release-resource/check"
)

func main() {
	json.NewEncoder(os.Stdout).Encode(check.CheckResponse{})
}
