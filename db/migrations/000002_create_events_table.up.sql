CREATE TABLE IF NOT EXISTS events
(
   id                       INTEGER  NOT NULL PRIMARY KEY
  ,equipment_id             VARCHAR(20)  NOT NULL
  ,sighting_date            TIMESTAMP  NOT NULL
  ,sighting_event_code      INTEGER  NOT NULL
  ,reporting_railroad_scac  VARCHAR(4) NOT NULL
  ,posting_date             TIMESTAMP  NOT NULL
  ,from_mark_id             VARCHAR(5)  NOT NULL
  ,load_empty_status        VARCHAR(1) NOT NULL
  ,sighting_claim_code      VARCHAR(1) NOT NULL
  ,sighting_event_code_text VARCHAR(19) NOT NULL
  ,train_id                 VARCHAR(7)
  ,train_alpha_code         VARCHAR(4) NOT NULL
  ,location_id              INT  NOT NULL
  ,waybill_id               INT  NOT NULL
);