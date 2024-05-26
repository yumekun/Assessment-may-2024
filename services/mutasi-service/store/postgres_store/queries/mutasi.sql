-- name: CreateMutasi :one
INSERT INTO mutasi (
    id,
    jenis_transaksi,
    nominal,
    nomor_rekening
) VALUES (
     $1, $2, $3,$4
) RETURNING *;




