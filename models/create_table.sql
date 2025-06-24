CREATE TABLE `user` (
    `id` bigint(20)  NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) NOT NULL COMMENT '用户ID',
    `username` varchar(64)  NOT NULL COMMENT '用户名',
    `password` varchar(64)  NOT NULL COMMENT '密码',
    `email` varchar(64) COMMENT '邮箱',
    `gender` tinyint(4) NOT NULL DEFAULT '0' COMMENT '性别',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `extend` varchar(128) COMMENT '扩展字段',
    `extend1` varchar(128) COMMENT '扩展字段1',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE,
    UNIQUE KEY `idx_username` (`username`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ;