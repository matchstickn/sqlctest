CREATE SEQUENCE IF NOT EXISTS tricks_id_sq 
    START WITH 1
    INCREMENT BY 1;

CREATE TABLE IF NOT EXISTS public.tricks (
    id BIGINT NOT NULL DEFAULT nextval('tricks_id_seq'),
    name VARCHAR(100) COLLATE pg_catalog."default",
    style integer,
    power boolean,
    CONSTRAINT tricks_pkey PRIMARY KEY (id),
    CONSTRAINT tricks_style_check CHECK (style <= 10)
);

CREATE TABLE IF NOT EXISTS spinners (
    UserID BIGINT GENERATED ALWAYS AS IDENTITY,
    Name VARCHAR(100) NOT NULL,
    Email VARCHAR(254) UNIQUE NOT NULL,
    Provider VARCHAR(32) NOT NULL,
    Tricks BIGINT[],
    ExpiresAt TIMESTAMP,
    AccessToken VARCHAR(255) NOT NULL,
    AccessTokenSecret VARCHAR(255),
    RefreshToken VARCHAR(255) NOT NULL,
    CONSTRAINT spinners_fkey_tricks FOREIGN KEY (Tricks) REFERENCES tricks(id),
    CONSTRAINT spinners_pkey PRIMARY KEY (UserID),
    CONSTRAINT spinners_Email_key UNIQUE (Email)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS spinners
    OWNER to postgres;