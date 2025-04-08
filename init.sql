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

CREATE TABLE Event (
    EventID INT NOT NULL,
    Name VARCHAR(255) NOT NULL,
    Date DATE NOT NULL,
    PRIMARY KEY (EventID)  
);

CREATE TABLE Championship (
    ChampionshipID INT NOT NULL,  
    EventID INT NOT NULL,
    Name VARCHAR(255) NOT NULL,
    PRIMARY KEY (ChampionshipID),  
    CONSTRAINT FK_Championship_Event 
    FOREIGN KEY (EventID) 
        REFERENCES Event(EventID)
);