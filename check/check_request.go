package check

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/henderjm/create-release-resource/concourse"
)

type CheckRequest struct {
	Version concourse.Version `json:"version"`
}

func NewCheckRequest(request []byte) (CheckRequest, error) {
	var checkRequest CheckRequest
	if err := json.NewDecoder(bytes.NewReader(request)).Decode(&checkRequest); err != nil {
		return CheckRequest{}, fmt.Errorf("Invalid parameters: %s\n", err)
	}

	return checkRequest, nil
}
