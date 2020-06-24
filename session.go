package goddit

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func OAuthLoginSession(username string, password string, useragent string, clientID string, clientSecret string) (*Session, error) {
	session := &Session{
		username:     username,
		password:     password,
		clientID:     clientID,
		clientSecret: clientSecret,
		useragent:    useragent,
	}

	tokenURL := normalURL + "/api/v1/access_token"

	formData := url.Values{
		"grant_type": {"password"},
		"username":   {username},
		"password":   {password},
	}

	req, err := http.NewRequest(POST, tokenURL, bytes.NewBufferString(formData.Encode()))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("HTTP Request error. URL: %s\nMethod:%s\nHost:%s\n", req.URL, req.Method, req.Host))
	}

	req.Header.Set("User-Agent", useragent)
	req.SetBasicAuth(clientID, clientSecret)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		error_string := fmt.Sprintf("HTTP ERROR\nStatus Code: %d\nError: %s", resp.StatusCode, resp.Status)
		if resp.StatusCode == 429 {
			return nil, errors.New(error_string + "\nPossible Fix: Make sure you have a unique user agent.")
		} else if resp.StatusCode == 401 {
			return nil, errors.New(error_string + "\nPossible Fix: Check to make sure your client ID and client Secret are both correct.")
		} else {
			return nil, errors.New(error_string)
		}
	}

	body, err := ioutil.ReadAll(resp.Body)
	bodyString := string(body)

	var oauthjson map[string]interface{}
	json.Unmarshal([]byte(bodyString), &oauthjson)

	token, token_ok := oauthjson["access_token"]
	expiretime, time_ok := oauthjson["expires_in"]
	if token_ok && time_ok {
		session.accesstoken = token.(string)
		session.expiretime = expiretime.(float64)
		go session.ExpireTimeCountdown()
	} else {
		return nil, errors.New(fmt.Sprintf("Cannot get Access Token Error: %s\nPossible Fix: Make sure Username and Password are correct.", oauthjson["error"]))
	}

	return session, nil
}

func (sess *Session) ExpireTimeCountdown() {
	expireTime := sess.expiretime
	for {
		time.Sleep(time.Second)
		expireTime -= 1
		if expireTime == 0 {
			for {
				session, err := OAuthLoginSession(
					sess.username,
					sess.password,
					sess.useragent,
					sess.clientID,
					sess.clientSecret,
				)

				if err != nil {
					continue
				} else {
					sess.accesstoken = session.accesstoken
					sess.ExpireTimeCountdown()
					break
				}
			}
		}
	}
}
