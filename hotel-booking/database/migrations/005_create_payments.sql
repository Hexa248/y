CREATE TABLE IF NOT EXISTS payments (
  id SERIAL PRIMARY KEY,
  booking_id INT NOT NULL,
  method VARCHAR(30) NOT NULL,
  status VARCHAR(20) NOT NULL,
  reference VARCHAR(80)
);
