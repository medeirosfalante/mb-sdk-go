package mercadobitcoin

import (
	"fmt"
	"strconv"
)

// Public struct
type Public struct {
	client *APIClient
	Crypto string
}

//Ticker - return a ticket exchange
type Ticker struct {
	High float64 `json:"high"`
	Low  float64 `json:"low"`
	Vol  float64 `json:"vol"`
	Last float64 `json:"last"`
	Buy  float64 `json:"buy"`
	Sell float64 `json:"sell"`
	Date int64   `json:"date"`
}

//TraderPagination - pagination in trade
type TraderPagination struct {
	CurrentPage    int32 `json:"current_page"`
	PageSize       int32 `json:"page_size"`
	RegistersCount int32 `json:"registers_count"`
	TotalPages     int32 `json:"total_pages"`
}

//Trader - struct for trade
type Trader struct {
	Type             string  `json:"type"`
	Amount           float64 `json:"amount"`
	UnitPrice        float32 `json:"unit_price"`
	ActiveOrderCode  string  `json:"active_order_code"`
	PassiveOrderCode string  `json:"passive_order_code"`
	Date             string  `json:"date"`
}

//Traders - Trader response
type Traders struct {
	Pagination TraderPagination `json:"pagination"`
	Trades     []Trader         `json:"trades"`
}

//TradeQuery - trade query
type TradeQuery struct {
	PageSize    int32 `json:"page_size"`
	CurrentPage int32 `json:"current_page"`
}

type Book struct {
	Asks []BookItem `json:"asks"`
	Bids []BookItem `json:"bids"`
}

type TradeResponse []TradeItem

type TradeItem struct {
	Date   int     `json:"date"`
	Price  float64 `json:"price"`
	Amount float64 `json:"amount"`
	Tid    int     `json:"tid"`
	Type   string  `json:"type"`
}

type BookItem []float64

func (c BookItem) Price() float64 {
	return c[0]
}
func (c BookItem) Amount() float64 {
	return c[1]
}

//Public - Create a new instance struct
func (c *APIClient) Public() *Public {
	return &Public{client: c}
}

// Ticker in exchange
func (p Public) Ticker(currency string) (*Ticker, *Error, error) {
	var responseData map[string]interface{}
	err, errAPI := p.client.Request("GET", fmt.Sprintf("/api/%s/ticker", currency), nil, nil, &responseData)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	ticker := responseData["ticker"].(map[string]interface{})
	response := &Ticker{}
	if response.High, err = strconv.ParseFloat(ticker["high"].(string), 64); err != nil {
		return nil, nil, fmt.Errorf("err:  %s", err)
	}
	if response.Low, err = strconv.ParseFloat(ticker["low"].(string), 64); err != nil {
		return nil, nil, err
	}
	if response.Vol, err = strconv.ParseFloat(ticker["vol"].(string), 64); err != nil {
		return nil, nil, err
	}
	if response.Last, err = strconv.ParseFloat(ticker["last"].(string), 64); err != nil {
		return nil, nil, err
	}
	if response.Buy, err = strconv.ParseFloat(ticker["buy"].(string), 64); err != nil {
		return nil, nil, err
	}
	if response.Sell, err = strconv.ParseFloat(ticker["sell"].(string), 64); err != nil {
		return nil, nil, err
	}
	response.Date = int64(ticker["date"].(float64))

	return response, nil, nil
}

// OrderBook - OrderBook in exchange
func (p Public) OrderBook(currency string) (*Book, *Error, error) {
	var response *Book
	err, errAPI := p.client.Request("GET", fmt.Sprintf("/api/%s/orderbook", currency), nil, nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}

// Trades - Trades in exchange
func (p Public) Trades(currency string) (*TradeResponse, *Error, error) {
	var response *TradeResponse
	err, errAPI := p.client.Request("GET", fmt.Sprintf("/api/%s/trades/", currency), nil, nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}
