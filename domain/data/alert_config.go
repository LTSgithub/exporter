package data

import (
	"context"
	"database/sql"
	"time"

	"github.com/lits01/xiaozhan/domain/data/model"

	"github.com/lits01/xiaozhan/pkg/util"
	"github.com/lits01/xiaozhan/repositories/generated"
	"github.com/pkg/errors"

	"github.com/go-logr/logr"
	"github.com/lits01/xiaozhan/pkg/configs"
)

type AlertConfig struct {
	config configs.Configuration
	log    *logr.Logger
	query  *generated.Queries
}

func NewAlertConfig(config configs.Configuration, log *logr.Logger, db *sql.DB) *AlertConfig {
	return &AlertConfig{
		config: config,
		log:    log,
		query:  generated.New(db),
	}
}

func (m *AlertConfig) List(ctx context.Context, req *model.AlertConfigListRequest) (*model.AlertConfigListResponse, error) {
	resp := &model.AlertConfigListResponse{}

	arg := generated.GetAlertConfigListByUserIdParams{
		UserID: req.UserId,
	}
	arg.Offset, arg.Limit = util.Paging(req.Page, req.PageSize)
	rets, err := m.query.GetAlertConfigListByUserId(ctx, arg)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resp.Items = make([]*model.AlertConfigInfo, len(rets))
	for k, v := range rets {
		resp.Items[k] = &model.AlertConfigInfo{
			Id:         v.ID,
			UserId:     v.UserID,
			StockCode:  v.StockCode,
			NotifyType: v.NotifyType,
			Price:      v.Price,
			Desc:       v.Desc,
			Deadline:   v.Deadline,
			CreateTime: v.CreateTime,
		}
	}

	return resp, nil
}

func (m *AlertConfig) Update(ctx context.Context, req *model.AlertConfigUpdateRequest) (*model.AlertConfigUpdateResponse, error) {
	resp := &model.AlertConfigUpdateResponse{}

	arg := generated.UpdateAlertConfigParams{
		Price:    req.Price,
		Deadline: req.Deadline,
		Desc:     req.Desc,
		ID:       req.Id,
	}
	if err := m.query.UpdateAlertConfig(ctx, arg); err != nil {
		return nil, errors.WithStack(err)
	}

	return resp, nil
}

func (m *AlertConfig) Create(ctx context.Context, req *model.AlertConfigCreateRequest) (*model.AlertConfigCreateResponse, error) {
	resp := &model.AlertConfigCreateResponse{}

	arg := generated.CreateAlertConfigParams{
		ID:         util.NewId(),
		UserID:     req.UserId,
		StockCode:  req.StockCode,
		NotifyType: req.NotifyType,
		Price:      req.Price,
		Deadline:   req.Deadline,
		CreateTime: time.Now().Unix(),
		Desc:       req.Desc,
	}
	if err := m.query.CreateAlertConfig(ctx, arg); err != nil {
		return nil, errors.WithStack(err)
	}

	return resp, nil
}

func (m *AlertConfig) Delete(ctx context.Context, req *model.AlertConfigDeleteRequest) (*model.AlertConfigDeleteResponse, error) {
	resp := &model.AlertConfigDeleteResponse{}

	if err := m.query.DeleteAlertConfig(ctx, req.Id); err != nil {
		return nil, errors.WithStack(err)
	}

	return resp, nil
}
