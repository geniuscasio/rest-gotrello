package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"encoding/json"
	"unicode/utf8"
	"regexp"
	"strconv"
	"strings"
	"errors"
	
	ends "github.com/geniuscasio/rest-gotrello/endpoints"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type inMsg struct {
	Messages []struct {
		Phone   int64  `json:"phone_number"`
		Extraid string `json:"extra_id"`
		Text    string `json:"text,omitempty"`
	} `json:"messages"`
	CallbackURL    string   `json:"callback_url"`
	StartTime      string   `json:"start_time"`
	Tag            string   `json:"tag"`
	Channels       []string `json:"channels"`
	ChannelOptions struct {
		Sms struct {
			AlphaName string `json:"alpha_name,omitempty"`
			TTL       int64  `json:"ttl,omitempty"`
		} `json:"sms,omitempty"`
		Viber struct {
			TTL     int64  `json:"ttl"`
			Img     string `json:"img"`
			Caption string `json:"caption"`
			Action  string `json:"action"`
		} `json:"viber"`
	} `json:"channel_options"`
}

type outMsg struct {
	Phone  string `json:"phone,omitempty"`
	Status string `json:"status"`
}

type IErrorCheck func(inMsg) error

const LATIN_MAX_SIZE = 765
const CYRILIC_MAX_SIZE = 365

func main() {
	router := mux.NewRouter()
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("Env POST must be set!")
		port = ":8005"
	} else {
		port = ":" + port
	}
	var dir string

	flag.StringVar(&dir, "dir", "./static", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	fmt.Println("Running server on port " + port)

	// GETs
	router.HandleFunc("/api/v1/login/", ends.Login).Methods("POST")
	router.HandleFunc("/api/v1/income/", ends.SecureEndpoint(ends.Get)).Methods("GET")
	router.HandleFunc("/api/v1/income/{id}", ends.SecureEndpoint(ends.Get)).Methods("GET")

	// POSTs
	router.HandleFunc("/api/v1/income/", ends.SecureEndpoint(ends.Create)).Methods("POST")
	
	router.HandleFunc("/test", sayHello).Methods("GET")

	// Static "/" must be last in code
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(dir))))

	log.Fatal(http.ListenAndServe(port, router))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	var in inMsg
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&in)
	if err != nil {
		fmt.Println(err)
		return
	}
	errorExists := false
	checkError := func(f IErrorCheck) {
		if errorExists == true {
			fmt.Println("Another one bite to dust")
			return
		}
		err := f(in)
		if err != nil {
			errorExists = true
			fmt.Println("Error", err)
			out := outMsg{Status: "Error " + err.Error()}
			var buf []byte
			buf, err = json.Marshal(out)
			_, err = w.Write(buf)
		}
	}
	checkError(checkNumber)
	checkError(checkText)
}

func checkText(i inMsg) error {
	var isViber bool
	var isSms bool
	var maximumSize int = 1000
	for _, b := range i.Channels {
		if b == "sms" {
			isSms = true
		}
		if b == "viber" {
			isViber = true
		}
	}
	for _, msg := range i.Messages {
		if isSms && isViber {
			if isCyrilic(msg.Text) {
				maximumSize = CYRILIC_MAX_SIZE
			} else {
				maximumSize = LATIN_MAX_SIZE
			}
		} else if isViber {
			maximumSize = 1000
		}
	}
	for _, msg := range i.Messages {
		if utf8.RuneCountInString(msg.Text) > maximumSize || utf8.RuneCountInString(msg.Text) == 0 {
			return fmt.Errorf("checkText: msg=%s minSize=1 maxSize=%d actualSize=%d", msg.Text, maximumSize, utf8.RuneCountInString(msg.Text))
		}
	}
	return nil
}

func isCyrilic(s string) bool {
	var re = regexp.MustCompile(`(?m)\p{Cyrillic}`)
	_isCyrilic := re.Match([]byte(s))
	return _isCyrilic
}

func checkNumber(i inMsg) error {
	for _, i := range i.Messages {
		n := strconv.Itoa(int(i.Phone))
		if len(n) != 12 || !strings.HasPrefix(n, "380") {
			return errors.New("checkNumber: invalid phone number")
		}
	}
	return nil
}
