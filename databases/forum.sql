DROP TABLE IF EXISTS `sessions`;

CREATE TABLE `sessions` (
  `session_id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `user_id` BIGINT UNSIGNED NOT NULL,
  `session_token` VARCHAR(255) NOT NULL UNIQUE,
  `expires_at` DATETIME NOT NULL,
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (`user_id`) REFERENCES `users` (`users_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `sessions` WRITE;

UNLOCK TABLES;

DROP TABLE IF EXISTS `admins`;

CREATE TABLE `admins` (
  `admin_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `admin_username` VARCHAR(255) NOT NULL,
  `admin_password_hash` VARCHAR(255) NOT NULL,
  `admin_email` VARCHAR(255) NOT NULL,
  `admin_cookie_id` INT DEFAULT NULL,
  `admin_key` VARCHAR(255) NOT NULL,
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`admin_id`),
  UNIQUE KEY `admin_username` (`admin_username`),
  UNIQUE KEY `admin_email` (`admin_email`),
  UNIQUE KEY `admin_key` (`admin_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `admins` WRITE;

UNLOCK TABLES;


DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `users_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,  -- Modification ici pour BIGINT UNSIGNED
  `users_username` VARCHAR(255) NOT NULL,
  `users_password_hash` VARCHAR(255) NOT NULL,
  `users_profile_picture` VARCHAR(255),
  `users_cookies_id` INT DEFAULT NULL,
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`users_id`),
  UNIQUE KEY `users_username` (`users_username`),
  UNIQUE KEY `users_password_hash` (`users_password_hash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `users` WRITE;

UNLOCK TABLES;

DROP TABLE IF EXISTS `guests`;

CREATE TABLE `guests` (
  `guests_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `guests_cookie_id` INT DEFAULT NULL,
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
  `last_visited_at` DATETIME DEFAULT NULL,
  PRIMARY KEY (`guests_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `guests` WRITE;

UNLOCK TABLES;

DROP TABLE IF EXISTS `posts`;

CREATE TABLE `posts` (
  `post_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `post_key` VARCHAR(255) NOT NULL,
  `post_image` BLOB,
  `post_comments` TEXT,
  `post_likes` INT DEFAULT '0',
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
  `user_id` BIGINT UNSIGNED DEFAULT NULL,  -- Modification pour BIGINT UNSIGNED
  PRIMARY KEY (`post_id`),
  UNIQUE KEY `post_key` (`post_key`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `posts_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`users_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `posts` WRITE;

UNLOCK TABLES;

DROP TABLE IF EXISTS `comments`;

CREATE TABLE `comments` (
  `comment_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `userID` INT UNSIGNED,
  `topicID` INT UNSIGNED,
  `postID` INT UNSIGNED,
  `content` TEXT,
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
  `status` VARCHAR(255),
  `visible` BOOLEAN,
  `comments_like` INT,
  `comments_dislike` INT,
  PRIMARY KEY (`comment_id`),
  FOREIGN KEY (`topicID`) REFERENCES `topics` (`topics_id`),
  FOREIGN KEY (`userID`) REFERENCES `users` (`users_id`)
);

LOCK TABLES `comments` WRITE;

UNLOCK TABLES;

DROP TABLE IF EXISTS `categories`;

CREATE TABLE `categories` (
  `categories_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `categories_name` VARCHAR(255),
  `categories_description` TEXT
  PRIMARY KEY (`categories_id`),
);

LOCK TABLES `categories` WRITE;

UNLOCK TABLES;

DROP TABLE IF EXISTS `topics`;

CREATE TABLE `topics` (
  `topics_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `topics_categoryID` INT UNSIGNED,
  `topics_userID` INT UNSIGNED,
  `topics_title` VARCHAR(255),
  `topics_content` TEXT,
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
  `topics_like` INT,
  `topics_dislike` INT,
  PRIMARY KEY (`topics_id`),
  FOREIGN KEY (`topics_categoryID`) REFERENCES `categories` (`categories_id`),
  FOREIGN KEY (`topics_userID`) REFERENCES `users` (`users_id`)
);

LOCK TABLES `topics` WRITE;

UNLOCK TABLES;

DROP TABLE IF EXISTS `topicsLikes`;

CREATE TABLE `topicsLikes` (
  `topicsLikes_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `topicID` INT UNSIGNED,
  `userID` INT UNSIGNED,
  PRIMARY KEY (`topicsLikes_id`),
  FOREIGN KEY (`topicID`) REFERENCES `topics` (`topics_id`),
  FOREIGN KEY (`userID`) REFERENCES `users` (`users_id`)
);

LOCK TABLES `topicsLikes` WRITE;

UNLOCK TABLES;

DROP TABLE IF EXISTS `topicsDislikes`;

CREATE TABLE `topicsDislikes` (
  `topicsDislikes_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `topicID` INT UNSIGNED,
  `userID` INT UNSIGNED,
  PRIMARY KEY (`topicsDislikes_id`),
  FOREIGN KEY (`topicID`) REFERENCES `topics` (`topics_id`),
  FOREIGN KEY (`userID`) REFERENCES `users` (`users_id`)
);

LOCK TABLES `topicsDislikes` WRITE;

UNLOCK TABLES;

DROP TABLE IF EXISTS `images`;

CREATE TABLE `images` (
    `images_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `url` VARCHAR(255),
    `filename` VARCHAR(255),
    `data` BLOB
    PRIMARY KEY (`images_id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `images` WRITE;

UNLOCK TABLES;

-- Insertion des données
INSERT INTO `admins` (`admin_username`, `admin_password_hash`, `admin_email`, `admin_key`) VALUES
('adminMar', '123', 'marino@ynov.com', 'yes');

INSERT INTO `users` (`users_username`, `users_password_hash`, `users_cookies_id`) VALUES 
('userMar', '1234', 2);
