CREATE TABLE IF NOT EXISTS cardealers 
( 
    id          bigserial PRIMARY KEY, 
    created_at  timestamp(0) with time zone NOT NULL DEFAULT NOW(), 
    updated_at  timestamp(0) with time zone NOT NULL DEFAULT NOW(), 
    title       text                        NOT NULL, 
    description text                        NOT NULL, 
    coordinates text                        NOT NULL, 
    address     text                        NOT NULL, 
    country     text                        NOT NULL 
); 
 
CREATE TABLE IF NOT EXISTS car 
( 
    id              bigserial PRIMARY KEY, 
    created_at      timestamp(0) with time zone NOT NULL DEFAULT NOW(), 
    updated_at      timestamp(0) with time zone NOT NULL DEFAULT NOW(), 
    title           text                        NOT NULL, 
    description     text, 
    year int 
); 
 
CREATE TABLE IF NOT EXISTS cardealers_and_car 
( 
    "id"         bigserial PRIMARY KEY, 
    "created_at" timestamp(0) with time zone NOT NULL DEFAULT NOW(), 
    "updated_at" timestamp(0) with time zone NOT NULL DEFAULT NOW(), 
    "cardealer" bigserial, 
    "car"       bigserial, 
    FOREIGN KEY (cardealers) 
        REFERENCES cardealers(id), 
    FOREIGN KEY (car) 
        REFERENCES car(id) 
);