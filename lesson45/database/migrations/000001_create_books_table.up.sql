CREATE TABLE books (
    book_id UUID DEFAULT GEN_RANDOM_UUID() PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    year_published INT,
    user_id UUID DEFAULT NULL,
    is_rental BOOLEAN DEFAULT FALSE
);
