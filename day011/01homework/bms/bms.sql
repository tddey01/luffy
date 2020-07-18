-- ------------
-- 这是注释
-- BMS
-- -----------
CREATE TABLE  book(
    id bigint(11) AUTO_INCREMENT     PRIMARY KEY (id),
    title VARCHAR(255) NOT NULL ,
    price DOUBLE(16,2) NOT NULL ,
   publisher_id bigint(22) NOT NULL DEFAULT 0
)ENGINE=InnoDB DEFAULT  CHARSET = utf8mb4;


CREATE TABLE publisher (
    id BIGINT(22) AUTO_INCREMENT PRIMARY KEY ,
    province VARCHAR(255) NOT NULL ,
    city VARCHAR(255) NOT NULL ,
)ENGINE =InnoDB DEFAULT CHARSET=utf8mb4;
