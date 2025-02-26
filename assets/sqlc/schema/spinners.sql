CREATE TABLE public.tricks (
    id BIGINT NOT NULL DEFAULT nextval('tricks_id_seq'::regclass),
    name character varying(100) COLLATE pg_catalog."default",
    style integer,
    power boolean,
    CONSTRAINT tricks_pkey PRIMARY KEY (id),
    CONSTRAINT tricks_style_check CHECK (style <= 10)
)

CREATE TABLE IF NOT EXISTS spinners.spinners
(
    UserID BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    Name VARCHAR(100) NOT NULL,
    Email VARCHAR(254) UNIQUE NOT NULL,
    Provider VARCHAR(32) NOT NULL,
    Tricks VARCHAR(100) REFERENCES tricks(name),
    ExpiresAt TIMESTAMP,
    AccessToken VARCHAR(255) NOT NULL,
    AccessTokenSecret VARCHAR(255),
    RefreshToken VARCHAR(255) NOT NULL,
    CONSTRAINT  spinners_pkey PRIMARY KEY (UserID),
    CONSTRAINT spinners_Email_key UNIQUE (Email)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS spinners.spinners
    OWNER to postgres;