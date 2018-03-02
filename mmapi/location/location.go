package location

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/user/mmapi/maxmind"
)

func GetLocationByIP(ipAddress string) (maxmind.Location, error) {
	ipInteger, err := convertIPToInt(ipAddress)
	if err != nil {
		return maxmind.Location{}, err
	}
	fmt.Printf("%v is the converted ip", ipInteger)
	location, err := maxmind.QueryForLocation(ipInteger)
	if err != nil {
		return maxmind.Location{}, err
	}
	return location, nil
}

func convertIPToInt(address string) (int, error) {
	fmt.Println("Address is " + address)
	address = strings.Replace(address, "\"", "", -1)
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
