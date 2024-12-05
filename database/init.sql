-- 创建数据库
CREATE DATABASE IF NOT EXISTS blog;
CREATE DATABASE IF NOT EXISTS blog_test;

-- 创建用户并授权
CREATE USER IF NOT EXISTS 'blog_user'@'localhost' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON blog.* TO 'blog_user'@'localhost';
GRANT ALL PRIVILEGES ON blog_test.* TO 'blog_user'@'localhost';
FLUSH PRIVILEGES; 