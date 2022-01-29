-- name: GetAlertConfigListByUserId :many
select * from alert_config where user_id = ? order by create_time limit ?,?;

-- name: CreateAlertConfig :exec
insert into alert_config (id,user_id,stock_code,notify_type,price,deadline,create_time,`desc`) values (?,?,?,?,?,?,?,?);

-- name: DeleteAlertConfig :exec
delete from alert_config where id = ?;

-- name: UpdateAlertConfig :exec
Update alert_config set price = ?,deadline = ?,`desc` = ? where id = ?;


