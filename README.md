
Auction House
===========================================

Please see full description in instructions.txt file 

===========================================

## How to run it 

1. Clone this repo 
2. Run
   `go run main.go`
3. An output.txt file should be created under the root directory.
4. The values in output.txt file are the result of the auctions read from input.txt file.


### Example:

Input:

10|1|SELL|toaster_1|10.00|20
12|8|BID|toaster_1|7.50
13|5|BID|toaster_1|12.50
15|8|SELL|tv_1|250.00|20
16
17|8|BID|toaster_1|20.00
18|1|BID|tv_1|150.00
19|3|BID|tv_1|200.00
20
21|3|BID|tv_1|300.00

Output:

20|toaster_1|8|SOLD|12.50|3|20.00|7.50
20|tv_1||UNSOLD|0|2|200.00|150.00
