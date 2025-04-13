-- +goose Up
-- On verification status
CREATE OR REPLACE FUNCTION update_status_on_verification()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.verification_status = 3 AND NEW.reception_start > NOW() THEN
        NEW.status := 3;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_status
BEFORE UPDATE ON tenders
FOR EACH ROW
WHEN (OLD.verification_status IS DISTINCT FROM NEW.verification_status)
EXECUTE FUNCTION update_status_on_verification();

-- Reception status
-- CREATE OR REPLACE FUNCTION schedule_status_update()
-- RETURNS VOID AS $$
-- BEGIN
--     INSERT INTO cron.job (schedule, command)
--     VALUES ('* * * * *', 'UPDATE tenders SET status = 4 WHERE reception_start <= NOW() AND verification_status = 3');
-- END;
-- $$ LANGUAGE plpgsql;

-- SELECT schedule_status_update();

CREATE OR REPLACE FUNCTION update_tender_receptionnotstarted()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.verification_status = 3 AND NEW.reception_start > NOW() THEN
        NEW.status := 3;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_receptionnotstarted
BEFORE UPDATE ON tenders
FOR EACH ROW
WHEN (OLD.verification_status IS DISTINCT FROM NEW.verification_status)
EXECUTE FUNCTION update_tender_receptionnotstarted();

-- +goose Down
-- On verification status
DROP TRIGGER IF EXISTS trigger_update_status ON tenders;
DROP FUNCTION IF EXISTS update_status_on_verification;

-- Reception status
DROP TRIGGER IF EXISTS trigger_receptionnotstarted ON tenders;
DROP FUNCTION IF EXISTS update_tender_receptionnotstarted;
-- DROP FUNCTION IF EXISTS schedule_status_update;

-- DELETE FROM cron.job WHERE command LIKE 'UPDATE tenders SET status = 4 WHERE reception_start <= NOW() AND verification_status = 3';