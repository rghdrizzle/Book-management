package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/rghdrizzle/go-bookstore/pkg/utils"
	"github.com/rghdrizzle/go-bookstore/pkg/models"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request){
	NewBook := models.GetallBook()
	res, _ := json.Marshal(NewBook)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
func GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookid := vars["bookid"]
	id , err := strconv.ParseInt(bookid,0,0)
	if err!= nil {
		fmt.Println("error while parsing")

	}
	bookdetails, _ := models.GetBookById(id)
	res , _ := json.Marshal(bookdetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res , _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookid := vars["bookid"]
	id , err := strconv.ParseInt(bookid,0,0)
	if err!=nil{
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(id)
	res , _ :=json.Marshal(book)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)



}
func UpdateBook(w http.ResponseWriter, r *http.Request){
	var updatebook = &models.Book{}
	utils.ParseBody(r,updatebook)
	vars := mux.Vars(r)
	bookid := vars["bookid"]
	id , err := strconv.ParseInt(bookid,0,0)
	if err!=nil{
		fmt.Println("error while parsing")
	}
	bookdetails , db:=models.GetBookById(id)
	if updatebook.Name!= ""{
		bookdetails.Name = updatebook.Name
	}     
	if updatebook.Author != ""{
		bookdetails.Author = updatebook.Author
	}
	if updatebook.Publication != ""{
		bookdetails.Publication= updatebook.Publication
	}
	db.Save(&bookdetails)
	res , _ :=json.Marshal(bookdetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}