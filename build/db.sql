-- users
DROP DATABASE IF EXISTS `mango_users`;
CREATE DATABASE `mango_users`;
GRANT ALL PRIVILEGES ON `mango_users`.* TO 'icarus'@'%';
USE `mango_users`;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
    `user_id`               BIGINT(20) UNSIGNED         NOT NULL AUTO_INCREMENT,
    `email`                 VARCHAR(100)                NOT NULL,
    `password`              VARCHAR(64)                 NOT NULL,
    `name`                  VARCHAR(50)                 NOT NULL,
    `slogan`                VARCHAR(150)                NOT NULL DEFAULT "",
    -- reader writer admin super
    `role`                  VARCHAR(20)                 NOT NULL DEFAULT "reader",
    `followers_count`       INT(16)                     NOT NULL DEFAULT 0,
    `following_count`       INT(16)                     NOT NULL DEFAULT 0,
    `like_count`            INT(16)                     NOT NULL DEFAULT 0,
    `dislike_count`         INT(16)                     NOT NULL DEFAULT 0,
    `posts_count`           INT(16)                     NOT NULL DEFAULT 0,
    -- online offline
    `online`                BOOLEAN                     NOT NULL DEFAULT FALSE,
    -- actived inactived forbidden deleted
    `account_status`        VARCHAR(20)                 NOT NULL DEFAULT "inactived",
    `created_at`            TIMESTAMP                   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at`            TIMESTAMP,
    `updated_at`            TIMESTAMP,
    PRIMARY KEY(`user_id`),
    UNIQUE (`email`),
    UNIQUE (`name`)
) ENGINE=InnoDB CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

DROP TABLE IF EXISTS `users_relations`;
CREATE TABLE `users_relations` (
    `users_relations_id`    BIGINT(20) UNSIGNED         NOT NULL AUTO_INCREMENT,
    `user_id`               BIGINT(20) UNSIGNED         NOT NULL,
    `related_user_id`       BIGINT(20) UNSIGNED         NOT NULL,
    -- user `relation` related_user
    -- follow block
    `relation`              VARCHAR(20)                 NOT NULL,
    `created_at`            TIMESTAMP                   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at`            TIMESTAMP,
    `updated_at`            TIMESTAMP,
    PRIMARY KEY(`users_relations_id`),
    KEY(`user_id`),
    KEY(`related_user_id`),
    KEY(`relation`)
) ENGINE=InnoDB CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

DROP TABLE IF EXISTS `users_settings`;
CREATE TABLE `users_settings` (
    `user_id`                   BIGINT(20) UNSIGNED         NOT NULL,
    `notification_follower`     BOOLEAN                     NOT NULL DEFAULT TRUE,
    `notification_following`    BOOLEAN                     NOT NULL DEFAULT TRUE,
    `notification_like`         BOOLEAN                     NOT NULL DEFAULT TRUE,
    `notification_reply`        BOOLEAN                     NOT NULL DEFAULT TRUE,
    `created_at`                TIMESTAMP                   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at`                TIMESTAMP,
    `updated_at`                TIMESTAMP,
    KEY(`user_id`)
) ENGINE=InnoDB CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;


-- categories
DROP DATABASE IF EXISTS `mango_categories`;
CREATE DATABASE `mango_categories`;
GRANT ALL PRIVILEGES ON `mango_categories`.* TO 'icarus'@'%';
USE `mango_categories`;

DROP TABLE IF EXISTS `categories`;
CREATE TABLE `categories` (
    `category_id`    BIGINT(20) UNSIGNED         NOT NULL AUTO_INCREMENT,
    `name`           VARCHAR(50)                 NOT NULL,
    `slogan`         VARCHAR(150)                NOT NULL,
    `description`    LONGTEXT                    NOT NULL,
    `parent`         BIGINT(20) UNSIGNED         NOT NULL DEFAULT 0,
    `posts_count`    BIGINT(20) UNSIGNED         NOT NULL DEFAULT 0,
    `created_at`     TIMESTAMP                   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at`     TIMESTAMP,
    `updated_at`     TIMESTAMP,
    PRIMARY KEY(`category_id`),
    KEY (`parent`)
) ENGINE=InnoDB CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;


-- tags
DROP DATABASE IF EXISTS `mango_tags`;
CREATE DATABASE `mango_tags`;
GRANT ALL PRIVILEGES ON `mango_tags`.* TO 'icarus'@'%';
USE `mango_tags`;

