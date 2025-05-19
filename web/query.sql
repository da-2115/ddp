-- name: GetMemberByID :one
SELECT * FROM Member
WHERE ArcheryAustraliaID = ?
LIMIT 1;

-- name: CreateMember :execresult
INSERT INTO Member (
    ArcheryAustraliaID, PasswordHash, FirstName, DateOfBirth, Gender, ClubRecorder
) VALUES (
    ?, ?, ?, ?, ?, ?
);

-- name: DeleteMember :exec
DELETE FROM Member
WHERE ArcheryAustraliaID = ?;

-- name: GetStagedEnds :many
SELECT *
FROM End en
JOIN `Range` ra ON ra.RangeID = en.RangeID
JOIN `Round` r ON r.RoundID = ra.RoundID
JOIN Event e ON e.EventID = r.EventID
WHERE en.Staged = TRUE;

-- name: StageEnd :exec
UPDATE End
SET Staged = FALSE
WHERE EndID = ?;

-- name: DeleteEnd :exec
DELETE FROM End
WHERE EndID = ?;

-- name: GetAllEvents :many
SELECT *
FROM Event
LIMIT ?
OFFSET ?;

-- name: GetRoundByEvent :many
SELECT *
FROM `Round`
WHERE EventID = ?
LIMIT ?
OFFSET ?;

-- name: GetRangeByRound :many
SELECT *
FROM `Range`
WHERE RoundID = ?
LIMIT ?
OFFSET ?;

-- name: GetEndByRound :many
SELECT *
FROM End
WHERE RangeID = ?
LIMIT ?
OFFSET ?;

-- name: GetScoreByEnd :many
SELECT *
FROM Score
WHERE EndID = ?;

-- name: GetRounds :many
SELECT  *
FROM Round r
JOIN Member m ON (
        (YEAR(CURDATE()) - YEAR(m.DateOfBirth) <= 14 AND r.Class IN ('Under14', 'Under16', 'Under18', 'Under21', 'Open')) OR
        (YEAR(CURDATE()) - YEAR(m.DateOfBirth) <= 16 AND r.Class IN ('Under16', 'Under18', 'Under21', 'Open')) OR
        (YEAR(CURDATE()) - YEAR(m.DateOfBirth) <= 18 AND r.Class IN ('Under18', 'Under21', 'Open')) OR
        (YEAR(CURDATE()) - YEAR(m.DateOfBirth) <= 21 AND r.Class IN ('Under21', 'Open')) OR
        (YEAR(CURDATE()) - YEAR(m.DateOfBirth) > 21 AND r.Class IN ('Open')) OR
        (YEAR(CURDATE()) - YEAR(m.DateOfBirth) >= 50 AND r.Class IN ('Open', '50Plus')) OR
        (YEAR(CURDATE()) - YEAR(m.DateOfBirth) >= 60 AND r.Class IN ('Open', '50Plus', '60Plus')) OR
        (YEAR(CURDATE()) - YEAR(m.DateOfBirth) >= 70 AND r.Class IN ('Open', '50Plus', '60Plus', '70Plus'))
    )
WHERE m.ArcheryAustraliaID = ? AND r.EventID = ? AND m.Gender = r.Gender
LIMIT ?
OFFSET ?;


-- name: GetEventsByID :many
SELECT *
FROM Event e
JOIN `Round` r ON e.EventID = r.EventID
JOIN `Range` ra ON r.RoundID = ra.RoundID
JOIN End en ON ra.RangeID = en.RangeID
WHERE en.ArcheryAustraliaID = ?
GROUP BY e.EventID
LIMIT ?
OFFSET ?;

-- name: GetRoundsByID :many
SELECT *
FROM `Round` r
JOIN Event e ON e.EventID = r.EventID
JOIN `Range` ra ON r.RoundID = ra.RoundID
JOIN End en ON ra.RangeID = en.RangeID
WHERE en.ArcheryAustraliaID = ? AND e.EventID = ?
GROUP BY r.RoundID
LIMIT ?
OFFSET ?;

-- name: GetRangesByID :many
SELECT *
FROM `Range` ra
JOIN `Round` r ON r.RoundID = ra.RoundID
JOIN Event e ON e.EventID = r.EventID
JOIN End en ON ra.RangeID = en.RangeID
WHERE en.ArcheryAustraliaID = ? AND e.EventID = ? AND r.RoundID = ?
GROUP BY ra.RangeID
LIMIT ?
OFFSET ?;

-- name: GetEndsByID :many
SELECT *
FROM End en
JOIN `Range` ra ON ra.RangeID = en.RangeID
JOIN `Round` r ON r.RoundID = ra.RoundID
JOIN Event e ON e.EventID = r.EventID
WHERE en.ArcheryAustraliaID = ? AND e.EventID = ? AND r.RoundID = ? AND en.RangeID = ?
GROUP BY en.EndID
LIMIT ?
OFFSET ?;

-- name: GetScoresByID :many
SELECT *
FROM Score s
JOIN End en ON s.EndID = en.EndID
JOIN `Range` ra ON ra.RangeID = en.RangeID
JOIN `Round` r ON r.RoundID = ra.RoundID
JOIN Event e ON e.EventID = r.EventID
WHERE en.ArcheryAustraliaID = ? AND e.EventID = ? AND r.RoundID = ? AND en.RangeID = ? AND en.EndID = ?
GROUP BY s.ScoreID
LIMIT ?
OFFSET ?;

-- name: GetPracticeEventsByID :many
SELECT *
FROM Event e
JOIN PracticeEvent pe ON e.EventID = pe.EventID
WHERE pe.ArcheryAustraliaID = ?
GROUP BY e.EventID
LIMIT ?
OFFSET ?;

-- name: CreateEvent :execresult
INSERT INTO Event (
    Name, Date
)
VALUES (
    ?, ?
);

-- name: CreateRound :execresult
INSERT INTO `Round` (
    EventID, Division, Class, Gender
)
VALUES (
    ?, ?, ?, ?
);

-- name: CreateRange :execresult
INSERT INTO `Range` (
    RoundID, Distance, TargetSize
)
VALUES (
    ?, ?, ?
);


-- name: CreateEnd :execresult
INSERT INTO End (
    RangeID, ArcheryAustraliaID, FinalScore, Staged
)
VALUES (
    ?, ?, ?, TRUE
);

-- name: CreateScore :execresult
INSERT INTO Score (
    EndID, ArrowNumber, Score
)
VALUES (
    ?, ?, ?
)
