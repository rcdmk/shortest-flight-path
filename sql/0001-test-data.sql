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

-- Copiando dados para a tabela routes.tb_airline: ~2 rows (aproximadamente)
/*!40000 ALTER TABLE `tb_airline` DISABLE KEYS */;
INSERT INTO `tb_airline` (`code`, `three_digit_code`, `name`, `country`, `active`) VALUES
	('AC', 'ACA', 'Air Canada', 'Canada', 1),
	('UA', 'UAL', 'United Airlines', 'United States', 1);
/*!40000 ALTER TABLE `tb_airline` ENABLE KEYS */;

-- Copiando dados para a tabela routes.tb_airport: ~5 rows (aproximadamente)
/*!40000 ALTER TABLE `tb_airport` DISABLE KEYS */;
INSERT INTO `tb_airport` (`iata3`, `name`, `city`, `country`, `latitude`, `longitude`, `active`) VALUES
	('JFK', 'John F Kennedy International Airport', 'New York', 'United States', 40.63980103, -73.77890015, 1),
	('LAX', 'Los Angeles International Airport', 'Los Angeles', 'United States', 33.94250107, -118.40799710, 1),
	('ORD', 'Chicago O\'Hare International Airport', 'Chicago', 'United States', 41.97859955, -87.90480042, 1),
	('YVR', 'Vancouver International Airport', 'Vancouver', 'Canada', 49.19390106, -123.18399810, 1),
	('YYZ', 'Lester B. Pearson International Airport', 'Toronto', 'Canada', 43.67720032, -79.63059998, 1);
/*!40000 ALTER TABLE `tb_airport` ENABLE KEYS */;

-- Copiando dados para a tabela routes.tb_route: ~0 rows (aproximadamente)
/*!40000 ALTER TABLE `tb_route` DISABLE KEYS */;
INSERT INTO `tb_route` (`airline_code`, `origin`, `destination`, `active`) VALUES
	('AC', 'JFK', 'YYZ', 1),
	('AC', 'LAX', 'YVR', 1),
	('AC', 'YVR', 'LAX', 1),
	('AC', 'YYZ', 'JFK', 1),
	('UA', 'JFK', 'LAX', 1),
	('UA', 'LAX', 'JFK', 1);
/*!40000 ALTER TABLE `tb_route` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
