-- Désactivation des vérifications de clés étrangères pour pouvoir supprimer les tables
SET foreign_key_checks = 0;

-- Suppression des tables si elles existent
DROP TABLE IF EXISTS `adminDashboardData`;
DROP TABLE IF EXISTS `postsLikes`;
DROP TABLE IF EXISTS `commentsLikes`;
DROP TABLE IF EXISTS `comments`;
DROP TABLE IF EXISTS `posts`;
DROP TABLE IF EXISTS `guests`;
DROP TABLE IF EXISTS `admins`;
DROP TABLE IF EXISTS `sessionsAdmins`;
DROP TABLE IF EXISTS `sessionsGuests`;
DROP TABLE IF EXISTS `sessionsUsers`;
DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `categories`;
DROP TABLE IF EXISTS `images`;

-- Réactivation des vérifications de clés étrangères
SET foreign_key_checks = 1;

-- Création des tables à partir d'ici comme dans le script précédent...
-- Table des images
CREATE TABLE `images` (
    `imageID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `url` VARCHAR(255),
    `filename` VARCHAR(255),
    `data` BLOB,
    PRIMARY KEY (`imageID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Catégories de posts
CREATE TABLE `categories` (
    `categoryID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `categoryName` VARCHAR(255),
    `categoryDescription` TEXT,
    PRIMARY KEY (`categoryID`),
    UNIQUE (`categoryName`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Utilisateurs
CREATE TABLE `users` (
    `userID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `userUsername` VARCHAR(255) NOT NULL,
    `userEmail` VARCHAR(255) NOT NULL,
    `userPasswordHash` VARCHAR(255) NOT NULL,
    `userProfilePicture` BIGINT UNSIGNED NOT NULL,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`userID`),
    UNIQUE (`userUsername`),
    UNIQUE (`userEmail`),
    FOREIGN KEY (`userProfilePicture`) REFERENCES `images` (`imageID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- Créer la table `admins`
CREATE TABLE `admins` (
    `adminID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `adminUsername` VARCHAR(255) NOT NULL,
    `adminPasswordHash` VARCHAR(255) NOT NULL,
    `adminEmail` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`adminID`),
    UNIQUE (`adminUsername`),
    UNIQUE (`adminEmail`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Créer la table `guests`
CREATE TABLE `guests` (
    `guestID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `last_visited_at` DATETIME DEFAULT NULL,
    PRIMARY KEY (`guestID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Posts
CREATE TABLE `posts` (
    `postID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `categoryID` BIGINT UNSIGNED NOT NULL,
    `postKey` VARCHAR(191) NOT NULL,
    `imageID` BIGINT UNSIGNED,
    `postComment` LONGTEXT,
    `postLike` BIGINT DEFAULT 0,
    `created_at` DATETIME(3) NULL,
    `userID` BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY (`postID`),
    CONSTRAINT `fk_posts_image` FOREIGN KEY (`imageID`) REFERENCES `images` (`imageID`) ON DELETE SET NULL ON UPDATE CASCADE,
    CONSTRAINT `fk_posts_category` FOREIGN KEY (`categoryID`) REFERENCES `categories` (`categoryID`),
    CONSTRAINT `fk_posts_user` FOREIGN KEY (`userID`) REFERENCES `users` (`userID`) ON DELETE CASCADE,
    CONSTRAINT `uni_posts_post_key` UNIQUE (`postKey`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Créer la table `sessionsUsers`
CREATE TABLE `sessionsUsers` (
    `sessionID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `userID` BIGINT UNSIGNED NOT NULL,
    `sessionToken` VARCHAR(191) NOT NULL UNIQUE,
    `expires_at` DATETIME NOT NULL,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`sessionID`),
    FOREIGN KEY (`userID`) REFERENCES `users` (`userID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Créer la table `sessionsAdmins`
CREATE TABLE `sessionsAdmins` (
    `sessionID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `adminID` BIGINT UNSIGNED NOT NULL,
    `sessionToken` VARCHAR(191) NOT NULL UNIQUE,
    `expires_at` DATETIME NOT NULL,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`sessionID`),
    FOREIGN KEY (`adminID`) REFERENCES `admins` (`adminID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Créer la table `sessionsGuests`
CREATE TABLE `sessionsGuests` (
    `sessionID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `guestID` BIGINT UNSIGNED NOT NULL,
    `sessionToken` VARCHAR(191) NOT NULL UNIQUE,
    `expires_at` DATETIME NOT NULL,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`sessionID`),
    FOREIGN KEY (`guestID`) REFERENCES `guests` (`guestID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- Commentaires
CREATE TABLE `comments` (
    `commentID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `userID` BIGINT UNSIGNED NOT NULL,
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `content` TEXT,
    `postID` BIGINT UNSIGNED NOT NULL,
    `status` ENUM('approuvé', 'en attente', 'rejeté') DEFAULT 'en attente',
    `visible` BOOLEAN DEFAULT TRUE,
    PRIMARY KEY (`commentID`),
    FOREIGN KEY (`userID`) REFERENCES `users` (`userID`),
    FOREIGN KEY (`postID`) REFERENCES `posts` (`postID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Likes sur les posts
CREATE TABLE `postsLikes` (
    `userID` BIGINT UNSIGNED NOT NULL,
    `postID` BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY (`userID`, `postID`),
    FOREIGN KEY (`userID`) REFERENCES `users` (`userID`) ON DELETE CASCADE,
    FOREIGN KEY (`postID`) REFERENCES `posts` (`postID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Likes sur les commentaires
CREATE TABLE `commentsLikes` (
    `userID` BIGINT UNSIGNED NOT NULL,
    `commentID` BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY (`userID`, `commentID`),
    FOREIGN KEY (`userID`) REFERENCES `users` (`userID`) ON DELETE CASCADE,
    FOREIGN KEY (`commentID`) REFERENCES `comments` (`commentID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Données du dashboard admin
CREATE TABLE `adminDashboardData` (
    `dashboardID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `adminID` BIGINT UNSIGNED NOT NULL,
    `totalUsers` BIGINT UNSIGNED DEFAULT 0,
    `totalPosts` BIGINT UNSIGNED DEFAULT 0,
    `totalComments` BIGINT UNSIGNED DEFAULT 0,
    `totalGuests` BIGINT UNSIGNED DEFAULT 0,
    `lastLogin` DATETIME(3) NULL,
    `generated_at` DATETIME(3) NULL,
    PRIMARY KEY (`dashboardID`),
    INDEX (`adminID`),
    FOREIGN KEY (`adminID`) REFERENCES `admins`(`adminID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- -------------------------------------------------------------------------
-- Créer l'image générique si elle n'existe pas encore
INSERT INTO images (url, filename, data)
SELECT 'default-image-url', 'default-image.jpg', NULL
FROM DUAL
WHERE NOT EXISTS (SELECT 1 FROM images WHERE filename = 'default-image.jpg');

-- Mettre à jour les posts sans image pour qu'ils pointent vers l'image générique
UPDATE posts
SET imageID = (SELECT imageID FROM images WHERE filename = 'default-image.jpg')
WHERE imageID IS NULL;

-- Suppression de l'image spécifique, par exemple, imageID = 3 (si elle existe)
DELETE FROM images WHERE imageID = 3;

-- ---- Images (profils et posts)
INSERT INTO images (url, filename) VALUES
   ('https://example.com/profile1.jpg', 'profile1.jpg'),
   ('https://example.com/profile2.jpg', 'profile2.jpg'),
   ('https://example.com/post1.jpg', 'post1.jpg');

-- Utilisateurs (avec imageID corrects = 1, 2)
INSERT INTO users (userUsername, userEmail, userPasswordHash, userProfilePicture) VALUES
   ('john_doe', 'john@example.com', 'hashed_password_1', 1),
   ('jane_doe', 'jane@example.com', 'hashed_password_2', 2);

-- Catégories
INSERT INTO categories (categoryName, categoryDescription) VALUES
   ('Hack', 'Tout sur la technologie'),
   ('Programmation', 'Art créatif et design');

-- Administrateurs
INSERT INTO admins (adminUsername, adminEmail, adminPasswordHash) VALUES
   ('admin1', 'admin1@example.com', 'adminpass1'),
   ('admin2', 'admin2@example.com', 'adminpass2');

-- Guests (aucun besoin de sessionID ici, car sessionsGuests utilise guestID, pas l’inverse)
INSERT INTO guests () VALUES (), ();

-- Sessions Users (userID = 1 et 2)
INSERT INTO sessionsUsers (userID, sessionToken, expires_at) VALUES
   (1, 'token1234', NOW() + INTERVAL 1 DAY),
   (2, 'token4567', NOW() + INTERVAL 1 DAY);

-- Sessions Admins (adminID = 1 et 2)
INSERT INTO sessionsAdmins (adminID, sessionToken, expires_at) VALUES
   (1, 'admintoken1', NOW() + INTERVAL 1 DAY),
   (2, 'admintoken2', NOW() + INTERVAL 1 DAY);

-- Sessions Guests (guestID = 1 et 2)
INSERT INTO sessionsGuests (guestID, sessionToken, expires_at) VALUES
   (1, 'guesttoken1', NOW() + INTERVAL 1 DAY),
   (2, 'guesttoken2', NOW() + INTERVAL 1 DAY);

-- Posts (imageID = 3 pour le premier post)
INSERT INTO posts (categoryID, postKey, imageID, postComment, userID) VALUES
   (1, 'hack-post-1', 3, 'Découvrez ce nouveau gadget !', 1),
   (2, 'prog-post-1', 2, 'Une peinture que j’ai faite ce week-end.', 2);

-- Commentaires
INSERT INTO comments (userID, postID, content, status, visible) VALUES
   (2, 1, 'Ça a l’air génial !', 'approuvé', TRUE),
   (1, 2, 'Beau travail !', 'approuvé', TRUE);

-- Likes sur les posts
INSERT INTO postsLikes (userID, postID) VALUES
   (1, 2), -- User 1 aime le post 2
   (2, 1); -- User 2 aime le post 1

-- Likes sur les commentaires
INSERT INTO commentsLikes (userID, commentID) VALUES
   (1, 1), -- User 1 aime le commentaire 1
   (2, 2); -- User 2 aime le commentaire 2

-- Données dashboard admin (adminID = 1)
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
