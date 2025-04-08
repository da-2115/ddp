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