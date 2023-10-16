package main

import (
	// "errors"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
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

	var req struct {
		ProductId string `json:"product_id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Address1  string `json:"address_1"`
		Address2  string `json:"address_2"`
		City      string `json:"city"`
		State     string `json:"state"`
		Zip       string `json:"zip"`
		Country   string `json:"country"`
	}

	err := json.NewDecoder(request.Body).Decode(&req);

	if err != nil {
		// log.Println(err)
		http.Error(responseWriter , err.Error() , http.StatusInternalServerError);
		return 
	}

	
	fmt.Println("the city ->" , req.City);

}


func handleHealth( responseWriter http.ResponseWriter , request *http.Request ){
	
	// fmt.Println("response -> " , &responseWriter , " " , "req -> " , &request);
	
	var resp []byte = []byte("Server is up and running !")

	_ , err :=  responseWriter.Write(resp);

	if err != nil{
		fmt.Println("err -> " , err);
	}


}

