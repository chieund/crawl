DROP TABLE IF EXISTS `articles`;
CREATE TABLE `articles` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `title` longtext COLLATE utf8mb4_unicode_ci,
    `slug` longtext COLLATE utf8mb4_unicode_ci,
    `image` longtext COLLATE utf8mb4_unicode_ci,
    `description` longtext COLLATE utf8mb4_unicode_ci,
    `content` longtext COLLATE utf8mb4_unicode_ci,
    `link` longtext COLLATE utf8mb4_unicode_ci,
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    `deleted_at` datetime(3) DEFAULT NULL,
    `viewed` bigint DEFAULT '0',
    `website_id` bigint DEFAULT NULL,
    `website_slug` longtext COLLATE utf8mb4_unicode_ci,
    `is_update_content` bigint DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

DROP TABLE IF EXISTS `tags`;
CREATE TABLE `tags` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `title` longtext COLLATE utf8mb4_unicode_ci,
    `slug` longtext COLLATE utf8mb4_unicode_ci,
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    `deleted_at` datetime(3) DEFAULT NULL,
    `hot` bigint DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

DROP TABLE IF EXISTS `article_tag`;
CREATE TABLE `article_tag` (
   `article_id` bigint NOT NULL,
   `tag_id` bigint NOT NULL,
   PRIMARY KEY (`article_id`,`tag_id`),
   KEY `fk_article_tag_tag` (`tag_id`),
   CONSTRAINT `fk_article_tag_article` FOREIGN KEY (`article_id`) REFERENCES `articles` (`id`),
   CONSTRAINT `fk_article_tag_tag` FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

DROP TABLE IF EXISTS `websites`;
CREATE TABLE `websites` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `title` longtext COLLATE utf8mb4_unicode_ci,
    `slug` longtext COLLATE utf8mb4_unicode_ci,
    `image` longtext COLLATE utf8mb4_unicode_ci,
    `description` longtext COLLATE utf8mb4_unicode_ci,
    `content` longtext COLLATE utf8mb4_unicode_ci,
    `link` longtext COLLATE utf8mb4_unicode_ci,
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    `deleted_at` datetime(3) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

INSERT INTO `websites` VALUES (1,'dev.to','dev-to','/images/icon/dev-to.png',NULL,NULL,NULL,NULL,NULL,NULL),(2,'freecodecamp.org','freecodecamp-org','/images/icon/freecodecam.jpeg',NULL,NULL,NULL,NULL,NULL,NULL),(3,'hashnode.com','hashnode-com','/images/icon/hashnode.jpeg',NULL,NULL,NULL,NULL,NULL,NULL),(5,'logrocket-com','logrocket-com','/images/icon/logrocket.png',NULL,NULL,NULL,NULL,NULL,NULL),(6,'infoq-com','infoq-com','/images/icon/infoq.png',NULL,NULL,NULL,NULL,NULL,NULL);