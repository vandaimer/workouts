CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE workout (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "week_number" int NOT NULL,
  distance int NOT NULL,
  time int NOT NULL,
  timestamp timestamp NOT NULL
)
