CREATE TABLE rentals (
    rental_id UUID DEFAULT GEN_RANDOM_UUID() PRIMARY KEY,
    book_id UUID NOT NULL,
    user_id UUID NOT NULL,
    rental_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    return_date TIMESTAMP,
    FOREIGN KEY (book_id) REFERENCES books(book_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);
