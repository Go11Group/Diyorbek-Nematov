CREATE TABLE author (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    birthdate DATE
);

CREATE TABLE book (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    genre VARCHAR(100),
    author_id INT,
    FOREIGN KEY (author_id) REFERENCES author(id)
);

INSERT INTO author (name, birthdate)
VALUES 
    ('Abdulla Qodiriy', '1907-02-04'),
    ('Alisher Navoiy', '1441-02-09'),
    ('Erkin Vohidov', '1940-09-01'),
    ('William Shakespeare', '1564-04-26'),
    ('Lev Tolstoy', '1828-09-09'),
    ('Mark Twain', '1835-11-30'),
    ('Aleksandr Pushkin', '1799-06-06');


INSERT INTO book (title, genre, author_id)
VALUES
    ('O''tgan kunlar', 'Roman', 1),
    ('Mehrobdan chayon', 'Roman', 1),
    ('Xamsa', 'Doston', 2),
    ('Nido', 'Doston', 3),
    ('Tarixi muli ajam', 'Doston', 2),
    ('Romeo va Juletta', 'Tradigiya', 4),
    ('Otello', 'tradigiya', 4),
    ('Urush va Tinchlik', 'Roman', 5),
    ('Shaxzoda va gado', 'Komidiya', 6),
    ('Missisipidagi avvalgi davrlar', 'Roman', 6),
    ('Ruslan va Lyudmila', 'Roman', 7),
    ('Kapitan qizi', 'Nasriy', 7);


SELECT a.name, COUNT(b.title) AS "Yozgan kitoblar soni"
FROM
    author AS a
INNER JOIN book AS b
    ON a.id = b.author_id
GROUP BY
    a.name


SELECT a.name, ARRAY_AGG(DISTINCT b.genre) AS "IJON Qilgan Janrlar"
FROM 
    author AS a
INNER JOIN book AS b
    ON a.id = b.author_id
GROUP BY a.name

SELECT a.name, b.genre, COUNT(b.genre) AS "Har bir janrda yozgan kitoblar soni"
FROM 
    author AS a
INNER JOIN book AS b
    ON a.id = b.author_id
GROUP BY a.name, b.genre

SELECT a.name, b.title AS "Kitob nomi"
FROM author AS a
LEFT JOIN book AS b
    ON a.id = b.author_id

SELECT a.name, b.title AS "Eng yoshi katta yozuvchining asarlari"
FROM 
    book AS b
INNER JOIN author AS a
    ON b.author_id = a.id
WHERE 
    a.birthdate = (SELECT min(birthdate) FROM author);