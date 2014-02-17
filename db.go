package main

import (
  "database/sql"
  "log"
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

func getRequestToken(s Services, requestTokenStr string) (*oauth.RequestToken, error) {
  log.Print( "requestion token " + requestTokenStr)
  var requestToken string
  var requestTokenSecret string

  err := s.storage.QueryRow("select request_token, request_token_secret from users where request_token = ?", requestTokenStr).Scan( &requestToken, &requestTokenSecret )

  var rt *oauth.RequestToken = &oauth.RequestToken{}

  switch {
  case err == sql.ErrNoRows:
    log.Fatal("no user with that token")
  case err != nil:
    log.Fatal(err)
  default:
    rt.Token = requestToken
    rt.Secret = requestTokenSecret
  }
  return rt, nil
}

func updateAccessToken(s Services, requestToken *oauth.RequestToken, atoken *oauth.AccessToken) error {
  tx, err := s.storage.Begin()
  if err != nil {
    return err
  }
  stmt,err := tx.Prepare("update users set oauth_token = ?, oauth_token_secret = ?, user_id = ?, screen_name = ? where request_token = ?")
  if err != nil {
    return err
  }
  defer stmt.Close()

  // TODO: check if user_id and screen_name exist

  _, err = stmt.Exec(atoken.Token, atoken.Secret, atoken.AdditionalData["user_id"], atoken.AdditionalData["screen_name"], requestToken.Token)
  if err != nil {
    return err
  }
  tx.Commit()

  return nil

}
