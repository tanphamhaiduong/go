-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE `role_feature` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `role_id` bigint(20) NOT NULL DEFAULT 0,
  `feature_id` bigint(20) NOT NULL DEFAULT 0,
  `created_by` varchar(64) NOT NULL DEFAULT '',
  `updated_by` varchar(64) NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_role_id_feature_id` (`role_id`,`feature_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS `role_feature`
