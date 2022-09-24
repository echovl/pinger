BEGIN;

CREATE TABLE
  IF NOT EXISTS hosts (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    url VARCHAR(1000) NOT NULL,
    mean INTEGER NOT NULL,
    last INTEGER NOT NULL,
    best INTEGER NOT NULL,
    worst INTEGER NOT NULL,
    PRIMARY KEY (id)
  );

COMMIT;
