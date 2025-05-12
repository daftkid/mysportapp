CREATE TABLE IF NOT EXISTS avatars(
    id serial PRIMARY KEY,
    path VARCHAR (256) NOT NULL,
    alt VARCHAR (120)
);

CREATE TABLE IF NOT EXISTS credentials(
    id serial PRIMARY KEY,
    email VARCHAR (64) NOT NULL,
    hpassword VARCHAR (120)
);

CREATE TABLE IF NOT EXISTS users(
    id serial PRIMARY KEY,
    name VARCHAR (64) NOT NULL,
    bio VARCHAR (256),
    avatar_id INTEGER,
    creds_id INTEGER,
    CONSTRAINT fk_avatar
        FOREIGN KEY (avatar_id)
            REFERENCES avatars(id),
    CONSTRAINT fk_credentials
        FOREIGN KEY (creds_id)
            REFERENCES credentials(id)
);

CREATE TABLE IF NOT EXISTS targets(
    id serial PRIMARY KEY,
    numOfDays INTEGER NOT NULL,
    numOfReps INTEGER NOT NULL,
    comment VARCHAR (64),
    user_id INTEGER,
    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
            REFERENCES users(id)
);