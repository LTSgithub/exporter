-- name: GetUserById :one
select * from user where id = ?;

-- name: GetUserByName :one
select * from user where username = ?;

-- name: CreateUser :exec
insert into user (username,password,description,create_time) values(?,?,?,?) ;

-- name: UpdateUser :exec
update user set username = ? and password = ? and description = ? ;




