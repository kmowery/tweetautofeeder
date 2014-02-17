package main

import (
  "github.com/mrjones/oauth"
)

func addRequestTokens(s Services, requestToken *oauth.RequestToken) error {
  tx, err := s.storage.Begin()
  if err != nil {
    return err
  }
  stmt,err := tx.Prepare("insert into users(id,request_token,request_token_secret) values(NULL,?,?)")
  if err != nil {
    return err
  }
  defer stmt.Close()

  _, err = stmt.Exec(requestToken.Token, requestToken.Secret)
  if err != nil {
    return err
  }
  tx.Commit()

  return nil
}
