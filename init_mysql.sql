use mysql;
UPDATE user SET plugin='mysql_native_password' WHERE User='root';
FLUSH PRIVILEGES;
ALTER USER 'root'@'localhost' IDENTIFIED BY '123456';
-- 创建数据库
create database `tiktok` default character set utf8 collate utf8_general_ci;
