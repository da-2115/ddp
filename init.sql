CREATE DATABASE ARCHERYDB;
USE DATABASE ARCHERYDB;

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
    EventID INT,
    Division INT,
    Class INT,
    Gender BOOL,
    FOREIGN KEY (EventID) REFERENCES Event(EventID),
    FOREIGN KEY (Division) REFERENCES Division(DivisionID),
    FOREIGN KEY (Class) REFERENCES Class(ClassID)
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

