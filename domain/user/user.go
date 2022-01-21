package user

import (
	"github.com/go-logr/logr"
	"github.com/lits01/xiaozhan/repositories/generated"
)

type User struct {
	queries *generated.Queries
	log     *logr.Logger
}

func NewUser(log *logr.Logger) *User {
	return &User{
		log: log,
	}
}

func (u *User) List() (map[string]*UserDetail, error) {
	resp := map[string]*UserDetail{}

	list := []*UserDetail{
		{UserId: "0000000001", DingDingUrl: "https://oapi.dingtalk.com/robot/send?access_token=1904f756ae73772add9a3ba8d6f3e8540328a1735bd05d3a705009b26ff93110"},
	}

	for _, v := range list {
		resp[v.UserId] = v
	}

	return resp, nil
}

func (u *User) Details() {

}

//func NewUser(db *sql.DB, log *logr.Logger) *User {
//	return &User{queries: generated.New(db), log: log}
//}
//
//func (m *User) CreateUser(ctx context.Context, params generated.CreateUserParams) error {
//
//	if err := m.queries.CreateUser(ctx, params); err != nil {
//		return errors.Wrap(err, "写数据库失败")
//	}
//
//	return nil
//}
