package util

import (
	"github.com/RiverDanceGit/yeepayGo"
	"github.com/RiverDanceGit/yeepayGo/enum"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
)

func getPostBody(params map[string]string) string {
	var list []string
	query := ""
	for key, val := range params {
		query = key + "=" + Urlencode(val)
		list = append(list, query)
	}
	sort.Strings(list)
	return strings.Join(list, "&")
}

func Post(url string, queryBody map[string]string, params map[string]string, headers map[string]string, logger yeepayGo.YeepayLoggerInterface) (HttpResponse, error) {
	startTime := time.Now()
	var httpResp HttpResponse

	postBody := ""
	if queryBody != nil {
		postBody = getPostBody(queryBody)
	}

	req, err := http.NewRequest(enum.HTTP_METHOD_POST, url, strings.NewReader(postBody))
	if err != nil {
		httpResp.SetStartTime(startTime)
		logger.Error(req.URL.String(), "|", headers, "|", postBody, "|", err)
		return httpResp, err
	}
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		httpResp.SetStartTime(startTime)
		logger.Error(req.URL.String(), "|", headers, "|", postBody, "|", err)
		return httpResp, err
	}
	defer resp.Body.Close()

	var bodyBytes []byte
	if resp.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(resp.Body)
	}

	httpResp.SetStartTime(startTime)
	httpResp.SetCode(resp.StatusCode)
	httpResp.SetBytes(bodyBytes)
	logger.Debug(req.URL.String(), "|", headers, "|", postBody, "|", resp.StatusCode, "|", httpResp.GetLatencyStr(), "|", string(bodyBytes))
	return httpResp, nil
}
