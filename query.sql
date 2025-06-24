

CREATE TABLE public.users (
	user_id bigserial NOT NULL,
	first_name text NULL,
	last_name text NULL,
	email text NOT NULL,
	address text NULL,
	date_of_birth date NULL,
	"password" text NOT NULL,
	last_login_date timestamptz NULL,
	jwt_token text NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT uni_users_email UNIQUE (email),
	CONSTRAINT users_pkey PRIMARY KEY (user_id)
);

CREATE TABLE public.genres (
	genre_id bigserial NOT NULL,
	"name" text NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT genres_pkey PRIMARY KEY (genre_id)
);

CREATE TABLE public.authors (
	author_id bigserial NOT NULL,
	"name" text NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT authors_pkey PRIMARY KEY (author_id)
);

CREATE TABLE public.books (
	book_id bigserial NOT NULL,
	genre_id int8 NULL,
	author_id int8 NULL,
	title text NULL,
	description text NULL,
	min_age_restriction int8 NULL,
	cover_url text NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT books_pkey PRIMARY KEY (book_id),
	CONSTRAINT fk_authors_books FOREIGN KEY (author_id) REFERENCES public.authors(author_id),
	CONSTRAINT fk_genres_books FOREIGN KEY (genre_id) REFERENCES public.genres(genre_id)
);

CREATE TABLE public.loans (
	loan_id bigserial NOT NULL,
	user_id int8 NULL,
	book_id int8 NULL,
	loan_date date NULL,
	due_date date NULL,
	return_date date NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT loans_pkey PRIMARY KEY (loan_id),
	CONSTRAINT fk_books_loans FOREIGN KEY (book_id) REFERENCES public.books(book_id),
	CONSTRAINT fk_users_loans FOREIGN KEY (user_id) REFERENCES public.users(user_id)
);