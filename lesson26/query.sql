CREATE TABLE
    Student (
        ID uuid DEFAULT gen_random_uuid () PRIMARY KEY,
        Name VARCHAR(50),
        Age int
    );

CREATE TABLE
    Course (
        ID uuid DEFAULT gen_random_uuid () PRIMARY KEY,
        Name VARCHAR(50)
    );

CREATE TABLE
    Student_course (
        ID uuid DEFAULT gen_random_uuid () PRIMARY KEY,
        student_id uuid NOT NULL,
        course_id uuid NOT NULL,
        FOREIGN KEY (student_id) REFERENCES Student (ID),
        FOREIGN KEY (course_id) REFERENCES Course (ID)
    );

CREATE TABLE
    Grade (
        ID uuid DEFAULT gen_random_uuid () PRIMARY KEY,
        grade int check (
            1 <= grade
            AND grade <= 5
        ),
        student_course_id uuid NOT NULl,
        FOREIGN KEY (student_course_id) REFERENCES Student_course (ID)
    );

INSERT INTO
    Student (Name, Age)
VALUES
    ('Alisher Karimov', 20),
    ('Dilnoza Otabekova', 21),
    ('Jasur Bekmurodov', 22),
    ('Nodira Tursunova', 23),
    ('Shavkat Mirzayev', 24),
    ('Zarina Qodirova', 25),
    ('Oybek Usmonov', 26),
    ('Malika Abdullayeva', 27),
    ('Sardor Ibragimov', 28),
    ('Gulnora Vohidova', 29);

INSERT INTO
    Course (Name)
VALUES
    ('Mathematics'),
    ('Physics'),
    ('Chemistry'),
    ('Biology'),
    ('Computer Science');

--course_id lar 
--11659c38-e035-4c7e-9d3f-63b47e8dfe92
--d521354a-ca52-4708-a722-792af269fdce
--a2e2005f-1230-4533-8570-418e0adbddc1
--6c17344d-6f76-4b39-846e-ce974b538db8
--1eae1118-4efd-4d08-a8fc-c4399be4be8e
--Student idlar 
--d74f64af-cd5d-4dbc-9ce0-271b8cbbe7be
--32537b1f-eff7-49bb-ba9c-7ab0fd1c7905
--30e9bece-c29c-4f2d-aeee-609dcb718fde
--2a62e8d4-796c-4bd9-a2dc-21b8d5f58166
--7e959cef-f616-4b22-a7a0-e11fa5c44409
--da4ae507-4a35-42dc-8d5d-eba573b64388
--1510632f-26d4-4b8d-8ea2-42a0f5af5771
--4bab7921-79f3-4d47-8107-e724913c4ce1
--a84294aa-9b25-4331-8c66-212678c34668
--38c80ef8-aa18-4f94-aa0b-1aa199f2d88a
-- Talabalarni kurslarga ro'yxatdan o'tkazish
INSERT INTO
    Student_course (student_id, course_id)
