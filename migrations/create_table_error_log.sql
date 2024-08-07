CREATE TABLE `error_log` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `error` TEXT NOT NULL,
    `request_id` INT NOT NULL
);