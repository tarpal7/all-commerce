CREATE TABLE  users
(
    id              serial          NOT NULL UNIQUE,
    username        VARCHAR (255)   NOT NULL UNIQUE,
    password_hash   VARCHAR (255)   NOT NULL 
);

CREATE TABLE allcommerce_lists
(
    id              serial          NOT NULL UNIQUE, 
    title           VARCHAR (255)   NOT NULL, 
    description     VARCHAR (255)   
);

CREATE TABLE allcommerce_items
(
    id              serial                                                  NOT NULL UNIQUE, 
    title           VARCHAR (255)                                           NOT NULL, 
    description     VARCHAR (255),   
    list_id         INT REFERENCES allcommerce_lists (id) on delete CASCADE NOT NULL
);

