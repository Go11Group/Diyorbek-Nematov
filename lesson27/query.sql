
CREATE TABLE company (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    address VARCHAR(255)
);

CREATE TABLE employee (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    position VARCHAR(100),
    company_id uuid REFERENCES company(id)
);

