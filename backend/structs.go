package main

// {"id":1,
// "image":"https://groupietrackers.herokuapp.com/api/images/queen.jpeg",
// "name":"Queen",
// "members":["Freddie Mercury","Brian May","John Daecon","Roger Meddows-Taylor","Mike Grose","Barry Mitchell","Doug Fogie"],
// "creationDate":1970,
// "firstAlbum":"14-12-1973",
// "locations":"https://groupietrackers.herokuapp.com/api/locations/1",
// "concertDates":"https://groupietrackers.herokuapp.com/api/dates/1",
// "relations":"https://groupietrackers.herokuapp.com/api/relation/1"},

type MusicBand struct {
	ID               int                 `json:"id"`
	Image            string              `json:"image"`
	Name             string              `json:"name"`
	Members          []string            `json:"members"`
	CreationDate     int                 `json:"creationDate"`
	FirstAlbum       string              `json:"firstAlbum"`
	LocationsDump    string              `json:"locations"`
	ConcertDatesDump string              `json:"concertDates"`
	Relations        string              `json:"relations"`
	Concerts         map[string][]string `json:"concerts"`
}

// Outgoing MusicBandDTO struct
type MusicBandDTO struct {
	ID                    int                 `json:"id"`
	Image                 string              `json:"image"`
	Name                  string              `json:"name"`
	Members               []string            `json:"members"`
	CreationDate          int                 `json:"creationDate"`
	FirstAlbum            string              `json:"firstAlbum"`
	ConcertsLocationDates map[string][]string `json:"concertsLocDates"`
}

// From relations data
// {
//	"id":1,
// 	"datesLocations":{"dunedin-new_zealand":["10-02-2020"],
//					 "georgia-usa":["22-08-2019"],
// 				 	"los_angeles-usa":["20-08-2019"],
// 					"nagoya-japan":["30-01-2019"],
// 					"north_carolina-usa":["23-08-2019"],
// 					"osaka-japan":["28-01-2020"],
//					"penrose-new_zealand":["07-02-2020"],
// 					"saitama-japan":["26-01-2020"]}
// }

type Concert struct {
	ID              int                 `json:"id"`
	LocationDateMap map[string][]string `json:"datesLocations"`
}

