-- hashed password is bcrypt 'password'
USE ARCHERYDB;
INSERT INTO Member (ArcheryAustraliaID, PasswordHash, FirstName, DateOfBirth, Gender, ClubRecorder, DefaultDivision) VALUES
('AA12345', '$2a$10$2EHV1xRkBeXr1mo4rDeptOpTTSmDXNmTPWXpwrVD5G44p8QHA0ojC', 'Alice', '1990-05-15', TRUE, FALSE, 'Compound'),
('AA12346', '$2a$10$2EHV1xRkBeXr1mo4rDeptOpTTSmDXNmTPWXpwrVD5G44p8QHA0ojC', 'Bob', '1985-08-22', FALSE, TRUE, 'Recurve'),
('AA12347', '$2a$10$2EHV1xRkBeXr1mo4rDeptOpTTSmDXNmTPWXpwrVD5G44p8QHA0ojC', 'Charlie', '1992-12-30', TRUE, TRUE, 'Compound'),
('AA12348', '$2a$10$2EHV1xRkBeXr1mo4rDeptOpTTSmDXNmTPWXpwrVD5G44p8QHA0ojC', 'Diana', '1988-03-10', FALSE, FALSE, 'Recurve'),
('AA12349', '$2a$10$2EHV1xRkBeXr1mo4rDeptOpTTSmDXNmTPWXpwrVD5G44p8QHA0ojC', 'Ethan', '1995-07-25', TRUE, FALSE, 'Compound'),
('AA12350', '$2a$10$2EHV1xRkBeXr1mo4rDeptOpTTSmDXNmTPWXpwrVD5G44p8QHA0ojC', 'Fiona', '1980-11-05', FALSE, TRUE,'Recurve');

INSERT INTO Event (Name, Date) VALUES
("Practice Event 1", NOW()),
("Practice Event 2", NOW()),
("Practice Event 3", NOW()),
("Actual Event 1", NOW()),
("Actual Event 2", NOW());

INSERT INTO PracticeEvent (EventID, ArcheryAustraliaID) VALUES
(1,'AA12345'),
(2,'AA12345'),
(3,'AA12345');
