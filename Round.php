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
