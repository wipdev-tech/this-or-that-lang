-- name: GetAllVotes :many
SELECT
    votes.lang1_id,
    languages1.name      AS lang1_Name,
    languages1.logo_slug AS lang1_slug,
    votes.lang2_id,
    languages2.name      AS lang2_Name,
    languages2.logo_slug AS lang2_slug,
    votes.choice
FROM votes
INNER JOIN languages languages1
ON languages1.id = votes.lang1_id
INNER JOIN languages languages2
ON languages2.id = votes.lang2_id;
