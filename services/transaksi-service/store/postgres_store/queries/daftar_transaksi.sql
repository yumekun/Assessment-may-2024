-- name: CreateTransaksi :one
INSERT INTO daftar_transaksi (
    id,
    jenis_transaksi,
    nominal,
    nomor_rekening
) VALUES (
    $1, $2, $3,$4
) RETURNING *;

-- name: GetDaftarTransaksi :many
SELECT * FROM daftar_akun 
WHERE nomor_rekening = $1;

