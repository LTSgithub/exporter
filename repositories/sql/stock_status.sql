-- name: GetNotUpdateStockList :many
select * from stock_status where update_time != ? limit 99;

-- name: GetStockStatusList :many
select * from stock_status ;

-- name: GetStockStatusCount :one
select count(code) from stock_status;

-- name: UpdateStockStatus :exec
update stock_status set sprice = ? , name = ? , update_time = ? where code = ?;

-- name: InsertStockStatus :exec
insert into stock_status(code,name, sprice,type,create_time,update_time) values (?,?,?,?,?,?);



