CREATE TABLE "daftar_pelanggan" (
  "id" bigserial PRIMARY KEY,
  "nama" varchar NOT NULL,
  "nik" varchar UNIQUE NOT NULL,
  "nomor_hp" varchar UNIQUE NOT NULL,
  "pin" varchar NOT NULL
);

CREATE TABLE "daftar_akun" (
  "id" bigserial,
  "id_pelanggan" bigint NOT NULL,
  "nomor_rekening" varchar PRIMARY KEY,
  "saldo" bigint NOT NULL
);

CREATE TABLE "daftar_transaksi" (
  "id" bigserial PRIMARY KEY,
  "nomor_rekening" varchar NOT NULL,
  "jenis_transaksi" varchar(1) NOT NULL,
  "nominal" bigint NOT NULL
);

CREATE TABLE "mutasi" (
  "id" bigserial PRIMARY KEY,
  "nomor_rekening" varchar NOT NULL,
  "jenis_transaksi" varchar(1) NOT NULL,
  "nominal" bigint NOT NULL
);


ALTER TABLE "daftar_akun" ADD FOREIGN KEY ("id_pelanggan") REFERENCES "daftar_pelanggan" ("id");
ALTER TABLE "daftar_transaksi" ADD FOREIGN KEY ("nomor_rekening") REFERENCES "daftar_akun" ("nomor_rekening");