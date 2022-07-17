# CREATE TABLE `user` (
#                         `id` bigint(20) NOT NULL AUTO_INCREMENT,
#                         `user_id` bigint(20) NOT NULL,
#                         `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
#                         `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
#                         `email` varchar(64) COLLATE utf8mb4_general_ci,
#                         `gender` tinyint(4) NOT NULL DEFAULT '0',
#                         `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
#                         `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
#                         PRIMARY KEY (`id`),
#                         UNIQUE KEY `idx_username` (`username`) USING BTREE,
#                         UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
# ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

DROP    TABLE IF EXISTS `communtity`;
CREATE TABLE `comuntity` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `communtity_id` int(10) unsigned NOT NULL,
    `communtity_name` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
    `introduction` varchar(256) COLLATE utf8mb4_general_ci NOT NULL,
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_community_id` (`communtity_id`),
    UNIQUE KEY `idx_community_name` (`communtity_name`)

)ENGINE=InnoDB DEFAULT CHARSET=UTF8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `comuntity` VALUES('1','1','Go','Golang','2022-07-16','2022-07-17');
INSERT INTO `comuntity` VALUES('2','2','RUST','RUST','2022-07-14','2022-07-15');