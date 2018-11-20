package endpoints

import (
	"fmt"
	"hash/fnv"
	"net/http"
	"time"
)

type mySession struct {
	expirationDate time.Time
	login          string
}

var sessions = make(map[string]mySession)
var users = map[string]string{
	"admin":     "-995833633",  //padmin
	"valentine": "-1749185786", //pvalentine
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

// CheckSession exits and not expired
func CheckSession(userName, sessionID string) (string, bool) {
	sessionInfo, ok := sessions[sessionID]
	message := "session not found"
	status := false
	if ok {
		message = "ok"
		status = true
		currentTime := time.Now()
		if currentTime.After(sessionInfo.expirationDate) {
			message = "session get old"
			status = false
		}
	}
	return message, status
}

// Login reads hash from request body, check if it valid, will create session for user and assign it to user cookie
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login invoked!")
	userName := r.FormValue("userName")
	userHash := r.FormValue("userHash")

	pass, ok := users[userName]

	if ok && pass == userHash {
		timestamp := time.Now()
		salt := timestamp.Format("2006-01-02 15:04:05")
		sessionID := fmt.Sprint(hash(userName + userHash + salt))

		fmt.Printf("userName=%s userHash =%s salt=%s hashed=%s \n", userName, userHash, salt, sessionID)

		expiration := timestamp.Add(10 * time.Second)
		sessions[sessionID] = mySession{expirationDate: expiration, login: userName}
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
