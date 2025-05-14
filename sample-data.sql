-- bcrypt hashed password is 'password'
USE ARCHERYDB;
INSERT INTO Member (ArcheryAustraliaID, PasswordHash, FirstName, DateOfBirth, Gender, ClubRecorder) VALUES
('AA12345', '$2a$10$2EHV1xRkBeXr1mo4rDeptOpTTSmDXNmTPWXpwrVD5G44p8QHA0ojC', 'Alice', '1990-05-15', "Male", TRUE),
('AA12346', '$2a$10$2EHV1xRkBeXr1mo4rDeptOpTTSmDXNmTPWXpwrVD5G44p8QHA0ojC', 'Bob', '1985-08-22', "Female", FALSE),
('AA12347', '$2a$10$2EHV1xRkBeXr1mo4rDeptOpTTSmDXNmTPWXpwrVD5G44p8QHA0ojC', 'Charlie', '1992-12-30', "Male", TRUE),
('AA12348', '$2a$10$2EHV1xRkBeXr1mo4rDeptOpTTSmDXNmTPWXpwrVD5G44p8QHA0ojC', 'Diana', '1988-03-10', "Female", FALSE),
('AA12349', '$2a$10$2EHV1xRkBeXr1mo4rDeptOpTTSmDXNmTPWXpwrVD5G44p8QHA0ojC', 'Ethan', '1995-07-25', "Male", FALSE),
('AA12350', '$2a$10$2EHV1xRkBeXr1mo4rDeptOpTTSmDXNmTPWXpwrVD5G44p8QHA0ojC', 'Fiona', '1980-11-05', "Female", TRUE);

INSERT INTO Event (Name, Date) VALUES
("Practice Event 1", NOW()),
("Practice Event 2", NOW()),
("Practice Event 3", NOW()),
("Actual Event 1", NOW());

INSERT INTO Event (Name, Date) VALUES -- needs to be seperate or LAST_INSERT_ID would be 1, as it is based on the first input on a multi-insert
("Actual Event 2", NOW());

SET @event_id = LAST_INSERT_ID(); -- this grabs the last auto_inc,

INSERT INTO `Round` (EventID, Division, Class, Gender) VALUES
(@event_id, "Compound", "Under14", "Male");

SET @round_id = LAST_INSERT_ID();

INSERT INTO `Range` (RoundID, Distance, TargetSize) VALUES
(@round_id, 30, 50);

SET @range_id = LAST_INSERT_ID();

INSERT INTO End (RangeID, ArcheryAustraliaID, Staged, FinalScore) VALUES
(@range_id, 'AA12345', FALSE, 20);

SET @end_id = LAST_INSERT_ID();

INSERT INTO Score (EndID, ArrowNumber, Score) VALUES
(@end_id, 1, '3'),
(@end_id, 2, '7'),
(@end_id, 3, 'X');

INSERT INTO PracticeEvent (EventID, ArcheryAustraliaID) VALUES
(1,'AA12345'),
(2,'AA12345'),
(3,'AA12345');

INSERT INTO Event (Name, Date) VALUES
("Actual Event 3", NOW());

SET @event_id = LAST_INSERT_ID(); -- this grabs the last auto_inc,

INSERT INTO `Round` (EventID, Division, Class, Gender) VALUES
(@event_id, "Recurve", "Under16", "Male");

SET @round_id = LAST_INSERT_ID();

INSERT INTO `Range` (RoundID, Distance, TargetSize) VALUES
(@round_id, 30, 50);

SET @range_id = LAST_INSERT_ID();

INSERT INTO End (RangeID, ArcheryAustraliaID, Staged, FinalScore) VALUES
(@range_id, 'AA12345', FALSE, 20);

SET @end_id = LAST_INSERT_ID();

INSERT INTO Score (EndID, ArrowNumber, Score) VALUES
(@end_id, 1, '3'),
(@end_id, 2, '7'),
(@end_id, 3, 'X');

INSERT INTO PracticeEvent (EventID, ArcheryAustraliaID) VALUES
(1,'AA12345'),
(2,'AA12345'),
(3,'AA12345');
