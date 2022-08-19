CREATE TABLE IF NOT EXISTS equipment
(
   id               INTEGER  NOT NULL PRIMARY KEY
  ,customer         VARCHAR(100) NOT NULL
  ,fleet            VARCHAR(100) NOT NULL
  ,equipment_id     VARCHAR(100)  NOT NULL
  ,equipment_status VARCHAR(1) NOT NULL
  ,date_added       TIMESTAMPTZ  NOT NULL
  ,date_removed     TIMESTAMPTZ
);