CREATE TABLE groups(
  group_id        VARCHAR(30) PRIMARY KEY,
  user_id         VARCHAR(30)   NOT NULL,
  created_at      DATETIME      NOT NULL,
  title           VARCHAR(50)   NOT NULL,
  body            VARCHAR(255)  NOT NULL,
  FOREIGN KEY(user_id) REFERENCES users(user_id)
);

CREATE UNIQUE INDEX group_index ON groups(group_id,user_id);
