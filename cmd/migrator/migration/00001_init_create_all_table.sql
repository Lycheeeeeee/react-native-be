-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE "public"."accounts" (
  "id" uuid NOT NULL PRIMARY KEY,
  "created_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
 
  "user_name" varchar(200),
  "digest_password" varchar,
  "type" varchar(200)
);

CREATE TABLE "public"."shops" (
  "id" uuid NOT NULL PRIMARY KEY,
  "created_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,
  
  "name" varchar(200),
  "address" varchar
);

CREATE TABLE "public"."drinks" (
  "id" uuid NOT NULL PRIMARY KEY,
  "created_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,

  "name" varchar(200),
  "description" varchar,
  "price" SMALLINT
);

CREATE TABLE "public"."details" (
  "id" uuid NOT NULL PRIMARY KEY,
  "created_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,

  "quantity" SMALLINT,
  "drink_id" uuid REFERENCES drinks(id)
);

CREATE TABLE "public"."orders" (
  "id" uuid NOT NULL PRIMARY KEY,
  "created_at" timestamptz DEFAULT now(),
  "deleted_at" timestamptz,

 "status" varchar(50),
 "order_time" timestamptz,
 "receive_time" smallint,
 "shop_id" uuid REFERENCES shops(id),
 "account_id" uuid REFERENCES accounts(id),
 "detail_id" uuid REFERENCES details(id),
 "total_price" INTEGER
);

