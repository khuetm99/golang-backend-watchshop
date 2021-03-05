-- +migrate Up

CREATE TABLE "products"
(
    "product_id" text NOT NULL PRIMARY KEY,
    "product_name" text NOT NULL,
    "quantity" integer NOT NULL,
    "sold_items" integer NOT NULL,
    "price" numeric NOT NULL,
    "cate_id" text NOT NULL,
    "product_image" text NOT NULL,
    "product_des" text NOT NULL,
    "deleted_at" timestamp with time zone,
    "created_at" timestamp with time zone default NOW(),
    "updated_at" timestamp with time zone default NOW(),
);


CREATE TABLE "collections" (
  "collection_id" varchar(255) PRIMARY KEY,
  "collection_name" varchar(255) NOT NULL,
  "created_at" TIMESTAMP NOT NULL,
  "updated_at" TIMESTAMP NOT NULL
);


-- +migrate Down
DROP TABLE products;
DROP TABLE collections;
DROP TABLE attributes;