package models

// Transaction struct
type Transaction struct {
	PaymentSellerDisplayName string `json:"paymentSellerDisplayName"`
	PaymentBuyerDisplayName  string `json:"paymentBuyerDisplayName"`
	PaymentDate              uint64 `json:"paymentDate"`
	StoryID                  string `json:"storyId"`
}

// Profile asd
type Profile struct {
	ProfileID          string `json:"profileId"`
	DisplayName        string `json:"displayName"`
	FirstName          string `json:"firstName"`
	LastName           string `json:"lastName"`
	Address            string `json:"address"`
	Email              string `json:"email"`
	IsMedia            bool   `json:"isMeadia"`
	IsProfessional     bool   `json:"isProfessional"`
	IsPress            bool   `json:"isPress"`
	SalesQuantity      int64  `json:"salesQuantity"`
	SalesAmount        int64  `json:"salesAmount"`
	WithdrawableAmount int64  `json:"withdrawableAmount"`
	CreateDate         int64  `json:"createDate"`
	// CompanyVatNumber   string `json:"companyVatNumber"`
}

// Withdrawals ..
type Withdrawals struct {
	// CashoutDetails       string `json:"cashoutDetails"`
	RequestAmount        int64  `json:"requestAmount"`
	RequestCompleted     bool   `json:"requestCompleted"`
	RequestCompletedDate int64  `json:"requestCompletedDate"`
	RequestUserID        string `json:"requestUser"`
	RequestDate          int64  `json:"requestDate"`
}

// TokenBody --
type TokenBody struct {
	email    string
	password string
	token    string
}
