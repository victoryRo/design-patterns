package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// Comenzando con la interfaz común, http.Handler
// Ya tenemos la interfaz común que decoraremos usando tipos anidados.
// Primero necesitamos crear nuestro tipo principal,
// que será el Handler que devuelva la oración Hello Decorator!:

type MyServer struct{}

func (m *MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Decorator!")
}

func LocalMainOne() {
	http.Handle("/", &MyServer{})

	fmt.Println("Run server port 3001 ...")
	log.Fatal(http.ListenAndServe(":3001", nil))
}

// --------------------------------------------------------------------------------

type LoggerMiddleware struct {
	Handler   http.Handler
	LogWriter io.Writer
}

func (l *LoggerMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(l.LogWriter, "Request URI: %s\n", r.RequestURI)
	fmt.Fprintf(l.LogWriter, "Host: %s\n", r.Host)
	fmt.Fprintf(l.LogWriter, "Content Length: %d\n", r.ContentLength)
	fmt.Fprintf(l.LogWriter, "Method: %s\n", r.Method)
	fmt.Fprintf(l.LogWriter, "-----------------------------------\n")
	l.Handler.ServeHTTP(w, r)
}

func LocalMainTwo() {
	// Hemos decorado MyServer con capacidades de registro sin modificarlo
	http.Handle("/", &LoggerMiddleware{
		Handler:   &MyServer{},
		LogWriter: os.Stdout,
	})

	fmt.Println("Run server port: 3003")
	log.Fatal(http.ListenAndServe(":3003", nil))
}

// --------------------------------------------------------------------------------

type SimpleAuthMiddleware struct {
	Handler  http.Handler
	User     string
	Password string
}

func (s *SimpleAuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, pass, ok := r.BasicAuth()
	fmt.Printf("user: %s pass: %s", user, pass)

	if ok {
		if user == s.User && pass == s.Password {
			s.Handler.ServeHTTP(w, r)
		} else {
			fmt.Fprintln(w, "user or password incorrect")
		}
	} else {
		fmt.Fprintln(w, "error trying to retrieve data from Basic auth")
	}
}

// --------------------------------------------------------------------------------

func LocalMain() {
	fmt.Println("Enter the type number of server you want to launch from the following:")
	fmt.Println("1: - Plain server")
	fmt.Println("2: - Server with logging")
	fmt.Println("3. - Server with logging and authentication")

	var selection int
	fmt.Fscanf(os.Stdin, "%d", &selection)

	var mySuperServer http.Handler

	switch selection {
	case 1:
		mySuperServer = new(MyServer)
	case 2:
		mySuperServer = &LoggerMiddleware{
			Handler:   new(MyServer),
			LogWriter: os.Stdout,
		}
	case 3:
		var user, password string

		fmt.Println("Enter user and password separated by a space")
		fmt.Fscanf(os.Stdin, "%s %s", &user, &password)

		mySuperServer = &LoggerMiddleware{
			Handler: &SimpleAuthMiddleware{
				Handler:  new(MyServer),
				User:     user,
				Password: password,
			},
			LogWriter: os.Stdout,
		}
	default:
		mySuperServer = new(MyServer)
	}

	http.Handle("/", mySuperServer)

	fmt.Println("Runing server at port: 3007")
	log.Fatal(http.ListenAndServe(":3007", nil))
}

// --------------------------------------------------------------------------------
