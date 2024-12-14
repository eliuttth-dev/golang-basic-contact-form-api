package main

import (
  "fmt"
  "log"
  "encoding/json"
  "net/http"
)

type ContactForm struct {
  Name    string  `json:"name"`
  Email   string  `json:"email"`
  Message string  `json:"message"`
}

var contactStore = make([]ContactForm, 0)

func submitContactForm(w http.ResponseWriter, r *http.Request){
  if r.Method != http.MethodPost {
    http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
    return
  }

  var form ContactForm
  decoder := json.NewDecoder(r.Body)
  err := decoder.Decode(&form)
  if err != nil {
    http.Error(w, "Invalid input", http.StatusBadRequest)
    return
  }

  contactStore = append(contactStore, form)

  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(map[string]string{"message": "Contact form submitted successfully"})
}

func main(){
  http.HandleFunc("/submit", submitContactForm)

  fmt.Println("Server running: http://localhost:3000")
  log.Fatal(http.ListenAndServe(":3000", nil))
}
