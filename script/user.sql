CREATE DATABASE `cms_user`;

CREATE TABLE `t_user_detail` (
    `id` int NOT NULL AUTO_INCREMENT COMMENT '主鍵ID',
    `username` varchar(64) DEFAULT '' COMMENT '用戶名',
    `password` varchar(64) DEFAULT '' COMMENT '密碼',
    `nickname` varchar(64) DEFAULT '' COMMENT '暱稱',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '創建時間',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新時間',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='cms用戶訊息';