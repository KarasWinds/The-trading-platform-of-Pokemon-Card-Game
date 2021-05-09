DROP DATABASE IF EXISTS `pokemon`;
CREATE DATABASE IF NOT EXISTS `pokemon`;
USE `pokemon`;

DROP TABLE IF EXISTS `cards`;
CREATE TABLE IF NOT EXISTS `cards` (
  `id` int(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(10),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `traders`;
CREATE TABLE IF NOT EXISTS `traders` (
  `id` int(20) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `orders`;
CREATE TABLE IF NOT EXISTS `orders` (
  `id` int(20) NOT NULL AUTO_INCREMENT,
  `price` float DEFAULT NULL,
  `type` varchar(10) DEFAULT NULL,
  `card_id` int(20) DEFAULT NULL,
  `trader_id` int(20) DEFAULT NULL,
  `status` TINYINT(1) NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_orders_card` (`card_id`),
  KEY `fk_orders_trader` (`trader_id`),
  CONSTRAINT `fk_orders_card` FOREIGN KEY (`card_id`) REFERENCES `cards` (`id`),
  CONSTRAINT `fk_orders_trader` FOREIGN KEY (`trader_id`) REFERENCES `traders` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `trades`;
CREATE TABLE IF NOT EXISTS `trades` (
  `id` int(20) NOT NULL AUTO_INCREMENT,
  `price` float DEFAULT NULL,
  `card_id` int(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_trades_card` (`card_id`),
  CONSTRAINT `fk_trades_card` FOREIGN KEY (`card_id`) REFERENCES `cards` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `cards` (`name`) VALUES
	('Pikachu'),
	('Bulbasaur'),
	('Charmander'),
	('Squirtle');
