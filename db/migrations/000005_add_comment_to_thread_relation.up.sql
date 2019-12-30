ALTER TABLE comments ADD `thread_id` varchar(255) NOT NULL;

ALTER TABLE comments ADD INDEX `thread_id` (`thread_id`);

ALTER TABLE comments ADD CONSTRAINT `comments_ibfk_1` FOREIGN KEY (`thread_id`) REFERENCES `threads` (`id`) ON DELETE CASCADE;