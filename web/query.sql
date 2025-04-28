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

-- name: GetEvents :many
SELECT *
FROM Event e
JOIN `Round` r ON e.EventID = r.EventID
JOIN `Range` ra ON r.RoundID = ra.RoundID
JOIN End en ON ra.RangeID = en.RangeID
WHERE en.ArcheryAustraliaID = ?
LIMIT ?
OFFSET ?;

-- name: GetRounds :many
SELECT *
FROM `Round` r
JOIN Event e ON e.EventID = r.EventID
JOIN `Range` ra ON r.RoundID = ra.RoundID
JOIN End en ON ra.RangeID = en.RangeID
WHERE en.ArcheryAustraliaID = ? AND e.EventID = ?
LIMIT ?
OFFSET ?;

-- name: GetRanges :many
SELECT *
FROM `Range` ra
JOIN `Round` r ON r.RoundID = ra.RoundID
JOIN Event e ON e.EventID = r.EventID
JOIN End en ON ra.RangeID = en.RangeID
WHERE en.ArcheryAustraliaID = ? AND e.EventID = ? AND r.RoundID = ?
LIMIT ?
OFFSET ?;

-- name: GetEnds :many
SELECT *
FROM End en
JOIN `Range` ra ON ra.RangeID = en.RangeID
JOIN `Round` r ON r.RoundID = ra.RoundID
JOIN Event e ON e.EventID = r.EventID
WHERE en.ArcheryAustraliaID = ? AND e.EventID = ? AND r.RoundID = ? AND en.RangeID = ?
LIMIT ?
OFFSET ?;

-- name: GetScores :many
SELECT *
FROM Score s
JOIN End en ON s.EndID = en.EndID
JOIN `Range` ra ON ra.RangeID = en.RangeID
JOIN `Round` r ON r.RoundID = ra.RoundID
JOIN Event e ON e.EventID = r.EventID
WHERE en.ArcheryAustraliaID = ? AND e.EventID = ? AND r.RoundID = ? AND en.RangeID = ? AND en.EndID = ?
LIMIT ?
OFFSET ?;

-- name: GetPracticeEvents :many
SELECT e.*
FROM Event e
JOIN PracticeEvent pe ON e.EventID = pe.EventID
WHERE pe.ArcheryAustraliaID = ?
LIMIT ?
OFFSET ?;
