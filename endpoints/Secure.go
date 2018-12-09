package endpoints

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"net/http"
	"time"
)

type CookieSession struct {
	ExpirationDate time.Time `json:"expirationDate"`
	Login          string    `json:"login"`
	ID             string    `json:"id"`
	Status         bool      `json:"status"`
	Message        string    `json:"message"`
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

func SecureEndpoint(f SecureHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. Verify session and create ResponseBox
		// 2. Pass ResponseBox to endpoint, get data and session info wrapped in struct ResponseBox
		// 3. Marshal ResponseBox and send as response
		s := GetSession(r)
		rb := ResponseBox{s, nil}
		if rb.Session.Status {
			f(r, &rb)
			json.NewEncoder(w).Encode(rb)
		} else {
			json.NewEncoder(w).Encode(rb)
		}
	}

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

// GetSession exits and not expired
func GetSession(r *http.Request) CookieSession {
	userName := ""
	cookie, err := r.Cookie("userName")
	if err == nil {
		userName = cookie.Value
	}

	session, ok := sessions[userName]
	session.Message = "session not found"
	session.Status = false
	if ok {
		session.Message = "ok"
		session.Status = true
		currentTime := time.Now()
		if currentTime.After(session.ExpirationDate) && session.Status {
			session.Message = "session get old"
			session.Status = false
		}
	}
	return session
}

// Login reads hash from request body, check if it valid, will create session for user and assign it to user cookie
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login invoked!")
	userName := r.FormValue("userName")
	userHash := r.FormValue("userHash")
	// Check username+password
	pass, ok := users[userName]

	if ok && pass == userHash {
		timestamp := time.Now()
		salt := timestamp.Format("2006-01-02 15:04:05")
		// Create session ID
		sessionID := fmt.Sprint(hash(userName + userHash + salt))
		fmt.Printf("userName=%s userHash =%s salt=%s hashed=%s \n", userName, userHash, salt, sessionID)
		expiration := timestamp.Add(10 * time.Minute)
		sessions[userName] = CookieSession{ExpirationDate: expiration, Login: userName, ID: sessionID}
		cookieSessionID := http.Cookie{
			Name:    "sessionID",
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