ALTER TABLE details ADD COLUMN order_id uuid;
ALTER TABLE "public"."details" ADD CONSTRAINT "details_order_id_fkey" FOREIGN KEY ("order_id") REFERENCES "public"."orders" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE orders ADD COLUMN simple_id varchar(50);
ALTER TABLE shops ADD COLUMN staff_account_id uuid;
ALTER TABLE "public"."shops" ADD CONSTRAINT "shops_fkey" FOREIGN KEY ("staff_account_id") REFERENCES "public"."accounts" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
INSERT INTO accounts (id,user_name,digest_password,type) VALUES 
('fe9c7f84-621b-48cf-8132-c79d13c10201','admin1','1111','staff'),
('fe9c7f84-621b-48cf-8132-c79d13c10202','admin2','2222','staff'),
('fe9c7f84-621b-48cf-8132-c79d13c10203','admin3','3333','staff');
INSERT INTO shops(id,name,address,staff_account_id) VALUES
('fe9c7f84-621b-48cf-8132-c79d13c10204','branch1','60-62 Cach Mang Thang 8 Phuong 6 Q3','fe9c7f84-621b-48cf-8132-c79d13c10201'),
('fe9c7f84-621b-48cf-8132-c79d13c10205','branch2','27 Han Thuyen Phuong Ben Nghe Q1','fe9c7f84-621b-48cf-8132-c79d13c10202'),
('fe9c7f84-621b-48cf-8132-c79d13c10206','branch3','97 Hai Ba Trung Ben Nghe Q1','fe9c7f84-621b-48cf-8132-c79d13c10203');
ALTER TABLE drinks ADD COLUMN assets varchar(50);
ALTER TABLE drinks ALTER COLUMN price SET DATA TYPE integer; 
INSERT INTO drinks (id,name,description,price,assets) VALUES ('e78cd1ba-2b45-42cb-9dd8-3ae6f14f93e1','Mocha Flavored','Sweet, ripe banana flavor and rich caramet',45000,'./assets/d1.jpg'),
('e78cd1ba-2b45-42cb-9dd8-3ae6f14f93e2','Cinnamon','Freshly ground cinnamon with a touch of nutmeg, ripe banana flavor and rich caramet',50000,'./assets/d2.jpg'),
('e78cd1ba-2b45-42cb-9dd8-3ae6f14f93e3','White Coffee','Oreos and vanilla ice cream',43000,'./assets/d3.jpg'),
('e78cd1ba-2b45-42cb-9dd8-3ae6f14f93e4','Dark Chocolate','Made with a single origin French chocolate',38000,'./assets/d4.jpg'),
('e78cd1ba-2b45-42cb-9dd8-3ae6f14f93e5','Espresso','Vanilla-roasted peaches, mixed with smooth creamy ice cream',50000,'./assets/d5.jpg'),
('e78cd1ba-2b45-42cb-9dd8-3ae6f14f93e6','Raspberry Swirl','Raspberry sorbet with creamy, rich vanilla ice cream',47000,'./assets/d6.jpg'),
('e78cd1ba-2b45-42cb-9dd8-3ae6f14f93e7','Strawberry','Fresh strawberries sweetened and blended with ice cream',64000,'./assets/d7.jpg'),
('e78cd1ba-2b45-42cb-9dd8-3ae6f14f93e8','Amarula Creme Brulee Coffee','Amarula Cream Liqueur, butterscotch schnapps,coffee and whipped cream',46000,'./assets/d8.jpg'),
('e78cd1ba-2b45-42cb-9dd8-3ae6f14f93e9','Monte Cristo Coffee','Kahlua, Grand Marnier, and coffee topped with whipped cream',42000,'./assets/d9.jpg'),
('e88cd1ba-2b45-42cb-9dd8-3ae6f14f93e1','Spanish Coffee','Tia Maria, St. Remy Brandy and coffee topped with whipped cream',38000,'./assets/d10.jpg'),
('e88cd1ba-2b45-42cb-9dd8-3ae6f14f93e2','Mocha Italia','Disaronno Amaretto, espresso and chocolate topped with whipped cream',45000,'./assets/d11.jpg');
INSERT INTO orders (id,status,order_time,receive_time,shop_id,account_id,simple_id) VALUES ('45dd60e4-82ba-4205-966b-0df3df198a9d','uncheck', '2019-01-14 15:31:32.72994+00','30','fe9c7f84-621b-48cf-8132-c79d13c10204','fe9c7f84-621b-48cf-8132-c79d13c10201','branch1');
INSERT INTO details(id,order_id,drink_id,quantity) VALUES ('a7169961-22b5-4e9c-bd41-a11c5db7b17d','45dd60e4-82ba-4205-966b-0df3df198a9d','e78cd1ba-2b45-42cb-9dd8-3ae6f14f93e1','12'),
('64f03071-0824-43ad-b74b-51bcf9e7555a','45dd60e4-82ba-4205-966b-0df3df198a9d','e78cd1ba-2b45-42cb-9dd8-3ae6f14f93e2','10'),
('64f03071-0824-43ad-b74b-51bcf9e7555b','45dd60e4-82ba-4205-966b-0df3df198a9d','e78cd1ba-2b45-42cb-9dd8-3ae6f14f93e3','14'),
('64f03071-0824-43ad-b74b-51bcf9e7555c','45dd60e4-82ba-4205-966b-0df3df198a9d','e78cd1ba-2b45-42cb-9dd8-3ae6f14f93e4','7'),
('64f03071-0824-43ad-b74b-51bcf9e7555d','45dd60e4-82ba-4205-966b-0df3df198a9d','e78cd1ba-2b45-42cb-9dd8-3ae6f14f93e5','6'),
('64f03071-0824-43ad-b74b-51bcf9e7555e','45dd60e4-82ba-4205-966b-0df3df198a9d','e78cd1ba-2b45-42cb-9dd8-3ae6f14f93e6','5'),
('64f03071-0824-43ad-b74b-51bcf9e7555f','45dd60e4-82ba-4205-966b-0df3df198a9d','e78cd1ba-2b45-42cb-9dd8-3ae6f14f93e7','4'),
('ff2c516a-0687-498f-9fb9-7c6a7ec32c0a','45dd60e4-82ba-4205-966b-0df3df198a9d','e78cd1ba-2b45-42cb-9dd8-3ae6f14f93e8','16'),
('e563de23-df3f-4dfc-b47d-e6bc6a52a8cb','45dd60e4-82ba-4205-966b-0df3df198a9d','e78cd1ba-2b45-42cb-9dd8-3ae6f14f93e9','18'),
('ecaff535-8865-4081-974d-7f125a1c864b','45dd60e4-82ba-4205-966b-0df3df198a9d','e88cd1ba-2b45-42cb-9dd8-3ae6f14f93e1','20'),
('455cdb93-835f-4ef5-a2a1-0238f078c3a4','45dd60e4-82ba-4205-966b-0df3df198a9d','e88cd1ba-2b45-42cb-9dd8-3ae6f14f93e2','10'),
('68ef9774-9280-480f-9b19-cf4550962482','45dd60e4-82ba-4205-966b-0df3df198a9d','e88cd1ba-2b45-42cb-9dd8-3ae6f14f93e2','17');



-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP SCHEMA public CASCADE;
CREATE SCHEMA public;

