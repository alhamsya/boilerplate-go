-- +migrate Down
DROP TABLE IF EXISTS `gorp_migrations`;

-- +migrate Up
CREATE TABLE IF NOT EXISTS `gorp_migrations` (
    `id` varchar(255) NOT NULL,
    `applied_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
)ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;
