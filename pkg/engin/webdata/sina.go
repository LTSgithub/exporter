package webdata

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-logr/logr"
	"github.com/lits01/xiaozhan/domain/common"
	"github.com/lits01/xiaozhan/pkg/configs"
	_type "github.com/lits01/xiaozhan/pkg/engin/type"
	"github.com/lits01/xiaozhan/pkg/net"
	"github.com/lits01/xiaozhan/pkg/util"
	"github.com/lits01/xiaozhan/repositories/generated"
	"github.com/pkg/errors"
)

type Webdata struct {
	config configs.Configuration
	log    *logr.Logger
	token  chan bool
}

func NewWebdata(conf configs.Configuration, log *logr.Logger) *Webdata {
	sina := &Webdata{
		config: conf,
		log:    log,
		token:  make(chan bool),
	}
	sina.run()

	return sina
}

func (m *Webdata) run() {
	for {
		m.token <- true
		time.Sleep(1 * time.Second)
	}
}

func (m *Webdata) getToken() {
	<-m.token
}

func (m *Webdata) GetDayPriceList(code string) []*_type.TV {
	var resp []*_type.TV

	return resp
}

func (m *Webdata) GetStockList() []string {
	var resp []string
	flag := `href="/q//go.php/vIR_StockSearch/key/`
	for i := 1; i <= 82; i++ {
		func() {
			m.log.Info("获取新浪数据", "index", i)
			ctx, cancel := common.NewDefaultContext()
			defer cancel()
			url := fmt.Sprintf("http://vip.stock.finance.sina.com.cn/q/go.php/vIR_CustomSearch/index.phtml?p=%v&sr_p=-1", i)
			m.getToken()
			data, err := net.SendHttp(ctx, http.MethodGet, url, nil)
			if err != nil {
				m.log.Error(err, "请求sina数据错误", "url", url)
				return
			}
			ss := strings.Split(string(data), " ")
			for _, s := range ss {
				if strings.Contains(s, flag) {
					l := len(s)
					c := s[l-13 : l-7]
					_, err := strconv.Atoi(c)
					if err != nil {
						m.log.Error(err, "解析的sina的stock code错误", "code", c)
						continue
					}
					resp = append(resp, c)
				}
			}
		}()
	}

	return resp
}

func (m *Webdata) GetStocksPrice(ctx context.Context, stocks []generated.StockStatus) (map[string]*generated.StockStatus, error) {
	resp := map[string]*generated.StockStatus{}

	var params string
	for _, v := range stocks {
		vPre := v.Code[:2]
		switch vPre {
		case "60":
			params += "sh" + v.Code + ","
		case "30", "00":
			params += "sz" + v.Code + ","
		default:
			m.log.Info("warning", "stock_code", v)
		}
	}
	params = strings.TrimRight(params, ",")
	totalUrl := fmt.Sprintf("%v%v", "http://hq.sinajs.cn/list=", params)

	m.getToken()
	out, err := net.SinaSendPost(ctx, totalUrl)
	if err != nil {
		return nil, errors.Wrap(err, totalUrl)
	}
	out = util.ConvertToString(out, "gbk", "utf-8")
	ss := strings.Split(out, ";")

	for _, v := range ss {
		if !strings.Contains(v, "var") {
			continue
		}
		CodeName := strings.Split(strings.Split(v, ",")[0], "=")
		code := CodeName[0][len(CodeName[0])-6:]
		name := strings.Split(CodeName[1], "\"")[1]
		info := &generated.StockStatus{Code: code, Name: name}
		info.Sprice, err = strconv.ParseFloat(strings.Split(v, ",")[3], 64)
		if err != nil {
			continue
		}
		resp[info.Code] = info
	}

	return resp, nil
}
