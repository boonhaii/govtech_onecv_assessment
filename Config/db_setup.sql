-- Already added locally, but present just for recording keeping purposes.

-- Creates the Database
-- Replace with api_db if need to create test database
CREATE DATABASE apidb;

-- Changes to the desired database
USE apidb;

-- Creates the Table Teachers
CREATE TABLE Teachers( 
  email VARCHAR(320), 
  PRIMARY KEY(email)
);

-- Creates the Table Students
CREATE TABLE Students(
  email VARCHAR(320), 
  suspended TINYINT, 
  PRIMARY KEY(email)
);

-- Creates the Table Registers
CREATE TABLE Registers( 
  t_email VARCHAR(320), 
  s_email VARCHAR(320), 
  FOREIGN KEY(t_email) REFERENCES Teachers(email), 
  FOREIGN KEY(s_email) REFERENCES Students(email), 
  PRIMARY KEY(t_email, s_email)
);