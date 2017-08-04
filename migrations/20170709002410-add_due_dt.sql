
-- +migrate Up
ALTER TABLE reminders
ADD due_dt DATETIME;

-- +migrate Down
CREATE TABLE t1_backup AS SELECT id,content,chat_id,created FROM reminders;
DROP TABLE reminders;
ALTER TABLE t1_backup RENAME TO reminders;
