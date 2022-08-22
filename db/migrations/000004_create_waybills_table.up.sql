CREATE TABLE IF NOT EXISTS waybills
(
  id                      INTEGER  NOT NULL PRIMARY KEY
  ,equipment_id           VARCHAR(10) NOT NULL
  ,waybill_date           TIMESTAMP NOT NULL
  ,waybill_number         INTEGER  NOT NULL
  ,created_date           TIMESTAMP  NOT NULL
  ,billing_road_mark_name VARCHAR(4) NOT NULL
  ,waybill_source_code    VARCHAR(1) NOT NULL
  ,load_empty_status      VARCHAR(1) NOT NULL
  ,origin_mark_name       VARCHAR(4) NOT NULL
  ,destination_mark_name  VARCHAR(4) NOT NULL
  ,sending_road_mark      VARCHAR(4) NOT NULL
  ,bill_of_lading_number  VARCHAR(30) NOT NULL
  ,bill_of_lading_date    TIMESTAMP NOT NULL
  ,equipment_weight       INTEGER  NOT NULL
  ,tare_weight            INTEGER  NOT NULL
  ,allowable_weight       INTEGER  NOT NULL
  ,dunnage_weight         INTEGER  NOT NULL
  ,equipment_weight_code  VARCHAR(1)
  ,commodity_code         INTEGER  NOT NULL
  ,commodity_description  VARCHAR(15) NOT NULL
  ,origin_id              VARCHAR(5)
  ,destination_id         INTEGER  NOT NULL
  ,routes                 VARCHAR(129) NOT NULL
  ,parties                VARCHAR(581) NOT NULL
)