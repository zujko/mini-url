CREATE TABLE profile (
    profile_id serial PRIMARY KEY,
    profile_name text NOT NULL
);

CREATE TABLE url (
    url_id serial PRIMARY KEY,
    short_url text UNIQUE,
    long_url text NOT NULL,
    user_id int REFERENCES profile (profile_id)
);

CREATE TABLE click (
    click_id serial PRIMARY KEY,
    click_date timestamp NOT NULL,
    country text,
    url_id int REFERENCES url (url_id) NOT NULL
);