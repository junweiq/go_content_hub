CREATE DATABASE IF NOT EXISTS `go_cms`;

USE `go_cms`;

CREATE TABLE IF NOT EXISTS `content`
(
    `id`              bigint    NOT NULL AUTO_INCREMENT COMMENT '主鍵ID',
    `title`           varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT '',
    `description`     text COLLATE utf8mb4_unicode_ci,
    `author`          varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '作者',
    `video_url`       varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT '',
    `thumbnail`       varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '封面url',
    `category`        varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '內容分類',
    `duration`        bigint                                  DEFAULT '0' COMMENT '內容時長',
    `resolution`      varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '分辨率 如720p、1080p',
    `fileSize`        bigint                                  DEFAULT '0',
    `format`          varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '文件格式 如mp4',
    `quality`         int                                     DEFAULT '0' COMMENT '視頻質量 1-高清 2-標清',
    `approval_status` int                                     DEFAULT '1' COMMENT '審核狀態 1-審核中 2-審核通過 3-審核不通過',
    `created_at`      timestamp NOT NULL                      DEFAULT CURRENT_TIMESTAMP,
    `updated_at`      timestamp NOT NULL                      DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='內容詳情';