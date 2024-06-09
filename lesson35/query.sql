
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);


CREATE TABLE problems (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);


CREATE TABLE solved_problems (
    id SERIAL PRIMARY KEY,
    problem_id INT NOT NULL,
    user_id INT NOT NULL,
    solution TEXT NOT NULL,
    FOREIGN KEY (problem_id) REFERENCES problems(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);


-- users jadvaliga 5 ta ma'lumot qo'shish
INSERT INTO users (username, email, password) VALUES
('alice', 'alice@example.com', 'password1'),
('bob', 'bob@example.com', 'password2'),
('charlie', 'charlie@example.com', 'password3'),
('david', 'david@example.com', 'password4'),
('eve', 'eve@example.com', 'password5');

-- problems jadvaliga 5 ta ma'lumot qo'shish
INSERT INTO problems (user_id, title, description) VALUES
(1, 'Reverse a string in Python', 'How can I reverse a string in Python?'),
(2, 'REST API best practices', 'What are best practices for designing a REST API?'),
(3, 'Implement binary search', 'How do I implement a binary search algorithm?'),
(4, 'Closure in JavaScript', 'What is a closure in JavaScript?'),
(5, 'Optimize SQL queries', 'How can I optimize SQL queries?');

-- solved_problems jadvaliga 5 ta ma'lumot qo'shish
INSERT INTO solved_problems (problem_id, user_id, solution) VALUES
(1, 1, 'Use slicing: s[::-1]'),
(2, 2, 'Use nouns for endpoints, stateless ops, proper HTTP codes'),
(3, 3, 'Divide search interval in half repeatedly'),
(4, 4, 'A function retaining access to lexical scope'),
(5, 5, 'Use indexing, efficient joins, analyze query plans');
