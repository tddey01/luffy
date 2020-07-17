-- ------------
-- 这是注释
-- BMS
-- -----------
CREATE TABLE  book(
    id bigint(11) AUTO_INCREMENT ,
    title VARCHAR(255) NOT NULL ,
    price DOUBLE(16,2) NOT NULL ,
 PRIMARY KEY (id)
)ENGINE=InnoDB DEFAULT  CHARSET = utf8mb4;


