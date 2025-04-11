DROP TABLE IF EXISTS `sessions`;

CREATE TABLE `sessions` (
  `sessionID` BIGINT UNSIGNED AUTO_INCREMENT,
  `userID` BIGINT UNSIGNED NOT NULL,
  `sessionToken` VARCHAR(255) NOT NULL UNIQUE,
  `expires_at` DATETIME NOT NULL,
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`sessionID`),
  FOREIGN KEY (`userID`) REFERENCES `users` (`userID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `sessions` WRITE;

UNLOCK TABLES;

DROP TABLE IF EXISTS `admins`;

CREATE TABLE `admins` (
  `adminID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `adminUsername` VARCHAR(255) NOT NULL,
  `adminPasswordHash` VARCHAR(255) NOT NULL,
  `adminEmail` VARCHAR(255) NOT NULL,
  `sessionID` INT DEFAULT NULL,
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`adminID`),
  UNIQUE KEY `adminUsername` (`adminUsername`),
  UNIQUE KEY `adminEmail` (`adminEmail`),
  FOREIGN KEY (`sessionID`) REFERENCES `sessions` (`sessionID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `admins` WRITE;

UNLOCK TABLES;


DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `userID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `userUsername` VARCHAR(255) NOT NULL,
  `userEmail` VARCHAR(255) NOT NULL,	
  `userPasswordHash` VARCHAR(255) NOT NULL,
  `userProfilePicture` VARCHAR(255),
  `sessionID` INT DEFAULT NULL,
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`userID`),
  UNIQUE KEY `userUsername` (`userUsername`),
  UNIQUE KEY `userEmail` (`userEmail`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `users` WRITE;

UNLOCK TABLES;

DROP TABLE IF EXISTS `guests`;

CREATE TABLE `guests` (
  `guestID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `sessionID` BIGINT UNSIGNED,
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
  `last_visited_at` DATETIME DEFAULT NULL,
  PRIMARY KEY (`guestID`);
  FOREIGN KEY (`sessionsID`) REFERENCES `sessions` (`sessionsID`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `guests` WRITE;

UNLOCK TABLES;

DROP TABLE IF EXISTS `posts`;

CREATE TABLE `posts` (
  `postID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `categoriesID` BIGINT UNSIGNED,
  `postKey` VARCHAR(255) NOT NULL,
  `imageID` BIGINT UNSIGNED,
  `postComment` TEXT,
  `postLike` INT DEFAULT '0',
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
  `userID` BIGINT UNSIGNED,
  PRIMARY KEY (`postID`),
  UNIQUE KEY `postKey` (`postKey`),
  FOREIGN KEY (`imageID`) REFERENCES `images` (`imageID`),
  FOREIGN KEY (`postLike`) REFERENCES `postsLikes` (`postLike`),
  FOREIGN KEY (`categoriesID`) REFERENCES `categories` (`categoriesID`),
  CONSTRAINT `posts_ibfk_2` FOREIGN KEY (`userID`) REFERENCES `users` (`userID`) ON DELETE CASCADE,
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `posts` WRITE;

UNLOCK TABLES;

DROP TABLE IF EXISTS `comments`;

CREATE TABLE `comments` (
  `commentID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `userID` BIGINT UNSIGNED,
  `postID` BIGINT UNSIGNED,
  `content` TEXT,
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
  `status` VARCHAR(255),
  `visible` BOOLEAN,
  PRIMARY KEY (`commentID`),,
  FOREIGN KEY (`userID`) REFERENCES `users` (`userID`),
  CONSTRAINT `posts_ibfk_2` FOREIGN KEY (`postID`) REFERENCES `posts` (`postID`) ON DELETE CASCADE,
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
-- ajouter commentLike si on a le temps

LOCK TABLES `comments` WRITE;

UNLOCK TABLES;

DROP TABLE IF EXISTS `categories`;

CREATE TABLE `categories` (
  `categoriesID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `categoriesName` VARCHAR(255),
  `categoriesDescription` TEXT,
  PRIMARY KEY (`categoriesID`),
  UNIQUE KEY `categoriesName` (`categoriesName`),
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `categories` WRITE;

UNLOCK TABLES;

DROP TABLE IF EXISTS `postsLikes`;

CREATE TABLE `postsLikes` (
  `userID` INT UNSIGNED,
  `postID` BIGINT UNSIGNED,
  `postLike` BOOLEAN DEFAULT 0,
  FOREIGN KEY (`postID`) REFERENCES `posts` (`postID`),
  FOREIGN KEY (`userID`) REFERENCES `users` (`userID`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `postsLikes` WRITE;

UNLOCK TABLES;

DROP TABLE IF EXISTS `images`;

CREATE TABLE `images` (
    `imageID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `url` VARCHAR(255),
    `filename` VARCHAR(255),
    `data` BLOB,
    PRIMARY KEY (`imageID`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

LOCK TABLES `images` WRITE;

UNLOCK TABLES;

-- Insertion des données
INSERT INTO `admins` (`adminUsername`, `adminPasswordHash`, `adminEmail`, `sessionID`) VALUES
('adminMar', '123', 'marino@ynov.com', 3);

INSERT INTO `users` (`userUsername`, `userPasswordHash`, `sessionID`) VALUES 
('userMar', '1234', 2);
