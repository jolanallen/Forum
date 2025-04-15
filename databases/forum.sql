-- Suppression des tables si elles existent
DROP TABLE IF EXISTS `adminDashboardData`;
DROP TABLE IF EXISTS `likes`;
DROP TABLE IF EXISTS `comments`;
DROP TABLE IF EXISTS `posts`;
DROP TABLE IF EXISTS `guests`;
DROP TABLE IF EXISTS `admins`;
DROP TABLE IF EXISTS `sessions`;
DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `categories`;
DROP TABLE IF EXISTS `images`;

-- Table des images
CREATE TABLE `images` (
   `imageID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
   `url` VARCHAR(255),
   `filename` VARCHAR(255),
   `data` BLOB,
   PRIMARY KEY (`imageID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Utilisateurs
CREATE TABLE `users` (
   `userID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
   `userUsername` VARCHAR(255) NOT NULL,
   `userEmail` VARCHAR(255) NOT NULL,
   `userPasswordHash` VARCHAR(255) NOT NULL,
   `userProfilePicture` BIGINT UNSIGNED,
   `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
   PRIMARY KEY (`userID`),
   UNIQUE (`userUsername`),
   UNIQUE (`userEmail`),
   FOREIGN KEY (`userProfilePicture`) REFERENCES `images` (`imageID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Sessions utilisateur
CREATE TABLE `sessions` (
   `sessionID` BIGINT UNSIGNED AUTO_INCREMENT,
   `userID` BIGINT UNSIGNED NOT NULL,
   `sessionToken` VARCHAR(255) NOT NULL UNIQUE,
   `expires_at` DATETIME NOT NULL,
   `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
   PRIMARY KEY (`sessionID`),
   FOREIGN KEY (`userID`) REFERENCES `users` (`userID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Administrateurs
CREATE TABLE `admins` (
   `adminID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
   `adminUsername` VARCHAR(255) NOT NULL,
   `adminPasswordHash` VARCHAR(255) NOT NULL,
   `adminEmail` VARCHAR(255) NOT NULL,
   `sessionID` BIGINT UNSIGNED DEFAULT NULL,
   `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
   PRIMARY KEY (`adminID`),
   UNIQUE (`adminUsername`),
   UNIQUE (`adminEmail`),
   FOREIGN KEY (`sessionID`) REFERENCES `sessions` (`sessionID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Visiteurs
CREATE TABLE `guests` (
   `guestID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
   `sessionID` BIGINT UNSIGNED,
   `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
   `last_visited_at` DATETIME DEFAULT NULL,
   PRIMARY KEY (`guestID`),
   FOREIGN KEY (`sessionID`) REFERENCES `sessions` (`sessionID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Catégories de posts
CREATE TABLE `categories` (
   `categoriesID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
   `categoriesName` VARCHAR(255),
   `categoriesDescription` TEXT,
   PRIMARY KEY (`categoriesID`),
   UNIQUE (`categoriesName`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Posts
CREATE TABLE `posts` (
   `postID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
   `categoriesID` BIGINT UNSIGNED,
   `postKey` VARCHAR(255) NOT NULL,
   `imageID` BIGINT UNSIGNED,
   `postComment` TEXT,
   `postLike` INT DEFAULT 0,
   `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
   `userID` BIGINT UNSIGNED,
   PRIMARY KEY (`postID`),
   FOREIGN KEY (`imageID`) REFERENCES `images` (`imageID`),
   FOREIGN KEY (`categoriesID`) REFERENCES `categories` (`categoriesID`),
   FOREIGN KEY (`userID`) REFERENCES `users` (`userID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Commentaires
CREATE TABLE `comments` (
   `commentID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
   `userID` BIGINT UNSIGNED,
   `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
   `content` TEXT,
   `postID` BIGINT UNSIGNED,
   `status` ENUM('approuvé', 'en attente', 'rejeté') DEFAULT 'en attente',
   `visible` BOOLEAN DEFAULT TRUE,
   `commentLike` INT DEFAULT 0,
   PRIMARY KEY (`commentID`),
   FOREIGN KEY (`userID`) REFERENCES `users` (`userID`),
   FOREIGN KEY (`postID`) REFERENCES `posts` (`postID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Likes
CREATE TABLE `likes` (
   `userID` BIGINT UNSIGNED NOT NULL,
   `postID` BIGINT UNSIGNED NULL,
   `commentID` BIGINT UNSIGNED NULL,
   `type` VARCHAR(255) NOT NULL,
   PRIMARY KEY (`userID`, `type`, `postID`, `commentID`),
   FOREIGN KEY (`userID`) REFERENCES `users` (`userID`) ON DELETE CASCADE,
   FOREIGN KEY (`postID`) REFERENCES `posts` (`postID`) ON DELETE CASCADE,
   FOREIGN KEY (`commentID`) REFERENCES `comments` (`commentID`) ON DELETE CASCADE,
   CHECK (
      (postID IS NOT NULL AND commentID IS NULL AND type = 'like') OR
      (commentID IS NOT NULL AND postID IS NULL AND type = 'like')
   )
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Données admin dashboard
CREATE TABLE `adminDashboardData` (
   `dashboardID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
   `adminID` BIGINT UNSIGNED NOT NULL,
   `totalUsers` BIGINT DEFAULT 0,
   `totalPosts` BIGINT DEFAULT 0,
   `totalComments` BIGINT DEFAULT 0,
   `totalGuests` BIGINT DEFAULT 0,
   `lastLogin` DATETIME DEFAULT CURRENT_TIMESTAMP,
   `generated_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
   PRIMARY KEY (`dashboardID`),
   FOREIGN KEY (`adminID`) REFERENCES `admins` (`adminID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Insertion de données
INSERT INTO images (url, filename) VALUES
   ('https://example.com/profile1.jpg', 'profile1.jpg'),
   ('https://example.com/profile2.jpg', 'profile2.jpg'),
   ('https://example.com/post1.jpg', 'post1.jpg');

INSERT INTO categories (categoriesName, categoriesDescription) VALUES
   ('Hack', 'Tout sur la technologie'),
   ('Programmation', 'Art créatif et design');

INSERT INTO admins (adminUsername, adminEmail, adminPasswordHash) VALUES
   ('admin1', 'admin1@example.com', 'adminpass1'),
   ('admin2', 'admin2@example.com', 'adminpass2');

INSERT INTO users (userUsername, userEmail, userPasswordHash, userProfilePicture) VALUES
   ('john_doe', 'john@example.com', 'hashed_password_1', 1),
   ('jane_doe', 'jane@example.com', 'hashed_password_2', 2);

INSERT INTO sessions (userID, sessionToken, expires_at) VALUES
   (1, 'token1234', NOW() + INTERVAL 1 DAY),
   (2, 'token4567', NOW() + INTERVAL 1 DAY);

INSERT INTO guests (sessionID) VALUES (1), (2);

INSERT INTO posts (categoriesID, postKey, imageID, postComment, userID) VALUES
   (1, 'hack-post-1', 3, 'Découvrez ce nouveau gadget !', 1),
   (2, 'prog-post-1', NULL, 'Une peinture que j’ai faite ce week-end.', 2);

INSERT INTO comments (userID, postID, content, status, visible) VALUES
   (2, 1, 'Ça a l’air génial !', 'approuvé', TRUE),
   (1, 2, 'Beau travail !', 'approuvé', TRUE);

-- Like sur un post
INSERT INTO likes (userID, postID, commentID, type) VALUES
   (1, 2, NULL, 'like'),
   (2, 1, NULL, 'like');

-- Like sur un commentaire
INSERT INTO likes (userID, postID, commentID, type) VALUES
   (1, NULL, 1, 'like');

-- Données dashboard
INSERT INTO adminDashboardData (
   adminID, totalUsers, totalPosts, totalComments, totalGuests, lastLogin
) VALUES (
   1,
   (SELECT COUNT(*) FROM users),
   (SELECT COUNT(*) FROM posts),
   (SELECT COUNT(*) FROM comments),
   (SELECT COUNT(*) FROM guests),
   NOW()
);
