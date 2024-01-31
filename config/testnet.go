package config

type Contracts struct {
	BidderRegistry         string
	ProviderRegistry       string
	PreconfCommitmentStore string
	Oracle                 string
}

var TestnetContracts = Contracts{
	BidderRegistry:         "0xa86a41b57Fb73f9118F84847574517258d29eAD0",
	ProviderRegistry:       "0x5960774AD41D03DAB4916a30bD2190f8b3b3b4b2",
	PreconfCommitmentStore: "0x7D1a4707e573D260581f3AB3f90f697Ab03fC6Dd",
	Oracle:                 "0xFA626Ad00244CC08e4E34A10d2d1Aa1E930AA6dC",
}
