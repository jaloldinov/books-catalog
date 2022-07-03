CREATE TABLE "author" (
  "id" varchar PRIMARY KEY,
  "firstname" varchar NOT NULL,
  "lastname" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "book_category" (
  "id" varchar PRIMARY KEY,
  "category_name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "book" (
  "id" varchar PRIMARY KEY,
  "book_name" varchar NOT NULL,
  "author_id" varchar NOT NULL,
  "category_id" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "book" ADD FOREIGN KEY ("author_id") REFERENCES "author" ("id");

ALTER TABLE "book" ADD FOREIGN KEY ("category_id") REFERENCES "book_category" ("id");
