CREATE TABLE group_user(
  id    INTEGER PRIMARY KEY AUTOINCREMENT,
  group_id        VARCHAR(30)   NOT NULL,
  user_id         VARCHAR(30)   NOT NULL,
  created_at      DATETIME      NOT NULL,
  invite          INTEGER       NOT NULL,
  FOREIGN KEY(group_id) REFERENCES groups(group_id),
  FOREIGN KEY(user_id)  REFERENCES users(user_id)
);

CREATE UNIQUE INDEX group_user_index ON group_user(group_id,user_id);
