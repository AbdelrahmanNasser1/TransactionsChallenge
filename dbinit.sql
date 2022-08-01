CREATE TABLE transaction(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  amount INT NOT NULL ,
  currency STRING NOT NULL,
  createdAt STRING NOT NULL
);