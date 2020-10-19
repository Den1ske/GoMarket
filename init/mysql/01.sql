CREATE DATABASE IF NOT EXISTS `go_market`;
CREATE DATABASE IF NOT EXISTS `go_market_test`;

CREATE USER 'root'@'localhost' IDENTIFIED BY 'local';
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%';