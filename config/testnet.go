package config

type Contracts struct {
	BidderRegistry         string
	ProviderRegistry       string
	PreconfCommitmentStore string
	Oracle                 string
}

var TestnetContracts = Contracts{
	BidderRegistry:         "0xa86a41b57Fb73f9118F84847574517258d29eAD0",
	ProviderRegistry:       "0x43009fbae169Ad4633737Ad3Ab2ED6A0C1E1eE4A",
	PreconfCommitmentStore: "0x34C326813051E95b9657f6B39931be50393C1202",
	Oracle:                 "0xB8721E3551F3d422602Eac3C89ccD26c75757bB7",
}
