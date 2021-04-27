CREATE TABLE posts
(
    id SERIAL NOT NULL,
	title varchar(200) not null,
    content varchar(2000) NOT NULL
)

insert into posts (title, content) values ('Second Post', 'Wow, it is the second post already!');
insert into posts (title, content) values ('Third Post', 'Ok, third one is here!');

CREATE TABLE comments
(
    id SERIAL NOT NULL,
    postId integer not null,
	author varchar(200) not null,
    content varchar(2000) NOT NULL
)

insert into comments (postId, author, content) values (1, 'Marc', 'Wow, nice!');
insert into comments (postId, author, content) values (1, 'Bob', 'Nice!');
insert into comments (postId, author, content) values (2, 'Viola', 'Keep it up!');
insert into comments (postId, author, content) values (3, 'Marc', 'So far so good!');
