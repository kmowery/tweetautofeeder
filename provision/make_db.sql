create table USERS (
  id INTEGER PRIMARY KEY ASC,
  request_token TEXT,
  request_token_secret TEXT,
  oauth_token TEXT,
  oauth_token_secret TEXT,
  user_id TEXT,
  screen_name TEXT,
  session_cookie TEXT
);

create table TWEETS (
  id integer PRIMARY KEY ASC,
  user_id TEXT,
  tweet TEXT,
  completed BOOL,
  FOREIGN KEY(user_id) REFERENCES users(user_id)
);
