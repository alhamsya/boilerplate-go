-- +migrate Down
DROP TABLE IF EXISTS `movie_mst_user`;

-- +migrate Up
CREATE TABLE IF NOT EXISTS `movie_mst_user` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `email` varchar(255) NOT NULL,
    `password` text NOT NULL,
    `status` int(100) NOT NULL DEFAULT 0,
    `created_at` datetime NOT NULL,
    `created_by` varchar(255) NOT NULL,
    `updated_at` datetime NULL,
    `updated_by` varchar(255) NULL,
    `deleted_at` datetime NULL,
    `deleted_by` varchar(255) NULL,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `idx-movie_mst_user-email`(`email`)
)ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;


