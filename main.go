package main

import (
	// "errors"
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/create-payment-intent" , handleCreatePaymentIntent);
	http.HandleFunc("/health" , handleHealth);
	
	log.Println("The server is starting ...")
	var err error = http.ListenAndServe("localhost:4242" , nil);
	
	if err != nil {
		log.Fatal(err);
	}

}


func handleCreatePaymentIntent( responseWriter http.ResponseWriter , request  *http.Request ){
	
	
	if request.Method != "POST" {
		http.Error(responseWriter , http.StatusText(http.StatusMethodNotAllowed) , http.StatusMethodNotAllowed )
		return
	}

	fmt.Println("The request method was POST and is OKAY !");
	
}


func handleHealth( responseWriter http.ResponseWriter , request *http.Request ){
	
	// fmt.Println("response -> " , &responseWriter , " " , "req -> " , &request);
	
	var resp []byte = []byte("Server is up and running !")

	_ , err :=  responseWriter.Write(resp);

	if err != nil{
		fmt.Println("err -> " , err);
	}


}

