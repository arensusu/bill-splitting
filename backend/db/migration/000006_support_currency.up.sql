ALTER TABLE "groups" ADD COLUMN "currency" varchar(10) DEFAULT "TWD";

ALTER TABLE "expenses"
ADD COLUMN "origin_amount" decimal,
ADD COLUMN "origin_currency" varchar(10) DEFAULT "TWD";

UPDATE "expenses" SET "origin_amount" = "amount";

