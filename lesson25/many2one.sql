CREATE TABLE CARS 
(
    ID uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name varchar NOT NULL,
    year date DEFAULT now(),
    brand varchar
);

CREATE TABLE Users
(
    ID uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name varchar NOT NULL
);

CREATE TABLE Users_cars 
(
    user_id uuid,
    car_id uuid,
    PRIMARY KEY (user_id, car_id),
    FOREIGN KEY (user_id) REFERENCES users(ID),
    FOREIGN KEY (car_id) REFERENCES cars(ID)
);


INSERT INTO cars (name, brand) VALUES
('Corolla', 'Toyota'),
('Civic', 'Honda'),
('Accord', 'Honda'),
('Camry', 'Toyota'),
('Model S', 'Tesla'),
('Golf', 'Volkswagen'),
('F-150', 'Ford'),
('X5', 'BMW'),
('C-Class', 'Mercedes-Benz'),
('3 Series', 'BMW');

INSERT INTO users (name) VALUES
('Akobir'),
('Bobur'),
('Umidjon'),
('Olim'),
('Shohruh'),
('Azizbek'),
('Shavkat'),
('Jasur'),
('Ravshan'),
('Jahongir');


-- Foydalanuvchilarga avtomobilni bog'lash
INSERT INTO users_cars (user_id, car_id) VALUES
('ab7e6eb3-c6da-4d12-ac99-2d1acfce0e5a', '1ca94172-6746-4b9d-9919-f7126f1cea5b'),
('e3f17ce4-568e-45f5-b4c9-be7825b9f9bf', '1ca94172-6746-4b9d-9919-f7126f1cea5b'),
('38261a08-2a84-4290-96ec-2800a7a6cbc4', '1ca94172-6746-4b9d-9919-f7126f1cea5b'),
('88ea4362-7d67-4b41-af21-1bad4e1c9a7c', '4474984d-1ec3-4c90-9991-a8e5005e66a8'),
('38261a08-2a84-4290-96ec-2800a7a6cbc4', '4474984d-1ec3-4c90-9991-a8e5005e66a8'),
('5c90add7-d14c-4ec6-81d4-871941454717', '4474984d-1ec3-4c90-9991-a8e5005e66a8'),
('f22af6fb-7c8e-4e9e-9043-20890745db90', 'f061562e-8903-4ad0-bf51-1174509a7751'),
('c228bf43-3c0e-4027-b8b1-77eee65db972', 'f061562e-8903-4ad0-bf51-1174509a7751'),
('ab7e6eb3-c6da-4d12-ac99-2d1acfce0e5a', 'f061562e-8903-4ad0-bf51-1174509a7751'),
('1b5c6cfe-db9a-4fa0-ab73-3bb8366fe4b4', 'b54f2770-a2e1-4744-98c8-ed14d5189969'),
('54828919-3aa4-4167-b7dc-e162499297e5', 'b54f2770-a2e1-4744-98c8-ed14d5189969'),
('e3f17ce4-568e-45f5-b4c9-be7825b9f9bf', 'b54f2770-a2e1-4744-98c8-ed14d5189969'),
('c228bf43-3c0e-4027-b8b1-77eee65db972', 'a804cc48-f9bc-4bdd-90b9-b58e899652aa'),
('88ea4362-7d67-4b41-af21-1bad4e1c9a7c', 'a804cc48-f9bc-4bdd-90b9-b58e899652aa'),
('1b5c6cfe-db9a-4fa0-ab73-3bb8366fe4b4', 'a804cc48-f9bc-4bdd-90b9-b58e899652aa'),
('38261a08-2a84-4290-96ec-2800a7a6cbc4', '4a94e6ef-e457-471c-9dce-be9d5c830012'),
('c228bf43-3c0e-4027-b8b1-77eee65db972', '4a94e6ef-e457-471c-9dce-be9d5c830012'),
('5c90add7-d14c-4ec6-81d4-871941454717', '4a94e6ef-e457-471c-9dce-be9d5c830012'),
('e3f17ce4-568e-45f5-b4c9-be7825b9f9bf', '40c6778c-e923-4746-beb5-b2822edc5108'),
('1b5c6cfe-db9a-4fa0-ab73-3bb8366fe4b4', '40c6778c-e923-4746-beb5-b2822edc5108'),
('54828919-3aa4-4167-b7dc-e162499297e5', '40c6778c-e923-4746-beb5-b2822edc5108'),
('88ea4362-7d67-4b41-af21-1bad4e1c9a7c', '91692233-08e5-4edd-9604-a952c4a89553'),
('8507c2a5-d24f-4355-be77-4b6637a1d437', '91692233-08e5-4edd-9604-a952c4a89553'),
('38261a08-2a84-4290-96ec-2800a7a6cbc4', '91692233-08e5-4edd-9604-a952c4a89553'),
('5c90add7-d14c-4ec6-81d4-871941454717', 'c2432c14-5686-406f-8203-32e3d2ecbbe6'),
('8507c2a5-d24f-4355-be77-4b6637a1d437', 'c2432c14-5686-406f-8203-32e3d2ecbbe6'),
('f22af6fb-7c8e-4e9e-9043-20890745db90', 'c2432c14-5686-406f-8203-32e3d2ecbbe6'),
('54828919-3aa4-4167-b7dc-e162499297e5', 'ebbc3b21-9f16-4175-a773-23760ffe4648'),
('f22af6fb-7c8e-4e9e-9043-20890745db90', 'ebbc3b21-9f16-4175-a773-23760ffe4648'),
('ab7e6eb3-c6da-4d12-ac99-2d1acfce0e5a', 'ebbc3b21-9f16-4175-a773-23760ffe4648');


SELECT 
    u.name AS user_name, 
    c.name AS car_name, 
    c.brand
FROM 
    users u
JOIN 
    users_cars uc ON u.id = uc.user_id
JOIN 
    cars c ON uc.car_id = c.id
ORDER BY 
    u.name;


