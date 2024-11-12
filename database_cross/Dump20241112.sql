-- MySQL dump 10.13  Distrib 8.0.38, for Win64 (x86_64)
--
-- Host: localhost    Database: crossborderpayment
-- ------------------------------------------------------
-- Server version	8.0.39

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `account`
--

DROP TABLE IF EXISTS `account`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `account` (
  `AccountID` int NOT NULL AUTO_INCREMENT,
  `UserID` int DEFAULT NULL,
  `EthereumAddress` varchar(255) DEFAULT NULL,
  `AccountType` varchar(50) DEFAULT NULL,
  `AccountStatus` varchar(50) DEFAULT 'Active',
  `CreatedAt` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `UpdatedAt` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `balance` decimal(18,2) DEFAULT '0.00',
  PRIMARY KEY (`AccountID`),
  KEY `fk_user` (`UserID`),
  CONSTRAINT `account_ibfk_1` FOREIGN KEY (`UserID`) REFERENCES `users` (`UserID`),
  CONSTRAINT `fk_user` FOREIGN KEY (`UserID`) REFERENCES `users` (`UserID`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `account`
--

LOCK TABLES `account` WRITE;
/*!40000 ALTER TABLE `account` DISABLE KEYS */;
INSERT INTO `account` VALUES (1,2,'0xDef456...','Receiver','Active','2024-10-05 06:43:05','2024-10-05 06:43:05',0.00),(2,26,'0x1bDf7Cacf577A52EC35c507f80a06607e604a255','Receiver','Active','2024-11-07 06:38:55','2024-11-07 06:38:55',0.00),(3,27,'0x48382B681bFAC0daC45BA73E0171189A8A497917','Receiver','Active','2024-11-11 09:20:32','2024-11-11 09:20:32',0.00);
/*!40000 ALTER TABLE `account` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transaction`
--

DROP TABLE IF EXISTS `transaction`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `transaction` (
  `TxID` varchar(255) NOT NULL,
  `SenderAccount` varchar(255) DEFAULT NULL,
  `ReceiverAccount` varchar(255) DEFAULT NULL,
  `Amount` decimal(10,2) DEFAULT NULL,
  `CurrencyType` varchar(3) DEFAULT NULL,
  `Timestamp` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `TransactionStatus` varchar(50) DEFAULT NULL,
  `Remarks` text,
  PRIMARY KEY (`TxID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transaction`
--

LOCK TABLES `transaction` WRITE;
/*!40000 ALTER TABLE `transaction` DISABLE KEYS */;
/*!40000 ALTER TABLE `transaction` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transactionrecord`
--

DROP TABLE IF EXISTS `transactionrecord`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `transactionrecord` (
  `TxID` varchar(255) NOT NULL,
  `SenderAccount` int DEFAULT NULL,
  `ReceiverAccount` int DEFAULT NULL,
  `Amount` decimal(18,8) DEFAULT NULL,
  `CurrencyType` varchar(10) DEFAULT NULL,
  `CurrencyUnit` varchar(10) DEFAULT NULL,
  `TxHash` varchar(255) DEFAULT NULL,
  `BlockHeight` int DEFAULT NULL,
  `Note` text,
  `Status` varchar(50) DEFAULT 'Pending',
  `Timestamp` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `CreatedAt` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`TxID`),
  KEY `fk_sender_account` (`SenderAccount`),
  KEY `fk_receiver_account` (`ReceiverAccount`),
  CONSTRAINT `fk_receiver_account` FOREIGN KEY (`ReceiverAccount`) REFERENCES `account` (`AccountID`),
  CONSTRAINT `fk_sender_account` FOREIGN KEY (`SenderAccount`) REFERENCES `account` (`AccountID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transactionrecord`
--

LOCK TABLES `transactionrecord` WRITE;
/*!40000 ALTER TABLE `transactionrecord` DISABLE KEYS */;
/*!40000 ALTER TABLE `transactionrecord` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transactionstatus`
--

DROP TABLE IF EXISTS `transactionstatus`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `transactionstatus` (
  `TxID` varchar(255) DEFAULT NULL,
  `Status` varchar(50) DEFAULT NULL,
  `Timestamp` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  KEY `TxID` (`TxID`),
  CONSTRAINT `transactionstatus_ibfk_1` FOREIGN KEY (`TxID`) REFERENCES `transaction` (`TxID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transactionstatus`
--

LOCK TABLES `transactionstatus` WRITE;
/*!40000 ALTER TABLE `transactionstatus` DISABLE KEYS */;
/*!40000 ALTER TABLE `transactionstatus` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `UserID` int NOT NULL AUTO_INCREMENT,
  `UserName` varchar(255) NOT NULL,
  `password_hash` varchar(255) NOT NULL,
  `Email` varchar(255) DEFAULT NULL,
  `CreatedAt` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `UpdatedAt` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `is_verified` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`UserID`),
  UNIQUE KEY `Email` (`Email`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'Alice','',NULL,'2024-10-05 06:39:29','2024-10-05 06:39:29','2024-10-05 08:55:47','2024-10-05 08:55:52',0),(2,'Bob','',NULL,'2024-10-05 06:39:29','2024-10-05 06:39:29','2024-10-05 08:55:47','2024-10-05 08:55:52',0),(4,'testuser','$2a$10$SBTpmIbtwDg925QDe4K6weofoLJLQGDrpKLSNVNh4dRjYQ6YpZXIW','testuser@example.com','2024-10-06 09:56:17','2024-10-06 09:56:17','2024-10-06 09:56:17','2024-10-06 09:56:17',0),(6,'newuser','$2a$10$Dxa6.rjcaOhkOONeCcfky.K6Ze6j4isp0nyiSlotU6l6VhUyX17mm','newuser@example.com','2024-10-06 10:28:51','2024-10-06 10:28:51','2024-10-06 10:28:52','2024-10-06 10:28:52',0),(7,'marjorie','$2a$10$U3x6q3dZyfX62aOFTCxeM.46US5FCz28UwqJJESIXNOxRbrg.WgPi','marjorie@example.com','2024-10-11 08:35:15','2024-10-11 08:35:15','2024-10-11 08:35:15','2024-10-11 08:35:15',0),(10,'mai','$2a$10$1iqWE7NNFKuO0WvJe9XgFO2bnh2rrVHQZRDYoSQ8uktA8vJzYsrFe','mai@example.com','2024-10-11 08:54:08','2024-10-11 08:54:08','2024-10-11 08:54:08','2024-10-11 08:54:08',0),(11,'marjoriemai','$2a$10$yJ49MOh3RoXc4uXHiYbf7.1jDfFr1pfH1ldtX92arEe6yh6VWJ01K','marjoriemai@example.com','2024-10-11 09:02:23','2024-10-11 09:02:23','2024-10-11 09:02:24','2024-10-11 09:02:24',0),(12,'test2','$2a$10$lt.pBlAYtJKRCY008AfwdOg1p03Bz8tMwYrws9x13hoRFa0fcGgfW','test2@example.com','2024-10-11 09:13:22','2024-10-11 09:13:22','2024-10-11 09:13:22','2024-10-11 09:13:22',0),(13,'sugar','$2a$10$3mB1IBsDFtmhAvi2nvCF/eJtK75dvzeuF4tYPAqzcbbVVr1VYsMBq','sugar11@example.com','2024-10-14 05:30:03','2024-10-14 05:30:03','2024-10-14 05:30:04','2024-10-14 05:30:04',0),(14,'test3','$2a$10$jdyfTSJ2BLd3d89ViJVIt.SVvJ83tSPZCH8rM2QukurlGJn1e3jgW','test3@example.com','2024-10-14 06:27:28','2024-10-14 06:27:28','2024-10-14 06:27:29','2024-10-14 06:27:29',0),(15,'test4','$2a$10$FWfO6DQtJPmuUA4/g0lrIuS4oy7ftnVrBBKw0sSukYP62OEsNBhAq','test4@example.com','2024-10-14 13:00:52','2024-10-14 13:00:52','2024-10-14 13:00:53','2024-10-14 13:00:53',0),(16,'test5','$2a$10$FmDfWmp935S2Hjw6naL7NuZDl0C4VRWb4a93idbj9/DAMChcFw27q','test5@example.com','2024-10-14 13:39:04','2024-10-14 13:39:04','2024-10-14 13:39:05','2024-10-14 13:39:05',0),(17,'test8','$2a$10$oW0SgQ2TxseXYOZkKR.qsODknwH8ObVZy8AsqlThEVBA0HVG2xhLG','test8@example.com','2024-10-15 05:52:14','2024-10-15 06:46:58','2024-10-15 06:46:59','2024-10-15 05:52:14',0),(20,'test9','$2a$10$eRTb6VKgtjMfuQrqk2LLXeST4b11JjfuuDCzdZf.2CJqOf2SgMs7q','test9@example.com','2024-10-15 07:26:08','2024-10-15 11:24:06','2024-10-15 11:24:07','2024-10-15 07:26:09',0),(21,'10test','$2a$10$mqb6SH.oHyDNr1dxprdXDOJMoyD.ZYX9nmnGWTxj3GI4LeqB4fsOC','test10@example.com','2024-10-15 11:33:30','2024-10-15 11:34:04','2024-10-15 11:34:05','2024-10-15 11:33:30',0),(22,'test1','$2a$10$ZE8J0QyNsHPOKGKqzhAu2.yE9TEHir63/XqNA8owYaUa.5rTfm/8y','test1@example.com','2024-10-18 07:31:52','2024-10-18 07:31:52','2024-10-18 07:31:53','2024-10-18 07:31:53',0),(23,'check','$2a$10$gN3ycKog.dDzG3Tzn85/3e9LcCa/MPIFANj3TV8n8hSn2SUbY.g/6','check@example','2024-10-18 13:16:55','2024-10-18 13:16:55','2024-10-18 13:16:56','2024-10-18 13:16:56',0),(24,'user1','$2a$10$F6GTsVcjdzl.PBS7aIQyVeTvy1cEo2ZpRjzyc7IzYXqCdHAnrNXvq','user1@example.com','2024-10-18 13:31:27','2024-10-18 13:31:27','2024-10-18 13:31:28','2024-10-18 13:31:28',0),(25,'check1','$2a$10$OKSGTXyAvbKdb4ZbyKY3EuYnpMRUSxr6ap0YzxN/Hid3pRb4K9upy','check1@example.com','2024-10-19 07:25:19','2024-10-19 14:04:38','2024-10-19 14:04:38','2024-10-19 07:25:19',0),(26,'maim','$2a$10$W.rDdXVbRwpZlYeOkpiWw.JLzbHIdsAefuIN0QMcCqG5gZY2kS6.C','mai1@example.com','2024-11-07 06:38:54','2024-11-07 06:39:57','2024-11-07 06:39:57','2024-11-07 06:38:55',0),(27,'testt','$2a$10$g7V3Ufamu9dCwJhNbhUouOuc0hRHWOXUQ/9JP31lUJX9DF5/IjBWy','test@test.com','2024-11-11 09:20:32','2024-11-11 09:40:19','2024-11-11 09:40:19','2024-11-11 09:20:32',0);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-11-12 14:01:57
