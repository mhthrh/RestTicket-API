package ServiceHandler

import (
	CustomersHandler "TicketManager/Customer"
	EventHandler "TicketManager/Event"
	TicketHandler "TicketManager/Ticket"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type TicketService struct {
	ticket   TicketHandler.Ticket
	customer CustomersHandler.Customer
	event    EventHandler.Event
}

func New() *TicketService {
	return &TicketService{
		ticket:   *TicketHandler.New(nil),
		customer: *CustomersHandler.New(nil),
		event:    *EventHandler.New(nil),
	}
}

func RunApi(endpoint string, t *TicketService) error {
	r := mux.NewRouter()
	RunApiOnRouter(r, *t)
	fmt.Println("Server Started ...")
	return http.ListenAndServe(endpoint, r)
}

func RunApiOnRouter(r *mux.Router, t TicketService) {
	handler := New()
	apiRouter := r.PathPrefix("/api/ticket").Subrouter()
	apiRouter.Methods("GET").Path("/buyTicket/{search}").HandlerFunc(handler.CheckTicket)
	apiRouter.Methods("GET").Path("/AllEvent").HandlerFunc(handler.AllEvents)
	apiRouter.Methods("POST").Path("/{operation:(?:signIn|signUp)}").HandlerFunc(handler.Login)

}

func (handler TicketService) CheckTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	search, ok := vars["search"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Search not Found.")
		return
	}

	handler.ticket.Check("")
	if handler.ticket.Status.Key != 0 {
		fmt.Println("Ticket not found!")
		w.WriteHeader(410)
		fmt.Fprintf(w, "Ticket by Serial %s not Found.", search)
		return
	}

	json.NewEncoder(w).Encode(handler.ticket.Status)

}
func (handler TicketService) AllEvents(w http.ResponseWriter, r *http.Request) {
	handler.event.Events()
	if handler.event.Status.Key != 0 {
		fmt.Println("Event not found!")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "No Event")
		return
	}
	json.NewEncoder(w).Encode(handler.event.Events())
}
func (handler TicketService) Login(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	op, ok := vars["operation"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "operation not Found.")
		return
	}

	err := json.NewDecoder(r.Body).Decode(&handler)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could not Decode request body by error: %v", err)
		return
	}

	switch strings.ToLower(op) {
	case "SignIn":
		handler.customer.SignIn()
		if handler.customer.Status.Key != 0 {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Could not find by error: %v", err)
			return
		}

	case "SignUp":
		handler.customer.SignUp()
		if handler.customer.Status.Key != 0 {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Could not Insert: %v", err)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(handler.customer)
}
