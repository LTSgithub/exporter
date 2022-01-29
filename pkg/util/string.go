package util

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/axgle/mahonia"
)

var (
	src   = rand.NewSource(time.Now().UnixNano())
	RandX = rand.New(src)
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

func NewId() string {
	now := time.Now().UTC().UnixNano()
	ticks := strconv.FormatInt(now/1000, 16)
	oldTicks := strconv.FormatInt(now/1000000+62135625600000, 16)
	return fmt.Sprintf("%s-%s-%s%s-%s-%s",
		oldTicks[:8],
		oldTicks[8:12],
		ticks[10:13],
		getRandomChars(1),
		getRandomChars(4),
		getRandomChars(12),
	)
}

func getRandomChars(charLen int) string {
	chars := "abcdef0123456789"
	res := make([]byte, charLen)
	for i := 0; i < charLen; i++ {
		idx := int64(math.Floor(1 + RandX.Float64()*16))
		res[i] = chars[idx-1]
	}

	return string(res)
}
