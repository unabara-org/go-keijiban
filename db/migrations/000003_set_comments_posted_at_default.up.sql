ALTER TABLE comments MODIFY  `posted_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP  ON UPDATE CURRENT_TIMESTAMP;
