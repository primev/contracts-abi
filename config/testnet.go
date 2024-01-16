package config

type Contracts struct {
	BidderRegistry         string
	ProviderRegistry       string
	PreconfCommitmentStore string
	Oracle                 string
}

var TestnetContracts = Contracts{
	BidderRegistry:         "0x62197Abd7672925c7606Bdf9931e42baCa6619AD",
	ProviderRegistry:       "0xeA73E67c2E34C4E02A2f3c5D416F59B76e7617fC",
	PreconfCommitmentStore: "0x39a591256f51a36e83F63252E727F5a5d8711236",
	Oracle:                 "0x8dfFDA0B27E848c69D92fF1F3E02c274a5801646",
}