DROP TABLE IF EXISTS `tags`;
CREATE TABLE `tags` (
    `tag_id`            BIGINT(20) UNSIGNED         NOT NULL AUTO_INCREMENT,
    `name`              VARCHAR(50)                 NOT NULL,
    `posts_count`       BIGINT(20) UNSIGNED         NOT NULL DEFAULT 0,
    `created_at`        TIMESTAMP                   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at`        TIMESTAMP,
    `updated_at`        TIMESTAMP,
    PRIMARY KEY(`tag_id`),
    UNIQUE (`name`)
) ENGINE=InnoDB CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;


-- posts
DROP DATABASE IF EXISTS `mango_posts`;
CREATE DATABASE `mango_posts`;
GRANT ALL PRIVILEGES ON `mango_posts`.* TO 'icarus'@'%';
USE `mango_posts`;

DROP TABLE IF EXISTS `posts`;
CREATE TABLE `posts` (
    `post_id`            BIGINT(20) UNSIGNED         NOT NULL AUTO_INCREMENT,
    `user_id`            BIGINT(20) UNSIGNED         NOT NULL,
    `user_ip`            BINARY(16)                  NOT NULL,
    `category_id`        BIGINT(20) UNSIGNED         NOT NULL,
    `title`              TEXT                        NOT NULL,
    `excerpt`            TEXT                        NOT NULL DEFAULT "",
    `content`            LONGTEXT                    NOT NULL DEFAULT "",
    -- question, vote, document, system_broadcast
    `class`              VARCHAR(20)                 NOT NULL,
    -- draft, pubish, private
    `status`             VARCHAR(20)                 NOT NULL,
    `allowed_user_role`  VARCHAR(20)                 NOT NULL,
    `comment_status`     VARCHAR(20)                 NOT NULL,
    `views_count`        INT(20) UNSIGNED            NOT NULL DEFAULT 0,
    `created_at`         TIMESTAMP                   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at`         TIMESTAMP,
    `updated_at`         TIMESTAMP,
    PRIMARY KEY(`post_id`),
    KEY (`user_id`),
    KEY (`class`),
    KEY (`category_id`)
) ENGINE=InnoDB CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

DROP TABLE IF EXISTS `posts_tags`;
CREATE TABLE `posts_tags` (
    `post_id`                 BIGINT(20) UNSIGNED         NOT NULL,
    `tag_id`                  BIGINT(20) UNSIGNED         NOT NULL,
    KEY (`post_id`),
    KEY (`tag_id`)
) ENGINE=InnoDB CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

DROP TABLE IF EXISTS `posts_reactions`;
CREATE TABLE `posts_reactions` (
    `user_id`                 BIGINT(20) UNSIGNED         NOT NULL,
    -- like dislike
    `reaction`                VARCHAR(20)                 NOT NULL,
    `post_id`                 BIGINT(20) UNSIGNED         NOT NULL,
    `created_at`              TIMESTAMP                   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at`              TIMESTAMP,
    `updated_at`              TIMESTAMP,
    KEY (`post_id`, `reaction`),
    KEY (`user_id`, `reaction`)
) ENGINE=InnoDB CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

DROP TABLE IF EXISTS `posts_priority`;
CREATE TABLE `posts_priority` (
    `post_id`                 BIGINT(20) UNSIGNED         NOT NULL,
    `post_weight`             INT(16) UNSIGNED            NOT NULL DEFAULT 0,
    `created_at`              TIMESTAMP                   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at`              TIMESTAMP,
    `updated_at`              TIMESTAMP,
    KEY (`post_id`),
    KEY (`post_weight`)
) ENGINE=InnoDB CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;


-- comments
DROP DATABASE IF EXISTS `mango_comments`;
CREATE DATABASE `mango_comments`;
GRANT ALL PRIVILEGES ON `mango_comments`.* TO 'icarus'@'%';
USE `mango_comments`;

DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
    `comment_id`          BIGINT(20) UNSIGNED         NOT NULL AUTO_INCREMENT,
    `post_id`             BIGINT(20) UNSIGNED         NOT NULL,
    `sender_user_id`      BIGINT(20) UNSIGNED         NOT NULL,
    `sender_user_ip`      BINARY(16)                  NOT NULL,
    `receiver_user_id`    BIGINT(20) UNSIGNED         NOT NULL,
    `content`             TEXT                        NOT NULL,
    `like_count`          INT(16)                     NOT NULL DEFAULT 0,
    `dislike_count`       INT(16)                     NOT NULL DEFAULT 0,
    `parent`              BIGINT(20) UNSIGNED         NOT NULL DEFAULT 0,
    `created_at`          TIMESTAMP                   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at`          TIMESTAMP,
    `updated_at`          TIMESTAMP,
    PRIMARY KEY(`comment_id`),
    KEY (`post_id`),
    KEY (`sender_user_id`),
    KEY (`receiver_user_id`),
    KEY (`parent`)
) ENGINE=InnoDB CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

DROP TABLE IF EXISTS `comments_reactions`;
CREATE TABLE `comments_reactions` (
    `comment_id`              BIGINT(20) UNSIGNED         NOT NULL,
    `reaction`                VARCHAR(20)                 NOT NULL,
    `user_id`                 BIGINT(20) UNSIGNED         NOT NULL,
    `created_at`              TIMESTAMP                   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at`              TIMESTAMP,
    `updated_at`              TIMESTAMP,
    KEY (`comment_id`, `reaction`),
    KEY (`user_id`, `reaction`)
) ENGINE=InnoDB CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;


-- notifications
DROP DATABASE IF EXISTS `mango_notifications`;
CREATE DATABASE `mango_notifications`;
GRANT ALL PRIVILEGES ON `mango_notifications`.* TO 'icarus'@'%';
USE `mango_notifications`;

-- push
DROP TABLE IF EXISTS `notifications`;
CREATE TABLE `notifications` (
    `notification_id`            BIGINT(20) UNSIGNED         NOT NULL AUTO_INCREMENT,
    -- read unread ignore
    `status`        VARCHAR(20)                 NOT NULL,
    -- system broadcast, comment reply, private chat, mention(@), like
    `class`         VARCHAR(20)                 NOT NULL,
    `user_id`       BIGINT(20) UNSIGNED         NOT NULL,
    `object_id`     BIGINT(20) UNSIGNED         NOT NULL,
    `created_at`    TIMESTAMP                   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at`    TIMESTAMP,
    `updated_at`    TIMESTAMP,
    PRIMARY KEY(notification_id),
    KEY (`user_id`),
    KEY (`class`, `object_id`, `status`)
) ENGINE=InnoDB CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

-- pull: to lazily generate notifications when user requests
DROP TABLE IF EXISTS `notifications_lazy`;
CREATE TABLE `notifications_lazy` (
    -- updates of following users; system broadcast
    `notification_type`       VARCHAR(20)                 NOT NULL,
    -- users_relations_id; user_id
    `object_id`               BIGINT(20) UNSIGNED         NOT NULL,
    `last_scan_time`          TIMESTAMP                   NOT NULL,
    `created_at`              TIMESTAMP                   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at`              TIMESTAMP,
    `updated_at`              TIMESTAMP,
    KEY (`notification_type`, `object_id`)
) ENGINE=InnoDB CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

-- validation center
DROP DATABASE IF EXISTS `mango_validations`;
CREATE DATABASE `mango_validations`;
GRANT ALL PRIVILEGES ON `mango_validations`.* TO 'icarus'@'%';
USE `mango_validations`;

DROP TABLE IF EXISTS `validations`;
CREATE TABLE `validations` (
    `validation_id`            BIGINT(20) UNSIGNED         NOT NULL AUTO_INCREMENT,
    -- email, third party service
    `class`         VARCHAR(20)                 NOT NULL,
    `object_id`     BIGINT(20) UNSIGNED         NOT NULL,
    `created_at`    TIMESTAMP                   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at`    TIMESTAMP,
    `updated_at`    TIMESTAMP,
    PRIMARY KEY(validation_id)
) ENGINE=InnoDB CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

DROP TABLE IF EXISTS `validations_email`;
CREATE TABLE `validations_email` (
    `validation_email_id`     BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    `code`                    VARCHAR(20)                 NOT NULL,
    `created_at`              TIMESTAMP                   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at`              TIMESTAMP,
    `updated_at`              TIMESTAMP,
    KEY (`validation_email_id`)
) ENGINE=InnoDB CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

-- storage
