package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// pokemon represents data about a record pokemon.
type pokemon struct {
	ID     string  `json:"id"`
	Name string  `json:"name"`
	Description string `json:"description"`
}

// pokemons slice to seed record pokemon data.
var pokemons = []pokemon{
	{ID: "1", Name: "Pikachu", Description: "Zap zap, bap bap"},
	{ID: "2", Name: "Charizard", Description: "Big and fiery"},
	{ID: "2", Name: "Snorlax", Description: "Lazy so and so"},
}

func main() {
	router := gin.Default()
	router.GET("/pokemon", getPokemon)
	router.GET("/pokemon/:id", getPokemonByID)
	router.POST("/pokemon", postPokemon)

	err := router.Run("localhost:8081")
	if err != nil {
		return
	}
}

// getPokemon responds with the list of all pokemon as JSON.
func getPokemon(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, pokemons)
}

// postPokemon adds an pokemon from JSON received in the request body.
func postPokemon(c *gin.Context) {
	var newPokemon pokemon

	// Call BindJSON to bind the received JSON to
	// newPokemon.
	if err := c.BindJSON(&newPokemon); err != nil {
		return
	}

	// Add the new pokemon to the slice.
	pokemons = append(pokemons, newPokemon)
	c.IndentedJSON(http.StatusCreated, newPokemon)
}

// getPokemonByID locates the pokemon whose ID value matches the id
// parameter sent by the client, then returns that pokemon as a response.
func getPokemonByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of pokemons, looking for
	// an pokemon whose ID value matches the parameter.
	for _, a := range pokemons {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "pokemon not found"})
}
