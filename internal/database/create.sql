CREATE TABLE IF NOT EXISTS "party-data"."party-data"
(
    id integer NOT NULL,
    persons integer[],
    persons_count integer,
    average_amount double precision,
    total_amount integer,
    CONSTRAINT "party-data_pkey" PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS "party-data"."party-data"
    OWNER to postgres;

COMMENT ON COLUMN "party-data"."party-data".persons
    IS 'ids of ''persons'' table';


CREATE TABLE IF NOT EXISTS "party-data".persons
(
    id integer NOT NULL,
    name text COLLATE pg_catalog."default" NOT NULL,
    spent integer,
    participants integer DEFAULT 1,
    balance double precision,
    indepted_to json,
    CONSTRAINT persons_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS "party-data".persons
    OWNER to postgres;