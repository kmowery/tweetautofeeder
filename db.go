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

func updateAccessToken(s Services, requestToken *oauth.RequestToken, atoken *oauth.AccessToken, sessionCookie string) error {
  tx, err := s.storage.Begin()
  if err != nil {
    return err
  }
  stmt,err := tx.Prepare("update users set oauth_token = ?, oauth_token_secret = ?, user_id = ?, screen_name = ?, session_cookie = ? where request_token = ?")
  if err != nil {
    return err
  }
  defer stmt.Close()

  // TODO: check if user_id and screen_name exist

  _, err = stmt.Exec(atoken.Token, atoken.Secret, atoken.AdditionalData["user_id"], atoken.AdditionalData["screen_name"], sessionCookie, requestToken.Token)
  if err != nil {
    return err
  }
  tx.Commit()

  return nil

}

func getUserWithCookie(s Services, sessionCookie string) (*User, error) {
  var user *User = &User{}

  err := s.storage.QueryRow("select user_id, screen_name, oauth_token, oauth_token_secret from users where session_cookie = ?", sessionCookie).Scan(
    &user.userId, &user.screenName, &user.atoken.Token, &user.atoken.Secret)

  switch {
  case err == sql.ErrNoRows:
    return nil, nil
  case err != nil:
    return nil, err
  default:
  }
  return user, nil
}

func addTweet(s Services, user User, tweet string) error {
  tx, err := s.storage.Begin()
  if err != nil {
    return err
  }
  stmt,err := tx.Prepare("insert into tweets(id,user_id,tweet,completed) values(NULL,?,?,0)")
  if err != nil {
    return err
  }
  defer stmt.Close()

  _, err = stmt.Exec(user.userId, tweet)
  if err != nil {
    return err
  }
  tx.Commit()

  return nil
}

func completeTweet(s Services, user *User, tweetid string) error {
  tx, err := s.storage.Begin()
  if err != nil {
    return err
  }
  stmt,err := tx.Prepare("update tweets set completed = 1 where user_id = ? and id = ?")
  if err != nil {
    return err
  }
  defer stmt.Close()

  _, err = stmt.Exec(user.userId, tweetid)
  if err != nil {
    return err
  }
  tx.Commit()

  return nil

}

func getTweet(s Services, user *User, tweetid string) (*Tweet, error) {
  var id string
  var tweet string

  err := s.storage.QueryRow("select id,tweet from tweets where user_id = ? and id = ?", user.userId, tweetid).Scan(&id, &tweet)

  switch {
  case err == sql.ErrNoRows:
    return nil, nil
  case err != nil:
    return nil, err
  default:
    return &Tweet{id,tweet}, nil
  }
}

func getTweets(s Services, user User, completed bool) ([]Tweet, error) {
  rows, err := s.storage.Query("select id,tweet from tweets where user_id = ? and completed = ?",
     user.userId, completed)

  switch {
  case err == sql.ErrNoRows:
    return nil, nil
  case err != nil:
    return nil, err
  default:
  }

  var tweets []Tweet

  for rows.Next() {
    var tweet string
    var id string
    err = rows.Scan(&id,&tweet)
    if(err == nil) {
      tweets = append(tweets, Tweet{id, tweet})
    } else {
      log.Println(err)

    }
  }
  return tweets, nil
}

func deleteTweet(s Services, user User, tweetid string) error {
  tx, err := s.storage.Begin()
  if err != nil {
    return err
  }
  stmt,err := tx.Prepare("delete from tweets where user_id = ? and id = ?")
  if err != nil {
    return err
  }
  defer stmt.Close()

  _, err = stmt.Exec(user.userId, tweetid)
  if err != nil {
    return err
  }
  tx.Commit()

  return nil

}

