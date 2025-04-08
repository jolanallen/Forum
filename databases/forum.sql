CREATE TABLE IF NOT EXISTS `sessions` (
  `session_id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `user_id` BIGINT UNSIGNED NOT NULL,
  `session_token` VARCHAR(255) NOT NULL UNIQUE,
  `expires_at` DATETIME NOT NULL,
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (`user_id`) REFERENCES `users` (`users_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `admins` (
  `admin_id` INT NOT NULL AUTO_INCREMENT,
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

CREATE TABLE IF NOT EXISTS `users` (
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

CREATE TABLE IF NOT EXISTS `guests` (
  `guests_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `guests_cookie_id` INT DEFAULT NULL,
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
  `last_visited_at` DATETIME DEFAULT NULL,
  PRIMARY KEY (`guests_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `posts` (
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

CREATE TABLE IF NOT EXISTS `comments` (
  `comment_id` INT UNSIGNED PRIMARY KEY,
  `userID` INT UNSIGNED,
  `topicID` INT UNSIGNED,
  `postID` INT UNSIGNED,
  `content` TEXT,
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
  `status` VARCHAR(255),
  `visible` BOOLEAN,
  `comments_like` INT,
  `comments_dislike` INT,
  FOREIGN KEY (`topicID`) REFERENCES `topics` (`topics_id`),
  FOREIGN KEY (`userID`) REFERENCES `users` (`users_id`)
);

CREATE TABLE IF NOT EXISTS `categories` (
  `categories_id` INT UNSIGNED PRIMARY KEY,
  `categories_name` VARCHAR(255),
  `categories_description` TEXT
);

CREATE TABLE IF NOT EXISTS `topics` (
  `topics_id` INT UNSIGNED PRIMARY KEY,
  `topics_categoryID` INT UNSIGNED,
  `topics_userID` INT UNSIGNED,
  `topics_title` VARCHAR(255),
  `topics_content` TEXT,
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
  `topics_like` INT,
  `topics_dislike` INT,
  FOREIGN KEY (`topics_categoryID`) REFERENCES `categories` (`categories_id`),
  FOREIGN KEY (`topics_userID`) REFERENCES `users` (`users_id`)
);

CREATE TABLE IF NOT EXISTS `topicsLikes` (
  `topicsLikes_id` INT UNSIGNED PRIMARY KEY,
  `topicID` INT UNSIGNED,
  `userID` INT UNSIGNED,
  FOREIGN KEY (`topicID`) REFERENCES `topics` (`topics_id`),
  FOREIGN KEY (`userID`) REFERENCES `users` (`users_id`)
);

CREATE TABLE IF NOT EXISTS `topicsDislikes` (
  `topicsDislikes_id` INT UNSIGNED PRIMARY KEY,
  `topicID` INT UNSIGNED,
  `userID` INT UNSIGNED,
  FOREIGN KEY (`topicID`) REFERENCES `topics` (`topics_id`),
  FOREIGN KEY (`userID`) REFERENCES `users` (`users_id`)
);

CREATE TABLE IF NOT EXISTS `images` (
    `images_id` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `url` VARCHAR(255),
    `filename` VARCHAR(255),
    `data` BLOB
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Insertion des données
INSERT INTO `admins` (`admin_username`, `admin_password_hash`, `admin_email`, `admin_cookie_id`, `admin_key`) VALUES
('adminMar', '123', 'marino@ynov.com', 1, 'yes');

INSERT INTO `users` (`users_username`, `users_password_hash`, `users_cookies_id`) VALUES 
('userMar', '1234', 2);
