DROP TABLE IF EXISTS `user_role`;
DROP TABLE IF EXISTS `role`;
DROP TABLE IF EXISTS `permission`;
DROP TABLE IF EXISTS `operation_record`;

CREATE TABLE `user_role` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `relation_id` VARCHAR(128) NOT NULL COMMENT '',
  `user_id` VARCHAR(128) NOT NULL COMMENT '用户唯一ID',
  `role_id` VARCHAR(128) NOT NULL COMMENT '角色唯一ID',
  `status` VARCHAR(32) NOT NULL COMMENT '映射状态',
  `expire_at` DATETIME NOT NULL COMMENT '过期时间',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` BIGINT UNSIGNED NOT NULL COMMENT 'soft deleted at',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `uniq_relation_id` (`relation_id`),
  UNIQUE INDEX `uniq_user_role` (`user_id`, `role_id`, `deleted_at`),
  INDEX `idx_role_id` (`role_id`, `deleted_at`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT '用户-角色映射表';

CREATE TABLE `role` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `role_id` VARCHAR(128) NOT NULL COMMENT '角色唯一ID',
  `domain` VARCHAR(64) NOT NULL COMMENT '',
  `owner_id` VARCHAR(128) NOT NULL COMMENT '',
  `manager_id_list` TEXT NULL COMMENT '',
  `parent_role_id` VARCHAR(128) NOT NULL COMMENT '父角色ID',
  `maximum_applicant` INT NOT NULL COMMENT '最多申请人数,用negative表示没有限制',
  `max_valid_day` INT NOT NULL COMMENT '最多有效天数,用negative表示没有限制',
  `name` VARCHAR(64) NOT NULL COMMENT '角色名称',
  `level` VARCHAR(64) NOT NULL COMMENT '角色等级',
  `status` VARCHAR(32) NOT NULL COMMENT '角色状态',
  `description` VARCHAR(256) NULL COMMENT '角色描述',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` BIGINT UNSIGNED NOT NULL COMMENT 'soft deleted at',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `uniq_role_id` (`role_id`),
  INDEX `idx_parent_role_id` (`parent_role_id`, `deleted_at`),
  INDEX `idx_domain` (`domain`, `deleted_at`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT '角色表';

CREATE TABLE `permission` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `permission_id` VARCHAR(128) NOT NULL COMMENT '权限ID',
  `role_id` VARCHAR(128) NOT NULL COMMENT '角色唯一ID',
  `resource` VARCHAR(128) NOT NULL COMMENT '资源名称',
  `action` VARCHAR(32) NOT NULL COMMENT '权限操作',
  `status` VARCHAR(32) NOT NULL COMMENT 'allow/deny/deleted',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` BIGINT UNSIGNED NOT NULL COMMENT 'soft deleted at',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `uniq_permission_id` (`permission_id`),
  UNIQUE INDEX `uniq_role_id` (`role_id`, `deleted_at`, `resource`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT '权限表';

CREATE TABLE `operation_record` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `record_id` VARCHAR(128) NOT NULL COMMENT '',
  `data_type` VARCHAR(32) NOT NULL COMMENT '',
  `data_id` VARCHAR(128) NOT NULL COMMENT '',
  `operation` VARCHAR(32) NOT NULL COMMENT '',
  `operator_id` VARCHAR(128) NOT NULL COMMENT '',
  `trace_id` VARCHAR(128) NOT NULL COMMENT '',
  `previous_value` TEXT NULL COMMENT '',
  `current_value` TEXT NULL COMMENT '',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `uniq_record_id` (`record_id`),
  INDEX `idx_data_id_created_at` (`data_id`, `created_at`),
  INDEX `idx_operator_id_created_at` (`operator_id`, `created_at`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT '权限表';