-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE `company` ADD COLUMN `status` VARCHAR(32) NOT NULL DEFAULT 'active', ADD INDEX `idx_status`(`status`);
ALTER TABLE `user` ADD COLUMN `status` VARCHAR(32) NOT NULL DEFAULT 'active', ADD INDEX `idx_status`(`status`);
ALTER TABLE `role` ADD COLUMN `status` VARCHAR(32) NOT NULL DEFAULT 'active', ADD INDEX `idx_status`(`status`);
ALTER TABLE `permission` ADD COLUMN `status` VARCHAR(32) NOT NULL DEFAULT 'active', ADD INDEX `idx_status`(`status`);
ALTER TABLE `status` ADD COLUMN `status` VARCHAR(32) NOT NULL DEFAULT 'active', ADD INDEX `idx_status`(`status`);
ALTER TABLE `section` ADD COLUMN `status` VARCHAR(32) NOT NULL DEFAULT 'active', ADD INDEX `idx_status`(`status`);
ALTER TABLE `status_group` ADD COLUMN `status` VARCHAR(32) NOT NULL DEFAULT 'active', ADD INDEX `idx_status`(`status`);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE `company` DROP COLUMN `status`;
ALTER TABLE `user` DROP COLUMN `status`;
ALTER TABLE `role` DROP COLUMN `status`;
ALTER TABLE `permission` DROP COLUMN `status`;
ALTER TABLE `status` DROP COLUMN `status`;
ALTER TABLE `section` DROP COLUMN `status`;
ALTER TABLE `status_group` DROP COLUMN `status`;