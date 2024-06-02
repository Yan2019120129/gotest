-- MySQL dump 10.13  Distrib 8.0.36, for Linux (x86_64)
--
-- Host: localhost    Database: my_admin
-- ------------------------------------------------------
-- Server version	8.0.36-0ubuntu0.23.10.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `goadmin_menu`
--

DROP TABLE IF EXISTS `goadmin_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `goadmin_menu` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int unsigned NOT NULL DEFAULT '0',
  `type` tinyint unsigned NOT NULL DEFAULT '0',
  `order` int unsigned NOT NULL DEFAULT '0',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `icon` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `uri` varchar(3000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `header` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `plugin_name` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `uuid` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_menu`
--

LOCK TABLES `goadmin_menu` WRITE;
/*!40000 ALTER TABLE `goadmin_menu` DISABLE KEYS */;
INSERT INTO `goadmin_menu` VALUES (1,0,1,2,'Admin','fa-tasks','',NULL,'',NULL,'2019-09-09 16:00:00','2019-09-09 16:00:00'),(2,1,1,2,'Users','fa-users','/info/manager',NULL,'',NULL,'2019-09-09 16:00:00','2019-09-09 16:00:00'),(3,1,1,3,'Roles','fa-user','/info/roles',NULL,'',NULL,'2019-09-09 16:00:00','2019-09-09 16:00:00'),(4,1,1,4,'Permission','fa-ban','/info/permission',NULL,'',NULL,'2019-09-09 16:00:00','2019-09-09 16:00:00'),(5,1,1,5,'Menu','fa-bars','/menu',NULL,'',NULL,'2019-09-09 16:00:00','2019-09-09 16:00:00'),(6,1,1,6,'Operation log','fa-history','/info/op',NULL,'',NULL,'2019-09-09 16:00:00','2019-09-09 16:00:00'),(7,0,1,1,'Dashboard','fa-bar-chart','/',NULL,'',NULL,'2019-09-09 16:00:00','2019-09-09 16:00:00'),(8,0,0,2,'用户','fa-user','/info/user','http://localhost:4100','',NULL,'2024-03-11 16:27:21','2024-03-11 16:27:38');
/*!40000 ALTER TABLE `goadmin_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `goadmin_operation_log`
--

DROP TABLE IF EXISTS `goadmin_operation_log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `goadmin_operation_log` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int unsigned NOT NULL,
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `input` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `admin_operation_log_user_id_index` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=70 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_operation_log`
--

LOCK TABLES `goadmin_operation_log` WRITE;
/*!40000 ALTER TABLE `goadmin_operation_log` DISABLE KEYS */;
INSERT INTO `goadmin_operation_log` VALUES (1,1,'/admin','GET','::1','','2024-03-11 16:15:42','2024-03-11 16:15:42'),(2,1,'/admin/info/generate/new','GET','::1','','2024-03-11 16:15:48','2024-03-11 16:15:48'),(3,1,'/admin/operation/_tool_choose_conn','POST','::1','','2024-03-11 16:15:52','2024-03-11 16:15:52'),(4,1,'/admin/operation/_tool_choose_table','POST','::1','','2024-03-11 16:18:40','2024-03-11 16:18:40'),(5,1,'/admin/info/generate/new','GET','::1','','2024-03-11 16:18:49','2024-03-11 16:18:49'),(6,1,'/admin/operation/_tool_choose_conn','POST','::1','','2024-03-11 16:18:51','2024-03-11 16:18:51'),(7,1,'/admin/operation/_tool_choose_table','POST','::1','','2024-03-11 16:18:54','2024-03-11 16:18:54'),(8,1,'/admin/new/generate','POST','::1','{\"__checkbox__hide_back_button\":[\"on\"],\"__checkbox__hide_continue_edit_check_box\":[\"on\"],\"__checkbox__hide_continue_new_check_box\":[\"on\"],\"__checkbox__hide_delete_button\":[\"on\"],\"__checkbox__hide_detail_button\":[\"on\"],\"__checkbox__hide_edit_button\":[\"on\"],\"__checkbox__hide_export_button\":[\"on\"],\"__checkbox__hide_filter_area\":[\"on\"],\"__checkbox__hide_filter_button\":[\"on\"],\"__checkbox__hide_new_button\":[\"on\"],\"__checkbox__hide_pagination\":[\"on\"],\"__checkbox__hide_query_info\":[\"on\"],\"__checkbox__hide_reset_button\":[\"on\"],\"__checkbox__hide_row_selector\":[\"on\"],\"__go_admin_previous_\":[\"http://localhost:4100/admin/login\"],\"__go_admin_t_\":[\"12040b53-08c2-4ef4-9835-66508afdc66e\"],\"conn\":[\"default\"],\"detail_description\":[\"\"],\"detail_display\":[\"0\"],\"detail_title\":[\"\"],\"extra_code\":[\"\"],\"extra_import_package[]\":[\"github.com/GoAdminGroup/go-admin/modules/db\"],\"field_canadd\":[\"n\",\"y\",\"y\",\"y\",\"y\",\"y\",\"y\",\"y\",\"y\"],\"field_canedit\":[\"y\",\"y\",\"y\",\"y\",\"y\",\"y\",\"y\",\"y\",\"y\"],\"field_db_type\":[\"Int\",\"Int\",\"Varchar\",\"Varchar\",\"Varchar\",\"Tinyint\",\"Tinyint\",\"Varchar\",\"Int\"],\"field_db_type_form\":[\"Int\",\"Int\",\"Varchar\",\"Varchar\",\"Varchar\",\"Tinyint\",\"Tinyint\",\"Varchar\",\"Int\"],\"field_default\":[\"\",\"\",\"\",\"\",\"\",\"\",\"\",\"\",\"\"],\"field_display\":[\"\",\"\",\"\",\"\",\"\",\"\",\"\",\"\",\"1\"],\"field_filterable\":[\"n\",\"n\",\"n\",\"n\",\"n\",\"n\",\"n\",\"n\",\"n\"],\"field_form_type_form\":[\"Default\",\"Number\",\"Text\",\"Text\",\"Text\",\"Number\",\"Number\",\"Text\",\"Number\"],\"field_head\":[\"Id\",\"Admin_id\",\"Name\",\"Alias\",\"Icon\",\"Sort\",\"Status\",\"Data\",\"Created_at\"],\"field_head_form\":[\"Id\",\"Admin_id\",\"Name\",\"Alias\",\"Icon\",\"Sort\",\"Status\",\"Data\",\"Created_at\"],\"field_hide\":[\"n\",\"n\",\"n\",\"n\",\"n\",\"n\",\"n\",\"n\",\"n\"],\"field_name\":[\"id\",\"admin_id\",\"name\",\"alias\",\"icon\",\"sort\",\"status\",\"data\",\"created_at\"],\"field_name_form\":[\"id\",\"admin_id\",\"name\",\"alias\",\"icon\",\"sort\",\"status\",\"data\",\"created_at\"],\"field_sortable\":[\"n\",\"n\",\"n\",\"n\",\"n\",\"n\",\"n\",\"n\",\"n\"],\"filter_form_layout\":[\"LayoutDefault\"],\"form_description\":[\"\"],\"form_title\":[\"\"],\"hide_back_button\":[\"n\"],\"hide_continue_edit_check_box\":[\"n\"],\"hide_continue_new_check_box\":[\"n\"],\"hide_delete_button\":[\"n\"],\"hide_detail_button\":[\"n\"],\"hide_edit_button\":[\"n\"],\"hide_export_button\":[\"n\"],\"hide_filter_area\":[\"n\"],\"hide_filter_button\":[\"n\"],\"hide_new_button\":[\"n\"],\"hide_pagination\":[\"n\"],\"hide_query_info\":[\"n\"],\"hide_reset_button\":[\"n\"],\"hide_row_selector\":[\"n\"],\"info_field_editable\":[\"n\",\"n\",\"n\",\"n\",\"n\",\"n\",\"n\",\"n\",\"n\"],\"package\":[\"tables\"],\"path\":[\"/home/programmer-yan/Documents/GoFile/gotest/frame/my-admin/tables\"],\"permission\":[\"n\"],\"pk\":[\"id\"],\"table\":[\"lang\"],\"table_description\":[\"\"],\"table_title\":[\"\"]}','2024-03-11 16:20:07','2024-03-11 16:20:07'),(9,1,'/admin/menu','GET','::1','','2024-03-11 16:24:28','2024-03-11 16:24:28'),(10,1,'/admin/menu/new','POST','::1','{\"__go_admin_previous_\":[\"/admin/menu\"],\"__go_admin_t_\":[\"04258849-e53b-456d-bf80-4fb78dcef2ab\"],\"header\":[\"http://localhost:4100\"],\"icon\":[\"fa-user\"],\"parent_id\":[\"0\"],\"plugin_name\":[\"\"],\"roles[]\":[\"1\",\"2\"],\"title\":[\"用户\"],\"uri\":[\"/admin/info/user\"]}','2024-03-11 16:27:21','2024-03-11 16:27:21'),(11,1,'/admin/menu/edit/show','GET','::1','','2024-03-11 16:27:27','2024-03-11 16:27:27'),(12,1,'/admin/menu/edit','POST','::1','{\"__go_admin_previous_\":[\"/admin/menu\"],\"__go_admin_t_\":[\"d61cd605-8c3b-40b1-8441-378e15f4a016\"],\"created_at\":[\"2024-03-12 00:27:21\"],\"header\":[\"http://localhost:4100\"],\"icon\":[\"fa-user\"],\"id\":[\"8\"],\"parent_id\":[\"0\"],\"plugin_name\":[\"\"],\"roles[]\":[\"1\",\"2\"],\"title\":[\"用户\"],\"updated_at\":[\"2024-03-12 00:27:21\"],\"uri\":[\"/info/user\"]}','2024-03-11 16:27:38','2024-03-11 16:27:38'),(13,1,'/admin/menu','GET','::1','','2024-03-11 16:27:44','2024-03-11 16:27:44'),(14,1,'/admin/menu','GET','::1','','2024-03-11 16:27:48','2024-03-11 16:27:48'),(15,1,'/admin/info/user','GET','::1','','2024-03-11 16:27:50','2024-03-11 16:27:50'),(16,1,'/admin/info/user','GET','::1','','2024-03-11 16:27:51','2024-03-11 16:27:51'),(17,1,'/admin/info/user','GET','::1','','2024-03-11 16:28:41','2024-03-11 16:28:41'),(18,1,'/admin/info/user/new','GET','::1','','2024-03-11 16:28:59','2024-03-11 16:28:59'),(19,1,'/admin/new/user','POST','::1','{\"__go_admin_previous_\":[\"/admin/info/user?__page=1\\u0026__pageSize=10\\u0026__sort=id\\u0026__sort_type=desc\"],\"__go_admin_t_\":[\"ca2cd202-3279-4ef0-8bc7-28b321d1f847\"],\"city\":[\"yulin\"],\"created_at\":[\"2024-03-12 00:00:00\"],\"gender\":[\"man\"],\"id\":[\"\"],\"ip\":[\"192.168.5.34\"],\"name\":[\"yan\"],\"phone\":[\"19907751429\"],\"updated_at\":[\"2024-03-12 00:00:00\"]}','2024-03-11 16:29:41','2024-03-11 16:29:41'),(20,1,'/admin/new/user','POST','::1','{\"__go_admin_previous_\":[\"/admin/info/user?__page=1\\u0026__pageSize=10\\u0026__sort=id\\u0026__sort_type=desc\"],\"__go_admin_t_\":[\"b1fab4d8-6a77-4e94-a9fe-f23ec57ff745\"],\"city\":[\"yulin\"],\"created_at\":[\"2024-03-12 00:00:00\"],\"gender\":[\"1\"],\"id\":[\"\"],\"ip\":[\"192.168.5.34\"],\"name\":[\"Yan\"],\"phone\":[\"19907751429\"],\"updated_at\":[\"2024-03-12 00:00:00\"]}','2024-03-11 16:30:16','2024-03-11 16:30:16'),(21,1,'/admin/new/user','POST','::1','{\"__go_admin_previous_\":[\"/admin/info/user?__page=1\\u0026__pageSize=10\\u0026__sort=id\\u0026__sort_type=desc\"],\"__go_admin_t_\":[\"dd4e01b8-0ef1-4089-a168-df9056366077\"],\"city\":[\"yulin\"],\"created_at\":[\"2024-03-12 00:00:00\"],\"gender\":[\"1\"],\"id\":[\"\"],\"ip\":[\"192.168.5.34\"],\"name\":[\"Yan\"],\"phone\":[\"\\\"19907751429\\\"\"],\"updated_at\":[\"2024-03-12 00:00:00\"]}','2024-03-11 16:30:50','2024-03-11 16:30:50'),(22,1,'/admin/new/user','POST','::1','{\"__go_admin_previous_\":[\"/admin/info/user?__page=1\\u0026__pageSize=10\\u0026__sort=id\\u0026__sort_type=desc\"],\"__go_admin_t_\":[\"995906e3-1df2-49cf-b668-2a7c8ae46ec1\"],\"city\":[\"yulin\"],\"created_at\":[\"2024-03-12 00:00:00\"],\"gender\":[\"1\"],\"id\":[\"\"],\"ip\":[\"192.168.5.34\"],\"name\":[\"Yan\"],\"phone\":[\"15577098792\"],\"updated_at\":[\"2024-03-12 00:00:00\"]}','2024-03-11 16:31:24','2024-03-11 16:31:24'),(23,1,'/admin/new/user','POST','::1','{\"__go_admin_previous_\":[\"/admin/info/user?__page=1\\u0026__pageSize=10\\u0026__sort=id\\u0026__sort_type=desc\"],\"__go_admin_t_\":[\"7bc6dbe7-9089-4aba-9597-cc7291532690\"],\"city\":[\"yulin\"],\"created_at\":[\"2024-03-12 00:00:00\"],\"gender\":[\"1\"],\"id\":[\"\"],\"ip\":[\"192.168.5.34\"],\"name\":[\"Yan\"],\"phone\":[\"1556\"],\"updated_at\":[\"2024-03-12 00:00:00\"]}','2024-03-11 16:31:48','2024-03-11 16:31:48'),(24,1,'/admin/info/user/edit','GET','::1','','2024-03-11 16:31:57','2024-03-11 16:31:57'),(25,1,'/admin/edit/user','POST','::1','{\"__go_admin_previous_\":[\"/admin/info/user?__page=1\\u0026__pageSize=10\\u0026__sort=id\\u0026__sort_type=desc\"],\"__go_admin_t_\":[\"2f213e7d-2596-4fa7-b4a6-9ad5e13ee302\"],\"city\":[\"yulin\"],\"created_at\":[\"2024-03-12 00:00:00\"],\"gender\":[\"1\"],\"id\":[\"1\"],\"ip\":[\"192.168.5.34\"],\"name\":[\"Yan\"],\"phone\":[\"15577098792\"],\"updated_at\":[\"2024-03-12 00:00:00\"]}','2024-03-11 16:32:07','2024-03-11 16:32:07'),(26,1,'/admin/edit/user','POST','::1','{\"__go_admin_previous_\":[\"/admin/info/user?__page=1\\u0026__pageSize=10\\u0026__sort=id\\u0026__sort_type=desc\"],\"__go_admin_t_\":[\"ce003ccc-f6a3-419a-8c3b-58cf32d81b95\"],\"city\":[\"yulin\"],\"created_at\":[\"2024-03-12 00:00:00\"],\"gender\":[\"1\"],\"id\":[\"1\"],\"ip\":[\"192.168.5.34\"],\"name\":[\"Yan\"],\"phone\":[\"15577098792\"],\"updated_at\":[\"2024-03-12 00:00:00\"]}','2024-03-11 16:33:59','2024-03-11 16:33:59'),(27,1,'/admin/info/user/edit','GET','::1','','2024-03-11 16:34:08','2024-03-11 16:34:08'),(28,1,'/admin/info/user','GET','::1','','2024-03-11 16:38:39','2024-03-11 16:38:39'),(29,1,'/admin/info/user','GET','::1','','2024-03-11 16:39:27','2024-03-11 16:39:27'),(30,1,'/admin/info/user','GET','::1','','2024-03-11 16:49:30','2024-03-11 16:49:30'),(31,1,'/admin/info/user','GET','::1','','2024-03-11 16:50:07','2024-03-11 16:50:07'),(32,1,'/admin/info/user','GET','::1','','2024-03-11 16:50:55','2024-03-11 16:50:55'),(33,1,'/admin/info/user','GET','::1','','2024-03-11 16:52:24','2024-03-11 16:52:24'),(34,1,'/admin/info/user','GET','::1','','2024-03-11 16:55:05','2024-03-11 16:55:05'),(35,1,'/admin/info/user','GET','::1','','2024-03-11 17:01:36','2024-03-11 17:01:36'),(36,1,'/admin/info/user/edit','GET','::1','','2024-03-11 17:01:50','2024-03-11 17:01:50'),(37,1,'/admin/info/user','GET','::1','','2024-03-11 17:01:53','2024-03-11 17:01:53'),(38,1,'/admin/info/site/edit','GET','::1','','2024-03-11 17:05:09','2024-03-11 17:05:09'),(39,1,'/admin/edit/site','POST','::1','{\"__checkbox__debug\":[\"on\"],\"access_assets_log_off\":[\"false\"],\"access_log_off\":[\"false\"],\"access_log_path\":[\"./logs/access.log\"],\"allow_del_operation_log\":[\"false\"],\"animation_delay\":[\"0.00\"],\"animation_duration\":[\"0.00\"],\"animation_type\":[\"\"],\"asset_url\":[\"\"],\"bootstrap_file_path\":[\"./bootstrap.go\"],\"color_scheme\":[\"\"],\"custom_403_html\":[\"\"],\"custom_404_html\":[\"\"],\"custom_500_html\":[\"\"],\"custom_foot_html\":[\"\"],\"custom_head_html\":[\"\"],\"debug\":[\"true\"],\"env\":[\"local\"],\"error_log_off\":[\"false\"],\"error_log_path\":[\"./logs/error.log\"],\"extra\":[\"\"],\"file_upload_engine\":[\"{\\\"name\\\":\\\"local\\\"}\"],\"footer_info\":[\"\"],\"go_mod_file_path\":[\"./go.mod\"],\"hide_app_info_entrance\":[\"false\"],\"hide_config_center_entrance\":[\"false\"],\"hide_plugin_entrance\":[\"false\"],\"hide_tool_entrance\":[\"false\"],\"id\":[\"1\"],\"info_log_off\":[\"false\"],\"info_log_path\":[\"./logs/info.log\"],\"language\":[\"en\"],\"logger_encoder_caller\":[\"full\"],\"logger_encoder_caller_key\":[\"caller\"],\"logger_encoder_duration\":[\"string\"],\"logger_encoder_encoding\":[\"console\"],\"logger_encoder_level\":[\"capitalColor\"],\"logger_encoder_level_key\":[\"level\"],\"logger_encoder_message_key\":[\"msg\"],\"logger_encoder_name_key\":[\"logger\"],\"logger_encoder_stacktrace_key\":[\"stacktrace\"],\"logger_encoder_time\":[\"iso8601\"],\"logger_encoder_time_key\":[\"ts\"],\"logger_level\":[\"0\"],\"logger_rotate_compress\":[\"false\"],\"logger_rotate_max_age\":[\"30\"],\"logger_rotate_max_backups\":[\"5\"],\"logger_rotate_max_size\":[\"10\"],\"login_logo\":[\"\"],\"login_title\":[\"GoAdmin\"],\"logo\":[\"\\u003cb\\u003eGo\\u003c/b\\u003eAdmin\"],\"mini_logo\":[\"\\u003cb\\u003eG\\u003c/b\\u003eA\"],\"no_limit_login_ip\":[\"false\"],\"operation_log_off\":[\"false\"],\"session_life_time\":[\"7200\"],\"sql_log\":[\"false\"],\"theme\":[\"sword\"],\"title\":[\"GoAdmin\"]}','2024-03-11 17:05:25','2024-03-11 17:05:25'),(40,1,'/admin/info/site','GET','::1','','2024-03-11 17:05:26','2024-03-11 17:05:26'),(41,1,'/admin/info/site/edit','GET','::1','','2024-03-11 17:05:26','2024-03-11 17:05:26'),(42,1,'/admin/info/user','GET','::1','','2024-03-11 17:05:29','2024-03-11 17:05:29'),(43,1,'/admin','GET','::1','','2024-06-01 18:59:37','2024-06-01 18:59:37'),(44,1,'/admin/info/manager','GET','::1','','2024-06-01 18:59:46','2024-06-01 18:59:46'),(45,1,'/admin/info/roles','GET','::1','','2024-06-01 18:59:49','2024-06-01 18:59:49'),(46,1,'/admin/info/permission','GET','::1','','2024-06-01 18:59:50','2024-06-01 18:59:50'),(47,1,'/admin/menu','GET','::1','','2024-06-01 18:59:51','2024-06-01 18:59:51'),(48,1,'/admin/hello','GET','::1','','2024-06-01 19:00:14','2024-06-01 19:00:14'),(49,1,'/admin','GET','::1','','2024-06-01 21:33:02','2024-06-01 21:33:02'),(50,1,'/admin/info/manager','GET','::1','','2024-06-01 21:33:09','2024-06-01 21:33:09'),(51,1,'/admin/info/roles','GET','::1','','2024-06-01 21:33:10','2024-06-01 21:33:10'),(52,1,'/admin/info/permission','GET','::1','','2024-06-01 21:33:11','2024-06-01 21:33:11'),(53,1,'/admin/menu','GET','::1','','2024-06-01 21:33:12','2024-06-01 21:33:12'),(54,1,'/admin/info/op','GET','::1','','2024-06-01 21:33:15','2024-06-01 21:33:15'),(55,1,'/admin/info/user','GET','::1','','2024-06-01 21:33:17','2024-06-01 21:33:17'),(56,1,'/admin/info/user','GET','::1','','2024-06-01 21:33:51','2024-06-01 21:33:51'),(57,1,'/admin/info/user','GET','::1','','2024-06-01 21:34:27','2024-06-01 21:34:27'),(58,1,'/admin/info/user','GET','::1','','2024-06-01 21:34:34','2024-06-01 21:34:34'),(59,1,'/admin/info/user','GET','::1','','2024-06-01 21:34:38','2024-06-01 21:34:38'),(60,1,'/admin/info/manager','GET','::1','','2024-06-01 21:34:40','2024-06-01 21:34:40'),(61,1,'/admin/info/user','GET','::1','','2024-06-01 21:34:45','2024-06-01 21:34:45'),(62,1,'/admin/info/user/edit','GET','::1','','2024-06-01 21:35:15','2024-06-01 21:35:15'),(63,1,'/admin','GET','::1','','2024-06-02 07:36:51','2024-06-02 07:36:51'),(64,1,'/admin/logout','GET','::1','','2024-06-02 07:38:02','2024-06-02 07:38:02'),(65,1,'/admin','GET','::1','','2024-06-02 07:38:04','2024-06-02 07:38:04'),(66,1,'/admin/logout','GET','::1','','2024-06-02 07:39:10','2024-06-02 07:39:10'),(67,1,'/admin','GET','::1','','2024-06-02 07:39:12','2024-06-02 07:39:12'),(68,1,'/admin/logout','GET','::1','','2024-06-02 07:40:10','2024-06-02 07:40:10'),(69,1,'/admin','GET','::1','','2024-06-02 07:43:32','2024-06-02 07:43:32');
/*!40000 ALTER TABLE `goadmin_operation_log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `goadmin_permissions`
--

DROP TABLE IF EXISTS `goadmin_permissions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `goadmin_permissions` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `slug` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `http_method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `http_path` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_permissions_name_unique` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_permissions`
--

LOCK TABLES `goadmin_permissions` WRITE;
/*!40000 ALTER TABLE `goadmin_permissions` DISABLE KEYS */;
INSERT INTO `goadmin_permissions` VALUES (1,'All permission','*','','*','2019-09-09 16:00:00','2019-09-09 16:00:00'),(2,'Dashboard','dashboard','GET,PUT,POST,DELETE','/','2019-09-09 16:00:00','2019-09-09 16:00:00'),(3,'users 查询','users_query','GET','/info/users','2024-03-11 16:09:52','2024-03-11 16:09:52'),(4,'users 编辑页显示','users_show_edit','GET','/info/users/edit','2024-03-11 16:09:52','2024-03-11 16:09:52'),(5,'users 新建记录页显示','users_show_create','GET','/info/users/new','2024-03-11 16:09:52','2024-03-11 16:09:52'),(6,'users 编辑','users_edit','POST','/edit/users','2024-03-11 16:09:52','2024-03-11 16:09:52'),(7,'users 新建','users_create','POST','/new/users','2024-03-11 16:09:52','2024-03-11 16:09:52'),(8,'users 删除','users_delete','POST','/delete/users','2024-03-11 16:09:52','2024-03-11 16:09:52'),(9,'users 导出','users_export','POST','/export/users','2024-03-11 16:09:52','2024-03-11 16:09:52');
/*!40000 ALTER TABLE `goadmin_permissions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `goadmin_role_menu`
--

DROP TABLE IF EXISTS `goadmin_role_menu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `goadmin_role_menu` (
  `role_id` int unsigned NOT NULL,
  `menu_id` int unsigned NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  KEY `admin_role_menu_role_id_menu_id_index` (`role_id`,`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_role_menu`
--

LOCK TABLES `goadmin_role_menu` WRITE;
/*!40000 ALTER TABLE `goadmin_role_menu` DISABLE KEYS */;
INSERT INTO `goadmin_role_menu` VALUES (1,1,'2019-09-09 16:00:00','2019-09-09 16:00:00'),(1,7,'2019-09-09 16:00:00','2019-09-09 16:00:00'),(2,7,'2019-09-09 16:00:00','2019-09-09 16:00:00'),(1,8,'2024-03-11 16:27:38','2024-03-11 16:27:38'),(2,8,'2024-03-11 16:27:38','2024-03-11 16:27:38');
/*!40000 ALTER TABLE `goadmin_role_menu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `goadmin_role_permissions`
--

DROP TABLE IF EXISTS `goadmin_role_permissions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `goadmin_role_permissions` (
  `role_id` int unsigned NOT NULL,
  `permission_id` int unsigned NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY `admin_role_permissions` (`role_id`,`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_role_permissions`
--

LOCK TABLES `goadmin_role_permissions` WRITE;
/*!40000 ALTER TABLE `goadmin_role_permissions` DISABLE KEYS */;
INSERT INTO `goadmin_role_permissions` VALUES (1,1,'2019-09-09 16:00:00','2019-09-09 16:00:00'),(1,2,'2019-09-09 16:00:00','2019-09-09 16:00:00'),(2,2,'2019-09-09 16:00:00','2019-09-09 16:00:00');
/*!40000 ALTER TABLE `goadmin_role_permissions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `goadmin_role_users`
--

DROP TABLE IF EXISTS `goadmin_role_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `goadmin_role_users` (
  `role_id` int unsigned NOT NULL,
  `user_id` int unsigned NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY `admin_user_roles` (`role_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_role_users`
--

LOCK TABLES `goadmin_role_users` WRITE;
/*!40000 ALTER TABLE `goadmin_role_users` DISABLE KEYS */;
INSERT INTO `goadmin_role_users` VALUES (1,1,'2019-09-09 16:00:00','2019-09-09 16:00:00'),(2,2,'2019-09-09 16:00:00','2019-09-09 16:00:00');
/*!40000 ALTER TABLE `goadmin_role_users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `goadmin_roles`
--

DROP TABLE IF EXISTS `goadmin_roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `goadmin_roles` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `slug` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_roles_name_unique` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_roles`
--

LOCK TABLES `goadmin_roles` WRITE;
/*!40000 ALTER TABLE `goadmin_roles` DISABLE KEYS */;
INSERT INTO `goadmin_roles` VALUES (1,'Administrator','administrator','2019-09-09 16:00:00','2019-09-09 16:00:00'),(2,'Operator','operator','2019-09-09 16:00:00','2019-09-09 16:00:00');
/*!40000 ALTER TABLE `goadmin_roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `goadmin_session`
--

DROP TABLE IF EXISTS `goadmin_session`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `goadmin_session` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `sid` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `values` varchar(3000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_session`
--

LOCK TABLES `goadmin_session` WRITE;
/*!40000 ALTER TABLE `goadmin_session` DISABLE KEYS */;
INSERT INTO `goadmin_session` VALUES (31,'47bc777d-de82-4e70-a9d2-473eccf1ed92','{\"user_id\":1}','2024-06-02 07:40:10','2024-06-02 07:40:10');
/*!40000 ALTER TABLE `goadmin_session` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `goadmin_site`
--

DROP TABLE IF EXISTS `goadmin_site`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `goadmin_site` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `value` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
  `description` varchar(3000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `state` tinyint unsigned NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=69 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_site`
--

LOCK TABLES `goadmin_site` WRITE;
/*!40000 ALTER TABLE `goadmin_site` DISABLE KEYS */;
INSERT INTO `goadmin_site` VALUES (1,'logger_encoder_time','iso8601',NULL,1,'2024-03-11 16:15:20','2024-03-11 16:15:20'),(2,'operation_log_off','false',NULL,1,'2024-03-11 16:15:20','2024-03-11 16:15:20'),(3,'prohibit_config_modification','false',NULL,1,'2024-03-11 16:15:20','2024-03-11 16:15:20'),(4,'open_admin_api','false',NULL,1,'2024-03-11 16:15:20','2024-03-11 16:15:20'),(5,'info_log_off','false',NULL,1,'2024-03-11 16:15:20','2024-03-11 16:15:20'),(6,'session_life_time','7200',NULL,1,'2024-03-11 16:15:20','2024-03-11 16:15:20'),(7,'access_log_path','./logs/access.log',NULL,1,'2024-03-11 16:15:20','2024-03-11 16:15:20'),(8,'animation_type','',NULL,1,'2024-03-11 16:15:20','2024-03-11 16:15:20'),(9,'custom_head_html','',NULL,1,'2024-03-11 16:15:20','2024-03-11 16:15:20'),(10,'auth_user_table','goadmin_users',NULL,1,'2024-03-11 16:15:20','2024-03-11 16:15:20'),(11,'custom_404_html','',NULL,1,'2024-03-11 16:15:20','2024-03-11 16:15:20'),(12,'logo','<b>Go</b>Admin',NULL,1,'2024-03-11 16:15:20','2024-03-11 16:15:20'),(13,'index_url','/',NULL,1,'2024-03-11 16:15:20','2024-03-11 16:15:20'),(14,'access_log_off','false',NULL,1,'2024-03-11 16:15:20','2024-03-11 16:15:20'),(15,'logger_rotate_compress','false',NULL,1,'2024-03-11 16:15:20','2024-03-11 16:15:20'),(16,'file_upload_engine','{\"name\":\"local\"}',NULL,1,'2024-03-11 16:15:20','2024-03-11 16:15:20'),(17,'domain','',NULL,1,'2024-03-11 16:15:20','2024-03-11 16:15:20'),(18,'access_assets_log_off','false',NULL,1,'2024-03-11 16:15:20','2024-03-11 16:15:20'),(19,'color_scheme','',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(20,'footer_info','',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(21,'animation_duration','0.00',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(22,'logger_encoder_caller_key','caller',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(23,'logger_encoder_message_key','msg',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(24,'info_log_path','./logs/info.log',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(25,'custom_foot_html','',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(26,'logger_level','0',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(27,'extra','',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(28,'hide_app_info_entrance','false',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(29,'bootstrap_file_path','./bootstrap.go',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(30,'debug','true',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(31,'error_log_off','false',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(32,'logger_encoder_name_key','logger',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(33,'hide_visitor_user_center_entrance','false',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(34,'language','en',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(35,'login_url','/login',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(36,'hide_plugin_entrance','false',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(37,'login_title','GoAdmin',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(38,'exclude_theme_components','null',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(39,'logger_encoder_stacktrace_key','stacktrace',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(40,'logger_encoder_duration','string',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(41,'sql_log','false',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(42,'logger_rotate_max_age','30',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(43,'custom_403_html','',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(44,'allow_del_operation_log','false',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(45,'url_prefix','admin',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(46,'title','GoAdmin',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(47,'logger_encoder_level_key','level',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(48,'logger_encoder_encoding','console',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(49,'asset_url','',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(50,'animation_delay','0.00',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(51,'site_off','false',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(52,'hide_tool_entrance','false',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(53,'app_id','UvYmPbYXEvql',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(54,'error_log_path','./logs/error.log',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(55,'go_mod_file_path','./go.mod',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(56,'asset_root_path','./public/',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(57,'login_logo','',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(58,'logger_rotate_max_backups','5',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(59,'logger_encoder_caller','full',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(60,'logger_encoder_time_key','ts',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(61,'logger_encoder_level','capitalColor',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(62,'no_limit_login_ip','false',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(63,'custom_500_html','',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(64,'theme','sword',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(65,'logger_rotate_max_size','10',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(66,'hide_config_center_entrance','false',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(67,'mini_logo','<b>G</b>A',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21'),(68,'env','local',NULL,1,'2024-03-11 16:15:21','2024-03-11 16:15:21');
/*!40000 ALTER TABLE `goadmin_site` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `goadmin_user_permissions`
--

DROP TABLE IF EXISTS `goadmin_user_permissions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `goadmin_user_permissions` (
  `user_id` int unsigned NOT NULL,
  `permission_id` int unsigned NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE KEY `admin_user_permissions` (`user_id`,`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_user_permissions`
--

LOCK TABLES `goadmin_user_permissions` WRITE;
/*!40000 ALTER TABLE `goadmin_user_permissions` DISABLE KEYS */;
INSERT INTO `goadmin_user_permissions` VALUES (1,1,'2019-09-09 16:00:00','2019-09-09 16:00:00'),(2,2,'2019-09-09 16:00:00','2019-09-09 16:00:00');
/*!40000 ALTER TABLE `goadmin_user_permissions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `goadmin_users`
--

DROP TABLE IF EXISTS `goadmin_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `goadmin_users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `remember_token` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_users_username_unique` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `goadmin_users`
--

LOCK TABLES `goadmin_users` WRITE;
/*!40000 ALTER TABLE `goadmin_users` DISABLE KEYS */;
INSERT INTO `goadmin_users` VALUES (1,'admin','$2a$10$PluOuJclSfRCf8BDHNrlWOxtzuxBosrLEXAv8Fs894BLwWISsHcQu','admin','','tlNcBVK9AvfYH7WEnwB1RKvocJu8FfRy4um3DJtwdHuJy0dwFsLOgAc0xUfh','2019-09-09 16:00:00','2019-09-09 16:00:00'),(2,'operator','$2a$10$rVqkOzHjN2MdlEprRflb1eGP0oZXuSrbJLOmJagFsCd81YZm0bsh.','Operator','',NULL,'2019-09-09 16:00:00','2019-09-09 16:00:00');
/*!40000 ALTER TABLE `goadmin_users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lang`
--

DROP TABLE IF EXISTS `lang`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `lang` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `admin_id` int unsigned NOT NULL DEFAULT '0' COMMENT '管理员ID',
  `name` varchar(50) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '名称',
  `alias` varchar(50) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '别名',
  `icon` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '图标',
  `sort` tinyint NOT NULL DEFAULT '99' COMMENT '排序',
  `status` tinyint NOT NULL DEFAULT '10' COMMENT '状态 -1禁用｜10启用',
  `data` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '数据',
  `created_at` int unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户语言';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lang`
--

LOCK TABLES `lang` WRITE;
/*!40000 ALTER TABLE `lang` DISABLE KEYS */;
/*!40000 ALTER TABLE `lang` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `gender` tinyint DEFAULT NULL,
  `city` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `ip` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `phone` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'Yan',1,'yulin','192.168.5.34','15577098792','2024-03-11 16:00:00','2024-03-11 16:00:00');
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

-- Dump completed on 2024-06-02 15:57:07
