-- name: CreateTransaksi :one
INSERT INTO daftar_transaksi (
    jenis_transaksi,
    nominal,
    nomor_rekening
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetDaftarTransaksi :many
SELECT * FROM daftar_akun 
WHERE nomor_rekening = $1;

