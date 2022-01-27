CREATE DATABASE go_practice;

\c go_practice

CREATE TABLE bank_codes (
  id SERIAL PRIMARY KEY,
  code INTEGER NOT NULL,
  name VARCHAR(45) NOT NULL
);

CREATE TABLE expedition_countries (
  id SERIAL PRIMARY KEY,
  code INTEGER NOT NULL,
  name VARCHAR(45) NOT NULL
);

CREATE TABLE product_segments (
  id SERIAL PRIMARY KEY,
  code INTEGER NOT NULL,
  name VARCHAR(45) NOT NULL
);

CREATE TABLE account_holders (
  id SERIAL PRIMARY KEY,
  name VARCHAR(150) NOT NULL,
  email VARCHAR(255) NOT NULL,
  alias VARCHAR(45) NOT NULL
);

CREATE TABLE credit_cards (
  id SERIAL PRIMARY KEY,
  card_number VARCHAR(16) NOT NULL,
  id_account_holder INTEGER NOT NULL,
  id_bank_code INTEGER NOT NULL,
  id_product_segment INTEGER NOT NULL,
  id_expedition_country INTEGER NOT NULL,
  CONSTRAINT fk_account_holder
      FOREIGN KEY(id_account_holder)
      REFERENCES account_holders(id),
  CONSTRAINT fk_bank_code
      FOREIGN KEY(id_bank_code)
      REFERENCES bank_codes(id),
  CONSTRAINT fk_product_segment
      FOREIGN KEY(id_product_segment)
      REFERENCES product_segments(id),
  CONSTRAINT fk_expedition_country
      FOREIGN KEY(id_expedition_country)
      REFERENCES expedition_countries(id)
);




INSERT INTO bank_codes (code, name) VALUES
    (11, 'Banca Bunsan'),
    (22, 'Go Bank'),
    (33, 'Aldo Institución Financiera'),
    (44, 'Angel Banca Privada'),
    (55, 'Hugo Finantial Services'),
    (66, 'BanHumberto'),
    (77, 'Gustavo Commercial Bank');

INSERT INTO expedition_countries (code, name) VALUES
    (5211, 'México');

INSERT INTO product_segments (code, name) VALUES
    (2, 'Tarjeta básica'),
    (4, 'Tarjeta departamental'),
    (6, 'Tarjeta clásica'),
    (8, 'Tarjeta Oro'),
    (0, 'Tarjeta Platino');
