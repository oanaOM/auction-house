// Package house provides functions for retrieving the inputed auctions,
// the winning bids and some stats
package house

import (
	"log"
	"strconv"
	"strings"
)

//Auction stores details of the incoming auction
type Auction struct {
	StartTime  int64
	UserID     int
	ActionType string
	Item       string
	Price      float64 //bid price for BIDS and reservePrice for sells
	CloseTime  int64
	Status     string
}

//Item stores esential details about an item
type Item struct {
	Name           string
	Price          float64
	UserID         int
	ExpirationTime int64
}

//Auctions store all auctions from the inpus file
var Auctions []Auction

//GetAuctions retrieves a list with all auctios
func GetAuctions(input []string) []Auction {
	var err error

	for _, line := range input {

		a := new(Auction)
		Auction := strings.Split(line, "|")

		for i, v := range Auction {
			switch i {
			case 0:
				a.StartTime, err = strconv.ParseInt(v, 10, 64)
				if err != nil {
					log.Fatal(err)
				}
			case 1:
				a.UserID, err = strconv.Atoi(v)
				if err != nil {
					log.Fatal(err)
				}
			case 2:
				a.ActionType = v
			case 3:
				a.Item = v
			case 4:
				a.Price, err = strconv.ParseFloat(v, 64)
				if err != nil {
					log.Fatal(err)
				}
			case 5:
				a.CloseTime, err = strconv.ParseInt(v, 10, 64)
				if err != nil {
					log.Fatal(err)
				}
			}
		}

		Auctions = append(Auctions, *a)
	}
	return Auctions
}

//ItemStats stores some stats for each item
type ItemStats struct {
	CloseTime int64
	Name      string
	BuyerID   int
	Status    string
	PaidPrice float64
	TotalBids int
	MaxPrice  float64
	MinPrice  float64
}

//iStats initializing stats for each item
var iStats []ItemStats

//initialize the max and min price of a valid bid
var maxP, minP float64

//outcome stores some stats for each item
var outcome []ItemStats

//WinnerBid retrieves the highest winning bid for an item
func WinnerBid(itm Item, bid []Auction) Auction {

	var winnerB Auction
	var tempItem ItemStats
	maxP = 0.0
	//initiate an array with all my prices
	var price []float64
	//initialize the highest price with 0
	winnerB.Price = 0
	// initialise an counter for each match item
	j := 1
	var z int

	for i := 0; i < len(bid); i++ {
		//skip the heartbeat messages
		if bid[i].UserID == 0 {
			continue
		}

		if bid[i].Item == itm.Name {

			if bid[i].StartTime < itm.ExpirationTime {
				//incrementor to count total valid bids for each item
				z++
				tempItem.Name = bid[i].Item
				price = append(price, bid[i].Price)
				if bid[i].Price > itm.Price && bid[i].Price > winnerB.Price {
					j++
					winnerB = bid[i]
				}
				tempItem.TotalBids = z

				if bid[i].Price > maxP {
					maxP = bid[i].Price
				}

				minP = price[0]
				if bid[i].Price < minP {
					minP = bid[i].Price

				}
				tempItem.MaxPrice = maxP
				tempItem.MinPrice = minP
			}

			//set a status for each bid
			if winnerB.UserID != 0 {
				winnerB.Status = "SOLD"
			} else {
				winnerB.Status = "UNSOLD"
			}

			//the buyer payes the price of the second highest
			if len(bid) > 2 && i > 0 && winnerB.UserID != 0 {
				winnerB.Price = bid[i-1].Price
			}

			winnerB.Item = bid[i].Item

		}

	}

	outcome = append(outcome, tempItem)

	return winnerB
}

//AddItem populates a new user
func (itm *Item) AddItem(new Auction) Item {
	itm.Name = new.Item
	itm.Price = new.Price
	itm.UserID = new.UserID
	itm.ExpirationTime = new.CloseTime

	return *itm
}

//GetStats computes stats for each item in auction house
func GetStats(winners []Auction, itemsList []Item) []ItemStats {

	var itmS ItemStats

	for i, winner := range winners {
		itmS.CloseTime = itemsList[i].ExpirationTime
		itmS.Name = winner.Item
		itmS.BuyerID = winner.UserID
		itmS.Status = winner.Status
		itmS.PaidPrice = winner.Price
		if outcome[i].Name == winner.Item {
			itmS.TotalBids = outcome[i].TotalBids
			itmS.MaxPrice = outcome[i].MaxPrice
			itmS.MinPrice = outcome[i].MinPrice
		}

		iStats = append(iStats, itmS)
	}
	return iStats
}

//CountItemBids how many bids where made for an item
func CountItemBids(name string, bids []Auction) int {
	var counter int
	for _, bid := range bids {
		if name == bid.Item {
			counter++
		}
	}
	return counter
}
