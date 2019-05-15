package main

import (
	"errors"
	"fmt"
	"hash/fnv"
	"net/http"
	"time"
)

type CookieSession struct {
	ExpirationDate time.Time `json:"expirationDate"`
	Login          string    `json:"login"`
}

type ResponseBox struct {
	Session CookieSession `json:"session"`
	Content interface{}   `json:"content"`
}

var sessions = make(map[string]CookieSession)
var users = map[string]string{
	"admin":     "-995833633",  //padmin
	"valentine": "-1749185786", //pvalentine
}

func AccessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
		session := GetSession(r)
		fmt.Printf("[AccessLog]%s - %s %s Session id = %s \n", r.RequestURI, r.Method, r.Host, session)
		next.ServeHTTP(w, r)
	})
}

func SecureEndoint(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
		if r.RequestURI == "/api/v1/login/" || r.RequestURI == "/" || r.RequestURI == "/js/login.js" || r.RequestURI == "/css/main.css" {
			next.ServeHTTP(w, r)
		} else {
			session := GetSession(r)
			ok, err := CheckSession(session)
			if ok {
				fmt.Printf("Session id = %s \n", session)
				next.ServeHTTP(w, r)
			} else {
				w.Write([]byte(err.Error()))
			}
		}
	})
}

// GetUserBySession returns Username by sessionId
func GetUserBySession(userName string) CookieSession {
	return sessions[userName]
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

// Logout will clear up session
func Logout(w http.ResponseWriter, r *http.Request) {
	sessionID, err := r.Cookie("sessionId")
	if err != nil {
		fmt.Println("Cookie not set!")
	} else {
		fmt.Println("User cookie sessionId=" + sessionID.Value)
	}
	delete(sessions, sessionID.Value)
}

// GetSession returns session string from *http.Request
func GetSession(r *http.Request) string {
	if r != nil {
		session, err := r.Cookie("session")
		if err != nil {
			fmt.Println("empty session coockie")
			return ""
		}
		return session.Value
	}
	return "request is nil"

}

// CheckSession returns true if given session string contains in session and not expired
func CheckSession(s string) (bool, error) {
	fmt.Println("CheckSession")
	ses, ok := sessions[s]
	if ok {
		if time.Now().Before(ses.ExpirationDate) {
			return true, nil
		}
		return false, errors.New("session expired")
	}
	return false, errors.New("session not found")
}

// Login reads hash from request body, check if it valid, will create session for user and assign it to user cookie
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login invoked!")
	userCreated := false
	userName := r.FormValue("userName")
	userHash := r.FormValue("userHash")
	// Check username+password

	fmt.Printf("пиздуем в бд за юзером %s с хешем %s\n", userName, userHash)
	pwd := getUser(userName)
	if pwd == "" {
		userCreated = createUser(userName, userHash)
		if !userCreated {
			panic("error creating user")
		}
	}

	if userCreated || pwd == userHash {
		timestamp := time.Now()
		salt := timestamp.Format("2006-01-02 15:04:05")
		// Create session ID
		sessionID := fmt.Sprint(hash(userName + userHash + salt))
		fmt.Printf("userName=%s userHash =%s salt=%s hashed=%s \n", userName, userHash, salt, sessionID)
		expiration := timestamp.Add(10 * time.Minute)
		sessions[sessionID] = CookieSession{ExpirationDate: expiration, Login: userName}
		cookieSessionID := http.Cookie{
			Name:    "session",
			Path:    "/",
			Value:   sessionID,
			Expires: expiration}

		cookieUserName := http.Cookie{
			Name:    "userName",
			Path:    "/",
			Value:   userName,
			Expires: expiration}

		http.SetCookie(w, &cookieSessionID)
		http.SetCookie(w, &cookieUserName)
		http.Redirect(w, r, "/income.html", http.StatusFound)
	} else {
		loginError := "Wrong login/password"
		http.Error(w, loginError, 500)
	}
}
