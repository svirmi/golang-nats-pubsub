package models

const MessageType_Ticker MessageType = "ticker"

type MessageType string

type TickerData struct {
	AskPrice  string `json:"ask_price"`
	BidPrice  string `json:"bid_price"`
	LastPrice string `json:"last_price"`
	Volume    string `json:"volume"`
	High      string `json:"high"`
	Low       string `json:"low"`
}

type TickerResponse struct {
	TradingPairId int64      `json:"trading_pair_id"`
	Data          TickerData `json:"data"`
}

type blocktradeWebsocketMessage struct {
	MessageType MessageType            `json:"message_type"`
	Payload     map[string]interface{} `json:"payload"`
}

type websocketMessage struct {
	Message []byte
	Error   error
}
