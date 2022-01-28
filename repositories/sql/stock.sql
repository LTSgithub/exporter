-- name: GetNotUpdateStockList :many
select * from stock where update_time != ? limit 99;

-- name: GetNotUpdateStockCodeList :many
select code from stock where update_time != ? limit 99;

-- name: GetStockList :many
select * from stock ;

-- name: GetStockCount :one
select count(code) from stock;

-- name: CreateStock :exec
insert into stock(code,name, type,create_time,update_time) values (?,?,?,?,?);