VALUES
    (
        'd74f64af-cd5d-4dbc-9ce0-271b8cbbe7be',
        '1eae1118-4efd-4d08-a8fc-c4399be4be8e'
    ), -- Alisher Karimov - Matematika
    (
        'd74f64af-cd5d-4dbc-9ce0-271b8cbbe7be',
        '11659c38-e035-4c7e-9d3f-63b47e8dfe92'
    ), -- Alisher Karimov - Fizika
    (
        '32537b1f-eff7-49bb-ba9c-7ab0fd1c7905',
        '1eae1118-4efd-4d08-a8fc-c4399be4be8e'
    ), -- Dilnoza Otabekova - Matematika
    (
        '32537b1f-eff7-49bb-ba9c-7ab0fd1c7905',
        'd521354a-ca52-4708-a722-792af269fdce'
    ), -- Dilnoza Otabekova - Kimyo
    (
        '30e9bece-c29c-4f2d-aeee-609dcb718fde',
        'a2e2005f-1230-4533-8570-418e0adbddc1'
    ), -- Jasur Bekmurodov - Biologiya
    (
        '30e9bece-c29c-4f2d-aeee-609dcb718fde',
        '6c17344d-6f76-4b39-846e-ce974b538db8'
    ), -- Jasur Bekmurodov - Kompyuter Ilmlari
    (
        '2a62e8d4-796c-4bd9-a2dc-21b8d5f58166',
        '1eae1118-4efd-4d08-a8fc-c4399be4be8e'
    ), -- Nodira Tursunova - Matematika
    (
        '2a62e8d4-796c-4bd9-a2dc-21b8d5f58166',
        '11659c38-e035-4c7e-9d3f-63b47e8dfe92'
    ), -- Nodira Tursunova - Fizika
    (
        '7e959cef-f616-4b22-a7a0-e11fa5c44409',
        'd521354a-ca52-4708-a722-792af269fdce'
    ), -- Shavkat Mirzayev - Kimyo
    (
        '7e959cef-f616-4b22-a7a0-e11fa5c44409',
        'a2e2005f-1230-4533-8570-418e0adbddc1'
    ), -- Shavkat Mirzayev - Biologiya
    (
        'da4ae507-4a35-42dc-8d5d-eba573b64388',
        '6c17344d-6f76-4b39-846e-ce974b538db8'
    ), -- Zarina Qodirova - Kompyuter Ilmlari
    (
        'da4ae507-4a35-42dc-8d5d-eba573b64388',
        '1eae1118-4efd-4d08-a8fc-c4399be4be8e'
    ), -- Zarina Qodirova - Matematika
    (
        '1510632f-26d4-4b8d-8ea2-42a0f5af5771',
        '11659c38-e035-4c7e-9d3f-63b47e8dfe92'
    ), -- Oybek Usmonov - Fizika
    (
        '1510632f-26d4-4b8d-8ea2-42a0f5af5771',
        'd521354a-ca52-4708-a722-792af269fdce'
    ), -- Oybek Usmonov - Kimyo
    (
        '4bab7921-79f3-4d47-8107-e724913c4ce1',
        'a2e2005f-1230-4533-8570-418e0adbddc1'
    ), -- Malika Abdullayeva - Biologiya
    (
        '4bab7921-79f3-4d47-8107-e724913c4ce1',
        '6c17344d-6f76-4b39-846e-ce974b538db8'
    ), -- Malika Abdullayeva - Kompyuter Ilmlari
    (
        'a84294aa-9b25-4331-8c66-212678c34668',
        '1eae1118-4efd-4d08-a8fc-c4399be4be8e'
    ), -- Sardor Ibragimov - Matematika
    (
        'a84294aa-9b25-4331-8c66-212678c34668',
        '11659c38-e035-4c7e-9d3f-63b47e8dfe92'
    ), -- Sardor Ibragimov - Fizika
    (
        '38c80ef8-aa18-4f94-aa0b-1aa199f2d88a',
        'd521354a-ca52-4708-a722-792af269fdce'
    ), -- Gulnora Vohidova - Kimyo
    (
        '38c80ef8-aa18-4f94-aa0b-1aa199f2d88a',
        'a2e2005f-1230-4533-8570-418e0adbddc1'
    );-- Gulnora Vohidova - Biologiya


---- student_course_id lar
--6ada79d4-1ecb-42b6-90bd-37ddd56ed384
--0061c633-8191-4e3c-b9a3-ef3b537cbe3e
--19e9f9e7-30f0-4f0b-bd44-044d86ac01a0
--225584ab-5b47-4da3-b9ff-d52e49ae211b
--dc4a8789-6def-4d20-a9d7-323a0c78f6e2
--a2e5ef8a-a13f-429a-8758-782e9701ce4c        
--d063b218-3e02-4616-9f45-69b4fe0d634c
--5ebc6bcd-012c-4a23-8920-401f107cb6e6
--5960e9af-9e52-413d-a754-08951592c104
--0f239c1c-534a-4320-8891-3d4c02753b04
--bf3ee79e-eaa2-42d4-94c0-4bd8c35819ad
--409d3eea-fb75-4586-af4e-0c909f8602fe
--4f618303-0f0b-41fa-9eda-595fb6a7e3d7
--deec0c32-0a7e-403a-94a7-368661f7ed0d
--c40e86a7-bb8b-4343-8cbe-0b639bb0b086
--dc18f5cf-24d9-4dda-b1ca-4f7a83c7fb5a
--1bd3fcc1-cd21-4a7c-a41e-76fb3964143c
--2b865ec6-de0b-4d2b-8829-0147e52b0a23
--cf189721-19b4-4055-8297-9a4358d346d4
--cb6810d3-da98-4d4a-895d-f034c9bda4be


-- Talabalar uchun baholarni qo'shish
INSERT INTO
    Grade (grade, student_course_id)
