ALTER TABLE "expenses"
DROP COLUMN "origin_amount",
DROP COLUMN "origin_currency";

ALTER TABLE "groups" DROP COLUMN "currency";
