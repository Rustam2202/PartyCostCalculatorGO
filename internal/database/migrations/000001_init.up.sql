CREATE TABLE IF NOT EXISTS persons (
	Id SERIAL PRIMARY KEY,
	Name VARCHAR(30) NOT NULL
);
CREATE TABLE IF NOT EXISTS events (
	Id SERIAL PRIMARY KEY,
	Name VARCHAR(30) NOT NULL,
	Date DATE,
	Total NUMERIC DEFAULT 0
);
CREATE TABLE IF NOT EXISTS pers_events (
	Id SERIAL PRIMARY KEY,
	Person INT,
	Event INT NOT NULL,
	Spent NUMERIC DEFAULT 0,
	Factor INT DEFAULT 1,
	FOREIGN KEY (Person) REFERENCES persons(Id) ON DELETE CASCADE,
	FOREIGN KEY (Event) REFERENCES events(Id) ON DELETE CASCADE
);