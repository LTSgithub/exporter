package util

import (
	"encoding/json"
	"github.com/axgle/mahonia"
)

func ToString(i interface{}) string {
	data, _ := json.Marshal(i)
	return string(data)
}
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
