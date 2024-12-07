-- create database
CREATE DATABASE IF NOT EXISTS blog;
CREATE DATABASE IF NOT EXISTS blog_test;

-- create user and grant privileges
SET @DB_USER = IFNULL(@DB_USER, 'blog_user');
SET @DB_PASSWORD = IFNULL(@DB_PASSWORD, '');

SET @create_user_query = CONCAT('CREATE USER IF NOT EXISTS ''', @DB_USER, '''@''localhost'' IDENTIFIED BY ''', @DB_PASSWORD, '''');
SET @grant_blog_query = CONCAT('GRANT ALL PRIVILEGES ON blog.* TO ''', @DB_USER, '''@''localhost''');
SET @grant_blog_test_query = CONCAT('GRANT ALL PRIVILEGES ON blog_test.* TO ''', @DB_USER, '''@''localhost''');

PREPARE stmt FROM @create_user_query;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;

PREPARE stmt FROM @grant_blog_query;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;

PREPARE stmt FROM @grant_blog_test_query;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;

FLUSH PRIVILEGES; 