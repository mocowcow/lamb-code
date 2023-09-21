-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- 主機： 127.0.0.1
-- 產生時間： 2023-09-21 17:31:16
-- 伺服器版本： 10.4.28-MariaDB
-- PHP 版本： 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 資料庫： `problem`
--

-- --------------------------------------------------------

--
-- 資料表結構 `problems`
--

CREATE TABLE `problems` (
  `id` int(11) NOT NULL,
  `title` varchar(50) NOT NULL,
  `content` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- 傾印資料表的資料 `problems`
--

INSERT INTO `problems` (`id`, `title`, `content`) VALUES
(1, 'a加b', '輸入兩個整數a和b，回傳a+b的總和。'),
(2, 'a乘b', '輸入兩個整數a和b，回傳a*b的總和。');

-- --------------------------------------------------------

--
-- 資料表結構 `testcases`
--

CREATE TABLE `testcases` (
  `id` int(11) NOT NULL,
  `problem_id` int(11) NOT NULL,
  `input` text NOT NULL,
  `output` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- 傾印資料表的資料 `testcases`
--

INSERT INTO `testcases` (`id`, `problem_id`, `input`, `output`) VALUES
(1, 1, '1 1', '2'),
(2, 1, '3 5', '8'),
(3, 2, '1 1', '1'),
(4, 2, '2 5', '10');

--
-- 已傾印資料表的索引
--

--
-- 資料表索引 `problems`
--
ALTER TABLE `problems`
  ADD PRIMARY KEY (`id`);

--
-- 資料表索引 `testcases`
--
ALTER TABLE `testcases`
  ADD PRIMARY KEY (`id`),
  ADD KEY `problem_id` (`problem_id`);

--
-- 在傾印的資料表使用自動遞增(AUTO_INCREMENT)
--

--
-- 使用資料表自動遞增(AUTO_INCREMENT) `problems`
--
ALTER TABLE `problems`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- 使用資料表自動遞增(AUTO_INCREMENT) `testcases`
--
ALTER TABLE `testcases`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
