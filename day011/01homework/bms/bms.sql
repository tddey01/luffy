-- ------------
-- 这是注释
-- BMS
-- -----------
CREATE TABLE book(
    id bigint(20) AUTO_INCREMENT PRIMARY KEY,
    title varchar(20) NOT NULL,
    price double(16,2) NOT NULL,
    publisher_id bigint(20) NOT NULL DEFAULT 0
)engine=InnoDB default charset=utf8mb4;

-- 出版社表
CREATE TABLE publisher(
    id bigint(20) AUTO_INCREMENT PRIMARY KEY,
    province varchar(20) NOT NULL,
    city varchar(20) NOT NULL
)engine=InnoDB default charset=utf8mb4;


-- SELECT book.id,book.price,book.title,publisher.province, publisher.city FROM book JOIN publisher ON book.publisher_id = publisher.id WHERE book.id=1;
