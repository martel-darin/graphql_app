CREATE TABLE users (id int PRIMARY KEY, name varchar(80) UNIQUE);
CREATE TABLE todos (
  id int PRIMARY KEY,
  text varchar(80),
  done boolean,
  user_id int REFERENCES users(id)
);