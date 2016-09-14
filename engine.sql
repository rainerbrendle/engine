/*
 * we want to create a clock as a singleton. Since the clock is basically a sequence,
 * the global clock is defined as a schema containing a sequence as the base structure
 */

CREATE SCHEMA IF NOT EXISTS clock;
CREATE SEQUENCE clock.tsn;


CREATE OR REPLACE FUNCTION clock.new_tsn() RETURNS BIGINT AS $$
   BEGIN
    
     RETURN nextval( 'clock.tsn' );

   END
$$ LANGUAGE plpgsql;





