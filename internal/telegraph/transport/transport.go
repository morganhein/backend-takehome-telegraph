package transport

import (
	"encoding/json"
	"fmt"
	"github.com/morganhein/backend-takehome-telegraph/internal/telegraph"
	"net/http"
	"strings"
)

type Transport interface {
	ListEquipment(w http.ResponseWriter, r *http.Request)
	ListEvents(w http.ResponseWriter, r *http.Request)
	ListLocations(w http.ResponseWriter, r *http.Request)
	ListWaybills(w http.ResponseWriter, r *http.Request)
	GetWaybill(w http.ResponseWriter, r *http.Request)
}

func NewHTTPTransport(srv telegraph.Service) *httpEndpoint {
	return &httpEndpoint{
		srv: srv,
	}
}

type httpEndpoint struct {
	srv telegraph.Service
}

func (t httpEndpoint) ListEquipment(w http.ResponseWriter, r *http.Request) {
	eq, err := t.srv.ListEquipment()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//marshal
	b, err := json.Marshal(eq)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	_, err = w.Write(b)
	if err != nil {
		// pipe probably broken, so don't try to write response again
		fmt.Println(err)
		return
	}
}

func (t httpEndpoint) ListEvents(w http.ResponseWriter, r *http.Request) {
	eq, err := t.srv.ListEvents()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//marshal
	b, err := json.Marshal(eq)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	_, err = w.Write(b)
	if err != nil {
		// pipe probably broken, so don't try to write response again
		fmt.Println(err)
		return
	}
}

func (t httpEndpoint) ListLocations(w http.ResponseWriter, r *http.Request) {
	eq, err := t.srv.ListLocations()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//marshal
	b, err := json.Marshal(eq)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	_, err = w.Write(b)
	if err != nil {
		// pipe probably broken, so don't try to write response again
		fmt.Println(err)
		return
	}
}

func (t httpEndpoint) ListWaybills(w http.ResponseWriter, r *http.Request) {
	eq, err := t.srv.ListWaybills()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//marshal
	b, err := json.Marshal(eq)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	_, err = w.Write(b)
	if err != nil {
		// pipe probably broken, so don't try to write response again
		fmt.Println(err)
		return
	}
}

func (t httpEndpoint) Getwaybill(w http.ResponseWriter, r *http.Request) {
	// I feel like there is probably some sub-router/demuxing that already exists in a library, which
	// would be an improvement over this.
	route := strings.TrimPrefix(r.URL.Path, "/waybills/")
	parts := strings.Split(route, "/")
	if len(parts) == 1 {
		parts = append(parts, "")
	}
	eq, err := t.srv.GetWaybill(parts[0], parts[1])
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//marshal
	b, err := json.Marshal(eq)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	_, err = w.Write(b)
	if err != nil {
		// pipe probably broken, so don't try to write response again
		fmt.Println(err)
		return
	}
}
