CREATE TABLE IF NOT EXISTS tasks (
              id SERIAL PRIMARY KEY,
              title TEXT NOT NULL,
              description TEXT NOT NULL,
              due_date DATE NOT NULL,
              completed BOOLEAN NOT NULL,
              created_at TIMESTAMP NOT NULL DEFAULT NOW()
);