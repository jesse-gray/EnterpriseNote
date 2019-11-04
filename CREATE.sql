--Need to have a database "enterprisedb"
--Drop all tables to start fresh
DROP SCHEMA public CASCADE;
CREATE SCHEMA public;

GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO public;

--CREATE tables
CREATE TABLE "user" (
	user_id					int					GENERATED ALWAYS AS IDENTITY,
	user_first_name			varchar(25)			NOT NULL,
	user_last_name			varchar(25),
	cookie_id				varchar(100),
	user_password			varchar(255)			NOT NULL,
	CONSTRAINT pk_user_id PRIMARY KEY (user_id)
)WITH ( 
  OIDS=FALSE 
);

CREATE TABLE note (
	note_id				int			GENERATED ALWAYS AS IDENTITY,
	note_text			text		NOT NULL,
	author_id			int			REFERENCES "user"(user_id) ON DELETE CASCADE,
	CONSTRAINT pk_note_id PRIMARY KEY (note_id) 
)WITH ( 
  OIDS=FALSE 
);

CREATE TABLE permissions (
	note_id				int			REFERENCES note(note_id) ON DELETE CASCADE,
	user_id				int			REFERENCES "user"(user_id) ON DELETE CASCADE,
	read_permission		bool		NOT NULL		DEFAULT false,
	write_permission	bool		NOT NULL		DEFAULT false,	
	CONSTRAINT pk_note_id_user_id PRIMARY KEY (note_id, user_id) 
)WITH ( 
  OIDS=FALSE 
);

CREATE TABLE favourites (
	author_id			int			REFERENCES "user"(user_id) ON DELETE CASCADE,
	favourite_id		int			REFERENCES "user"(user_id) ON DELETE CASCADE,	
	read_permission		bool		NOT NULL		DEFAULT false,
	write_permission	bool		NOT NULL		DEFAULT false,	
	CONSTRAINT pk_user_id_favourite_id PRIMARY KEY (author_id, favourite_id) 
)WITH ( 
  OIDS=FALSE 
);

--Mock Data
INSERT INTO "user" (user_first_name, user_last_name, user_password)
VALUES 	('John', 'Acers', 'password'),
		('James', 'Greene', 'password'),
		('Tony', 'Hions', 'password'),
		('Charlotte', 'Huffing', 'password'),
		('Tim', 'Odens', 'password'),
		('Sharon', 'Tomkins', 'password');
		
INSERT INTO note (note_text, author_id)
VALUES 	('This is sample text for the first note', 1),
		('This is some more sample text, however this is for the second note', 2),
		('More text for more notes', 4),
		('Random words for another note', 3),
		('Small note', 6),
		('Its hard to make fake notes', 3),
		('Is this the last note?', 5),
		('it was not, but perhaps this one is.', 1);
		
INSERT INTO permissions
VALUES 	(2, 1, true, false),
		(2, 3, true, true),
		(2, 4, true, false),
		(3, 1, true, true),
		(4, 1, true, false),
		(4, 6, true, false),
		(5, 1, true, false),
		(5, 2, true, true),
		(5, 3, true, false),
		(5, 4, true, false),
		(5, 5, true, true),
		(6, 2, true, false),
		(6, 4, true, false),
		(7, 6, true, false),
		(7, 2, true, false),
		(7, 3, true, true),
		(8, 3, true, false),
		(8, 5, true, true);

INSERT INTO favourites
VALUES 	(1, 3, true, false),
		(1, 5, true, true),
		(1, 6, true, true),
		(2, 1, true, false),
		(2, 5, true, true),
		(3, 2, true, false),
		(4, 1, true, true),
		(4, 2, true, true),
		(4, 6, true, true),
		(5, 2, true, false),
		(5, 4, true, true),
		(6, 1, true, false),
		(6, 2, true, true),
		(6, 3, true, false),
		(6, 5, true, true);
