-- CREATE USER tidybeaver WITH PASSWORD 'tidybeaver';
CREATE DATABASE "TidyBeaverLogs"
    WITH
    OWNER = tidybeaver
    ENCODING = 'UTF8'
   LC_COLLATE = 'en_US.UTF-8'
   LC_CTYPE = 'en_US.UTF-8'
    LOCALE_PROVIDER = 'libc'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;
--GRANT ALL ON DATABASE "TidyBeaverLogs" TO tidybeaver WITH GRANT OPTION;
