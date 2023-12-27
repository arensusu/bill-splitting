CREATE TABLE "users" (
  "id" varchar PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "groups" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(255) UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "group_invitations" (
  "code" varchar PRIMARY KEY,
  "group_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "group_members" (
  "group_id" bigint,
  "user_id" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("group_id", "user_id")
);

CREATE TABLE "expenses" (
  "id" bigserial PRIMARY KEY,
  "group_id" bigint NOT NULL,
  "payer_id" varchar NOT NULL,
  "amount" bigint NOT NULL,
  "description" varchar(255) NOT NULL,
  "date" timestamptz NOT NULL,
  "is_settled" boolean NOT NULL DEFAULT false
);

CREATE TABLE "user_expenses" (
  "expense_id" bigint,
  "user_id" varchar,
  "share" bigint NOT NULL,
  PRIMARY KEY ("expense_id", "user_id")
);

CREATE TABLE "settlements" (
  "group_id" bigserial NOT NULL,
  "payer_id" varchar NOT NULL,
  "payee_id" varchar NOT NULL,
  "amount" bigint NOT NULL,
  "is_confirmed" boolean NOT NULL DEFAULT false,
  PRIMARY KEY ("group_id", "payer_id", "payee_id"),
  CHECK (payer_id != payee_id)
);

ALTER TABLE "group_invitations" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id");

ALTER TABLE "group_members" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id");

ALTER TABLE "group_members" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "expenses" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id");

ALTER TABLE "expenses" ADD FOREIGN KEY ("payer_id") REFERENCES "users" ("id");

ALTER TABLE "user_expenses" ADD FOREIGN KEY ("expense_id") REFERENCES "expenses" ("id");

ALTER TABLE "user_expenses" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "settlements" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id");

ALTER TABLE "settlements" ADD FOREIGN KEY ("payer_id") REFERENCES "users" ("id");

ALTER TABLE "settlements" ADD FOREIGN KEY ("payee_id") REFERENCES "users" ("id");
