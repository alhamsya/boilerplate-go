
-- +migrate Up
CREATE TABLE IF NOT EXISTS `hst_log` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `endpoint` text,
    `request` text NOT NULL,
    `response` text NOT NULL,
    `source_data` varchar(255) NOT NULL,
    `created_at` datetime NOT NULL,
    `created_by` varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
    )ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;

-- +migrate Down
DROP TABLE `hst_log`;
