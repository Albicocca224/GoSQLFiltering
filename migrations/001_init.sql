CREATE TABLE IF NOT EXISTS users (
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(100) NOT NULL,
    email      VARCHAR(150) NOT NULL UNIQUE,
    gender     VARCHAR(10)  NOT NULL CHECK (gender IN ('male', 'female', 'other')),
    birth_date DATE         NOT NULL
);

CREATE TABLE IF NOT EXISTS user_friends (
    user_id   INTEGER REFERENCES users(id) ON DELETE CASCADE,
    friend_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, friend_id),
    CONSTRAINT no_self_friend CHECK (user_id <> friend_id)
);



INSERT INTO users (name, email, gender, birth_date) VALUES
  ('Alice Johnson',   'alice@example.com',   'female', '1995-03-12'),
  ('Bob Smith',       'bob@example.com',     'male',   '1993-07-24'),
  ('Carol White',     'carol@example.com',   'female', '1997-11-05'),
  ('David Brown',     'david@example.com',   'male',   '1990-01-30'),
  ('Eva Davis',       'eva@example.com',     'female', '1998-06-18'),
  ('Frank Miller',    'frank@example.com',   'male',   '1992-09-09'),
  ('Grace Wilson',    'grace@example.com',   'female', '1996-04-22'),
  ('Henry Moore',     'henry@example.com',   'male',   '1994-12-03'),
  ('Iris Taylor',     'iris@example.com',    'female', '1999-02-14'),
  ('Jack Anderson',   'jack@example.com',    'male',   '1991-08-07'),
  ('Karen Thomas',    'karen@example.com',   'female', '1993-05-25'),
  ('Leo Jackson',     'leo@example.com',     'male',   '1995-10-16'),
  ('Mona Harris',     'mona@example.com',    'female', '1997-07-08'),
  ('Nate Martin',     'nate@example.com',    'male',   '1990-03-19'),
  ('Olivia Garcia',   'olivia@example.com',  'female', '1998-01-27'),
  ('Paul Martinez',   'paul@example.com',    'male',   '1992-11-11'),
  ('Quinn Robinson',  'quinn@example.com',   'other',  '1996-06-30'),
  ('Rachel Clark',    'rachel@example.com',  'female', '1994-09-04'),
  ('Sam Rodriguez',   'sam@example.com',     'male',   '1999-04-15'),
  ('Tina Lewis',      'tina@example.com',    'female', '1991-12-22'),
  ('Uma Lee',         'uma@example.com',     'female', '1993-02-08'),
  ('Victor Walker',   'victor@example.com',  'male',   '1995-08-17'),
  ('Wendy Hall',      'wendy@example.com',   'female', '1997-05-29'),
  ('Xander Allen',    'xander@example.com',  'male',   '1990-10-01'),
  ('Yara Young',      'yara@example.com',    'female', '1998-03-13')
ON CONFLICT DO NOTHING;


INSERT INTO user_friends VALUES (1,2),(2,1);
INSERT INTO user_friends VALUES (1,3),(3,1);
INSERT INTO user_friends VALUES (1,4),(4,1);
INSERT INTO user_friends VALUES (1,5),(5,1);
INSERT INTO user_friends VALUES (1,6),(6,1);

INSERT INTO user_friends VALUES (2,3),(3,2);
INSERT INTO user_friends VALUES (2,4),(4,2);
INSERT INTO user_friends VALUES (2,5),(5,2);
INSERT INTO user_friends VALUES (2,17),(17,2);

INSERT INTO user_friends VALUES (6,7),(7,6);
INSERT INTO user_friends VALUES (6,8),(8,6);
INSERT INTO user_friends VALUES (6,9),(9,6);
INSERT INTO user_friends VALUES (6,10),(10,6);

INSERT INTO user_friends VALUES (7,8),(8,7);
INSERT INTO user_friends VALUES (7,9),(9,7);
INSERT INTO user_friends VALUES (7,10),(10,7);
INSERT INTO user_friends VALUES (7,18),(18,7);

INSERT INTO user_friends VALUES (11,12),(12,11);
INSERT INTO user_friends VALUES (11,13),(13,11);
INSERT INTO user_friends VALUES (13,14),(14,13);
INSERT INTO user_friends VALUES (14,15),(15,14);
INSERT INTO user_friends VALUES (15,16),(16,15);
INSERT INTO user_friends VALUES (19,20),(20,19);
INSERT INTO user_friends VALUES (21,22),(22,21);
INSERT INTO user_friends VALUES (23,24),(24,23);
INSERT INTO user_friends VALUES (24,25),(25,24);
