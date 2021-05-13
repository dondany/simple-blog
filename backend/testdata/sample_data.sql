DROP TABLE IF EXISTS posts;
CREATE TABLE posts
(
    id SERIAL NOT NULL,
	title varchar(200) not null,
    content varchar(2000) NOT NULL,
    likes int default 0
)

insert into posts (title, content, likes) values ('First Post', 'Yo!', 1);
insert into posts (title, content, likes) values ('Second Post', 'Wow, it is the second post already!', 5);
insert into posts (title, content, likes) values ('Third Post', 'Ok, third one is here!', 12);

DROP TABLE IF EXISTS comments;
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
