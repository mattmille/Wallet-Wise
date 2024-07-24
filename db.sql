BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "category" (
	"category_id"	INTEGER,
	"category_name"	TEXT NOT NULL UNIQUE,
	"category_limit"	REAL NOT NULL,
	PRIMARY KEY("category_id" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "expense" (
	"expense_id"	INTEGER,
	"expense_epoch"	REAL NOT NULL,
	"expense_amount"	REAL NOT NULL,
	"expense_description"	TEXT,
	"expense_category_id"	INTEGER,
	PRIMARY KEY("expense_id" AUTOINCREMENT),
	FOREIGN KEY("expense_category_id") REFERENCES "category"("category_id")
);
INSERT INTO "category" VALUES (1,'Fast Food',300.59);
INSERT INTO "category" VALUES (2,'Online Shopping',150.23);
INSERT INTO "expense" VALUES (1,1641043200,23.45,'McDonald''s',1);
INSERT INTO "expense" VALUES (2,1647614400,95.23,'Clothes',2);
COMMIT;
