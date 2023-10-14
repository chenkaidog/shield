DROP TABLE IF EXISTS `account`;
DROP TABLE IF EXISTS `user`;
DROP TABLE IF EXISTS `login_record`;

CREATE TABLE `account`
(
    `id`         BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `account_id` VARCHAR(128)    NOT NULL COMMENT 'unique account id',
    `username`   VARCHAR(64)     NOT NULL COMMENT '用户帐号',
    `password`   VARCHAR(256)    NOT NULL COMMENT '密码md5',
    `salt`       VARCHAR(256)    NOT NULL COMMENT '盐',
    `status`     VARCHAR(32)     NOT NULL COMMENT '帐号状态',
    `created_at` DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE INDEX `uniq_account_id` (`account_id`),
    UNIQUE INDEX `uniq_username` (`username`),
    INDEX `idx_created_at` (`created_at`)
) ENGINE = INNODB
  DEFAULT CHARSET = utf8mb4
    COMMENT '用户帐号表';

CREATE TABLE `user`
(
    `id`          BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `user_id`     VARCHAR(128)    NOT NULL COMMENT '用户唯一ID',
    `account_id`  VARCHAR(128)    NOT NULL COMMENT 'unique account id'
    `name`        VARCHAR(64)     NOT NULL COMMENT '用户名称',
    `gender`      VARCHAR(16)     NOT NULL COMMENT '',
    `phone`       VARCHAR(32)     NOT NULL COMMENT '用户手机号码',
    `email`       VARCHAR(64)     NOT NULL COMMENT '用户邮箱',
    `description` VARCHAR(256)    NULL COMMENT '用户描述',
    `created_at`  DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE INDEX `uniq_user_id` (`user_id`),
    UNIQUE INDEX `uniq_account_id` (`account_id`),
    INDEX `idx_created_at` (`created_at`)
) ENGINE = INNODB
  DEFAULT CHARSET = utf8mb4
    COMMENT '用户信息表';

CREATE TABLE `login_record`
(
  `id`        BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `account_id`  VARCHAR(128)    NOT NULL COMMENT 'unique account id',
  `login_at`  DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '',
  `ipv4`      VARCHAR(16)     NOT NULL COMMENT '',
  `device`    VARCHAR(64)     NOT NULL COMMENT '',
  `reason`    VARCHAR(64)     NOT NULL COMMENT '',
  `status`    VARCHAR(32)     NOT NULL COMMENT '',
  PRIMARY KEY (`id`),
  INDEX `idx_username_login_at` (`account_id`, `login_at`)
) ENGINE = INNODB
  DEFAULT CHARSET = utf8mb4
    COMMENT '';