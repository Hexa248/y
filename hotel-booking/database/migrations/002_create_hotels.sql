CREATE TABLE IF NOT EXISTS hotels (
  id SERIAL PRIMARY KEY,
  city_id INT NOT NULL,
  name VARCHAR(200) NOT NULL,
  address TEXT NOT NULL,
  description TEXT,
  image_url TEXT,
  rating NUMERIC(2,1)
);
