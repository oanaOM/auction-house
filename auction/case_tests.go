package auction

var auctions = []Auction{
	// Auction{StartTime: 10, UserID: 1, ActionType: "SELL", Item: "toaster_1", Price: 10.00, CloseTime: 20.0},
	Auction{StartTime: 12, UserID: 8, ActionType: "BID", Item: "toaster_1", Price: 7.50},
	Auction{StartTime: 13, UserID: 5, ActionType: "BID", Item: "toaster_1", Price: 12.50},
	// Auction{StartTime: 15, UserID: 8, ActionType: "SELL", Item: "tv_1", Price: 250.00, CloseTime: 22.0},
	// Auction{StartTime: 16},
	Auction{StartTime: 17, UserID: 8, ActionType: "BID", Item: "toaster_1", Price: 20.00},
	Auction{StartTime: 18, UserID: 1, ActionType: "BID", Item: "tv_1", Price: 150.00},
	Auction{StartTime: 19, UserID: 3, ActionType: "BID", Item: "tv_1", Price: 200.00},
	// Auction{StartTime: 20},
	Auction{StartTime: 21, UserID: 3, ActionType: "BID", Item: "tv_1", Price: 300.00},
}

var bidsTest = []Auction{
	Auction{StartTime: 12, UserID: 8, ActionType: "BID", Item: "toaster_1", Price: 7.50},
	Auction{StartTime: 13, UserID: 5, ActionType: "BID", Item: "toaster_1", Price: 12.50},
	Auction{StartTime: 17, UserID: 8, ActionType: "BID", Item: "toaster_1", Price: 20.00},
	Auction{StartTime: 18, UserID: 1, ActionType: "BID", Item: "tv_1", Price: 150.00},
	Auction{StartTime: 19, UserID: 3, ActionType: "BID", Item: "tv_1", Price: 200.00},
	Auction{StartTime: 21, UserID: 3, ActionType: "BID", Item: "tv_1", Price: 300.00},
}

var sellItemsTest = []Item{
	Item{Name: "toaster_1", Price: 10.00, UserID: 8, ExpirationTime: 20},
	Item{Name: "tv_1", Price: 250.00, UserID: 8, ExpirationTime: 20},
}

var winnersItemsTest = []Auction{
	Auction{StartTime: 17, UserID: 8, ActionType: "BID", Item: "toaster_1", Price: 12.5, CloseTime: 0, Status: "SOLD"},
	Auction{StartTime: 0, UserID: 0, ActionType: "BID", Item: "tv_1", Price: 0, CloseTime: 0, Status: "UNSOLD"},
}

var itemStatsTests = []struct {
	startTime    int64
	closeTime    int64
	name         string
	userID       int
	status       string
	price        float64
	numberBids   int
	highestPrice float64
	lowestPrice  float64
}{
	{
		startTime:    10,
		closeTime:    20,
		name:         "toaster_1",
		userID:       8,
		status:       "SOLD",
		price:        12.50,
		numberBids:   3,
		highestPrice: 20.00,
		lowestPrice:  7.50,
	},
	{
		startTime:    15,
		closeTime:    20,
		name:         "tv_1",
		userID:       0,
		status:       "UNSOLD",
		price:        0,
		numberBids:   2,
		highestPrice: 200.00,
		lowestPrice:  150.00,
	},
}
