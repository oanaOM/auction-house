package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/oanaOM/auction-tracker/house"
)

//bids stores all the auction of type bid
var bids []house.Auction

//store all items that are for sold
var sellItems []house.Item

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

	//get the list of all inputs
	house.GetAuctions(inputs)

	//initialise an item that will be for sell
	var sellItem house.Item

	fmt.Println("=== Incoming auctions: ")
	for _, a := range house.Auctions {
		//skip the hearthbeat moments
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
	//WinnerBids stores the status for each item at the end of the auction
	var results []house.Auction
	var totalBids []int

	//fmt.Println(bids)
	for _, itm := range sellItems {
		results = append(results, house.WinnerBid(itm, bids))

		//calculate total bids per item
		t := house.CountItemBids(itm.Name, bids)
		totalBids = append(totalBids, t)

	}

	finalStats := house.GetStats(results, sellItems)

	// #### format the stats for each item as string
	var output string

	for _, result := range finalStats {
		output += fmt.Sprintf("%v|%v|%v|%v|%.2f|%d|%.2f|%.2f\n", result.CloseTime, result.Name, result.BuyerID, result.Status, result.PaidPrice, result.TotalBids, result.MaxPrice, result.MinPrice)

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
