

-- Create the daftar_pelanggan table
CREATE TABLE IF NOT EXISTS "daftar_pelanggan" (
  "id" bigserial PRIMARY KEY,
  "nama" varchar(255) NOT NULL,
  "nik" varchar(255) UNIQUE NOT NULL,
  "nomor_hp" varchar(255) UNIQUE NOT NULL,
  "pin" varchar(255) NOT NULL
);

-- Create the daftar_akun table
CREATE TABLE IF NOT EXISTS "daftar_akun" (
  "id" bigserial PRIMARY KEY,
  "id_pelanggan" bigint NOT NULL,
  "nomor_rekening" varchar(255) UNIQUE NOT NULL,
  "saldo" bigint NOT NULL,
  FOREIGN KEY ("id_pelanggan") REFERENCES "daftar_pelanggan" ("id")
);

-- Create the daftar_transaksi table
CREATE TABLE IF NOT EXISTS "daftar_transaksi" (
  "id" bigserial PRIMARY KEY,
  "nomor_rekening" varchar(255) NOT NULL,
  "jenis_transaksi" char(1) NOT NULL,
  "nominal" bigint NOT NULL,
  FOREIGN KEY ("nomor_rekening") REFERENCES "daftar_akun" ("nomor_rekening")
);

-- Create the mutasi table
CREATE TABLE IF NOT EXISTS "mutasi" (
  "id" bigserial PRIMARY KEY,
  "nomor_rekening" varchar(255) NOT NULL,
  "jenis_transaksi" char(1) NOT NULL,
  "nominal" bigint NOT NULL
);