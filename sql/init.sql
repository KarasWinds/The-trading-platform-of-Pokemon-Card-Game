DROP TABLE IF EXISTS `card`;
CREATE TABLE IF NOT EXISTS `card` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;

INSERT INTO `card` (`ID`, `name`) VALUES
	(1, 'Pikachu'),
	(2, 'Bulbasaur'),
	(3, 'Charmander'),
	(4, 'Squirtle');

DROP TABLE IF EXISTS `buy_orders`;
CREATE TABLE IF NOT EXISTS `buy_orders` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `trader` varchar(50) NOT NULL,
  `card_ID` int(11) NOT NULL,
  `price` float NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `sell_orders`;
CREATE TABLE IF NOT EXISTS `sell_orders` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `trader` varchar(50) NOT NULL,
  `card_ID` int(11) NOT NULL,
  `price` float NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `trades`;
CREATE TABLE IF NOT EXISTS `trades` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `buy` varchar(50) DEFAULT NULL,
  `sell` varchar(50) DEFAULT NULL,
  `card_ID` int(11) DEFAULT NULL,
  `price` float DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;