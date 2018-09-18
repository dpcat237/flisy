package router

import (
	"crypto/tls"
	"net/http"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"

	"gitlab.com/dpcat237/flisy/src/router/controller"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type router struct {
	srv      *http.Server
	certFile string
	keyFile  string
}

func New(cc controller.Collector, port, certFile, keyFile string) *router {
	routes := []route{
		/** Seat **/
		{
			"List seats",
			"GET",
			"/seat",
			cc.UsCnt.ValidationMiddleware(cc.StCnt.GetSeats),
		},
		{
			"Assign next available seat",
			"PUT",
			"/seat/assign",
			cc.UsCnt.ValidationMiddleware(cc.StCnt.AssignSeat),
		},
		{
			"Get seat",
			"GET",
			"/seat/{index}",
			cc.UsCnt.ValidationMiddleware(cc.StCnt.GetSeat),
		},

		/** User **/
		{
			"Generate token", // Implemented for demo to be able make required requests with SSL
			"GET",
			"/user/token",
			cc.UsCnt.GenerateToken,
		},
	}

	return create(routes, port, certFile, keyFile)
}

func (r *router) Init() {
	logger.Fatal(r.srv.ListenAndServeTLS(r.certFile, r.keyFile))
}

func create(routes []route, port, certFile, keyFile string) *router {
	var r router
	r.certFile = certFile
	r.keyFile = keyFile
	hdl := mux.NewRouter().StrictSlash(true)
	hdl.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		w.Write([]byte("This is an example server.\n"))
	})

	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		hdl.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	r.srv = &http.Server{
		Addr:         ":" + port,
		Handler:      hdl,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}

	return &r
}
