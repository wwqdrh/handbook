-- upgrade --
CREATE TABLE IF NOT EXISTS `aerich` (
    `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `version` VARCHAR(255) NOT NULL,
    `app` VARCHAR(20) NOT NULL,
    `content` JSON NOT NULL
) CHARACTER SET utf8mb4;
CREATE TABLE IF NOT EXISTS `admin` (
    `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `username` VARCHAR(255) NOT NULL UNIQUE,
    `password` VARCHAR(255) NOT NULL,
    `nickname` VARCHAR(255) NOT NULL,
    `mobile` VARCHAR(20) NOT NULL,
    `is_used` SMALLINT NOT NULL  COMMENT 'nouse: 0\nisuse: 1',
    `is_deleted` SMALLINT NOT NULL  COMMENT 'nodelete: 0\nisdelete: 1',
    `created_at` DATETIME(6) NOT NULL  DEFAULT CURRENT_TIMESTAMP(6),
    `created_user` VARCHAR(255) NOT NULL,
    `updated_at` DATETIME(6) NOT NULL  DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    `updated_user` VARCHAR(255) NOT NULL
) CHARACTER SET utf8mb4;
CREATE TABLE IF NOT EXISTS `admin_menu` (
    `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `admin_id` INT NOT NULL,
    `menu_id` INT NOT NULL,
    `created_at` DATETIME(6) NOT NULL  DEFAULT CURRENT_TIMESTAMP(6),
    `created_user` VARCHAR(255) NOT NULL,
    KEY `idx_admin_menu_admin_i_862f05` (`admin_id`)
) CHARACTER SET utf8mb4;
CREATE TABLE IF NOT EXISTS `authorized` (
    `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `business_key` VARCHAR(32) NOT NULL,
    `business_secret` VARCHAR(60) NOT NULL,
    `business_developer` VARCHAR(60) NOT NULL,
    `remark` VARCHAR(255) NOT NULL,
    `is_used` SMALLINT NOT NULL  COMMENT 'nouse: 0\nisuse: 1',
    `is_deleted` SMALLINT NOT NULL  COMMENT 'nodelete: 0\nisdelete: 1',
    `created_at` DATETIME(6) NOT NULL  DEFAULT CURRENT_TIMESTAMP(6),
    `created_user` VARCHAR(255) NOT NULL,
    `updated_at` DATETIME(6) NOT NULL  DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    `updated_user` VARCHAR(255) NOT NULL
) CHARACTER SET utf8mb4;
CREATE TABLE IF NOT EXISTS `authorized_api` (
    `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `business_key` VARCHAR(32) NOT NULL,
    `method` VARCHAR(30) NOT NULL,
    `api` VARCHAR(100) NOT NULL,
    `is_deleted` SMALLINT NOT NULL  COMMENT 'nodelete: 0\nisdelete: 1',
    `created_at` DATETIME(6) NOT NULL  DEFAULT CURRENT_TIMESTAMP(6),
    `created_user` VARCHAR(255) NOT NULL,
    `updated_at` DATETIME(6) NOT NULL  DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    `updated_user` VARCHAR(255) NOT NULL
) CHARACTER SET utf8mb4;
CREATE TABLE IF NOT EXISTS `menu` (
    `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `pid` INT NOT NULL,
    `name` VARCHAR(32) NOT NULL,
    `link` VARCHAR(100) NOT NULL,
    `icon` VARCHAR(60) NOT NULL,
    `level` SMALLINT NOT NULL,
    `sort` SMALLINT NOT NULL,
    `is_used` SMALLINT NOT NULL  COMMENT 'nouse: 0\nisuse: 1',
    `is_deleted` SMALLINT NOT NULL  COMMENT 'nodelete: 0\nisdelete: 1',
    `created_at` DATETIME(6) NOT NULL  DEFAULT CURRENT_TIMESTAMP(6),
    `created_user` VARCHAR(255) NOT NULL,
    `updated_at` DATETIME(6) NOT NULL  DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    `updated_user` VARCHAR(255) NOT NULL
) CHARACTER SET utf8mb4;
CREATE TABLE IF NOT EXISTS `menu_action` (
    `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `menu_id` INT NOT NULL,
    `method` VARCHAR(30) NOT NULL,
    `api` VARCHAR(100) NOT NULL,
    `is_deleted` SMALLINT NOT NULL  COMMENT 'nodelete: 0\nisdelete: 1',
    `created_at` DATETIME(6) NOT NULL  DEFAULT CURRENT_TIMESTAMP(6),
    `created_user` VARCHAR(255) NOT NULL,
    `updated_at` DATETIME(6) NOT NULL  DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    `updated_user` VARCHAR(255) NOT NULL
) CHARACTER SET utf8mb4;
