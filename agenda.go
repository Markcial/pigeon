package Pigeon

import (
    "github.com/flosch/pongo2"
    "net/http"
)

type Value struct {
    value string
}

type Group struct {
    name string
}

type contact struct {
    id int
    phone, email Value
    name string
    group Group
}

func NewContact() contact {
    entity := contact {
        group: Group{},
        email: Value{},
        phone: Value{},
    }

    return entity
}

// Agenda
type Agenda struct {
    contacts []contact
    groups []Group
}

func NewAgenda() Agenda {
    entity := Agenda{
        contacts: make([]contact, 0),
        groups: make([]Group, 0),
    }
    return entity
}

func (a *Agenda) Add (c contact) {
    a.contacts = append(a.contacts, c)
}

func (a *Agenda) All() []contact {
    return a.contacts
}

var agenda = NewAgenda()
var listContactsTpl = pongo2.Must(pongo2.FromFile("templates/agenda/index.pongo.html"))
var editContactTpl = pongo2.Must(pongo2.FromFile("templates/contact/edit.pongo.html"))


func updateContactHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("whoo update"))

}

func editContactHandler(w http.ResponseWriter, r *http.Request) {
   err := editContactTpl.ExecuteWriter(pongo2.Context{"agenda": agenda, "query": r.FormValue("query")}, w)
   if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
   }
}

func addContactHandler(w http.ResponseWriter, r *http.Request) {

  w.Write([]byte("whoo add"))
}

func displayContactsHandler(w http.ResponseWriter, r *http.Request) {
   err := listContactsTpl.ExecuteWriter(pongo2.Context{"agenda": agenda}, w)
   if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
   }
}

func init() {
    r.Get("/contacts", displayContactsHandler)
    r.Post("/contact/edit/{id:[0-9]+}", updateContactHandler)
    r.Get("/contact/edit/{id:[0-9]+}", editContactHandler)
    r.Post("/contact/add", addContactHandler)
    r.Get("/contact/add", addContactHandler)

    contact := contact{name:"name", id:1}
    agenda.Add(contact)
}
