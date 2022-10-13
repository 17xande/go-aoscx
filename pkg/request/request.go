package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Eg: http://192.168.0.1/rest/v1/
const UrlAPI = "%s://%s/rest/%s/%s"

type Request struct {
	Protocol   string
	IP         string
	APIVersion string
	Path       string
}

func New(ip, path string) *Request {
	return &Request{
		Protocol:   "http",
		IP:         ip,
		APIVersion: "v8",
		Path:       path,
	}
}

func (r *Request) genUrl() string {
	return fmt.Sprintf(UrlAPI, r.Protocol, r.IP, r.APIVersion, r.Path)
}

func (r *Request) Get() (string, error) {
	url := r.genUrl()
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("can't create get request: %w", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("can't get VLAN info: %w", err)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("can't read response body: %w", err)
	}

	pretty, err := jsonIndent(resBody)
	if err != nil {
		return "", err
	}

	return pretty, nil
}

func jsonIndent(data []byte) (string, error) {
	pretty := bytes.Buffer{}
	if err := json.Indent(&pretty, data, "", "  "); err != nil {
		return "", fmt.Errorf("can't pretty marshal data: %w", err)
	}

	return pretty.String(), nil
}
