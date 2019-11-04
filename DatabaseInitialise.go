package main

import (
	"log"
)

// construct the database using sql execution

func contructDB() {
	db := opendb() // assume database is already created we are just intialising the tables
	defer db.Close()

	_tables :=
		`--Need to have a database "enterprisedb"
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
	
	--Create a table on the program side to store passwords for login, passwords at this stage are unprotected
	
	CREATE TABLE pasword (
		pword_id int NOT NULL UNIQUE AUTO_INCREMENT,
		user_id int,
		password varchar(25),
		PRIMARY KEY (pword_id)
	);
	
	ALTER TABLE pword
		ADD CONSTRAINT 	user_id FOREIGN KEY (user_id)
		REFERENCES "user" (user_id);`

	_, err := db.Exec(_tables)
	if err != nil {
		log.Fatal()
	}

	_data :=
		`--Mock Data
	INSERT INTO "user" (user_first_name, user_last_name)
	VALUES 	('John', 'Smith'),
			('Sharon', 'Tomkins');
			
	INSERT INTO note (note_text, author_id)
	VALUES 	('This is sample text for the first note', 1),
			('This is some more sample text, however this is for the second note', 2);
			
	INSERT INTO permissions
	VALUES 	(2, 1, true, false);
	
	-- sample passwords
	INSERT INTO pword (user_id, pword)
	VAlUES 	(1, 'password'),
			(2, 'password');`

	_, err = db.Exec(_data)
	if err != nil {
		log.Fatal()
	}
}
