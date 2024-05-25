-- name: GetDaftarAkun :one
SELECT * FROM daftar_akun
WHERE nomor_rekening = $1;

-- name: UpdateSaldo :one
UPDATE daftar_akun
SET saldo = $2
WHERE nomor_rekening = $1
RETURNING *;