use work_thor;

CREATE TABLE `content`
(
    `id`           bigint(20) AUTO_INCREMENT,
    `uid`          bigint(20)   DEFAULT 0  NOT NULL COMMENT '用户ID',
    `content_date` int(10)      DEFAULT 0  NOT NULL COMMENT '创建日报的时间，精确到日 20210609',
    `content_tag`  varchar(500) DEFAULT '' NOT NULL COMMENT '标签，存json数据',
    `content_text` text         NOT NULL COMMENT '文本内容',
    `content_type` tinyint(5)   DEFAULT 0  NOT NULL COMMENT '内容类型，枚举值：1、日报；2、周报',
    `deleted`      tinyint(4)   DEFAULT 0  NOT NULL COMMENT '删除状态：0、未删除，1、已删除',
    `c_time`       int(11)      DEFAULT 0  NOT NULL COMMENT '创建时间',
    `m_time`       int(11)      DEFAULT 0  NOT NULL COMMENT '更新',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_uid_cd_ct` (`uid`, `content_date`, `content_type`),
    KEY `idx_cd` (`content_date`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='内容表';

CREATE TABLE `user`
(
    `id`         bigint(20) UNSIGNED AUTO_INCREMENT COMMENT '主键ID',
    `uid`        bigint(20)   DEFAULT 0  NOT NULL COMMENT 'uid',
    `user_name`  varchar(255) DEFAULT '' NOT NULL COMMENT '用户名',
    `email`      varchar(255) DEFAULT '' NOT NULL COMMENT '邮件',
    `phone`      varchar(255) DEFAULT '' NOT NULL COMMENT '手机号',
    `avatar`     varchar(255) DEFAULT '' NOT NULL COMMENT '头像',
    `sn`         varchar(255) DEFAULT '' NOT NULL COMMENT '名',
    `given_name` varchar(255) DEFAULT '' NOT NULL COMMENT '姓',
    `deleted`    tinyint(4)   DEFAULT 0  NOT NULL COMMENT '删除状态：0、未删除，1、已删除',
    `c_time`     int(11)      DEFAULT 0  NOT NULL COMMENT '创建时间',
    `m_time`     int(11)      DEFAULT 0  NOT NULL COMMENT '更新',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_un` (`user_name`),
    KEY `idx_uid` (`uid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户表';
