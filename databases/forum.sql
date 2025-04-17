-- Disable foreign key checks to allow dropping tables
SET foreign_key_checks = 0;

-- Drop the tables if they exist
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

-- Re-enable foreign key checks
SET foreign_key_checks = 1;

-- Create tables starting from here, similar to the previous script...
-- Table for images
CREATE TABLE `images` (
    `imageID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `url` VARCHAR(255),
    `filename` VARCHAR(255),
    `data` BLOB,
    PRIMARY KEY (`imageID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Categories of posts
CREATE TABLE `categories` (
    `categoryID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `categoryName` VARCHAR(255),
    `categoryDescription` TEXT,
    PRIMARY KEY (`categoryID`),
    UNIQUE (`categoryName`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Users table
CREATE TABLE `users` (
    `userID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `userUsername` VARCHAR(255) NOT NULL,
    `userEmail` VARCHAR(255) NOT NULL,
    `userPasswordHash` VARCHAR(255) NOT NULL,
    `userProfilePicture` BIGINT UNSIGNED NOT NULL,
    `createdAt` DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`userID`),
    UNIQUE (`userUsername`),
    UNIQUE (`userEmail`),
    FOREIGN KEY (`userProfilePicture`) REFERENCES `images` (`imageID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Create the admins table
CREATE TABLE `admins` (
    `adminID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `adminUsername` VARCHAR(255) NOT NULL,
    `adminPasswordHash` VARCHAR(255) NOT NULL,
    `adminEmail` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`adminID`),
    UNIQUE (`adminUsername`),
    UNIQUE (`adminEmail`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Create the guests table
CREATE TABLE `guests` (
    `guestID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `createdAt` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `lastVisitedAt` DATETIME DEFAULT NULL,
    PRIMARY KEY (`guestID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Posts table
CREATE TABLE `posts` (
    `postID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `categoryID` BIGINT UNSIGNED NOT NULL,
    `postKey` VARCHAR(191) NOT NULL,
    `imageID` BIGINT UNSIGNED,
    `postComment` LONGTEXT,
    `postLike` BIGINT DEFAULT 0,
    `createdAt` DATETIME(3) NULL,
    `userID` BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY (`postID`),
    CONSTRAINT `fk_posts_image` FOREIGN KEY (`imageID`) REFERENCES `images` (`imageID`) ON DELETE SET NULL ON UPDATE CASCADE,
    CONSTRAINT `fk_posts_category` FOREIGN KEY (`categoryID`) REFERENCES `categories` (`categoryID`),
    CONSTRAINT `fk_posts_user` FOREIGN KEY (`userID`) REFERENCES `users` (`userID`) ON DELETE CASCADE,
    CONSTRAINT `uni_posts_post_key` UNIQUE (`postKey`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Create sessionsUsers table
CREATE TABLE `sessionsUsers` (
    `sessionID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `userID` BIGINT UNSIGNED NOT NULL,
    `sessionToken` VARCHAR(191) NOT NULL UNIQUE,
    `expiresAt` DATETIME NOT NULL,
    `createdAt` DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`sessionID`),
    FOREIGN KEY (`userID`) REFERENCES `users` (`userID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Create sessionsAdmins table
CREATE TABLE `sessionsAdmins` (
    `sessionID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `adminID` BIGINT UNSIGNED NOT NULL,
    `sessionToken` VARCHAR(191) NOT NULL UNIQUE,
    `expiresAt` DATETIME NOT NULL,
    `createdAt` DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`sessionID`),
    FOREIGN KEY (`adminID`) REFERENCES `admins` (`adminID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Create sessionsGuests table
CREATE TABLE `sessionsGuests` (
    `sessionID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `guestID` BIGINT UNSIGNED NOT NULL,
    `sessionToken` VARCHAR(191) NOT NULL UNIQUE,
    `expiresAt` DATETIME NOT NULL,
    `createdAt` DATETIME DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`sessionID`),
    FOREIGN KEY (`guestID`) REFERENCES `guests` (`guestID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Comments table
CREATE TABLE `comments` (
    `commentID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `userID` BIGINT UNSIGNED NOT NULL,
    `createdAt` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `content` TEXT,
    `postID` BIGINT UNSIGNED NOT NULL,
    `status` ENUM('approved', 'pending', 'rejected') DEFAULT 'pending',
    `visible` BOOLEAN DEFAULT TRUE,
    `commentLike` BIGINT DEFAULT 0,
    PRIMARY KEY (`commentID`),
    FOREIGN KEY (`userID`) REFERENCES `users` (`userID`),
    FOREIGN KEY (`postID`) REFERENCES `posts` (`postID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Likes on posts
CREATE TABLE `postsLikes` (
    `userID` BIGINT UNSIGNED NOT NULL,
    `postID` BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY (`userID`, `postID`),
    FOREIGN KEY (`userID`) REFERENCES `users` (`userID`) ON DELETE CASCADE,
    FOREIGN KEY (`postID`) REFERENCES `posts` (`postID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Likes on comments
CREATE TABLE `commentsLikes` (
    `userID` BIGINT UNSIGNED NOT NULL,
    `commentID` BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY (`userID`, `commentID`),
    FOREIGN KEY (`userID`) REFERENCES `users` (`userID`) ON DELETE CASCADE,
    FOREIGN KEY (`commentID`) REFERENCES `comments` (`commentID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Admin dashboard data table
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
-- Insert a default image if it does not exist
INSERT INTO images (url, filename, data)
SELECT 'default-image-url', 'default-image.jpg', NULL
FROM DUAL
WHERE NOT EXISTS (SELECT 1 FROM images WHERE filename = 'default-image.jpg');

-- Update posts without images to point to the default image
UPDATE posts
SET imageID = (SELECT imageID FROM images WHERE filename = 'default-image.jpg')
WHERE imageID IS NULL;

-- Delete a specific image, for example, imageID = 3 (if it exists)
DELETE FROM images WHERE imageID = 3;

-- ---- Insert sample images (profile and post images)
INSERT INTO images (url, filename) VALUES
   ('https://example.com/profile1.jpg', 'profile1.jpg'),
   ('https://example.com/profile2.jpg', 'profile2.jpg'),
   ('https://example.com/post1.jpg', 'post1.jpg');

-- Insert sample users (with correct imageIDs = 1, 2)
INSERT INTO users (userUsername, userEmail, userPasswordHash, userProfilePicture) VALUES
   ('john_doe', 'john@example.com', '$2a$10$f2GTlyF/9n.kAqYjRWH2GeQ9Iw62SxAd29IH6OaXhxLAbMM/FxYyO', 1),
   ('jane_doe', 'jane@example.com', 'hashed_password_2', 2);

-- Insert sample categories
INSERT INTO categories (categoryName, categoryDescription) VALUES
   ('hack', 'All about technology'),
   ('prog', 'Creative art and design');

-- Insert sample administrators
INSERT INTO admins (adminUsername, adminEmail, adminPasswordHash) VALUES
   ('admin1', 'admin1@example.com', 'adminpass1'),
   ('admin2', 'admin2@example.com', 'adminpass2');

-- Insert sample guests (no need for sessionID here, since sessionsGuests uses guestID)
INSERT INTO guests () VALUES (), ();

-- Insert sample session data for users (userID = 1 and 2)
INSERT INTO sessionsUsers (userID, sessionToken, expiresAt) VALUES
   (1, 'token1234', NOW() + INTERVAL 1 DAY),
   (2, 'token4567', NOW() + INTERVAL 1 DAY);

-- Insert sample session data for administrators (adminID = 1 and 2)
INSERT INTO sessionsAdmins (adminID, sessionToken, expiresAt) VALUES
   (1, 'admintoken1', NOW() + INTERVAL 1 DAY),
   (2, 'admintoken2', NOW() + INTERVAL 1 DAY);

-- Insert sample session data for guests (guestID = 1 and 2)
INSERT INTO sessionsGuests (guestID, sessionToken, expiresAt) VALUES
   (1, 'guesttoken1', NOW() + INTERVAL 1 DAY),
   (2, 'guesttoken2', NOW() + INTERVAL 1 DAY);

-- Insert sample posts (imageID = 3 for the first post)
INSERT INTO posts (categoryID, postKey, imageID, postComment, userID) VALUES
   (1, 'hack-post-1', 3, 'Check out this new gadget!', 1),
   (2, 'prog-post-1', 2, 'A painting I made this weekend.', 2);

-- Insert sample comments
INSERT INTO comments (userID, postID, content, status, visible) VALUES
   (2, 1, 'That looks amazing!', 'approved', TRUE),
   (1, 2, 'Great work!', 'approved', TRUE);

-- Insert sample likes on posts
INSERT INTO postsLikes (userID, postID) VALUES
   (1, 2), -- User 1 likes post 2
   (2, 1); -- User 2 likes post 1

-- Insert sample likes on comments
INSERT INTO commentsLikes (userID, commentID) VALUES
   (1, 1), -- User 1 likes comment 1
   (2, 2); -- User 2 likes comment 2

-- Insert admin dashboard data (adminID = 1)
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
