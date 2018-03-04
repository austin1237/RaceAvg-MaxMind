package maxmind

type Block struct {
	StartIPNum int `csv:"startIpNum"`
	EndIPNum   int `csv:"endIpNum"`
	LocID      int `csv:"locId"`
}

type DBLocation struct {
	LocID      int     `csv:"locId"`
	Country    string  `csv:"country"`
	Region     string  `csv:"region"`
	City       string  `csv:"city"`
	PostalCode string  `csv:"postalCode"`
	Latitude   float64 `csv:"latitude"`
	Longitude  float64 `csv:"longitude"`
	MetroCode  string  `csv:"metroCode"`
	AreaCode   string  `csv:"areaCode"`
}

type MaxMindDB struct {
	BlockMap       map[int]Block
	SortedStartIps []int
	LocationMap    map[int]DBLocation
}

type ClientLocation struct {
	LocationID int     `json:"locationId"`
	Country    string  `json:"country"`
	Region     string  `json:"region"`
	City       string  `json:"city"`
	PostalCode string  `json:"postalCode"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	MetroCode  string  `json:"metroCode"`
	AreaCode   string  `json:"areaCode"`
}
