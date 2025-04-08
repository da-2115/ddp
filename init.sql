CREATE DATABASE ARCHERYDB;
USE DATABASE ARCHERYDB;

CREATE TABLE Member
(
    ArcheryAustraliaID NVARCHAR( 255 ) NOT NULL,
    FirstName NVARCHAR( 255 ) NOT NULL,
    DateOfBirth DATE NOT NULL,
    Gender BOOL NOT NULL,
    DefaultBowType NVARCHAR( 255 ) NOT NULL,
    PRIMARY KEY ( ArcheryAustraliaID )
);

CREATE TABLE PracticeEvent
(
    PracticeID INT NOT NULL,
    EventID INT NOT NULL,
    ArcheryAustraliaID NVARCHAR( 255 ) NOT NULL,
    PRIMARY KEY (PracticeID),
    FOREIGN KEY (EventID) REFERENCES Event(EventID),
    FOREIGN KEY (ArcheryAustraliaID) REFERENCES Member(ArcheryAustraliaID)
);

CREATE TABLE End
(
    EndID INT NOT NULL,
    RangeID INT NOT NULL,
    ArcheryAustraliaID NVARCHAR( 255 ) NOT NULL,
    FinalScore INT,
    PRIMARY KEY (EndID),
    FOREIGN KEY (RangeID) REFERENCES Range(RangeID),
    FOREIGN KEY (ArcheryAustraliaID) REFERENCES Member(ArcheryAustraliaID)
);

CREATE TABLE Score
(
    ScoreID INT NOT NULL,
    EndID INT NOT NULL,
    ArrowNumber INT,
    Score INT,
    PRIMARY KEY (ScoreID),
    FOREIGN KEY (EndID) REFERENCES End(EndID)
);

