CREATE TABLE employees (
	id SERIAL PRIMARY KEY,
	name VARCHAR(100),
	position VARCHAR(100),
	salary DECIMAL(10, 2)
);

INSERT INTO employees (name, position, salary) VALUES ('Ali Valiyev', 'Manager', 60000.00),
	('Olim Murodov', 'Developer', 55000.00),
	('Gulnoza Karimova', 'Designer', 52000.00),
	('Sanjar Akramov', 'Tester', 50000.00),
	('Dilshod Rahimov', 'Analyst', 58000.00),
	('Shahnoza Usmonova', 'Developer', 54000.00),
	('Akmal Tursunov', 'Project Manager', 62000.00),
	('Akmal Tursunov', 'Project Manager', 62000.00),
	('Nodira Ergasheva', 'HR Manager', 57000.00),
	('Javlon Nizamov', 'Support Engineer', 51000.00),
	('Ravshan Oripov', 'DevOps Engineer', 59000.00);


