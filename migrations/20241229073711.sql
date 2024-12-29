-- Create "meme_coins" table
CREATE TABLE "public"."meme_coins" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" character varying(50) NOT NULL,
  "description" text NULL,
  "popularity_score" bigint NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_meme_coins_name" UNIQUE ("name")
);
-- Create index "idx_meme_coins_deleted_at" to table: "meme_coins"
CREATE INDEX "idx_meme_coins_deleted_at" ON "public"."meme_coins" ("deleted_at");
