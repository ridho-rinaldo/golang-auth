-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               10.1.32-MariaDB - mariadb.org binary distribution
-- Server OS:                    Win32
-- HeidiSQL Version:             12.0.0.6468
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Dumping database structure for auth_apps
CREATE DATABASE IF NOT EXISTS `auth_apps` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `auth_apps`;

-- Dumping structure for table auth_apps.users
CREATE TABLE IF NOT EXISTS `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` longtext,
  `email` varchar(191) DEFAULT NULL,
  `password` longblob,
  `phone` longtext,
  `gender` longtext,
  `address` longtext,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=latin1;

-- Dumping data for table auth_apps.users: ~5 rows (approximately)
INSERT INTO `users` (`id`, `username`, `email`, `password`, `phone`, `gender`, `address`) VALUES
	(1, 'Ridho Rinaldo', 'rinaldo@mail.com', _binary 0x243261243134244345454a6430506369746151336141444d384c586365705156766144654476396c4b68747376746d32485456527636767566513975, '08723123456', 'Laki - Laki', 'Jl. Mondatery No 12 Padukuhan Sujepolo'),
	(2, 'Memedd Sadana', 'sadna@mail.com', _binary 0x2432612431342456717a764d46704f6d4d4361394958526d385277632e782e78524a4a32465078465a774973763043493953774b6f4c33574859416d, '08992345679', 'Laki - Laki', 'Jl. Pati Sumber Harjo No 123'),
	(3, 'Hera Rotane', 'rotane@mail.com', _binary 0x24326124313424374e3241572f466977646f393069487676646d54384f4f6e5179584769556c7a784f425163642f47706e2f644a744a78514e596865, '08213245631', 'Perempuan', 'Jl. Rupa rupawan 31, Sarung Bantal'),
	(4, 'Giodan Suhena', 'suhegio@mail.com', _binary 0x2432612431342452626458584f4a4939513572394259374a443754654f5442796c68566c39754679552e38536c584a4c56474e667850555830364953, '08892315578', 'Perempuan', 'Jl. Siantar Suda Dekat'),
	(5, 'Yora Naslo', 'naslo@mail.com', _binary 0x24326124313424357141707763727650305a5864765970617374474f656d444a2e2f4275594b7a6374476c545445574d6138746c4d6b6f5563364b2e, '08120938145', 'Laki - Laki', 'Jl. Perbatas Sulaga');

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
