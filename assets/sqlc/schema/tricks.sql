CREATE TABLE IF NOT EXISTS public.tricks
(
    id BIGINT NOT NULL DEFAULT nextval('tricks_id_seq'::regclass),
    name character varying(100) COLLATE pg_catalog."default",
    style integer,
    power boolean,
    CONSTRAINT tricks_pkey PRIMARY KEY (id),
    CONSTRAINT tricks_style_check CHECK (style <= 10)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.tricks
    OWNER to postgres;