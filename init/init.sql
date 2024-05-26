

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
  "id" bigserial PRIMARY KEY,
  "nomor_rekening" varchar(255) NOT NULL,
  "jenis_transaksi" varchar(255) NOT NULL,
  "nominal" bigint NOT NULL
);