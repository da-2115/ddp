CREATE DATABASE ARCHERYDB;
USE ARCHERYDB;

CREATE TABLE Member
(
    ArcheryAustraliaID NVARCHAR( 255 ) NOT NULL,
    FirstName NVARCHAR( 255 ),
    DateOfBirth DATE,
    Gender BOOL,
    DefaultBowType NVARCHAR( 255 ),
    PRIMARY KEY ( ArcheryAustraliaID )
);

CREATE TABLE Round (
    RoundID INT PRIMARY KEY,
    EventID INT NOT NULL,
    Division VARCHAR(255) NOT NULL,
    Class VARCHAR(255) NOT NULL,
    Gender BOOL,
    FOREIGN KEY (EventID) REFERENCES Event(EventID),
    FOREIGN KEY (Division) REFERENCES Division(DivisionID),
    FOREIGN KEY (Class) REFERENCES Class(ClassID)
);

CREATE TABLE Range (
    RangeID INT PRIMARY KEY,
    RoundID INT NOT NULL,
    Staged BOOL NOT NULL,
    Distance INT NOT NULL,
    TargetSize INT NOT NULL,
    FOREIGN KEY (RoundID) REFERENCES Round(RoundID)
);
CREATE TABLE IF NOT EXISTS PracticeEvent
(
    PracticeID INT NOT NULL,
    EventID INT NOT NULL,
    ArcheryAustraliaID NVARCHAR( 255 ) NOT NULL,
    PRIMARY KEY (PracticeID),
    FOREIGN KEY (EventID) REFERENCES Event(EventID),
    FOREIGN KEY (ArcheryAustraliaID) REFERENCES Member(ArcheryAustraliaID)
);

CREATE TABLE IF NOT EXISTS End
(
    EndID INT NOT NULL,
    RangeID INT NOT NULL,
    ArcheryAustraliaID NVARCHAR( 255 ) NOT NULL,
    FinalScore INT,
    PRIMARY KEY (EndID),
    FOREIGN KEY (RangeID) REFERENCES Range(RangeID),
    FOREIGN KEY (ArcheryAustraliaID) REFERENCES Member(ArcheryAustraliaID)
);

CREATE TABLE IF NOT EXISTS Score
(
    ScoreID INT NOT NULL,
    EndID INT NOT NULL,
    ArrowNumber INT,
    Score INT,
    PRIMARY KEY (ScoreID),
    FOREIGN KEY (EndID) REFERENCES End(EndID)
);
