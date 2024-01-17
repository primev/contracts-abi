package config

type Contracts struct {
	BidderRegistry         string
	ProviderRegistry       string
	PreconfCommitmentStore string
	Oracle                 string
}

var TestnetContracts = Contracts{
	BidderRegistry:         "0xb1692ab820BF196d54e21Fbb1607931558B1d077",
	ProviderRegistry:       "0xeA73E67c2E34C4E02A2f3c5D416F59B76e7617fC",
	PreconfCommitmentStore: "0x8C8571F28a90813f8C468E5dbCF3B53c94aCE580",
	Oracle:                 "0x8a965147eD06B3faF672a71bED3689161B8109c3",
}
