CREATE TABLE IF NOT EXISTS locations
(
  id         INTEGER  NOT NULL PRIMARY KEY
  ,city      VARCHAR(9) NOT NULL
  ,city_long VARCHAR(20) NOT NULL
  ,station   VARCHAR(17) NOT NULL
  ,fsac      INTEGER  NOT NULL
  ,scac      VARCHAR(4) NOT NULL
  ,splc      INTEGER  NOT NULL
  ,state     VARCHAR(2) NOT NULL
  ,time_zone VARCHAR(2) NOT NULL
  ,longitude NUMERIC(11,6) NOT NULL
  ,latitude  NUMERIC(9,6) NOT NULL
  ,country   VARCHAR(2) NOT NULL
);