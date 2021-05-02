CREATE DATABASE IF NOT EXISTS `pokemon`;
USE `pokemon`;

CREATE TABLE IF NOT EXISTS `order_buy` (
  `buy_ID` int(50) unsigned NOT NULL AUTO_INCREMENT,
  `trader_ID` smallint(5) unsigned NOT NULL,
  `card_ID` tinyint(3) unsigned NOT NULL,
  `price` float unsigned NOT NULL,
  `time` datetime NOT NULL,
  PRIMARY KEY (`buy_ID`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `order_sell` (
  `sell_ID` int(50) unsigned NOT NULL AUTO_INCREMENT,
  `trader_ID` smallint(5) unsigned NOT NULL,
  `card_ID` tinyint(3) unsigned NOT NULL,
  `price` float unsigned NOT NULL,
  `time` datetime NOT NULL,
  PRIMARY KEY (`sell_ID`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `trades` (
  `trade_ID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `buy_ID` int(50) unsigned NOT NULL,
  `sell_ID` int(50) unsigned NOT NULL,
  `price` float unsigned NOT NULL,
  `card_id` tinyint(3) unsigned NOT NULL,
  `time` datetime NOT NULL,
  PRIMARY KEY (`trade_ID`) USING BTREE,
  KEY `FK_trades_order_buy` (`buy_ID`),
  KEY `FK_trades_order_sell` (`sell_ID`),
  CONSTRAINT `FK_trades_order_buy` FOREIGN KEY (`buy_ID`) REFERENCES `order_buy` (`buy_ID`),
  CONSTRAINT `FK_trades_order_sell` FOREIGN KEY (`sell_ID`) REFERENCES `order_sell` (`sell_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
