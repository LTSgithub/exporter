// Code generated by sqlc. DO NOT EDIT.
// source: user.sql

package generated

import (
	"context"
	"time"
)

const createUser = `-- name: CreateUser :exec
insert into user (username,password,description,create_time) values(?,?,?,?)
`

type CreateUserParams struct {
	Username    string    `db:"username" json:"username"`
	Password    string    `db:"password" json:"password"`
	Description string    `db:"description" json:"description"`
	CreateTime  time.Time `db:"create_time" json:"createTime"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.exec(ctx, q.createUserStmt, createUser,
		arg.Username,
		arg.Password,
		arg.Description,
		arg.CreateTime,
	)
	return err
}

const getUserById = `-- name: GetUserById :one
select id, username, password, description, create_time from user where id = ?
`

func (q *Queries) GetUserById(ctx context.Context, id int32) (User, error) {
	row := q.queryRow(ctx, q.getUserByIdStmt, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Description,
		&i.CreateTime,
	)
	return i, err
}

const getUserByName = `-- name: GetUserByName :one
select id, username, password, description, create_time from user where username = ?
`

func (q *Queries) GetUserByName(ctx context.Context, username string) (User, error) {
	row := q.queryRow(ctx, q.getUserByNameStmt, getUserByName, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Description,
		&i.CreateTime,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :exec
update user set username = ? and password = ? and description = ?
`

type UpdateUserParams struct {
	Username    string `db:"username" json:"username"`
	Password    string `db:"password" json:"password"`
	Description string `db:"description" json:"description"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.exec(ctx, q.updateUserStmt, updateUser, arg.Username, arg.Password, arg.Description)
	return err
}
