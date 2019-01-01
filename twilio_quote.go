package main  

import(
    "fmt"
    "net/http"
    "io/ioutil"
    "log"
    "encoding/json"
    "os"
    "math/rand"
    "time"  
    "strings"
    "net/url"    
)

type quoteObject struct{
	Quote string
}

func main(){
    //from your Twilio account, enter your accountSid and authToken
    accountSid := "ACXXXX"
    authToken := "XXXXXX"
    twilioUrl := "https://api.twilio.com/2010-04-01/Accounts/" +  accountSid  + "/Messages.json"
    client := &http.Client{}
    request, _ := http.NewRequest("GET", "https://andruxnet-random-famous-quotes.p.rapidapi.com/?count=23&cat=famous", nil)
    //register with RAPIDAPI and get your key and insert it in YOUR_RAPIDAPI_KEY
    request.Header.Set("X-RapidAPI-Key", "YOUR_RAPIDAPI_KEY")
    response, _ := client.Do(request)
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatalf("ERROR: %s", err)
    }
    var store []quoteObject
    getQuote := json.Unmarshal(body, &store)
    if getQuote != nil {
 	fmt.Println("error", getQuote)
 	os.Exit(1)
      } 
    //this is to randomly generate the quote
    rand.Seed(time.Now().Unix())

    //this is essentially the quote from the RapidApi 
    fmt.Println(store[rand.Intn(len(store))].Quote)
    msgData := url.Values{}
    //replace it with the phone number you intend to send to, format: "+1408xxxxxx" 
    msgData.Set("To", "+1408XXXXXXX")
    //replace it with the phone number you register with your Twilio Account, format: "+1408xxxxxx" 
    msgData.Set("From", "+1408XXXXXXX")
    msgData.Set("Body", "Quote of the day: " + store[rand.Intn(len(store))].Quote)
    msgDataReader := *strings.NewReader(msgData.Encode())
    client1 := &http.Client{}
    req, _ := http.NewRequest("POST", twilioUrl, &msgDataReader)
    req.SetBasicAuth(accountSid,authToken)
    req.Header.Add("Accept", "application/json")
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    resp, _ := client1.Do(req)
    if (resp.StatusCode >= 200 && resp.StatusCode < 300) {
	var data map[string]interface{}
  	decoder := json.NewDecoder(resp.Body)
  	err := decoder.Decode(&data)
  	if (err == nil) {
  		fmt.Println(data["sid"])
  		}	
	}   
 }
 
