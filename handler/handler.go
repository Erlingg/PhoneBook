package handler

import (
	"PhoneBook/interface1"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type phoneBooksHandlers struct {
	srv interface1.ContactService
}


func NewPhoneBookService(s interface1.ContactService) *phoneBooksHandlers {
	return &phoneBooksHandlers{srv: s}
}
func (h phoneBooksHandlers) AddContactHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	contact := interface1.Contact{}
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		log.Panicln(err)
		return
	}
	if err := h.srv.AddContactService(contact); err != nil {
		//_, _ =w.Write(err)
	}
}

func (h phoneBooksHandlers) GetContactHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	contact := h.srv.GetContactService(ps.ByName("id"))
	b, err := json.Marshal(contact)
	if err != nil {
		log.Panicln(err)
		return
	}
	_, _ = w.Write(b)

}

func (h phoneBooksHandlers) GetContactListHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	list := h.srv.GetContactListService()
	b, err := json.Marshal(list)
	if err != nil {
		log.Panicln(err)
		return
	}
	_, _ = w.Write(b)
}

func (h phoneBooksHandlers) UpdateContactHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	contact := interface1.Contact{}
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		log.Panicln(err)
		return
	}
	id := ps.ByName("id")
	h.srv.UpdateContactService(id, contact)
}

func (h *phoneBooksHandlers) DeleteContactHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	h.srv.DeleteContactService(id)
}

