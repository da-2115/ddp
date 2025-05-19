-- Create limited-access user for web application
CREATE USER 'webapp_user'@'%' IDENTIFIED BY 'WebAppDB2025!';

-- Select, Insert
GRANT SELECT, INSERT ON ARCHERYDB.Event TO 'webapp_user'@'%';
GRANT SELECT, INSERT ON ARCHERYDB.Round TO 'webapp_user'@'%';
GRANT SELECT, INSERT ON ARCHERYDB.Range TO 'webapp_user'@'%';
GRANT SELECT, INSERT ON ARCHERYDB.PracticeEvent TO 'webapp_user'@'%';
GRANT SELECT, INSERT ON ARCHERYDB.Championship TO 'webapp_user'@'%';

-- Select, Insert, Update, Delete
GRANT SELECT, INSERT, UPDATE, DELETE ON ARCHERYDB.End TO 'webapp_user'@'%';
GRANT SELECT, INSERT, UPDATE, DELETE ON ARCHERYDB.Score TO 'webapp_user'@'%';

-- Select only
GRANT SELECT ON ARCHERYDB.Member TO 'webapp_user'@'%';
