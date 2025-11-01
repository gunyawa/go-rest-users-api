CREATE TABLE users (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  city TEXT NOT NULL
);

CREATE TABLE orders (
  id INTEGER PRIMARY KEY,
  user_id INTEGER NOT NULL,
  amount INTEGER NOT NULL,
  FOREIGN KEY(user_id) REFERENCES users(id)
);

INSERT INTO users (id, name, city) VALUES
(1, 'Alice', 'Almaty'),
(2, 'Bob', 'Astana'),
(3, 'Eve', 'Almaty');

INSERT INTO orders (user_id, amount) VALUES
(1, 5),
(1, 7),
(3, 3);
