CREATE TABLE IF NOT EXISTS ride(
    id VARCHAR(26) NOT NULL PRIMARY KEY,
    user_id VARCHAR(26) NOT NULL,
    source_latitude DOUBLE PRECISION NOT NULL,
    source_longitude DOUBLE PRECISION NOT NULL,
    destination_latitude DOUBLE PRECISION NOT NULL,
    destination_longitude DOUBLE PRECISION NOT NULL,
    status varchar(15) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL
);