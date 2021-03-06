package webapi

type vspInfoResponse struct {
	Timestamp     int64   `json:"timestamp" binding:"required"`
	PubKey        []byte  `json:"pubkey" binding:"required"`
	FeePercentage float64 `json:"feepercentage" binding:"required"`
	VspClosed     bool    `json:"vspclosed" binding:"required"`
	Network       string  `json:"network" binding:"required"`
}

type FeeAddressRequest struct {
	Timestamp  int64  `json:"timestamp" binding:"required"`
	TicketHash string `json:"tickethash" binding:"required"`
}

type feeAddressResponse struct {
	Timestamp  int64             `json:"timestamp" binding:"required"`
	FeeAddress string            `json:"feeaddress" binding:"required"`
	FeeAmount  float64           `json:"feeamount" binding:"required"`
	Expiration int64             `json:"expiration" binding:"required"`
	Request    FeeAddressRequest `json:"request" binding:"required"`
}

type PayFeeRequest struct {
	Timestamp   int64             `json:"timestamp" binding:"required"`
	TicketHash  string            `json:"tickethash" binding:"required"`
	FeeTx       string            `json:"feetx" binding:"required"`
	VotingKey   string            `json:"votingkey" binding:"required"`
	VoteChoices map[string]string `json:"votechoices" binding:"required"`
}

type payFeeResponse struct {
	Timestamp int64         `json:"timestamp" binding:"required"`
	Request   PayFeeRequest `json:"request" binding:"required"`
}

type SetVoteChoicesRequest struct {
	Timestamp   int64             `json:"timestamp" binding:"required"`
	TicketHash  string            `json:"tickethash" binding:"required"`
	VoteChoices map[string]string `json:"votechoices" binding:"required"`
}

type setVoteChoicesResponse struct {
	Timestamp int64                 `json:"timestamp" binding:"required"`
	Request   SetVoteChoicesRequest `json:"request" binding:"required"`
}

type TicketStatusRequest struct {
	Timestamp  int64  `json:"timestamp" binding:"required"`
	TicketHash string `json:"tickethash" binding:"required"`
}

type ticketStatusResponse struct {
	Timestamp       int64               `json:"timestamp" binding:"required"`
	TicketConfirmed bool                `json:"ticketconfirmed" binding:"required"`
	FeeTxReceived   bool                `json:"feetxreceived" binding:"required"`
	FeeTxBroadcast  bool                `json:"feetxbroadcast" binding:"required"`
	FeeConfirmed    bool                `json:"feeconfirmed" binding:"required"`
	FeeTxHash       string              `json:"feetxhash" binding:"required"`
	VoteChoices     map[string]string   `json:"votechoices" binding:"required"`
	Request         TicketStatusRequest `json:"request" binding:"required"`
}
