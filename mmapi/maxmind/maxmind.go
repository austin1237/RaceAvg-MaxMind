package maxmind

import (
	"errors"
	"fmt"
	"io/ioutil"
	"sort"

	"github.com/jszwec/csvutil"
)

type Block struct {
	StartIPNum int `csv:"startIpNum"`
	EndIPNum   int `csv:"endIpNum"`
	LocID      int `csv:"locId"`
}

type Location struct {
	LocID      int     `csv:"locId" json:"locId"`
	Country    string  `csv:"country" json:"country"`
	Region     string  `csv:"region" json:"region"`
	City       string  `csv:"city" json:"city"`
	PostalCode string  `csv:"postalCode" json:"postalCode"`
	Latitude   float64 `csv:"latitude" json:"latitude"`
	Longitude  float64 `csv:"longitude" json:"longitude"`
	MetroCode  string  `csv:"metroCode" json:"metroCode"`
	AreaCode   string  `csv:"areaCode" json:"areaCode"`
}

type MaxMindMemory struct {
	BlockMap       map[int]Block
	SortedStartIps []int
	LocationMap    map[int]Location
}

var mmMemory MaxMindMemory

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func LoadCSVs() {
	var blocks []Block
	var locations []Location

	fmt.Println("setting up db")

	bytes, err := ioutil.ReadFile("GeoLiteCity-Blocks.csv")
	check(err)
	if err := csvutil.Unmarshal(bytes, &blocks); err != nil {
		fmt.Println("error:", err)
	}

	bytes, err = ioutil.ReadFile("GeoLiteCity-Location.csv")
	check(err)
	if err := csvutil.Unmarshal(bytes, &locations); err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("blocks finished loading %+v items", len(blocks))
	fmt.Printf("locations finished loading %+v items", len(locations))
	sortedStartIP, blockMap := sortByStartingIP(blocks)
	locationMap := sortByLocID(locations)
	mmMemory.BlockMap = blockMap
	mmMemory.LocationMap = locationMap
	mmMemory.SortedStartIps = sortedStartIP
}

func sortByLocID(locations []Location) map[int]Location {
	locationMap := make(map[int]Location)
	for _, element := range locations {
		locationMap[element.LocID] = element
	}
	return locationMap
}

func sortByStartingIP(blocks []Block) ([]int, map[int]Block) {

	blockMap := make(map[int]Block)
	startIPs := make([]int, 0)
	for _, element := range blocks {
		blockMap[element.StartIPNum] = element
		startIPs = append(startIPs, element.StartIPNum)
	}
	sort.Ints(startIPs)

	fmt.Printf("number of unquie elements %v \n", len(blockMap))
	return startIPs, blockMap
}

// Implementing the between query here
// https://dev.maxmind.com/geoip/legacy/csv/#SQL_Queries
func QueryForLocation(ipInteger int) (Location, error) {
	endPosition := sort.Search(len(mmMemory.SortedStartIps), func(i int) bool { return mmMemory.SortedStartIps[i] >= ipInteger })

	for i := 0; i < endPosition; i++ {
		startIndex := mmMemory.SortedStartIps[i]
		if mmMemory.BlockMap[startIndex].EndIPNum >= ipInteger {
			locID := mmMemory.BlockMap[startIndex].LocID
			location := mmMemory.LocationMap[locID]
			fmt.Printf("foundTheLocationId here %v \n", i)
			fmt.Printf("locID is %v \n", mmMemory.BlockMap[startIndex].LocID)
			fmt.Printf("location is %v \n", location)
			return location, nil
		}

	}
	err := errors.New("No location found")
	return Location{}, err
}
