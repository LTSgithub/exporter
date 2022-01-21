package str

import (
	"encoding/json"
	"fmt"
	"hash/crc32"
	"math"
	"math/rand"
	"strconv"
	"time"
)

var (
	src   = rand.NewSource(time.Now().UnixNano())
	RandX = rand.New(src)
)

func StringHashUint32(s string) uint32 {
	return crc32.ChecksumIEEE([]byte(s))
}

func ToString(i interface{}) string {
	if i == nil {
		return ""
	}
	data, _ := json.Marshal(i)
	return string(data)
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
