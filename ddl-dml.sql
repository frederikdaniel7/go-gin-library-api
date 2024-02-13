create table books (
	id bigserial primary key,
	title varchar not null,
	book_description varchar not null,
	quantity int not null,
	cover varchar not null,
	author_id bigint not null,
	created_at timestamp not null default now(),
	updated_at timestamp not null default now(),
	deleted_at timestamp,
	foreign key(author_id) references author(id)
);


create table author (
	id bigserial primary key,
	author_name varchar not null
);

drop table books;

insert into books (title, book_description,quantity,cover, author_id) values
('a','desc a', 100, 'soft', 1),
('b','desc b', 150, 'hard', 1),
('c','desc c', 200, 'hard', 2),
('d','desc d', 250, 'soft', 2),
('e','desc e', 300, 'soft', 3),
('f','desc f', 350, 'hard', 3),
('g','desc g', 400, 'soft', 2),
('h','desc h', 450, 'soft', 2),
('i','desc i', 500, 'soft', 1),
('j','desc j', 550, 'soft', 1);