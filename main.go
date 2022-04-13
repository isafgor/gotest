package main

import (
  "fmt"
  "net/http"
  "html/template"
)

type User struct {
  Name string
  Age uint16
  Money int16
  Avg_grades, Happiness float64
  Hobbies []string
}

func (u User) getAllInfo () string {
  return fmt.Sprintf("Username is: %s. He is %d and he " +
     "has money equal: %d", u.Name, u.Age, u.Money)
}

func (u *User) setNewName (n string) {
  u.Name = n
}

func home_page(w http.ResponseWriter, r *http.Request) {
  bob := User{"Bob", 25, -50, 4.2, 0.8, []string{"soccer", "chess", "hookah"}}
  bob.setNewName("Igor")
  // fmt.Fprintf(w, bob.getAllInfo())

  tmpl, _ := template.ParseFiles("templates/index.html")
  tmpl.Execute(w, bob)
}

func about_page(w http.ResponseWriter, r *http.Request) {
  about := User{"Искандер", 23, -50, 4, 1, []string{"Настолки", "Кальян", "TypeScript"}}

  tmpl, _ := template.ParseFiles("templates/about.html")
  tmpl.Execute(w, about)
}

func err_page(w http.ResponseWriter, r *http.Request) {

  tmpl, _ := template.ParseFiles("templates/err.html")
  tmpl.Execute(w, nil)
}

func handleRequest() {
  http.HandleFunc("/", home_page)
  http.HandleFunc("/about/", about_page)
  http.HandleFunc("/err/", err_page)
  http.ListenAndServe(":8000", nil)
}

func main() {
  handleRequest()
}
