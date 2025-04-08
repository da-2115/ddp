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
