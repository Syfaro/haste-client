package haste

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type HasteResponse struct {
	Key string `json:"key"`
}

func (resp *HasteResponse) GetLink(haste *Haste) string {
	return haste.Host + "/" + resp.Key
}

type Haste struct {
	Host string
}

func NewHaste(host string) *Haste {
	return &Haste{
		Host: host,
	}
}

func (haste *Haste) UploadString(data string) (*HasteResponse, error) {
	return haste.UploadBuffer(bytes.NewBuffer([]byte(data)))
}

func (haste *Haste) UploadBytes(data []byte) (*HasteResponse, error) {
	return haste.UploadBuffer(bytes.NewBuffer(data))
}

func (haste *Haste) UploadBuffer(data *bytes.Buffer) (*HasteResponse, error) {
	req, err := http.NewRequest("POST", "http://paste.syfaro.net/documents", data)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiResp HasteResponse
	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		return nil, err
	}

	return &apiResp, nil
}
