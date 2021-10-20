CREATE DATABASE IF NOT EXISTS gochat_db;
USE gochat_db;

CREATE TABLE IF NOT EXISTS comment (
  id          INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  name        VARCHAR(256) NOT NULL,
  context     TEXT NOT NULL,
  created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO comment values (1, 'Winning Ticker', 'ウィニングチケットは1996年のダービー馬で、柴田政人を男にした。', '1993-05-30 15:30:00', '1993-05-30 15:35:00');
INSERT INTO comment values (2, 'Jungle Pocket', 'ジャングルポケットは2001年のダービー馬でペリエを背にジャパンカップも勝つ。', '2001-05-27 15:30:00', '1993-05-30 15:36:00');

CREATE TABLE IF NOT EXISTS `session` (
  `session_key` char(64) NOT NULL,
  `session_data` blob,
  `session_expiry` int(11) unsigned NOT NULL,
  `created_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `updated_at`  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`session_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;