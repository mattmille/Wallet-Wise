BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "category" (
	"category_id"	INTEGER NOT NULL,
	"category_name"	TEXT NOT NULL UNIQUE,
	"category_limit"	INTEGER NOT NULL,
	PRIMARY KEY("category_id" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "expense" (
	"expense_id"	INTEGER,
	"expense_amount"	INTEGER NOT NULL DEFAULT 0,
	"expense_category_id"	INTEGER,
	"expense_description"	TEXT NOT NULL DEFAULT 'n/a',
	"expense_date"	INTEGER,
	PRIMARY KEY("expense_id" AUTOINCREMENT),
	FOREIGN KEY("expense_category_id") REFERENCES "category"("category_id")
);
INSERT INTO "category" VALUES (1,'Fast Food',300);
INSERT INTO "category" VALUES (2,'Online Shopping',150);
INSERT INTO "expense" VALUES (1,22,1,'McDonald''s',NULL);
INSERT INTO "expense" VALUES (2,38,2,'Clothes',NULL);
COMMIT;
