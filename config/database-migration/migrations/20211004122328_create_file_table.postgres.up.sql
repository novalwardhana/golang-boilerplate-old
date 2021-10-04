CREATE TABLE public.files (
  id serial NOT NULL,
  directory varchar(255) NOT NULL,
  name varchar(255) NOT NULL,
  size varchar(255) NOT NULL,
  created_at timestamp with time zone NULL,
  updated_at timestamp with time zone NULL,
  CONSTRAINT files_pkey PRIMARY KEY (id)
);