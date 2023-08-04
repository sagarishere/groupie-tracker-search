package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

var MusicBandDTOs []MusicBandDTO

func main() {

	var logFile *os.File

	// create log file
	logFile, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// setup log to print the line number as well as to output to file and console both
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(io.MultiWriter(os.Stdout, logFile))

	fetchJsonData()
	log.Println("Starting the backend server")

	// Create a new ServeMux
	nakedMux := http.NewServeMux()

	// Attach the middleware
	mux := middleware(nakedMux)

	// create routes
	nakedMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is the backend, the frontend is at http://localhost:3000"))
	})

	// pass processed data on the route /data
	nakedMux.HandleFunc("/data", sendProcessedData)

	// listen and serve on port 8080
	http.ListenAndServe(":8080", mux)

}

// sendProcessedData sends the processed data to the frontend
// it is called when the frontend makes a GET request to /data
func sendProcessedData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log.Println("Sending processed data to frontend")

	data, err := json.Marshal(MusicBandDTOs)
	if err != nil {
		log.Println(err)
		return
	}

	w.Write(data)
}

// fetchJsonData fetches the json data from https://groupietrackers.herokuapp.com/api/artists
// and stores it in MusicBandDTOs
func fetchJsonData() {
	// Make HTTP GET request
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Println(err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return
	}

	// Unmarshal the JSON data
	var MusicBands []MusicBand
	err = json.Unmarshal(body, &MusicBands)
	if err != nil {
		log.Println(err)
		return
	}

	// fetch relations data for Concerts
	for i, artist := range MusicBands {
		// Make HTTP GET request
		response, err := http.Get(artist.Relations)
		if err != nil {
			log.Println(err)
			return
		}
		defer response.Body.Close()

		// Read the response body
		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Println(err)
			return
		}

		// Unmarshal the JSON data
		var thisConcert Concert
		err = json.Unmarshal(body, &thisConcert)
		if err != nil {
			log.Println(err)
			return
		}

		// add Concerts to MusicBands
		MusicBands[i].Concerts = thisConcert.LocationDateMap
	}

	// convert MusicBands to MusicBandDTOs

	for _, artist := range MusicBands {
		var thisMusicBandDTO MusicBandDTO
		thisMusicBandDTO.ID = artist.ID
		thisMusicBandDTO.Image = artist.Image
		thisMusicBandDTO.Name = artist.Name
		thisMusicBandDTO.Members = artist.Members
		thisMusicBandDTO.CreationDate = artist.CreationDate
		thisMusicBandDTO.FirstAlbum = artist.FirstAlbum
		thisMusicBandDTO.ConcertsLocationDates = artist.Concerts
		MusicBandDTOs = append(MusicBandDTOs, thisMusicBandDTO)
	}
}
