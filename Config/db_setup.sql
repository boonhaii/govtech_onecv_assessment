-- Already added locally, but present just for recording keeping purposes.

CREATE TABLE Teachers( 
  email VARCHAR(320), 
  PRIMARY KEY(email)
);

CREATE TABLE Students(
  email VARCHAR(320), 
  suspended TINYINT, 
  PRIMARY KEY(email)
);

CREATE TABLE Registers( 
  t_email VARCHAR(320), 
  s_email VARCHAR(320), 
  FOREIGN KEY(t_email) REFERENCES Teachers(email), 
  FOREIGN KEY(s_email) REFERENCES Students(email), 
  PRIMARY KEY(t_email, s_email)
);