CREATE TABLE user_auth (
  id      INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
  user_id       VARCHAR(30)     NOT NULL,
  expires       DATETIME        NOT NULL,
  session       VARCHAR(50)     NOT NULL,
  FOREIGN KEY(user_id) REFERENCES users(user_id)
);