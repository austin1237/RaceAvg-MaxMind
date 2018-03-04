package maxmind

import (
	"errors"
	"fmt"
	"io/ioutil"
	"sort"

	"github.com/jszwec/csvutil"
)

var mmDB MaxMindDB

// If the csvs can't be loaded to build the DB hard crash
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func MustLoadCSVs() {
	var blocks []Block
	var locations []DBLocation

	fmt.Println("setting up db")

	bytes, err := ioutil.ReadFile("GeoLiteCity-Blocks.csv")
	check(err)
	err = csvutil.Unmarshal(bytes, &blocks)
	check(err)

	bytes, err = ioutil.ReadFile("GeoLiteCity-Location.csv")
	check(err)
	err = csvutil.Unmarshal(bytes, &locations)
	check(err)
	sortedStartIP, blockMap := sortByStartingIP(blocks)
	locationMap := sortByLocID(locations)
	mmDB.BlockMap = blockMap
	mmDB.LocationMap = locationMap
	mmDB.SortedStartIps = sortedStartIP
	fmt.Println("db loaded")
}

func sortByLocID(locations []DBLocation) map[int]DBLocation {
	locationMap := make(map[int]DBLocation)
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

	return startIPs, blockMap
}

// Implementing the between query found here https://dev.maxmind.com/geoip/legacy/csv/#SQL_Queries
// Find the location where ipInteger >= startIp && ipInteger <= endIpNum
func QueryForLocation(ipInteger int) (DBLocation, error) {
	endPosition := sort.Search(len(mmDB.SortedStartIps), func(i int) bool { return mmDB.SortedStartIps[i] >= ipInteger })

	for i := 0; i < endPosition; i++ {
		startIndex := mmDB.SortedStartIps[i]
		if mmDB.BlockMap[startIndex].EndIPNum >= ipInteger {
			locID := mmDB.BlockMap[startIndex].LocID
			location := mmDB.LocationMap[locID]
			return location, nil
		}

	}
	err := errors.New("No location found")
	return DBLocation{}, err
}
