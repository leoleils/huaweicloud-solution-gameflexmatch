package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Https struct {
	HttpClient	*http.Client
}

func (h *Https) Close(resp *http.Response) {
	err := resp.Body.Close()
	if err != nil {
		Logger.Errorf("close resp body error: %+v", err)
	}
}

func (h *Https) Get(url string, queryParams map[string]string, token string) (body []byte, err error){
	
	// 设置client对象
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
        return nil, err
    }
	query := req.URL.Query()
	req.Header.Add("Auth-Token", token)
	for key, value := range queryParams {
		query.Add(key, value)
	}
	req.URL.RawQuery = query.Encode()
	var resp *http.Response
	h.HttpClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			DisableKeepAlives: true,
		},
	}
    resp, err = h.HttpClient.Do(req)
	if err != nil {
		return nil, err
    }
	defer resp.Body.Close()
	h.HttpClient.CloseIdleConnections()
	

	return io.ReadAll(resp.Body)
} 

func (h *Https) Post(url string, data interface{}, token string) (body []byte, err error) {
	encoderData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	param := bytes.NewBuffer([]byte(encoderData))
	req, err := http.NewRequest("POST", url, param)
	req.Close = true
	// req.Cookie()
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Auth-Token", token)
	if err != nil {
        return nil, err
    }
	h.HttpClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			DisableKeepAlives: true,
		},
	}
	resp, err := h.HttpClient.Do(req)
	if err != nil {
		return nil, err
    }
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func (h *Https) Put(url string, data interface{}) (body []byte, err error) {
	encoderData, _ := json.Marshal(data)
	param := bytes.NewBuffer([]byte(encoderData))
	fmt.Printf("PUT %s use %+v \n", url, param)
	req, err := http.NewRequest("PUT", url, param)
	req.Header.Add("Content-Type", "application/json")
	req.Close = true
	if err != nil {
        return nil, err
    }
	h.HttpClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			DisableKeepAlives: false,
		},
	}
	resp, err := h.HttpClient.Do(req)
	fmt.Printf("response state: %s, code: %d\n", resp.Status, resp.StatusCode)
	if err != nil {
		return nil, err
    }
	defer h.Close(resp)

	return io.ReadAll(resp.Body)
}


func (h *Https) Delete(url string, data interface{}) (body []byte, err error) {
	encoderData, _ := json.Marshal(data)
	param := bytes.NewBuffer([]byte(encoderData))
	fmt.Printf("DELETE %s use %+v \n", url, param)
	req, err := http.NewRequest("DELETE", url, param)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
        return nil, err
    }
	h.HttpClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			DisableKeepAlives: false,
		},
	}
	resp, err := h.HttpClient.Do(req)
	fmt.Printf("response state: %s, code: %d\n", resp.Status, resp.StatusCode)
	if err != nil {
		return nil, err
    }
	defer h.Close(resp)

	return io.ReadAll(resp.Body)
}