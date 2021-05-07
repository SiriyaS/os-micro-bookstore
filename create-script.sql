-- DBMS: MySQL

-- category
CREATE TABLE `category` (
  `id` int(2) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4;

-- authors
CREATE TABLE `authors` (
  `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(70) NOT NULL,
  `email` varchar(30) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- publishers
CREATE TABLE `publishers` (
  `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(70) NOT NULL,
  `address` varchar(50) NOT NULL,
  `telephone` varchar(20) DEFAULT NULL,
  `website` varchar(50) NOT NULL,
  `email` varchar(30) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- books
CREATE TABLE `books` (
  `ISBN` varchar(13) NOT NULL,
  `name` varchar(100) DEFAULT NULL,
  `author` int(6) unsigned DEFAULT NULL,
  `unit_price` decimal(6,2) unsigned DEFAULT NULL,
  `publish_year` int(4) NOT NULL,
  `publisher` int(6) unsigned DEFAULT NULL,
  `edition` int(3) NOT NULL,
  `category` int(2) unsigned DEFAULT NULL,
  PRIMARY KEY (`ISBN`),
  KEY `author` (`author`),
  KEY `category` (`category`),
  KEY `publisher` (`publisher`),
  CONSTRAINT `books_ibfk_1` FOREIGN KEY (`author`) REFERENCES `authors` (`id`),
  CONSTRAINT `books_ibfk_2` FOREIGN KEY (`category`) REFERENCES `category` (`id`),
  CONSTRAINT `books_ibfk_3` FOREIGN KEY (`publisher`) REFERENCES `publishers` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- users
CREATE TABLE `users` (
  `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(70) NOT NULL,
  `email` varchar(30) NOT NULL,
  `address` varchar(50) NOT NULL,
  `telephone` varchar(10) NOT NULL,
  `username` varchar(30) NOT NULL,
  `password` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- order_main
CREATE TABLE `order_main` (
  `order_no` varchar(10) NOT NULL,
  `user` int(6) unsigned DEFAULT NULL,
  `order_date` datetime DEFAULT NULL,
  `grand_total` decimal(8,2) unsigned DEFAULT NULL,
  PRIMARY KEY (`order_no`),
  KEY `user` (`user`),
  CONSTRAINT `order_main_ibfk_1` FOREIGN KEY (`user`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- order_detail
CREATE TABLE `order_detail` (
  `order_no` varchar(10) NOT NULL,
  `order_seq` int(3) unsigned NOT NULL,
  `book_isbn` varchar(13) NOT NULL,
  `quantity` int(3) unsigned DEFAULT NULL,
  `unit_price` decimal(6,2) unsigned DEFAULT NULL,
  `total` decimal(8,2) unsigned DEFAULT NULL,
  KEY `order_no` (`order_no`),
  KEY `book_isbn` (`book_isbn`),
  CONSTRAINT `order_detail_ibfk_1` FOREIGN KEY (`order_no`) REFERENCES `order_main` (`order_no`),
  CONSTRAINT `order_detail_ibfk_2` FOREIGN KEY (`book_isbn`) REFERENCES `books` (`ISBN`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;