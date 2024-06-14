CREATE TABLE Users (
    user_id UUID DEFAULT GEN_RANDOM_UUID() PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    birthday TIMESTAMP,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE Courses (
    course_id UUID DEFAULT GEN_RANDOM_UUID() PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0 
);

CREATE TABLE Lessons (
    lesson_id UUID DEFAULT GEN_RANDOM_UUID() PRIMARY KEY,
    course_id UUID NOT NULL,
    title VARCHAR(100) NOT NULL,
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0,
    FOREIGN KEY (course_id) REFERENCES courses(course_id)
);

CREATE TABLE Enrollments (
    enrollment_id UUID DEFAULT GEN_RANDOM_UUID() PRIMARY KEY,
    user_id UUID NOT NULL,
    course_id UUID NOT NULL,
    enrollment_date TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (course_id) REFERENCES courses(course_id),
    UNIQUE (user_id, course_id)
);

-- course_idlar 
-- 0b2064f1-e913-40e0-98c7-566ae2b79c63
-- abd92518-7ffd-49d5-a040-2c8216affc6f
-- 6f69147e-cda1-4d9f-90f7-eb3d13e9becf
-- 97dfe8d2-e620-4084-9d8c-bda63d37b5d7
-- 82a50acd-4bf7-411c-9000-d34210e8eeb2
-- 82a50acd-4bf7-411c-9000-d34210e8eeb2
-- 8d85b302-7027-44cf-8b9c-d3e32898415a
-- 8071333b-6e63-47eb-bb92-768c7225657e
-- 8071333b-6e63-47eb-bb92-768c7225657e
-- 8071333b-6e63-47eb-bb92-768c7225657e


-- user_id lar 
-- bce43ad0-6c90-4c1d-b42d-2f33428ba74d 
-- 476e0241-ec61-4cba-86fc-7ed9e428255d 
-- 21c38157-c238-4b71-ae4d-34bb8584d73f 
-- 5bda1923-78fa-4174-bb15-550349fb2e5b 
-- 481607c0-e440-4243-aee7-9bf835e229b0 
-- 84191ea0-3228-41b7-b311-257132e20904 
-- 5193ef08-8416-46a9-99b7-57a1c5d030a6 
-- 4bfa2548-1689-4a15-b2ad-2e4cdb8ec5df 
-- 01d88145-90d3-4a44-8779-beea73439bdd 
-- 8aae3057-dd73-4c43-95cb-8a08971a4a40 
-- 181c4d26-4242-48d0-8bbc-23b9991f54d0 
-- c9c3f048-62f5-4569-b6ed-700db3b6cfb9 
-- 6448affd-3f53-48d0-a56b-6987ef967cc6 
-- 79f6cda2-3627-4e28-aa36-c25e439e50c1 
-- 05aa4480-e26a-453f-af92-841d812e1379 
-- ca410666-c151-47fc-90b9-489c5c74458b 
-- 64a50688-06bd-4f18-b28a-09f15e143f66 
-- 73805823-ff79-490b-b06d-3fdf57d9800c 
-- 4c716f5a-e45c-4667-ab20-653d1c154a1d 
-- 18099c90-6a7f-412a-b6aa-905dc12a66c2 
-- b4f6bdf6-52ff-4318-8676-c70d7dbeae95 
-- 3f2ae335-ba59-4ac8-b0cd-a9a101e5fbce 
-- 0a121c3c-3e50-4074-83e9-4755d2c04785 
-- e025e5f3-6234-4af3-a14b-384a49bb2929 
-- 45328c01-64c4-4806-9d65-3e8c621bec92 
-- f5acaa59-2645-403b-8d4c-6706cb5f2c0d 
-- bb23a837-4f06-4345-8093-0928c91e9e97 
-- 5c162156-9bf4-4b3f-88df-3bfc01484ca5 
-- 26bb0961-ed2c-4552-a240-35aa0c4e0a12 
-- 9d05bc08-9fc1-462d-a7b4-dd04a4fc013c 


SELECT 
    c.course_id,
    c.title,
    COUNT(c.course_id) AS enrollment_count
FROM 
    courses AS c 
INNER JOIN
    enrollments AS e ON c.course_id = e.course_id
WHERE 
    e.enrollment_date BETWEEN 
GROUP BY
    c.course_id, c.title
HAVING
    COUNT(c.course_id) = (
        SELECT 
            MAX(enroll_count) 
        FROM (
            SELECT 
                COUNT(e.course_id) AS enroll_count
            FROM 
                courses AS c 
            INNER JOIN
                enrollments AS e ON c.course_id = e.course_id
            GROUP BY
                c.course_id
        ) AS counts
    )


SELECT
    u.user_id,
    c.course_id,
    c.title,
    c.description
FROM
    users AS u 
INNER JOIN 
    enrollments AS e ON u.user_id = e.user_id
INNER JOIN
    courses AS c ON e.course_id = c.course_id
WHERE 
    u.deleted_at = 0 AND c.deleted_at = 0 AND e.deleted_at = 0 