CREATE TABLE IF NOT EXISTS "order_info" (
  "order_uid" varchar PRIMARY KEY NOT NULL,
  "track_number" varchar NOT NULL,
  "entry" varchar,
  "delivery" jsonb,
  "items" json,
  "locale" varchar,
  "internal_signature" varchar,
  "customer_id" varchar,
  "delivery_service" varchar,
  "shardkey" varchar,
  "sm_id" int,
  "date_created" timestamptz NOT NULL DEFAULT (now()),
  "oof_shard" varchar
);

CREATE TABLE IF NOT EXISTS "payment" (
  "transaction" varchar PRIMARY KEY NOT NULL,
  "request_id" varchar,
  "currency" varchar,
  "provider" varchar,
  "amount" int,
  "payment_dt" bigint,
  "bank" varchar,
  "delivery_cost" int,
  "goods_total" int,
  "custom_fee" int
);

ALTER TABLE "payment" ADD FOREIGN KEY ("transaction") REFERENCES "order_info" ("order_uid");
