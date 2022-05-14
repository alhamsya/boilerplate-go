package spotify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/alhamsya/boilerplate-go/lib/managers/custom_error"
)

func (s *Spotify) GetUser() (responseObject *Profile, err error) {
	endpoint := fmt.Sprintf("%s/me", s.Cfg.External["omdb"].Endpoint)
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, customError.WrapFlag(err, "http", "request")
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer BQAys7CwOZMuNx_0wmscPQVu9TvTR8Bv4L6IGqHZhHQDAdrp6PzJXo93tls8r0qmlcSHJC40SK7BLgmBl1X0O5sLuabcT4AOzeFsijnSUGKGa25vi3y17vXm4a0lAS_Z90Iv7YCDOG5-HcuPkq3F9QTXBQ")

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
		return nil, customError.WrapFlag(err, "http", "Do")
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, customError.WrapFlag(err, "ioutil", "ReadAll")
	}

	err = json.Unmarshal(bodyBytes, &responseObject)
	if err != nil {
		return nil, customError.WrapFlag(err, "json", "Unmarshal")
	}

	return responseObject, nil
}
