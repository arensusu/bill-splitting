ALTER TABLE "expenses"
DROP COLUMN "origin_amount" decimal,
DROP COLUMN "origin_currency" varchar(10) DEFAULT "TWD";

ALTER TABLE "groups" DROP COLUMN "currency" varchar(10) DEFAULT "TWD";
