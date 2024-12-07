-- create database
CREATE DATABASE IF NOT EXISTS blog;
CREATE DATABASE IF NOT EXISTS blog_test;

-- create user and grant privileges
CREATE USER IF NOT EXISTS 'blog_user'@'localhost' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON blog.* TO 'blog_user'@'localhost';
GRANT ALL PRIVILEGES ON blog_test.* TO 'blog_user'@'localhost';
FLUSH PRIVILEGES; 