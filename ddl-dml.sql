create table books (
	id bigserial primary key,
	title varchar not null,
	book_description varchar not null,
	quantity int not null,
	cover varchar not null,
	created_at timestamp not null,
	updated_at timestamp not null,
	deleted_at timestamp
);

drop table books;

insert into books (title, book_description,quantity,cover, created_at,updated_at) values
('a', 'desc a', 100, "soft", now(), now()),
('b','desc b', 150, "hard", now(), now()),
('c','desc c', 200, "hard", now(), now()),
('d','desc d', 250, "soft", now(), now()),
('e','desc e', 300, "soft", now(), now()),
('f','desc f', 350, "hard", now(), now()),
('g','desc g', 400, "soft", now(), now()),
('h','desc h', 450, "soft", now(), now()),
('i','desc i', 500, "soft", now(), now()),
('j','desc j', 550, "soft", now(), now());
