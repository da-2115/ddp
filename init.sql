CREATE DATABASE ARCHERYDB;
USE DATABASE ARCHERYDB;

CREATE TABLE Member
(
    ArcheryAustraliaID NVARCHAR(255),
    FirstName NVARCHAR(255),
    DateOfBirth DATE,
    Gender BOOL,
    DefaultBowType NVARCHAR(255),
    PRIMARY KEY (ArcheryAustraliaID)
);