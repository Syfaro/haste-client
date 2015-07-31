package haste_test

import (
	"bytes"
	"github.com/syfaro/haste-client"
	"testing"
)

func TestNewHaste(t *testing.T) {
	hasteClient := haste.NewHaste("http://paste.syfaro.net")

	if hasteClient.Host != "http://paste.syfaro.net" {
		t.Fail()
	}
}

func TestGetLink(t *testing.T) {
	hasteClient := haste.NewHaste("http://paste.syfaro.net")
	hasteResp := &haste.Response{
		Key: "qwerty",
	}

	if hasteResp.GetLink(hasteClient) != "http://paste.syfaro.net/qwerty" {
		t.Fail()
	}
}

func TestFetch(t *testing.T) {
	hasteClient := haste.NewHaste("http://paste.syfaro.net")
	data, err := hasteClient.Fetch("uyenekemev")

	if err != nil {
		t.Error("Unable to load page")
		t.Fail()
	}

	if data != "hello" {
		t.Error("Response is bad")
		t.Fail()
	}
}

func TestUploadString(t *testing.T) {
	hasteClient := haste.NewHaste("http://paste.syfaro.net")
	hasteResp, err := hasteClient.UploadString("haste-client golang string test")

	if err != nil {
		t.Error("Unable to load page")
		t.Fail()
	}

	if hasteResp.Key == "" {
		t.Error("Key was not set")
		t.Fail()
	}

	t.Log(hasteResp.Key)

	data, err := hasteClient.Fetch(hasteResp.Key)
	if err != nil {
		t.Error("Unable to fetch")
		t.Fail()
	}

	if data != "haste-client golang string test" {
		t.Error("Data was different")
		t.Fail()
	}
}

func TestUploadBytes(t *testing.T) {
	hasteClient := haste.NewHaste("http://paste.syfaro.net")
	hasteResp, err := hasteClient.UploadBytes([]byte("haste-client golang bytes test"))

	if err != nil {
		t.Error("Unable to load page")
		t.Fail()
	}

	if hasteResp.Key == "" {
		t.Error("Key was not set")
		t.Fail()
	}

	t.Log(hasteResp.Key)

	data, err := hasteClient.Fetch(hasteResp.Key)
	if err != nil {
		t.Error("Unable to fetch")
		t.Fail()
	}

	if data != "haste-client golang bytes test" {
		t.Error("Data was different")
		t.Fail()
	}
}

func TestUploadBuffer(t *testing.T) {
	hasteClient := haste.NewHaste("http://paste.syfaro.net")
	hasteResp, err := hasteClient.UploadBuffer(bytes.NewBufferString("haste-client golang buffer test"))

	if err != nil {
		t.Error("Unable to load page")
		t.Fail()
	}

	if hasteResp.Key == "" {
		t.Error("Key was not set")
		t.Fail()
	}

	t.Log(hasteResp.Key)

	data, err := hasteClient.Fetch(hasteResp.Key)
	if err != nil {
		t.Error("Unable to fetch")
		t.Fail()
	}

	if data != "haste-client golang buffer test" {
		t.Error("Data was different")
		t.Fail()
	}
}
