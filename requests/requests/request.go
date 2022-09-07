package requests

import (
	"encoding/json"            // to decode our json
	"fmt"                      // to log our erros or status
	"github.com/gin-gonic/gin" // gin framework
	"net/http"                 // to handle our api endpoint
)

// practice route
func FirstRequest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello from requests",
	})
}

/*
The pokemon api retruns a an arrary of results
with indivudual pokemon objects
*/
// this takes care of indivudual pokemon objects
type Pokemon struct {
	Name string `json:name`
	Url  string `json:url`
}

// this takes car of the result array that holds our data
type Results struct {
	Results []Pokemon `json:results`
}

/*
This function takes an api endpoint and returns a
Json object containing our pokemons
also there a set Cookie in there as well se if that worked
*/
func SecondRequest(c *gin.Context) {
	// we declare of Result struct here
	var jsonResult Results
	// api endpoint
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon?offset=10&limit=10")
	if err != nil {
		// if error retrun an http error
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error happened while reading url",
		})
	}
	defer resp.Body.Close()
	// setting up cookie
	cookie, err := c.Cookie("gin_cookie")
	if err != nil {
		cookie = "NotSet"
		c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
	}
	// print cookie status
	fmt.Printf("Cookie value %s \n", cookie)
	// decode our json reponse into are structs
	json.NewDecoder(resp.Body).Decode(&jsonResult)
	// return our json value
	c.JSON(http.StatusOK, jsonResult)
}

func ThirdRequest(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "main wesbite",
	})
}
