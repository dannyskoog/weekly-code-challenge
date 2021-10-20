/**
 * Code challenge week 39, 2021 (language: PostgreSQL, playground: https://www.db-fiddle.com/f/wTNiMarLAbNkN8dFT1pvUt/1)
 *
 * Description:
 *     An engineer at a car company gets the assignment to build a system that can help persist sales digitally. The engineer therefore
 *     starts looking at the database layer and quickly realises that it would be a good idea to keep track of when records are updated.
 *     And since the engineer is quite well-wandered within PostgreSQL, using a trigger to achieve this seems reasonable. One of the system
 *     requirements, though, is that cars are supposed to be marked as sold in batches and that if anything goes wrong with the batch update
 *     all changes need to be rolled back. This seamlessly brings the engineer’s mind to transactions. So the engineer gets to work and
 *     creates an early database layer POC that simulates marking cars as sold in batch, also taking into account some other logic taking
 *     time in between the updates. But upon executing the sql statements the engineer gets extremely confused. The thing is that the
 *     engineer expected the updated_at value (as seen as “Record updated at time (UTC)“) to hold a timestamp of when a record update was
 *     applied to the database. But that doesn’t seem to be the case.
 *
 * Questions:
 *     1. Why isn’t the updated_at timestamp matching the time of when a record update is applied to the database?
 *     2. How can the sql statements be adjusted to result in an updated_at timestamp matching the time of when a record update is applied to the database?
 */

CREATE TABLE IF NOT EXISTS car (id INT GENERATED ALWAYS AS IDENTITY, brand VARCHAR NOT NULL, sold bool DEFAULT false NOT NULL, updated_at timestamp);
TRUNCATE car;
INSERT INTO car (brand) VALUES ('Toyota');
INSERT INTO car (brand) VALUES ('Mercedes');

CREATE OR REPLACE FUNCTION public.updated_at_trigger_func()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER updated_at_trigger
BEFORE UPDATE ON public.car
                  FOR EACH ROW
                  EXECUTE PROCEDURE public.updated_at_trigger_func();

CREATE TEMP TABLE IF NOT EXISTS tt (vc timestamp);
TRUNCATE TABLE tt;
INSERT INTO tt VALUES (now());

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
    to_char(now(), 'HH24:MI:SS') AS "After transaction time (UTC)"
FROM car;
