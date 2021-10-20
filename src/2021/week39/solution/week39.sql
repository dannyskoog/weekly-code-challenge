CREATE TABLE IF NOT EXISTS car (id INT GENERATED ALWAYS AS IDENTITY, brand VARCHAR NOT NULL, sold bool DEFAULT false NOT NULL, updated_at timestamp);
TRUNCATE car;
INSERT INTO car (brand) VALUES ('Toyota');
INSERT INTO car (brand) VALUES ('Mercedes');
​
CREATE OR REPLACE FUNCTION public.updated_at_trigger_func()
    RETURNS TRIGGER AS $$
BEGIN
    UPDATE car SET updated_at = clock_timestamp() WHERE id = NEW.id;
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;
​
DROP TRIGGER IF EXISTS updated_at_trigger
  ON public.car;
CREATE CONSTRAINT TRIGGER updated_at_trigger
AFTER UPDATE ON public.car
    DEFERRABLE INITIALLY DEFERRED
    FOR EACH ROW
    WHEN (OLD.updated_at IS NOT DISTINCT FROM NEW.updated_at)
    EXECUTE PROCEDURE public.updated_at_trigger_func();
​
CREATE TEMP TABLE IF NOT EXISTS tt (vc timestamp);
TRUNCATE TABLE tt;
INSERT INTO tt VALUES (clock_timestamp());
​
BEGIN;
SELECT pg_sleep(1);
UPDATE car SET sold = true WHERE id = 1;
SELECT pg_sleep(1);
UPDATE car SET sold = true WHERE id = 2;
SELECT pg_sleep(1);
COMMIT;
                       
SELECT 
    (SELECT to_char(vc, 'HH24:MI:SS') FROM tt LIMIT 1) AS "Before transaction time (UTC)",
    to_char(updated_at, 'HH24:MI:SS') AS "Record updated at time (UTC)",
    to_char(clock_timestamp(), 'HH24:MI:SS') AS "After transaction time (UTC)"
FROM car;

/* or equivalent playground solution: https://www.db-fiddle.com/f/qKNT7M1Mkt643v41HHsNSK/0 */
