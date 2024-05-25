
-- name: GetPelanggan :one
SELECT * FROM daftar_Pelanggan
WHERE id = $1;