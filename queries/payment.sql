-- name: CreatePaymentLink :one
INSERT INTO PAYMENT_LINKS (
  DESCRIPTION,
  AMOUNT,
  REFERENCE
)VALUES (
  $1,
  $2,
  $3
) RETURNING *;

-- name: GetPaymentLink :one
SELECT * FROM PAYMENT_LINKS
WHERE reference = $1 LIMIT 1;
