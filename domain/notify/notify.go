package notify

import (
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type Notify struct {
	log *logr.Logger
}

func NewNotify(log *logr.Logger) *Notify {
	return &Notify{
		log: log,
	}
}

func (n *Notify) SendDingMsg(alertConfigs []*AlertConfig) error {

	for _, v := range alertConfigs {
		if err := n.sendDingMsg(v); err != nil {
			n.log.Error(err, "")
		}
	}

	return nil
}

func (n *Notify) sendDingMsg(alertConfig *AlertConfig) error {

	if alertConfig == nil {
		return nil
	}

	webHook := alertConfig.DingDingUrl
	content := `{"msgtype": "text",
		"text": {"content": "` + alertConfig.Mag + `"}
	}`

	req, err := http.NewRequest(http.MethodPost, webHook, strings.NewReader(content))
	if err != nil {
		return errors.Wrap(err, "")
	}

	client := &http.Client{}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.Wrap(err, "")
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "")
	}

	n.log.Info(string(data))

	return nil
}
