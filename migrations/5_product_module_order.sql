-- +migrate Up
CREATE TABLE "card"
(
    "order_id" TEXT NOT NULL,
    "product_id" TEXT NOT NULL,
    "product_name" TEXT NOT NULL,
    "product_image" TEXT NOT NULL,
    "quantity" INTEGER DEFAULT 1 NOT NULL,
    "price" NUMERIC NOT NULL
);
CREATE TABLE "orders"
(
    "user_id" TEXT NOT NULL,
    "order_id" TEXT NOT NULL,
    "status" TEXT NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NOT NULL
);
-- +migrate Down
DROP TABLE orders;
DROP TABLE card;