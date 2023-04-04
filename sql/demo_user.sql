CREATE TABLE `demo_user` (
                             `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                             `user_name` varchar(32) NOT NULL COMMENT '用户名',
                             `password` varchar(128) NOT NULL COMMENT '密码',
                             `email` varchar(32) NOT NULL DEFAULT '' COMMENT '邮箱',
                             `avatar` varchar(256) NOT NULL DEFAULT '' COMMENT '头像',
                             `created_at` datetime NOT NULL,
                             `updated_at` datetime NOT NULL,
                             PRIMARY KEY (`id`),
                             UNIQUE KEY `uk_uname` (`user_name`),
                             KEY `idx_email` (`email`),
                             KEY `idx_created_at` (`created_at`),
                             KEY `idx_updated_at` (`updated_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;