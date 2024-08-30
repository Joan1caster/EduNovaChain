package models

type Metadata struct {
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Content string `json:"content"`
}

type NFTData struct {
	WalletAddress string   `json:"walletAddress"`
	Metadata      Metadata `json:"metadata"`
}
