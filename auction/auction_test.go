package auction

import (
	"fmt"
	"testing"
)

var WinnerBids []Auction

func TestWinnerBid(t *testing.T) {
	var WinnerBids []Auction

	for i, item := range sellItemsTest {

		WinnerBids = append(WinnerBids, WinnerBid(item, bidsTest))

		// check received price
		if WinnerBids[i].Status != itemStatsTests[i].status {
			t.Fatalf("FAIL: item %v should have status %v. Instead we got the status %v", WinnerBids[i].Item, itemStatsTests[i].status, WinnerBids[i].Status)
		}

		//check received price
		if WinnerBids[i].Price != itemStatsTests[i].price {
			t.Fatalf("FAIL: item %v should cost %v. Instead we got the price of %v", WinnerBids[i].Item, itemStatsTests[i].price, WinnerBids[i].Price)
		}

		//check buyer ID
		if WinnerBids[i].UserID != itemStatsTests[i].userID {
			t.Fatalf("FAIL: item %v was bought by %v. Instead we got the buyer ID of %v", WinnerBids[i].Item, itemStatsTests[i].userID, WinnerBids[i].UserID)
		}

	}

}

func TestAddItem(t *testing.T) {

	for i, a := range auctions {
		if a.ActionType == "SELL" {
			a.Item = sellItemsTest[i].Name
			a.Price = sellItemsTest[i].Price
			a.UserID = sellItemsTest[i].UserID
			a.CloseTime = sellItemsTest[i].ExpirationTime
		}
	}
}

func TestGetItemsStats(t *testing.T) {

	resultItemsStas := GetItemsStats(winnersItemsTest, sellItemsTest)

	fmt.Println(resultItemsStas)

	for i, result := range resultItemsStas {

		if result.closeTime != itemStatsTests[i].closeTime {
			t.Errorf("Closing price expected be %v. We got %v .", itemStatsTests[i].closeTime, result.closeTime)
		}
		if result.itemName != itemStatsTests[i].name {
			t.Errorf("Item name expected be %v. We got %v .", itemStatsTests[i].name, result.itemName)
		}
		if result.buyerID != itemStatsTests[i].userID {
			t.Errorf("Buyer id expected be %v. We got %v .", itemStatsTests[i].userID, result.buyerID)
		}
		if result.status != itemStatsTests[i].status {
			t.Errorf("Item status status expected be %v. We got %v .", itemStatsTests[i].status, result.status)
		}

	}

}
