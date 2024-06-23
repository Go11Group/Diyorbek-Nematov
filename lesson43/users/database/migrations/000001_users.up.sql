CREATE TABLE Users (
    id UUID DEFAULT GEN_RANDOM_UUID() PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    phone VARCHAR(30) NOT NULL,
    age INTEGER
)
