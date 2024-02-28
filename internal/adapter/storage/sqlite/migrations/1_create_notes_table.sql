CREATE TABLE "notes" (
	"id"	TEXT NOT NULL UNIQUE,
	"title"	TEXT NOT NULL,
	"content"	TEXT NOT NULL,
	"encrypted"	NUMERIC NOT NULL,
	"encryption_key"	BLOB,
	PRIMARY KEY("id")
);
