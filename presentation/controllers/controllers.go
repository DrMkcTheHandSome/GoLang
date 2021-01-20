package controllers

import(
	"github.com/gorilla/mux"
	"log"
	"net/http"
	services "services"
	httpSwagger "github.com/swaggo/http-swagger"
	)


	func HandleRequests() {
		InitializeRoutes()
	 }
	 
	 func InitializeRoutes(){
		InitRoutesByGorillaMux()
	 }
	 
	 func InitRoutesByGorillaMux(){
		myRouter := mux.NewRouter().StrictSlash(true)
		myRouter.HandleFunc("/", services.HomePage).Methods("GET")
		myRouter.HandleFunc("/migration", services.CreateDatabaseSchema).Methods("POST")
		myRouter.HandleFunc("/products", services.ReturnAllProducts).Methods("GET")
		myRouter.HandleFunc("/product", services.CreateNewProduct).Methods("POST")
		myRouter.HandleFunc("/product/{id}", services.UpdateProduct).Methods("PUT")
		myRouter.HandleFunc("/product/{id}", services.DeleteProduct).Methods("DELETE")
		myRouter.HandleFunc("/product/{id}",services.ReturnSingleProduct).Methods("GET")
		myRouter.HandleFunc("/user", services.CreateNewUser).Methods("POST")
		myRouter.HandleFunc("/users", services.ReturnAllUsers).Methods("GET")
		myRouter.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
		log.Fatal(http.ListenAndServe(":9000", myRouter))

		/*
		myRouter.HandleFunc("/user/loginViaGoogle", loginUserViaGoogle).Methods("GET")
		myRouter.HandleFunc("/user/login", loginUserWithPassword).Methods("POST")
		myRouter.HandleFunc("/googlecallback", handleGoogleCallback).Methods("GET")
		*/
	 }
	 