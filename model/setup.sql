USE mysql;

SET GLOBAL max_allowed_packet = 1024 * 1024 * 256;

CREATE TABLE STORE (
    key VARCHAR(1000) PRIMARY KEY NOT NULL,
    value VARCHAR(1000) 
);