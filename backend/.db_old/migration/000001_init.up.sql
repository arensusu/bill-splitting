CREATE TABLE "users" (
  "id" varchar PRIMARY KEY,
  "username" varchar(255) UNIQUE NOT NULL
);

CREATE TABLE "groups" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(255) NOT NULL
);

CREATE TABLE "members" (
  "id" SERIAL PRIMARY KEY,
  "group_id" int NOT NULL,
  "user_id" varchar NOT NULL
);

CREATE TABLE "group_invitations" (
  "code" varchar PRIMARY KEY,
  "group_id" int NOT NULL
);

CREATE TABLE "expenses" (
  "id" SERIAL PRIMARY KEY,
  "member_id" int NOT NULL,
  "amount" decimal NOT NULL,
  "description" varchar(255) NOT NULL,
  "date" timestamptz NOT NULL,
  "is_settled" bool DEFAULT false NOT NULL
);

CREATE TABLE "settlements" (
  "payer_id" int,
  "payee_id" int,
  "amount" decimal NOT NULL,
  PRIMARY KEY ("payer_id", "payee_id"),
  CHECK ("payer_id" != "payee_id")
);

ALTER TABLE "members" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id");

ALTER TABLE "members" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "group_invitations" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id");

ALTER TABLE "expenses" ADD FOREIGN KEY ("member_id") REFERENCES "members" ("id");

ALTER TABLE "settlements" ADD FOREIGN KEY ("payer_id") REFERENCES "members" ("id");

ALTER TABLE "settlements" ADD FOREIGN KEY ("payee_id") REFERENCES "members" ("id");
