package net

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type HttpRequestOption struct {
	Header map[string]string
}

func getHttpRequestOption(opts ...HttpRequestOption) *HttpRequestOption {
	defaultRequest := &HttpRequestOption{
		Header: map[string]string{
			"Content-Type": "application/json",
		},
	}

	if len(opts) == 0 {
		return defaultRequest
	}

	opt := opts[0]
	if opt.Header == nil {
		opt.Header = map[string]string{"Content-Type": "application/json"}
	}

	if _, ok := opt.Header["Content-Type"]; !ok {
		opt.Header["Content-Type"] = "application/json"
	}

	return &opt
}

func SendHttp(ctx context.Context, method string, url string, in interface{}, opts ...HttpRequestOption) ([]byte, error) {
	opt := getHttpRequestOption(opts...)

	data, err := json.Marshal(in)
	if err != nil {
		return nil, errors.Wrap(err, "json序列化失败")
	}

	request, err := http.NewRequestWithContext(ctx, method, url, strings.NewReader(string(data)))
	if err != nil {
		return nil, errors.Wrapf(err, "创建请求失败，url[%v]", url)
	}
	for k, v := range opt.Header {
		request.Header.Set(k, v)
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, errors.Wrapf(err, "请求失败,url[%v]", url)
	}
	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	if response.StatusCode != http.StatusOK {
		return nil, errors.Errorf("请求失败,url[%v],status_code[%v]", url, response.StatusCode)
	}
	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "解析返回值失败,url[%v]", url)
	}

	return data, nil
}

func deaHttpResult(data []byte, out interface{}) error {
	st := &struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}{}

	if err := json.Unmarshal(data, st); err != nil {
		fmt.Println("http return", string(data))
		return errors.Wrap(err, "数据异常")
	}
	if st.Code != 0 {
		return errors.Errorf("code,%v]", st.Code)
	}

	if err := json.Unmarshal(data, out); err != nil {
		fmt.Println("http return", string(data))
		return errors.Wrap(err, "数据异常")
	}

	return nil
}

func SinaSendPost(ctx context.Context, url string) (string, error) {

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return "", errors.Wrapf(err, "创建请求失败，url[%v]", url)
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", errors.Wrapf(err, "请求失败,url[%v]", url)
	}
	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	if response.StatusCode != http.StatusOK {
		return "", errors.Errorf("请求失败,url[%v],status_code[%v]", url, response.StatusCode)
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", errors.Wrapf(err, "解析返回值失败,url[%v]", url)
	}

	return string(data), err
}
