package main

import (
	"context"
	"github.com/axgle/mahonia"
	"github.com/lits01/xiaozhan/pkg/net"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func get() {

	out, err := net.SinaSendPost(context.Background(), "http://hq.sinajs.cn/list=sz002307,sh600928")
	if err != nil {
		log.Fatalln(err)
	}

	ss := strings.Split(out, ";")
	for _, v := range ss {
		log.Println(ConvertToString(v, "gbk", "utf-8"))
	}
}

func SendDingMsg(msg string) {
	//请求地址模板
	//webHook := `https://oapi.dingtalk.com/robot/send?access_token=04c381fc31944ad2905f31733e31fa15570ae12efc857062dab16b605a369e4c`
	webHook := `https://oapi.dingtalk.com/robot/send?access_token=1904f756ae73772add9a3ba8d6f3e8540328a1735bd05d3a705009b26ff93110`
	content := `{"msgtype": "text",
		"text": {"content": "` + msg + `"}
	}`
	//创建一个请求
	req, err := http.NewRequest("POST", webHook, strings.NewReader(content))
	if err != nil {
		log.Println(err)
		return
	}

	client := &http.Client{}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	//关闭请求
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println(resp.StatusCode)
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(string(data))
}

func main() {

	SendDingMsg("stock")
}
