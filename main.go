
package main

import (
	"PhoneBook/handler"
	"PhoneBook/repository"
	"PhoneBook/service"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

//var ID = 1
//
//type Contact struct {
//	id    int
//	name  string
//	phone string
//}
//
////type Contacts []Contact
//var contacts = []Contact{
//	{
//		1,
//		"FIRST",
//		"111",
//	},
//}
func main() {
	phonebook := repository.NewPhoneBookRepository()
	srv := service.NewPhoneBookService(phonebook)
	handlers := handler.NewPhoneBookService(srv)

	router := httprouter.New()
	router.PUT("/contact/:id", handlers.UpdateContactHandler)
	router.GET("/contact/:id", handlers.GetContactHandler)
	router.GET("/contacts/list", handlers.GetContactListHandler)
	router.POST("/contact", handlers.AddContactHandler)
	router.DELETE("/contact/:id", handlers.DeleteContactHandler)

	fmt.Println("Server is listening...")
	log.Fatal(http.ListenAndServe("localhost:8181", router))
}