ALTER TABLE users ADD CONSTRAINT unique_username UNIQUE (user_name);
ALTER TABLE users ADD CONSTRAINT unique_email UNIQUE (email);