VALUES
    (5, '6ada79d4-1ecb-42b6-90bd-37ddd56ed384'), -- Alisher Karimov - Matematika
    (4, '0061c633-8191-4e3c-b9a3-ef3b537cbe3e'), -- Alisher Karimov - Fizika
    (3, '19e9f9e7-30f0-4f0b-bd44-044d86ac01a0'), -- Dilnoza Otabekova - Matematika
    (2, '225584ab-5b47-4da3-b9ff-d52e49ae211b'), -- Dilnoza Otabekova - Kimyo
    (5, 'dc4a8789-6def-4d20-a9d7-323a0c78f6e2'), -- Jasur Bekmurodov - Biologiya
    (4, 'a2e5ef8a-a13f-429a-8758-782e9701ce4c'), -- Jasur Bekmurodov - Kompyuter Ilmlari
    (3, 'd063b218-3e02-4616-9f45-69b4fe0d634c'), -- Nodira Tursunova - Matematika
    (2, '5ebc6bcd-012c-4a23-8920-401f107cb6e6'), -- Nodira Tursunova - Fizika
    (5, '5960e9af-9e52-413d-a754-08951592c104'), -- Shavkat Mirzayev - Kimyo
    (4, '0f239c1c-534a-4320-8891-3d4c02753b04'), -- Shavkat Mirzayev - Biologiya
    (3, 'bf3ee79e-eaa2-42d4-94c0-4bd8c35819ad'), -- Zarina Qodirova - Kompyuter Ilmlari
    (2, '409d3eea-fb75-4586-af4e-0c909f8602fe'), -- Zarina Qodirova - Matematika
    (5, '4f618303-0f0b-41fa-9eda-595fb6a7e3d7'), -- Oybek Usmonov - Fizika
    (4, 'deec0c32-0a7e-403a-94a7-368661f7ed0d'), -- Oybek Usmonov - Kimyo
    (3, 'c40e86a7-bb8b-4343-8cbe-0b639bb0b086'), -- Malika Abdullayeva - Biologiya
    (2, 'dc18f5cf-24d9-4dda-b1ca-4f7a83c7fb5a'), -- Malika Abdullayeva - Kompyuter Ilmlari
    (5, '1bd3fcc1-cd21-4a7c-a41e-76fb3964143c'), -- Sardor Ibragimov - Matematika
    (4, '2b865ec6-de0b-4d2b-8829-0147e52b0a23'), -- Sardor Ibragimov - Fizika
    (3, 'cf189721-19b4-4055-8297-9a4358d346d4'), -- Gulnora Vohidova - Kimyo
    (2, 'cb6810d3-da98-4d4a-895d-f034c9bda4be'); -- Gulnora Vohidova - Biologiya



-- 3 - savol
SELECT
    c.name,
    ROUND(AVG(g.grade), 2) AS "O'rtacha baho"
FROM
    student_course AS sc
    LEFT JOIN course AS c ON sc.course_id = c.id
    LEFT JOIN grade AS g ON sc.id = g.student_course_id
GROUP BY
    c.name

    
-- 5-savol
SELECT
    c.name,
    ROUND(AVG(g.grade), 2) AS "O'rtacha baho"
FROM
    course c
    JOIN student_course sc ON c.id = sc.course_id
    JOIN grade g ON sc.id = g.student_course_id
GROUP BY
    c.name
ORDER BY
    "O'rtacha baho" DESC
LIMIT
    1;


-- 4-savol
SELECT
    c.name,
    ARRAY_AGG (s.name),
    ARRAY_AGG (s.age)
FROM
    course as c
    LEFT JOIN student_course as sc ON sc.course_id = c.id
    LEFT JOIN student as s ON sc.student_id = s.id
    INNER JOIN (
        SELECT
            MIN(s.age) as min_age,
            c.name
        FROM
            course AS c
            LEFT JOIN student_course AS sc ON c.id = sc.course_id
            LEFT JOIN student AS s ON s.id = sc.student_id
        GROUP BY
            c.name
    ) AS ms ON ms.min_age = s.age
    AND ms.name = c.name
GROUP BY
    c.name


    -- 2 - savol
SELECT
    c.name,
    ARRAY_AGG (s.name),
    ARRAY_AGG (g.grade)
FROM
    student_course AS sc
    LEFT JOIN student AS s ON sc.student_id = s.id
    LEFT JOIN course AS c ON sc.course_id = c.id
    LEFT JOIN grade AS g ON g.student_course_id = sc.id
    INNER JOIN (
        SELECT
            Max(g.grade) as max_grade,
            c.name
        FROM
            student_course AS sc
            LEFT JOIN student AS s ON sc.student_id = s.id
            LEFT JOIN course AS c ON sc.course_id = c.id
            LEFT JOIN grade AS g ON g.student_course_id = sc.id
        GROUP BY
            c.name
    ) as ms ON ms.max_grade = g.grade
    AND ms.name = c.name
GROUP BY
    c.name