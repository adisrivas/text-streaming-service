CREATE TABLE `third_provider` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `type` TINYINT(4) NOT NULL COMMENT '0->outgoing, 1->incoming',
    `created_at` INT(11) NOT NULL,
    `request_id` INT NOT NULL,
    `user_id` INT NOT NULL COMMENT '0->system generated',
    `is_available` TINYINT(4) NOT NULL DEFAULT '1',
);