CREATE TABLE IF NOT EXISTS transport_info (
    bus_number VARCHAR(20) PRIMARY KEY,
    location VARCHAR(200) NOT NULL,
    status BOOLEAN DEFAULT FALSE
);
