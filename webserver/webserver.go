package webserver

import (
	"encoding/json"
	"fmt"
	"github.com/StCredZero/rsrvtrst_thing/fib"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	mux *mux.Router
	Fibber *fib.Fibber
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func makeFunctionHandler(varname string, fn func(n uint64) uint64) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		inputString := mux.Vars(req)[varname]
		input, err :=strconv.ParseUint(inputString, 10, 64)
		if err != nil {
			respondWithError(rw, http.StatusInternalServerError, err.Error())
		} else {
			result := fn(input)
			resultStr := strconv.FormatUint(result, 10)
			fmt.Fprintf(rw, resultStr)
		}
	}
}

func makeClearHandler(fn func()) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		fn()
		respondWithJSON(rw, http.StatusOK, "OK")
	}
}

func (s *Server) Start() {
	s.mux = mux.NewRouter()
	s.mux.PathPrefix("/ordinal/{ordinal_n}").Handler(makeFunctionHandler("ordinal_n", s.Fibber.FibonaciOrdinal))
	s.mux.PathPrefix("/cardinality_less/{cardinality_x}").Handler(makeFunctionHandler("cardinality_x", s.Fibber.CardinalityLessThan))
	s.mux.PathPrefix("/clear").Handler(makeClearHandler(s.Fibber.Initialize))

	err := http.ListenAndServe(":8080", s.mux.Router)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

