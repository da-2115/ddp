-- drop database ARCHERYDB;
CREATE DATABASE IF NOT EXISTS ARCHERYDB;
USE ARCHERYDB;

CREATE TABLE IF NOT EXISTS Member
(
    ArcheryAustraliaID NVARCHAR(255) PRIMARY KEY NOT NULL,
    PasswordHash VARCHAR(255) NOT NULL,
    FirstName NVARCHAR(255) NOT NULL,
    DateOfBirth DATE NOT NULL,
    Gender ENUM('Male', 'Female') NOT NULL,
    ClubRecorder BOOL NOT NULL
);

CREATE TABLE IF NOT EXISTS Event (
    EventID INT AUTO_INCREMENT PRIMARY KEY,
    Name NVARCHAR(255) NOT NULL,
    Date DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS Championship (
    ChampionshipID INT PRIMARY KEY NOT NULL ,  
    EventID INT NOT NULL,
    Name NVARCHAR(255) NOT NULL,
    FOREIGN KEY (EventID) REFERENCES Event(EventID)
 );
 
CREATE TABLE IF NOT EXISTS Round (
    RoundID INT AUTO_INCREMENT PRIMARY KEY,
    EventID INT NOT NULL,
    Class ENUM('Under14', 'Under16', 'Under18', 'Under21', 'Open', '50Plus', '60Plus', '70Plus') NOT NULL,
    Division ENUM ('Recurve', 'Compound', 'RecurveBarebow', 'CompoundBarebow', 'Longbow') NOT NULL,
    Gender ENUM('Male', 'Female') NOT NULL,
    FOREIGN KEY (EventID) REFERENCES Event(EventID)
);

CREATE TABLE IF NOT EXISTS `Range` 
(
    RangeID INT AUTO_INCREMENT PRIMARY KEY,
    RoundID INT NOT NULL,
    Distance INT NOT NULL,
    TargetSize INT NOT NULL,
    FOREIGN KEY (RoundID) REFERENCES `Round`(RoundID)
);

CREATE TABLE IF NOT EXISTS PracticeEvent
(
    PracticeID INT AUTO_INCREMENT PRIMARY KEY,
    EventID INT NOT NULL,
    ArcheryAustraliaID NVARCHAR(255) NOT NULL,
    FOREIGN KEY (ArcheryAustraliaID) REFERENCES Member(ArcheryAustraliaID),
    FOREIGN KEY (EventID) REFERENCES Event(EventID)
);

CREATE TABLE IF NOT EXISTS End
(
    EndID INT AUTO_INCREMENT PRIMARY KEY,
    RangeID INT NOT NULL,
    ArcheryAustraliaID NVARCHAR( 255 ) NOT NULL,
    FinalScore INT NOT NULL,
    Staged BOOL NOT NULL,
    FOREIGN KEY (RangeID) REFERENCES `Range`(RangeID),
    FOREIGN KEY (ArcheryAustraliaID) REFERENCES Member(ArcheryAustraliaID)
);

CREATE TABLE IF NOT EXISTS Score
(
    ScoreID INT AUTO_INCREMENT PRIMARY KEY,
    EndID INT NOT NULL,
    ArrowNumber INT NOT NULL,
    Score NVARCHAR(255) NOT NULL,
    FOREIGN KEY (EndID) REFERENCES End(EndID)
);
