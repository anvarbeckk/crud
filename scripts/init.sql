CREATE TABLE IF NOT EXISTS blogs {
    id SERIAL PRIMARY KEY,
    title VARCHAR(250) NOT NULL,
    content text not NULL
};

CREATE TABLE IF NOT EXISTS news {
    id SERIAL PRIMARY KEY,
    title varchar(255) not null,
    content text not null
}
