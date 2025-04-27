-- name: GetMemberByID :one
SELECT * FROM Member
WHERE ArcheryAustraliaID = ?
LIMIT 1;

-- name: CreateMember :execresult
INSERT INTO Member (
    ArcheryAustraliaID, PasswordHash, FirstName, DateOfBirth, Gender, ClubRecorder, DefaultDivision
) VALUES (
    ?, ?, ?, ?, ?, ?, ?
);

-- name: DeleteMember :exec
DELETE FROM Member
WHERE ArcheryAustraliaID = ?;
