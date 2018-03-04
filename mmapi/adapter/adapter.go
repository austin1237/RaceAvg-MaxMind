package adapter

import (
	"errors"
	"strconv"
	"strings"

	"github.com/user/mmapi/maxmind"
)

func ClientIPToDb(address string) (int, error) {
	splitAddress := strings.Split(address, ".")
	splitNumbers := []int{}
	for _, element := range splitAddress {
		num, err := strconv.Atoi(element)
		if err != nil {
			err = errors.New("Invalid IP address " + element)
			return 0, err
		}
		splitNumbers = append(splitNumbers, num)
	}
	// equation found here https://dev.maxmind.com/geoip/legacy/csv/#Integer_IPv4_Representation
	integerIP := (16777216 * splitNumbers[0]) + (65536 * splitNumbers[1]) + (256 * splitNumbers[2]) + splitNumbers[3]
	return integerIP, nil
}

func DbLocationToClient(dbLocation maxmind.DBLocation) maxmind.ClientLocation {
	clientLoc := maxmind.ClientLocation{
		LocationID: dbLocation.LocID,
		Country:    dbLocation.Country,
		Region:     dbLocation.Region,
		City:       dbLocation.City,
		Latitude:   dbLocation.Latitude,
		Longitude:  dbLocation.Longitude,
		MetroCode:  dbLocation.MetroCode,
	}
	return clientLoc
}
