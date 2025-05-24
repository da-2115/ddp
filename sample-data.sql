-- bcrypt hashed password is 'password'
USE ARCHERYDB;
INSERT INTO Member (ArcheryAustraliaID, PasswordHash, FirstName, DateOfBirth, Gender, ClubRecorder) VALUES
('AA12345', '$2a$10$2EHV1xRkBeXr1mo4rDeptOpTTSmDXNmTPWXpwrVD5G44p8QHA0ojC', 'Bob', '2015-05-15', "Male", TRUE), -- Privileged User
('AA12346', '$2a$10$2EHV1xRkBeXr1mo4rDeptOpTTSmDXNmTPWXpwrVD5G44p8QHA0ojC', 'Alice', '1985-08-22', "Female", FALSE); -- Normal User

INSERT INTO Event (EventName, Date) VALUES -- needs to be seperate or LAST_INSERT_ID would be 1, as it is based on the first input on a multi-insert
("Summer Event 1", NOW());

SET @event_id = LAST_INSERT_ID(); -- this grabs the last auto_inc,

INSERT INTO `Round` (EventID, Division, Class, Gender) VALUES
(@event_id, "Compound", "Open", "Male");

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

INSERT INTO End (RangeID, ArcheryAustraliaID, Staged, FinalScore) VALUES
(@range_id, 'AA12346', Staged, 31);

SET @end_id = LAST_INSERT_ID();

INSERT INTO Score (EndID, ArrowNumber, Score) VALUES
(@end_id, 1, '3'),
(@end_id, 2, '2'),
(@end_id, 3, '5'),
(@end_id, 4, '7'),
(@end_id, 5, '4'),
(@end_id, 6, 'X');

INSERT INTO `Round` (EventID, Division, Class, Gender) VALUES
(@event_id, "Compound", "Under21", "Male");

SET @range_id = LAST_INSERT_ID();

INSERT INTO `Range` (RoundID, Distance, TargetSize) VALUES
(@round_id, 50, 30);

SET @range_id = LAST_INSERT_ID();

INSERT INTO End (RangeID, ArcheryAustraliaID, Staged, FinalScore) VALUES
(@range_id, 'AA12345', TRUE, 20);

SET @end_id = LAST_INSERT_ID();

INSERT INTO Score (EndID, ArrowNumber, Score) VALUES
(@end_id, 1, '3'),
(@end_id, 2, '7'),
(@end_id, 3, 'X');

INSERT INTO Event (EventName, Date) VALUES
("Summer Event 2", NOW());

SET @event_id = LAST_INSERT_ID(); -- this grabs the last auto_inc,

-- Championship
INSERT INTO Championship (EventID, ChampionshipName) VALUES
(@event_id, "Major Championship");

INSERT INTO `Round` (EventID, Division, Class, Gender) VALUES
(@event_id, "Recurve", "50Plus", "Female");

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

-- Create a new event
INSERT INTO Event (EventName, Date) VALUES
("Summer Event 3", NOW());

SET @event_id = LAST_INSERT_ID(); -- Get the last inserted event ID

-- Create a new round for the event
INSERT INTO `Round` (EventID, Division, Class, Gender) VALUES
(@event_id, "Compound", "Open", "Female");

SET @round_id = LAST_INSERT_ID(); -- Get the last inserted round ID

-- Create a new range for the round
INSERT INTO `Range` (RoundID, Distance, TargetSize) VALUES
(@round_id, 30, 50);

SET @range_id = LAST_INSERT_ID(); -- Get the last inserted range ID

-- Create ends and scores for AA12345
INSERT INTO End (RangeID, ArcheryAustraliaID, Staged, FinalScore) VALUES
(@range_id, 'AA12345', FALSE, 20);

SET @end_id = LAST_INSERT_ID(); -- Get the last inserted end ID

INSERT INTO Score (EndID, ArrowNumber, Score) VALUES
(@end_id, 1, '4'),
(@end_id, 2, '6'),
(@end_id, 3, 'X');

INSERT INTO End (RangeID, ArcheryAustraliaID, Staged, FinalScore) VALUES
(@range_id, 'AA12345', TRUE, 22);

SET @end_id = LAST_INSERT_ID(); -- Get the last inserted end ID

INSERT INTO Score (EndID, ArrowNumber, Score) VALUES
(@end_id, 1, '5'),
(@end_id, 2, '7'),
(@end_id, 3, 'X');

-- Create a new round for the event
INSERT INTO `Round` (EventID, Division, Class, Gender) VALUES
(@event_id, "Recurve", "Open", "Male");

SET @round_id = LAST_INSERT_ID(); -- Get the last inserted round ID

-- Create a new range for the round
INSERT INTO `Range` (RoundID, Distance, TargetSize) VALUES
(@round_id, 30, 50);

SET @range_id = LAST_INSERT_ID(); -- Get the last inserted range ID

-- Create ends and scores for AA12346
INSERT INTO End (RangeID, ArcheryAustraliaID, Staged, FinalScore) VALUES
(@range_id, 'AA12346', FALSE, 21);

SET @end_id = LAST_INSERT_ID(); -- Get the last inserted end ID

INSERT INTO Score (EndID, ArrowNumber, Score) VALUES
(@end_id, 1, '3'),
(@end_id, 2, '8'),
(@end_id, 3, 'X');

INSERT INTO End (RangeID, ArcheryAustraliaID, Staged, FinalScore) VALUES
(@range_id, 'AA12346', TRUE, 28);

SET @end_id = LAST_INSERT_ID(); -- Get the last inserted end ID

INSERT INTO Score (EndID, ArrowNumber, Score) VALUES
(@end_id, 1, '6'),
(@end_id, 2, '7'),
(@end_id, 3, '5'),
(@end_id, 4, 'X');

INSERT INTO Event (EventName, Date) VALUES
("Practice Event", NOW());

SET @event_id = LAST_INSERT_ID();

INSERT INTO `Round` (EventID, Division, Class, Gender) VALUES
(@event_id, "Recurve", "Open", "Male");

SET @round_id = LAST_INSERT_ID(); -- Get the last inserted round ID

SET @round_id = LAST_INSERT_ID();

INSERT INTO `Range` (RoundID, Distance, TargetSize) VALUES
(@round_id, 30, 50);

SET @range_id = LAST_INSERT_ID();

INSERT INTO End (RangeID, ArcheryAustraliaID, Staged, FinalScore) VALUES
(@range_id, 'AA12345', FALSE, 20);

SET @end_id = LAST_INSERT_ID();

INSERT INTO Score (EndID, ArrowNumber, Score) VALUES
(@end_id, 1, '4'),
(@end_id, 2, '6'),
(@end_id, 3, 'X');

INSERT INTO PracticeEvent (EventID, ArcheryAustraliaID) VALUES
(@event_id, 'AA12345');
