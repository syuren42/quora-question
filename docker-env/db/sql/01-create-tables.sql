CREATE USER operator IDENTIFIED BY 'testtest' ;
CREATE USER viewer IDENTIFIED BY 'testtest';
CREATE USER app IDENTIFIED BY 'testtest';

CREATE TABLE questions (
  id int(11) NOT NULL ,
  content varchar(1024) default '',
  PRIMARY KEY (ID)
);

