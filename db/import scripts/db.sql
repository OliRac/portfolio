-- Database: portfolio_dev

-- DROP DATABASE portfolio_dev;

CREATE DATABASE portfolio_dev
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'English_United States.1252'
    LC_CTYPE = 'English_United States.1252'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

COMMENT ON DATABASE portfolio_dev
    IS 'Will mostly contain strings of text to be used on my portfolio.';