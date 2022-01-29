CREATE TABLE `user` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户id',
    `username` varchar(100) NOT NULL DEFAULT '' COMMENT '用户名',
    `password` varchar(100) NOT NULL DEFAULT '' COMMENT '密码',
    `description` varchar(1000) NOT NULL DEFAULT '' COMMENT '描述信息',
    `create_time` time NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';

CREATE TABLE `app` (
    `app_status` longtext NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COMMENT='app状态表';

CREATE TABLE `stock` (
    `code` varchar(6) NOT NULL,
    `name` varchar(100) NOT NULL DEFAULT '',
    `type` varchar(10) NOT NULL DEFAULT '',
    `create_time` varchar(100) NOT NULL DEFAULT '',
    `update_time` varchar(100) NOT NULL DEFAULT '',
    PRIMARY KEY (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `alert_config` (
    `user_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
    `stock_code` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
    `notify_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
    `price` double NOT NULL DEFAULT '0',
    `id` varchar(36) NOT NULL DEFAULT '',
    `deadline` bigint NOT NULL,
    `create_time` bigint NOT NULL DEFAULT '0',
    `desc` varchar(100) NOT NULL DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;




