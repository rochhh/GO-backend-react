package main

import (
	// "errors"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"io"
	"github.com/stripe/stripe-go/v75"
	"github.com/stripe/stripe-go/v75/paymentintent"
)

func main() {

	stripe.Key = "<key>"
    // remove the above before pushing to github 

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

	err := json.NewDecoder(request.Body).Decode(&req); // convert the json to Go code , hence .Decode()

	if err != nil {
		// log.Println(err)
		http.Error(responseWriter , err.Error() , http.StatusInternalServerError);
		return 
	}


	params := &stripe.PaymentIntentParams{
		Amount: stripe.Int64(calculateOrderAmount(req.ProductId)),
		Currency: stripe.String(string(stripe.CurrencyUSD)),  // try removing one string as to know more 
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	paymentIntent , err := paymentintent.New(params);   // req to stripe server 

	if err != nil{
		http.Error( responseWriter , err.Error() , http.StatusInternalServerError)
	}

	fmt.Println("paymentIntent.ClientSecret ->" , paymentIntent.ClientSecret)  // this is the secret key which revolves between the FE , BE and stripe server 

	var resp struct {
		ClientSecret string `json:"clientSecret"`
	}

	resp.ClientSecret = paymentIntent.ClientSecret //send the value to the person requesting this (FE)

	var buf bytes.Buffer

	err = json.NewEncoder(&buf).Encode(resp)

	if err != nil {
		http.Error(responseWriter , err.Error() , http.StatusInternalServerError);

	}

	responseWriter.Header().Set("Content-Type" , "application/json");

	_ , err = io.Copy(responseWriter , &buf); // the user will receive this responseWriter which is copied in buffer


	if err != nil {
		fmt.Println(err);
	}


}


func handleHealth( responseWriter http.ResponseWriter , request *http.Request ){
	
	// fmt.Println("response -> " , &responseWriter , " " , "req -> " , &request);
	
	var resp []byte = []byte("Server is up and running !")

	_ , err :=  responseWriter.Write(resp);

	if err != nil{
		fmt.Println("err -> " , err);
	}


}

func calculateOrderAmount(productId string) int64 {
	switch productId{
	case "Forever Pants":
		return 26000

	case "Foreever Shirt":
			return 15500

	case "Forever Shorts":
return 30000
	
	}

	return 0
}