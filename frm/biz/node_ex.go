package biz

import "fmt"

func GetKrakenKey(nodeCode string) string {
	return fmt.Sprintf("KRAKEN_%v", nodeCode)
}

func GetHallKey(nodeCode string) string {
	return fmt.Sprintf("HALL_%v", nodeCode)
}
