package server

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/sessions"
	"github.com/jgunnink/bluebox"
)

var cookieStore = sessions.NewCookieStore([]byte("gelatinoid-reparable-salmon"))

const cookieName = "_bluebox_session"

type funcDetails struct {
	Handler     string
	Description string
}

// Server is the container for the server instance
type Server struct {
	*http.Server
}

// New returns a new instance of the Server
func New(port string, as bluebox.AuthService, us bluebox.UserService) *Server {
	mw := NewMiddleware(us)
	ac := NewAuthController(mw, as, us)
	uc := NewUserController(mw, us)
	masterHandler := NewMasterHandler(ac, uc)

	s := &http.Server{
		Addr:    port,
		Handler: masterHandler,
	}
	return &Server{
		Server: s,
	}
}

// Run will begin the server
func (s *Server) Run() error {
	return s.ListenAndServe()
}

// Shutdown will gracefully shutdown the Server
func (s *Server) Shutdown() error {
	return s.Shutdown()
}

// MasterHandler is a collection of all the service Handlers.
type MasterHandler struct {
	AuthController *AuthController
	UserController *UserController
	ProxyHandler   func(http.ResponseWriter, *http.Request)
	FileServer     *http.ServeMux
}

// NewMasterHandler returns the top level handler
func NewMasterHandler(ac *AuthController, uc *UserController) *MasterHandler {
	log.Println("Run fileserver")
	fileshttp := http.NewServeMux()
	fileshttp.Handle("/", http.FileServer(http.Dir("./public/")))

	return &MasterHandler{
		AuthController: ac,
		UserController: uc,
		FileServer:     fileshttp,
	}
}

// ServeHTTP delegates a request to the appropriate subhandler.
func (h *MasterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Header.Get("X-API") != "":
		h.HandleAPI(w, r)
	default:
		h.FileServer.ServeHTTP(w, r)
	}
}

// HandleAPI will handle all API calls
func (h *MasterHandler) HandleAPI(w http.ResponseWriter, r *http.Request) {
	log.Println("Run API server")
	switch {
	case strings.HasPrefix(r.URL.Path, "/users"):
		log.Println("Run UserController")
		h.UserController.ServeHTTP(w, r)
	case strings.HasPrefix(r.URL.Path, "/auth"):
		log.Println("Run AuthController")
		h.AuthController.ServeHTTP(w, r)
	default:
		http.NotFound(w, r)
	}
}

// HandleReactProxy will serve the react app
func (h *MasterHandler) HandleReactProxy(w http.ResponseWriter, r *http.Request) {
	log.Println("Run react proxy")
	if os.Getenv("PROD") == "" {
		log.Println(r.Method, r.URL.Path)
		h.ProxyHandler(w, r)
	} else {
		log.Println(r.Method, r.URL.Path)
		h.FileServer.ServeHTTP(w, r)
	}
}
