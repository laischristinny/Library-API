package routers

import (
	"LibraryAPI-GitFlow/src/controller"
	"net/http"
)

var BooksRouters = []Router{
	{
		URI:                "/books",
		Method:             http.MethodPost,
		Function:             controller.CreateBook,
	
	},
	{
		URI:                "/books/{bookID}",
		Method:             http.MethodGet,
		Function:             controller.SearchBook,
	
	},
	{
		URI:                "/books",
		Method:             http.MethodGet,
		Function:             controller.GetAllBooks,
	
	},
	{
		URI:                "/books/checkout/{bookID}",
		Method:             http.MethodPatch,
		Function:             controller.CheckoutBook,
		
	},
	{
		URI:                "/books/return/{bookID}",
		Method:             http.MethodPatch,
		Function:             controller.ReturnBook,
		
	},
}