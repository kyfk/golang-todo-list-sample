-- create "todos" table
CREATE TABLE `todos` (`id` char(36) NOT NULL, `deleted_at` datetime NOT NULL DEFAULT "0000-00-00 00:00:00", `title` varchar(255) NOT NULL, `desc` varchar(255) NOT NULL, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
