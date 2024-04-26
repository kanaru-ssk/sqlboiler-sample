CREATE TYPE role AS ENUM ('OWNER', 'MEMBER');

CREATE TABLE IF NOT EXISTS team_member(
  id SERIAL PRIMARY KEY,
  team_id INT REFERENCES teams(id),
  user_id INT REFERENCES users(id),
  user_role role NOT NULL
);