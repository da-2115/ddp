-- setup-security.sql
-- Create limited database users for web access

-- Create a limited-access user for the web application
CREATE USER 'webapp_user'@'%' IDENTIFIED BY 'WebAppDB2025!';

-- Grant only necessary permissions to specific tables
GRANT SELECT, INSERT, UPDATE ON ARCHERYDB.Member TO 'webapp_user'@'%';
GRANT SELECT ON ARCHERYDB.Event TO 'webapp_user'@'%';
GRANT SELECT, INSERT ON ARCHERYDB.End TO 'webapp_user'@'%';
GRANT SELECT, INSERT ON ARCHERYDB.Score TO 'webapp_user'@'%';
