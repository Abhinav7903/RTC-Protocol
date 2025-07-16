-- Create rooms table using SERIAL ID
CREATE TABLE rooms (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  room_type VARCHAR(10) CHECK (room_type IN ('single', 'group')) NOT NULL DEFAULT 'single',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create participants
CREATE TABLE room_participants (
  id SERIAL PRIMARY KEY,
  room_id INTEGER NOT NULL REFERENCES rooms(id) ON DELETE CASCADE,
  display_name VARCHAR(50),
  joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create signals
CREATE TABLE signals (
  id SERIAL PRIMARY KEY,
  room_id INTEGER NOT NULL REFERENCES rooms(id) ON DELETE CASCADE,
  sender_id INTEGER, -- now references room_participants.id
  signal_type VARCHAR(20) CHECK (signal_type IN ('offer', 'answer', 'candidate','ice')) NOT NULL,
  signal_payload JSONB NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
