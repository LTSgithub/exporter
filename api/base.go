package api

import (
	"fmt"
	"net/http"

	"github.com/lits01/xiaozhan/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
)

type HttpResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Base struct {
	ctx *gin.Context
	log *logr.Logger
}

func (m *Base) UnmarshalRequest(req interface{}) error {
	if req == nil {
		return errors.Errorf("")
	}
	var err error

	switch m.ctx.Request.Method {
	case http.MethodGet:
		err = m.ctx.BindQuery(req)
	case http.MethodPost:
		err = m.ctx.BindJSON(req)
	}

	m.log.Info("request", "url", m.ctx.Request.RequestURI,
		"method", m.ctx.Request.Method,
		"clientIp", m.ctx.ClientIP(),
		"request", util.ToString(req),
	)

	return errors.Wrap(err, "解析请求失败")
}

func (m *Base) Response(resp interface{}, err error) {
	httpResp := HttpResponse{
		Code:    0,
		Message: "",
		Data:    resp,
	}

	if err != nil {
		m.log.Error(err, "")
		httpResp.Code = -1
		httpResp.Message = fmt.Sprintf("%+v", err)
		return
	}

	m.ctx.String(http.StatusOK, util.ToString(httpResp))
}
