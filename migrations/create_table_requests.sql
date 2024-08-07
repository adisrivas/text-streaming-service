CREATE TABLE `requests` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `provider` TINYINT(4) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `user_id` INT NOT NULL COMMENT '0->system generated',
    `prompt` TEXT NOT NULL
);