package omdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"time"
)

func (o *OMDB) GetListMovie(search string, page int64) (responseObject *OMDBList, err error) {
	endpoint := fmt.Sprintf("%s?apikey=%s&s=%s&page=%d", o.Cfg.External["omdb"].Endpoint, o.Cfg.External["omdb"].Key, search, page)
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	timeout := time.Duration(10) * time.Second
	transport := &http.Transport{
		ResponseHeaderTimeout: timeout,
		Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, timeout)
		},
		DisableKeepAlives: true,
	}
	client := &http.Client{
		Transport: transport,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyBytes, &responseObject)
	if err != nil {
		return nil, err
	}

	return responseObject, nil
}

func (o *OMDB) GetDetailMovie(movieID string) (responseObject *OMDBDetail, err error) {
	endpoint := fmt.Sprintf("%s?apikey=%s&i=%s&plot=full", o.Cfg.External["omdb"].Endpoint, o.Cfg.External["omdb"].Key, movieID)
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	timeout := time.Duration(10) * time.Second
	transport := &http.Transport{
		ResponseHeaderTimeout: timeout,
		Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, timeout)
		},
		DisableKeepAlives: true,
	}
	client := &http.Client{
		Transport: transport,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyBytes, &responseObject)
	if err != nil {
		return nil, err
	}

	status, err := strconv.ParseBool(responseObject.Response)
	if err != nil {
		return nil, fmt.Errorf("response from api any problem")
	}

	if !status {
		return nil, fmt.Errorf("request from internal service to third party any problem")
	}

	return responseObject, nil
}
