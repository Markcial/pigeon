package Pigeon

import (
    "github.com/flosch/pongo2"
    "net/http"
    "time"
)

type group struct {
    name string
}

type user struct {
   displayName string
   email string
   password string
   group group
}

// Agenda
type agenda struct {
    users []*user
    groups []*group
}

func NewAgenda() agenda {
    entity := agenda{
        users: make([]*user, 0),
        groups: make([]*group, 0),
    }
    return entity
}

func (a *agenda) FindUser (email string) *user {
    for _, user := range a.users {
       if user.email == email {
          return user
       }
    }

    return nil
}

func (a *agenda) AddUser (u *user) {
    a.users = append(a.users, u)
}

func (a *agenda) AddGroup (g *group) {
    a.groups = append(a.groups, g)
}

func (a *agenda) Contacts() []*user {
    return a.users
}

func (a *agenda) Groups () []*group {
    return a.groups
}

var agendaEntity = NewAgenda()
var listContactsTpl = pongo2.Must(pongo2.FromFile("templates/agenda/index.pongo.html"))
var editContactTpl = pongo2.Must(pongo2.FromFile("templates/contact/edit.pongo.html"))
var loginTpl = pongo2.Must(pongo2.FromFile("templates/index/login.pongo.html"))
var dashboardTpl = pongo2.Must(pongo2.FromFile("templates/index/dashboard.pongo.html"))

func updateContactHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("whoo update"))

}

func editContactHandler(w http.ResponseWriter, r *http.Request) {
   err := editContactTpl.ExecuteWriter(pongo2.Context{"agenda": agendaEntity, "query": r.FormValue("query")}, w)
   if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
   }
}

func addContactHandler(w http.ResponseWriter, r *http.Request) {

  w.Write([]byte("whoo add"))
}

func displayContactsHandler(w http.ResponseWriter, r *http.Request) {
   err := listContactsTpl.ExecuteWriter(pongo2.Context{"agenda": agendaEntity}, w)
   if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
   }
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
    case "GET":
      err := loginTpl.ExecuteWriter(pongo2.Context{}, w)
      if err != nil {
          http.Error(w, err.Error(), http.StatusInternalServerError)
      }
    case "POST":
      email := r.FormValue("email")
      password := r.FormValue("password")
      userEntity := agendaEntity.FindUser(email)
      if userEntity == nil {
        userEntity := &user{email:email, password:password}
        agendaEntity.AddUser(userEntity)

        expiration := time.Now().Add(365 * 24 * time.Hour)
        cookie := http.Cookie{Name: "username", Value: userEntity.email, Expires: expiration}
        http.SetCookie(w, &cookie)
        http.Redirect(w, r, "/dashboard", http.StatusFound)
      } else if userEntity.password == password {
         http.Error(w, "Unauthorized", http.StatusUnauthorized)
      }
      //fmt.Println(r.FormValue("password"))
  }
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
  // read cookie
  cookie, _ := r.Cookie("username")
  // missing cookie? redirect to login
  if cookie == nil {
    http.Redirect(w, r, "/login", http.StatusFound)
  }

  err := dashboardTpl.ExecuteWriter(pongo2.Context{"agenda": agendaEntity}, w)
  if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func init() {

    r.Get("/dashboard", dashboardHandler)
    r.Get("/contacts", displayContactsHandler)
    r.Post("/contact/edit/{id:[0-9]+}", updateContactHandler)
    r.Get("/contact/edit/{id:[0-9]+}", editContactHandler)
    r.Post("/contact/add", addContactHandler)
    r.Get("/contact/add", addContactHandler)

    r.Get("/echo", echoHandler)
    r.Get("/message/compose", composeMessageHandler)
    r.Get("/message/view", viewMessageHandler)
    r.Get("/message/add", addMessageHandler)

    r.Get("/login", loginHandler)
    r.Post("/login", loginHandler)
}
