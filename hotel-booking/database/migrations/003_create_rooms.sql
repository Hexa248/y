CREATE TABLE IF NOT EXISTS rooms (
  id SERIAL PRIMARY KEY,
  hotel_id INT NOT NULL,
  type VARCHAR(20) NOT NULL,
  name VARCHAR(200) NOT NULL,
  price INT NOT NULL,
  beds INT NOT NULL,
  capacity INT NOT NULL,
  facilities JSONB
);
