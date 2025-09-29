-- 创建数据库
CREATE DATABASE IF NOT EXISTS melody_cure CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 使用数据库
USE melody_cure;

-- 创建用户并授权
CREATE USER IF NOT EXISTS 'melody_cure'@'%' IDENTIFIED BY '12345678';
GRANT ALL PRIVILEGES ON melody_cure.* TO 'melody_cure'@'%';
FLUSH PRIVILEGES;

-- 设置时区
SET time_zone = '+08:00';