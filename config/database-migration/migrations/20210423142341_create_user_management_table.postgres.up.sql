CREATE TABLE public.users (
	id serial NOT NULL,
	name varchar(80) NOT NULL,
  username varchar(80) NOT NULL,
	email varchar(255) NOT NULL,
	"password" varchar(255) NOT NULL,
  is_active bool NOT NULL DEFAULT true,
	created_at timestamp with time zone NULL,
	updated_at timestamp with time zone NULL,
	deleted_at timestamp with time zone NULL,
  CONSTRAINT users_pkey PRIMARY KEY (id),
  CONSTRAINT users_username_unique UNIQUE (username),
	CONSTRAINT users_email_unique UNIQUE (email)
);
