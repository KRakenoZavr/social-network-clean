CREATE TABLE user_follow(
  id    INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id1        VARCHAR(30)   NOT NULL,
  user_id2        VARCHAR(30)   NOT NULL,
  created_at      DATETIME      NOT NULL,
  invite          INTEGER       NOT NULL,
  FOREIGN KEY(user_id1)  REFERENCES  users(user_id),
  FOREIGN KEY(user_id2)  REFERENCES  users(user_id)
);
