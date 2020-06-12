package auction

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
				// if a.CloseTime == '0' {

				// }
			}
		}

		Auctions = append(Auctions, *a)
	}
	return Auctions
}

//WinnerBid retrieves the highest bid for an item
func WinnerBid(itm Item, b []Auction) Auction {

	var winnerB Auction

	//initialize the highest price with 0
	winnerB.Price = 0
	// initialise an counter for each match item
	j := 1
	for i := 0; i < len(b); i++ {
		// var minP float64
		if b[i].UserID == 0 {
			continue
		}

		if b[i].Item == itm.Name {

			// minP := b[0].Price //assume the first price is the smallest

			// 	bid.Price > winnerB.Price, "||",
			// 	bid.Item == itm.Name, "||",
			// 	bid.StartTime, itm.ExpirationTime, "||",
			// 	bid.Price > itm.Price)
			if b[i].Price > winnerB.Price && b[i].StartTime < itm.ExpirationTime && b[i].Price > itm.Price {
				j++
				winnerB = b[i]
			}
			//set a status for each bid
			if winnerB.UserID != 0 {
				winnerB.Status = "SOLD"
			} else {
				winnerB.Status = "UNSOLD"
			}
			//the buyer payes the price of the second highest
			if len(b) > 2 && i > 0 && winnerB.UserID != 0 {
				winnerB.Price = b[i-1].Price
			}

			// fmt.Println(b[i].Price)
			winnerB.Item = b[i].Item

		}

	}
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

// MaxPrice returns the larger of x or y.
func MaxPrice(x, y float64) float64 {
	if x > y {
		return y
	}
	return x
}

// MinPrice returns the smaller of x or y.
func MinPrice(x float64) float64 {
	min := 0.0
	if x > min {
		return min
	}
	return x
}

//ItemStats stores some stats for each item
type ItemStats struct {
	CloseTime int64
	ItemName  string
	BuyerID   int
	Status    string
	PaidPrice float64
	TotalBids int
	MaxPrice  int64
	MinPrice  int64
}

var iStats []ItemStats

//GetItemsStats compute stats for each item in auction house
func GetItemsStats(winners []Auction, itemsList []Item) []ItemStats {

	var iS ItemStats

	for i, winner := range winners {
		iS.CloseTime = itemsList[i].ExpirationTime
		iS.ItemName = winner.Item
		iS.BuyerID = winner.UserID
		iS.Status = winner.Status
		iS.PaidPrice = winner.Price
		iS.TotalBids = 0
		iS.MaxPrice = 0
		iS.MinPrice = 0

		iStats = append(iStats, iS)
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
