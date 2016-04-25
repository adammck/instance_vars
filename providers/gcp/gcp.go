package gcp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

type GCP struct {
}

const attrsURL = "http://metadata.google.internal/computeMetadata/v1/instance/attributes/?recursive=true"

var isUpper = regexp.MustCompile(`^[0-9A-Z_]+$`)

func New() *GCP {
	return &GCP{}
}

func (g *GCP) Get() (map[string]string, error) {
	vars := map[string]string{}

	client := &http.Client{
		Timeout: 1 * time.Second,
	}
	req, err := http.NewRequest("GET", attrsURL, nil)
	req.Header.Add("Metadata-Flavor", "Google")
	res, err := client.Do(req)
	if err != nil {
		return vars, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return vars, err
	}

	json.Unmarshal(body, &vars)

	// Remove keys which don't look like env vars.
	for k, _ := range vars {
		if !isUpper.MatchString(k) {
			delete(vars, k)
		}
	}

	return vars, nil
}
