package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/oanaOM/auction-tracker/auction"
)

var bids []auction.Auction
var sellItems []auction.Item

func main() {

	//Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Failing to open :%v", err)
	}
	//Read the file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	//Initialise a content variable that will hold our file content
	var inputs []string

	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}

	//get the list of all listed
	auction.GetAuctions(inputs)

	// #######################

	var sellItem auction.Item

	fmt.Println("=== Incoming auctions: ")
	for _, a := range auction.Auctions {
		//skip the hearthbeat moments
		fmt.Printf("%+v\n", a)
		if a.UserID != 0 {
			if a.ActionType == "BID" {
				bids = append(bids, a)
			}
			if a.ActionType == "SELL" {
				sellItem.AddItem(a)
				sellItems = append(sellItems, sellItem)
			}

		}
	}

	var WinnerBids []auction.Auction
	var totalBids []int

	//fmt.Println(bids)
	for _, itm := range sellItems {
		WinnerBids = append(WinnerBids, auction.WinnerBid(itm, bids))

		//calculate total bids per item
		var t int
		t = auction.CountItemBids(itm.Name, bids)
		totalBids = append(totalBids, t)

	}

	resultItemsStas := auction.GetItemsStats(WinnerBids, sellItems)

	// #### format the stats for each item as string
	var output string

	for _, result := range resultItemsStas {
		output += fmt.Sprintf("|%v|%v|%v|%v|%v|%v|%v|%v|\n", result.CloseTime, result.ItemName, result.BuyerID, result.Status, result.PaidPrice, result.TotalBids, result.MaxPrice, result.MinPrice)

	}
	//create a new file
	f, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	n2, err := f.WriteString(output)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Succesfully wrote %v bytes to output.txt. \n", n2)
}
