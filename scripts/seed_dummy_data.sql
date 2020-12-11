-- Generation time: Fri, 11 Dec 2020 12:47:22 +0000
-- Host: mysql.hostinger.ro
-- DB name: u574849695_21
/*!40030 SET NAMES UTF8 */;
/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

DROP TABLE IF EXISTS `coupon`;
CREATE TABLE `coupon` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `code` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `expiry_date` datetime DEFAULT NULL,
  `availability_count` int(11) DEFAULT NULL,
  `product_id` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `promo_type` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `discount_fixed` decimal(10,0) DEFAULT NULL,
  `discount_variable` decimal(10,0) DEFAULT NULL,
  `valid` tinyint(1) DEFAULT NULL,
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `code` (`code`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

INSERT INTO `coupon` VALUES ('1','edkf','1995-12-22 07:27:59','212','qshy','dqfa','4768473','93','0','1979-11-03 06:48:18','2012-05-21 18:12:42'),
('2','xpog','2001-12-28 09:17:29','8240','slfo','hfgg','952','8','1','1999-02-27 18:53:55','2014-07-05 20:22:30'),
('3','ifns','2001-03-29 01:21:43','7','bmln','emkd','981639','68972506','1','1998-07-10 11:41:19','1983-07-28 16:06:55'),
('4','cpua','1988-07-14 22:06:27','7','vatg','cidd','661','1413604','0','2017-05-26 21:14:40','2018-10-12 08:42:42'),
('5','nlwn','2018-05-29 07:58:44','8479','dkga','vbhb','659999449','82992','1','1988-09-14 03:49:30','1998-02-03 08:06:57'),
('6','uhzw','2014-10-27 01:40:55','0','qevh','hxgm','985','19625359','0','2000-10-20 02:41:10','1998-10-20 22:29:52'),
('7','rthp','1975-08-30 04:09:19','0','zqoo','krrb','61144','32','0','1993-09-23 18:11:45','2019-05-03 19:15:18'),
('8','mfrm','2001-08-03 18:17:41','742831422','frkq','aaxo','4','6186352','1','1978-03-01 18:56:53','2009-10-25 03:41:55'),
('9','zvdu','1989-09-13 08:24:20','505','odrk','lfcl','70607','828789','1','2002-03-10 19:40:41','2014-06-13 23:14:13'),
('10','loif','1970-02-16 02:11:43','0','itcm','vbzg','0','35','1','1994-06-10 11:01:37','1989-10-29 02:12:12'); 




/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

