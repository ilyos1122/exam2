CREATE TABLE "category" (
    "id" UUID NOT NULL PRIMARY KEY,
    "category_title" VARCHAR(46) NOT NULL,
    "image" VARCHAR,
    "parent_id" UUID REFERENCES "category" ("id"),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "product" (
    "id" UUID NOT NULL PRIMARY KEY,
    "product_id" VARCHAR NOT NULL,
    "title" VARCHAR(46) NOT NULL,
    "description" VARCHAR NOT NULL,
    "price" NUMERIC NOT NULL,
    "product_image" VARCHAR(255) NOT NULL,
    "category_id" UUID NOT NULL REFERENCES category(id),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE clients (
    "id" UUID NOT NULL PRIMARY KEY,
    "first_name" VARCHAR(32),
    "last_name" VARCHAR(64),
    "phone" VARCHAR(20),
    "photo" VARCHAR,
    "date_of_birth" VARCHAR,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE branches(
    "id" UUID NOT NULL PRIMARY KEY,
    "name" VARCHAR(64),
    "phone" VARCHAR(20),
    "photo" VARCHAR,
    "work_start_hour" VARCHAR,
    "work_end_hour" VARCHAR ,
    "address" VARCHAR,
    "delivery_price" NUMERIC DEFAULT 10000,
    "status" VARCHAR DEFAULT 'active',
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE orders(
    "id" UUID NOT NULL PRIMARY KEY, 
    "order_id" VARCHAR NOT NULL,
    "client_id" UUID  NOT NULL REFERENCES clients(id),
    "branch_id" UUID NOT NULL REFERENCES branches(id),
    "delivery_address" VARCHAR NOT NULL,
    "delivery_price" NUMERIC default 10000,
    "total_count" BIGINT,
    "total_price" NUMERIC,
    "status" VARCHAR(20) DEFAULT 'new',
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP

);

CREATE TABLE order_products(
    "order_product_id" UUID NOT NULL PRIMARY KEY,
    "order_id" VARCHAR NOT NULL,
    "product_id" UUID NOT NULL REFERENCES product(id),
    "discount_typ" VARCHAR,
    "discount_amount" NUMERIC,
    "quantity" INT,
    "sum" NUMERIC,
    "price" numeric,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


