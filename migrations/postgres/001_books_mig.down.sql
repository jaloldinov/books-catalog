DROP TABLE IF EXISTS "book"
DROP TABLE IF EXISTS "book_category"
DROP TABLE IF EXISTS "author"

ALTER TABLE "book" DROP CONSTRAINT IF EXISTS "fk_book_category";

ALTER TABLE "book" DROP CONSTRAINT IF EXISTS "fk_book_author";

