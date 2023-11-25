CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar(255) UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "groups" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(255) UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "group_members" (
  "group_id" int,
  "user_id" int,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("group_id", "user_id")
);

CREATE TABLE "expenses" (
  "id" bigserial PRIMARY KEY,
  "group_id" int NOT NULL,
  "payer_id" int NOT NULL,
  "amount" decimal NOT NULL,
  "description" varchar(255) NOT NULL,
  "date" timestamptz NOT NULL
);

CREATE TABLE "user_expenses" (
  "expense_id" int,
  "user_id" int,
  "share" decimal NOT NULL,
  PRIMARY KEY ("expense_id", "user_id")
);

CREATE TABLE "settlements" (
  "id" bigserial PRIMARY KEY,
  "payer_id" int NOT NULL,
  "payee_id" int NOT NULL,
  "amount" decimal NOT NULL,
  "date" timestamptz NOT NULL
  CHECK (payer_id != payee_id)
);

ALTER TABLE "group_members" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id");

ALTER TABLE "group_members" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "expenses" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id");

ALTER TABLE "expenses" ADD FOREIGN KEY ("payer_id") REFERENCES "users" ("id");

ALTER TABLE "user_expenses" ADD FOREIGN KEY ("expense_id") REFERENCES "expenses" ("id");

ALTER TABLE "user_expenses" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "settlements" ADD FOREIGN KEY ("payer_id") REFERENCES "users" ("id");

ALTER TABLE "settlements" ADD FOREIGN KEY ("payee_id") REFERENCES "users" ("id");
