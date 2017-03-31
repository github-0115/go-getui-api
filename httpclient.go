package getui

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

func Get(url string, auth_token string) (string, error) {

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("authtoken", auth_token)
	req.Header.Add("Charset", CHARSET)
	req.Header.Add("Content-Type", CONTENT_TYPE_JSON)

	client := &http.Client{
		Timeout: DEFAULT_CONNECTION_TIMEOUT * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	result, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func BytePost(url string, auth_token string, bodyByte []byte) (string, error) {
	body := bytes.NewBuffer(bodyByte)

	req, err := http.NewRequest("POST", url, body)
	req.Header.Add("authtoken", auth_token)
	req.Header.Add("Charset", CHARSET)
	req.Header.Add("Content-Type", CONTENT_TYPE_JSON)

	client := &http.Client{
		Timeout: DEFAULT_CONNECTION_TIMEOUT * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	result, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func Delete(url string, auth_token string, bodyByte []byte) (string, error) {

	body := bytes.NewBuffer(bodyByte)

	req, err := http.NewRequest("DELETE", url, body)
	req.Header.Add("authtoken", auth_token)
	req.Header.Add("Charset", CHARSET)
	req.Header.Add("Content-Type", CONTENT_TYPE_JSON)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	result, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "", err
	}

	return string(result), nil
}
