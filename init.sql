CREATE DATABASE IF NOT EXISTS ARCHERYDB;
USE ARCHERYDB;

CREATE TABLE IF NOT EXISTS Class (
    ClassID NVARCHAR(255) PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS Division (
    DivisionID NVARCHAR(255) PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS Member
(
    ArcheryAustraliaID NVARCHAR(255) PRIMARY KEY NOT NULL,
    FirstName NVARCHAR(255),
    DateOfBirth DATE,
    Gender BOOL,
    DefaultDivision NVARCHAR(255),
    FOREIGN KEY (DefaultDivision) REFERENCES Division(DivisionID)
);

CREATE TABLE IF NOT EXISTS Event (
    EventID INT AUTO_INCREMENT PRIMARY KEY,
    Name NVARCHAR(255) NOT NULL,
    Date DATE NOT NULL
);


CREATE TABLE Championship (
    ChampionshipID INT PRIMARY KEY NOT NULL ,  
    EventID INT NOT NULL,
    Name NVARCHAR(255) NOT NULL,
    FOREIGN KEY (EventID) REFERENCES Event(EventID)
 );
 
CREATE TABLE Round (
    RoundID INT AUTO_INCREMENT PRIMARY KEY,
    EventID INT NOT NULL,
    Division NVARCHAR(255) NOT NULL,
    Class NVARCHAR(255) NOT NULL,
    Gender BOOL,
    FOREIGN KEY (EventID) REFERENCES Event(EventID),
    FOREIGN KEY (Division) REFERENCES Division(DivisionID),
    FOREIGN KEY (Class) REFERENCES Class(ClassID)
);

CREATE TABLE IF NOT EXISTS `Range` 
(
    RangeID INT AUTO_INCREMENT PRIMARY KEY,
    RoundID INT NOT NULL,
    Staged BOOL NOT NULL,
    Distance INT NOT NULL,
    TargetSize INT NOT NULL,
    FOREIGN KEY (RoundID) REFERENCES `Round`(RoundID)
);

CREATE TABLE PracticeEvent
(
    PracticeID INT AUTO_INCREMENT PRIMARY KEY,
    EventID INT NOT NULL,
    ArcheryAustraliaID NVARCHAR(255) NOT NULL,
    FOREIGN KEY (ArcheryAustraliaID) REFERENCES Member(ArcheryAustraliaID),
    FOREIGN KEY (EventID) REFERENCES Event(EventID)
);

CREATE TABLE IF NOT EXISTS End
(
    EndID INT PRIMARY KEY NOT NULL,
    RangeID INT NOT NULL,
    ArcheryAustraliaID NVARCHAR( 255 ) NOT NULL,
    FinalScore INT NOT NULL,
    FOREIGN KEY (RangeID) REFERENCES `Range`(RangeID),
    FOREIGN KEY (ArcheryAustraliaID) REFERENCES Member(ArcheryAustraliaID)
);

CREATE TABLE IF NOT EXISTS Score
(
    ScoreID INT PRIMARY KEY NOT NULL,
    EndID INT NOT NULL,
    ArrowNumber INT NOT NULL,
    Score INT NOT NULL,
    FOREIGN KEY (EndID) REFERENCES End(EndID)
);
