/*
 * First Schema setup for power data management
 * 
 */


/* 
 * A basic schema for power data
 */

CREATE SCHEMA IF NOT EXISTS power;


/* for playing around we have one table with a string as key and another string as value
 *
 * TSN is primary key
 * key is not null (unique index to be created later)
 * value is text (to be replaced to JSON data or others
 */
CREATE TABLE power.data (
    tsn bigint primary key,
    key text   unique,
    value text
);

/* Functions for put and get
 *
 * put always creates new entry with new TSN
 * (most simple approach, can be improved)
 * 
 * get needs to select the latest entry per key (subselect needed)
 */

CREATE OR REPLACE FUNCTION power.put( _key text, _value text) RETURNS VOID AS $$
   BEGIN

     INSERT INTO power.data( tsn, key, value) 
     VALUES( nextval( 'clock.tsn' ), _key, _value ); 

   END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION power.get( _key text) 
     RETURNS TABLE( ret_tsn bigint, ret_key text, ret_value text ) AS $$
   BEGIN

     RETURN QUERY
       SELECT tsn, key, value FROM power.data 
           WHERE tsn = (select max(tsn) from power.data where key = _key ); 

   END;
$$ LANGUAGE plpgsql;


/*
 * we want to create a clock as a singleton. Since the clock is basically a sequence,
 * the global clock is defined as a schema containing a sequence as the base structure
 */

CREATE SCHEMA IF NOT EXISTS clock;
CREATE SEQUENCE clock.tsn;


CREATE OR REPLACE FUNCTION clock.new_tsn() RETURNS BIGINT AS $$
   BEGIN
    
     RETURN nextval( 'clock.tsn' );

   END;
$$ LANGUAGE plpgsql;




