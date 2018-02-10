-- phpMyAdmin SQL Dump
-- version 4.7.7
-- https://www.phpmyadmin.net/
--
-- Host: db
-- Generation Time: 2018 年 2 月 10 日 13:20
-- サーバのバージョン： 10.1.30-MariaDB-1~jessie
-- PHP Version: 7.1.9

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `eniwa`
--

-- --------------------------------------------------------

--
-- テーブルの構造 `groups`
--

CREATE TABLE `groups` (
  `id` int(11) NOT NULL,
  `group_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `start` datetime NOT NULL,
  `dead` datetime NOT NULL,
  `state` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- テーブルの構造 `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `password` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- テーブルの構造 `user_groups`
--

CREATE TABLE `user_groups` (
  `user_id` int(11) NOT NULL,
  `group_id` int(11) NOT NULL,
  `goal_price` int(11) NOT NULL,
  `current_price` int(11) NOT NULL,
  `goal_desc` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `join_flag` tinyint(4) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `groups`
--
ALTER TABLE `groups`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user_groups`
--
ALTER TABLE `user_groups`
  ADD KEY `user_group_fk_1` (`user_id`),
  ADD KEY `user_group_fk_2` (`group_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `groups`
--
ALTER TABLE `groups`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- ダンプしたテーブルの制約
--

--
-- テーブルの制約 `user_groups`
--
ALTER TABLE `user_groups`
  ADD CONSTRAINT `user_group_fk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `user_group_fk_2` FOREIGN KEY (`group_id`) REFERENCES `groups` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
