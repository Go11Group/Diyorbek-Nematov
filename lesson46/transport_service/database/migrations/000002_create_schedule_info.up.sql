CREATE TABLE IF NOT EXISTS schedule (
    bus_number VARCHAR(20) NOT NULL,
    stop VARCHAR(100) NOT NULL,
    arrival_time TIMESTAMP NOT NULL,
    FOREIGN KEY (bus_number) REFERENCES transport_info(bus_number)
);
