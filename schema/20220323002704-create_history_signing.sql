-- +migrate Down
DROP TABLE IF EXISTS `movie_hst_signing`;

-- +migrate Up
CREATE TABLE IF NOT EXISTS `movie_hst_signing` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `email` varchar(255) NOT NULL,
    `created_at` datetime NOT NULL,
    `created_by` varchar(255) NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `idx-movie_hst_signing-email`(`email`)
)ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;
