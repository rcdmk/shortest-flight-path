-- --------------------------------------------------------
-- Servidor:                     127.0.0.1
-- Versão do servidor:           5.5.45 - MySQL Community Server (GPL)
-- OS do Servidor:               Win64
-- HeidiSQL Versão:              9.4.0.5125
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- Copiando estrutura para tabela routes.tb_airline
CREATE TABLE IF NOT EXISTS `tb_airline` (
  `code` varchar(2) NOT NULL,
  `three_digit_code` varchar(3) NOT NULL,
  `name` varchar(30) NOT NULL,
  `country` varchar(50) NOT NULL,
  `active` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`code`),
  UNIQUE KEY `three_digit_code` (`three_digit_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Exportação de dados foi desmarcado.
-- Copiando estrutura para tabela routes.tb_airport
CREATE TABLE IF NOT EXISTS `tb_airport` (
  `iata3` varchar(3) NOT NULL,
  `name` varchar(75) NOT NULL,
  `city` varchar(50) NOT NULL,
  `country` varchar(50) NOT NULL,
  `latitude` decimal(11,9) NOT NULL,
  `longitude` decimal(12,9) NOT NULL,
  `active` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`iata3`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Exportação de dados foi desmarcado.
-- Copiando estrutura para tabela routes.tb_route
CREATE TABLE IF NOT EXISTS `tb_route` (
  `airline_code` varchar(2) NOT NULL,
  `origin` varchar(3) NOT NULL,
  `destination` varchar(3) NOT NULL,
  `active` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`airline_code`,`origin`,`destination`),
  KEY `FK_tb_route_tb_airport` (`origin`),
  KEY `FK_tb_route_tb_airport_2` (`destination`),
  CONSTRAINT `FK_tb_route_tb_airline` FOREIGN KEY (`airline_code`) REFERENCES `tb_airline` (`code`) ON DELETE CASCADE,
  CONSTRAINT `FK_tb_route_tb_airport` FOREIGN KEY (`origin`) REFERENCES `tb_airport` (`iata3`) ON DELETE CASCADE,
  CONSTRAINT `FK_tb_route_tb_airport_2` FOREIGN KEY (`destination`) REFERENCES `tb_airport` (`iata3`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Exportação de dados foi desmarcado.
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
