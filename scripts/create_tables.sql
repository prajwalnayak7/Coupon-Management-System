Use cms;

CREATE TABLE `coupon` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `code` varchar(255) UNIQUE NOT NULL,
  `expiry_date` datetime,
  `availability_count` int,
  `product_id` varchar(255),
  `promo_type` varchar(255),
  `discount_fixed` decimal,
  `discount_variable` decimal,
  `valid` boolean,
  `created_at` datetime DEFAULT (now()),
  `updated_at` datetime DEFAULT (now())
);

CREATE TABLE `order` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `coupon_id` int,
  `client_id` int UNIQUE NOT NULL,
  `status` varchar(255),
  `created_at` varchar(255) DEFAULT (now()) COMMENT 'When order created'
);

ALTER TABLE `order` ADD FOREIGN KEY (`coupon_id`) REFERENCES `coupon` (`id`);
