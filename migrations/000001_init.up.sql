CREATE TABLE IF NOT EXISTS persons (
	id SERIAL PRIMARY KEY,
	name VARCHAR(30) NOT NULL
);
CREATE INDEX persons_name_idx ON persons (name);
CREATE TABLE IF NOT EXISTS events (
	id SERIAL PRIMARY KEY,
	name VARCHAR(30) NOT NULL,
	date DATE,
	total NUMERIC DEFAULT 0
);
CREATE INDEX events_name_idx ON events (name);
CREATE TABLE IF NOT EXISTS persons_events (
	id SERIAL PRIMARY KEY,
	person_id INT,
	event_id INT NOT NULL,
	spent NUMERIC DEFAULT 0,
	factor INT DEFAULT 1,
	FOREIGN KEY (person_id) REFERENCES persons(Id) ON DELETE CASCADE,
	FOREIGN KEY (event_id) REFERENCES events(Id) ON DELETE CASCADE
);