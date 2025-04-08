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
    TargetNumber INT,
    Score INT,
    X INT
);

