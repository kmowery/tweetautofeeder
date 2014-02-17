package main

import (
  "net/http"
  "net/url"
  "encoding/base64"
  "regexp"
  "crypto/rand"
  "crypto/hmac"
  "crypto/sha1"
  "fmt"
  //"time"
  "sort"
  "strings"
  "github.com/hoisie/mustache"
)

func generateRandomString() string {
  var r *regexp.Regexp;
  var err error;

  r,err = regexp.Compile( `[^\w]` )

  b := make([]byte, 32)
  _, err = rand.Read(b)
  if(err != nil ) {
    // not sure what to do here...
    fmt.Println("couldn't read random bytes...")
  }

  return r.ReplaceAllString(base64.StdEncoding.EncodeToString(b), "")
}

func generateOAuthHeader(base_url url.URL, method string, data map[string][]string) string {

  oauth_parameters := make(map[string][]string)

  //TODO not actually right
  //oauth_parameters["oauth_callback"] = []string{base_url};

  //oauth_parameters["oauth_consumer_key"] = []string{API_CONSUMER_KEY};
  //oauth_parameters["oauth_nonce"] = []string{generateRandomString()};
  //oauth_parameters["oauth_signature_method"] = []string{"HMAC-SHA1"};
  //oauth_parameters["oauth_timestamp"] = []string{fmt.Sprintf("%d", time.Now().Unix())};
  //oauth_parameters["oauth_version"] = []string{"1.0"};

  //oauth_parameters["oauth_callback"] = []string{"http://localhost/sign-in-with-twitter/"}
  //oauth_parameters["oauth_consumer_key"] = []string{"cChZNFj6T5R0TigYB9yd1w"}
  //oauth_parameters["oauth_nonce"] = []string{"ea9ec8429b68d6b77cd5600adbbb0456"}
  //oauth_parameters["oauth_signature"] = []string{"F1Li3tvehgcraF8DMJ7OyxO4w9Y%3D"}
  //oauth_parameters["oauth_signature_method"] = []string{"HMAC-SHA1"}
  //oauth_parameters["oauth_timestamp"] = []string{"1318467427"}
  //oauth_parameters["oauth_version"] = []string{"1.0"}

  for key, value := range base_url.Query() {
    oauth_parameters[url.QueryEscape(key)] = value
  }

  if(data != nil) {
    for key, value := range data {
      // TODO fix if we need more values
      oauth_parameters[url.QueryEscape(key)] = value
    }
  }


  oauth_parameters["oauth_consumer_key"] = []string{"xvz1evFS4wEEPTGEFPHBog"}
  oauth_parameters["oauth_nonce"] = []string{"kYjzVBB8Y0ZFabxSWbWovY3uYSQ2pTgmZeNu2VS4cg"}
  oauth_parameters["oauth_signature_method"] = []string{"HMAC-SHA1"}
  oauth_parameters["oauth_timestamp"] = []string{"1318622958"}
  oauth_parameters["oauth_token"] = []string{"370773112-GmHxMAgYyLbNEtIKZeRNFsMKPR9EyMZeS9weJAEb"}
  oauth_parameters["oauth_version"] = []string{"1.0"}

  keys := make([]string, 0, len(oauth_parameters))
  for k := range oauth_parameters {
      keys = append(keys, k)
  }
  sort.Strings(keys)

  plus_regex, _ := regexp.Compile(`\+`)

  ////TODO: percent encode keys
  parameter_str := ""
  for index,param := range keys {
    parameter_str += param + "=" + plus_regex.ReplaceAllString(
        url.QueryEscape(oauth_parameters[param][0]) , "%20")
    if (index < len(oauth_parameters)-1 ) {
      parameter_str += "&";
    }
  }

  fmt.Printf("parameter_str: %s\n\n", parameter_str)

  // use base url instead of full url with query bits
  signature_base_str := fmt.Sprintf("%s&%s&%s",
      method,
      url.QueryEscape(strings.Split( base_url.String(), "?")[0] ),
      url.QueryEscape(parameter_str))

  fmt.Printf("signature base: %s\n", signature_base_str)

  // TODO: doesn't support doing anything as a user yet, fix later
  signing_key := fmt.Sprintf("%s&", API_CONSUMER_SECRET)

  signing_key = fmt.Sprintf("%s&%s", "kAcSOqF21Fu85e7zjz7ZN2U4ZRhfV3WpwPAoE3Z7kBw", "LswwdoUaIvS8ltyTt5jkRh4J50vUPVVHtR2YPi5kE")

  mac := hmac.New(sha1.New, []byte(signing_key))
  mac.Write([]byte(signature_base_str))

  hmac_bytes := mac.Sum(nil)
  hmac_base64 := base64.StdEncoding.EncodeToString(hmac_bytes)

  oauth_parameters["oauth_signature"] = []string{hmac_base64}
  oauth_str := "Oauth ";

  for param := range oauth_parameters {
    oauth_str += fmt.Sprintf(`%s="%s"`, param, url.QueryEscape(oauth_parameters[param][0]));
    oauth_str += ", ";
  }

  //ugly
  return strings.TrimSuffix(oauth_str, ", ")
}

func makeRequestFor(url url.URL, method string, data map[string][]string) *http.Request {
  oauth_str := generateOAuthHeader(url, method, data)

  fmt.Printf("oauth str: %s\n\n", oauth_str)

  // Step 1: make POST to twitter
  request, _ := http.NewRequest("method", url.Path, nil)
  request.Header.Add("Authorization", oauth_str)

  return request
}



func loginHandler(w http.ResponseWriter, r *http.Request) {

  //client := &http.Client{};
  
  url,_ := url.Parse("https://api.twitter.com/1/statuses/update.json?include_entities=true")

  // Step 1: make POST to twitter
  makeRequestFor(*url, "POST", map[string][]string{"status": {"Hello Ladies + Gentlemen, a signed OAuth request!"}} )
  //makeRequestFor("https://api.twitter.com/oauth/request_token", "POST")
  //request := makeRequestFor("https://api.twitter.com/oauth/request_token", "POST")
  //resp,err := client.Do(request)

  //if(err != nil) {
  //  fmt.Printf("error fetching page: %d", err)
  //}

  //fmt.Printf("Status code: %d\n", resp.StatusCode)
  //fmt.Printf("Headers:     %s\n", resp.Header)

  data := mustache.RenderFile("/usr/share/tweetautofeeder/templates/blog_main.must", map[string]string{"thing":"places"})
  w.Write([]byte(data))
  return
}

