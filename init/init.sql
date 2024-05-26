

-- Create the daftar_pelanggan table
CREATE TABLE  "daftar_pelanggan" (
  "id" varchar PRIMARY KEY,
  "nama" varchar(255) NOT NULL,
  "nik" varchar(255) UNIQUE NOT NULL,
  "nomor_hp" varchar(255) UNIQUE NOT NULL,
  "pin" varchar(255) NOT NULL
);

-- Create the daftar_akun table
CREATE TABLE  "daftar_akun" (
  "id" varchar PRIMARY KEY,
  "id_pelanggan" varchar NOT NULL,
  "nomor_rekening" varchar(255) UNIQUE NOT NULL,
  "saldo" bigint NOT NULL,
  FOREIGN KEY ("id_pelanggan") REFERENCES "daftar_pelanggan" ("id")
);

-- Create the daftar_transaksi table
CREATE TABLE  "daftar_transaksi" (
  "id" varchar PRIMARY KEY,
  "nomor_rekening" varchar(255) NOT NULL,
  "jenis_transaksi" varchar(255) NOT NULL,
  "nominal" bigint NOT NULL,
  FOREIGN KEY ("nomor_rekening") REFERENCES "daftar_akun" ("nomor_rekening")
);

-- Create the mutasi table
CREATE TABLE  "mutasi" (
  "id" varchar PRIMARY KEY,
  "nomor_rekening" varchar(255) NOT NULL,
  "jenis_transaksi" varchar(255) NOT NULL,
  "nominal" bigint NOT NULL
);

INSERT INTO "public"."daftar_pelanggan" ("id","nama","nik","nomor_hp","pin") VALUES ('001','person A','nomor235','yeywiir','2301');
INSERT INTO "public"."daftar_pelanggan" ("id","nama","nik","nomor_hp","pin") VALUES ('002','person B','nomor236','yeywieeew','1000');
INSERT INTO "public"."daftar_akun" ("id","id_pelanggan","nomor_rekening","saldo") VALUES ('001','001','A001',500000);
INSERT INTO "public"."daftar_akun" ("id","id_pelanggan","nomor_rekening","saldo") VALUES ('002','002','A002',300000);