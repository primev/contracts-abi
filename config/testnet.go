package config

type Contracts struct {
	BidderRegistry         string
	ProviderRegistry       string
	PreconfCommitmentStore string
	Oracle                 string
}

var TestnetContracts = Contracts{
	BidderRegistry:         "0x5340b92E261141D6B4D0DC6F847667E5e4A63544",
	ProviderRegistry:       "0xeA73E67c2E34C4E02A2f3c5D416F59B76e7617fC",
	PreconfCommitmentStore: "0x451656c1E7eDf82397EBE04f38819c9970AA3658",
	Oracle:                 "0x943685b6999626D2323526ce3fb6EF1786Ee8Ee3",
}
