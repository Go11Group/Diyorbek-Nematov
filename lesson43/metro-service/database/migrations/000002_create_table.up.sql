CREATE TABLE cards (
    id UUID DEFAULT GEN_RANDOM_UUID() PRIMARY KEY,
    user_id UUID NOT NULL,
    number VARCHAR(20) UNIQUE NOT NULL
);

CREATE TABLE station (
    id UUID DEFAULT GEN_RANDOM_UUID() PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE terminal (
    id UUID DEFAULT GEN_RANDOM_UUID() PRIMARY KEY,
    station_id UUID NOT NULL,
    FOREIGN KEY (station_id) REFERENCES station(id)
);

CREATE TABLE transactions (
    id UUID DEFAULT GEN_RANDOM_UUID() PRIMARY KEY,
    card_id UUID NOT NULL,
    terminal_id UUID DEFAULT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    type transaction_type NOT NULL,
    FOREIGN KEY (card_id) REFERENCES cards(id),
    FOREIGN KEY (terminal_id) REFERENCES terminal(id)
);
