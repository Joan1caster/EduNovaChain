package models

import "time"

type Metadata struct {
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Content string `json:"content"`
}

type IpfsResponseData struct {
	IpfsHash    string    `json:"IpfsHash"`
	PinSize     int       `json:"PinSize"`
	Timestamp   time.Time `json:"Timestamp"`
	IsDuplicate bool      `json:"isDuplicate"`
}

type Region struct {
	RegionID                string `json:"regionId"`
	CurrentReplicationCount int    `json:"currentReplicationCount"`
	DesiredReplicationCount int    `json:"desiredReplicationCount"`
}

type Row struct {
	ID            string    `json:"id"`
	IPFSPinHash   string    `json:"ipfs_pin_hash"`
	Size          int       `json:"size"`
	UserID        string    `json:"user_id"`
	DatePinned    time.Time `json:"date_pinned"`
	DateUnpinned  time.Time `json:"date_unpinned"`
	Metadata      Metadata  `json:"metadata"`
	Regions       []Region  `json:"regions"`
	MimeType      string    `json:"mime_type"`
	NumberOfFiles int       `json:"number_of_files"`
}

type IpfsData struct {
	Count int   `json:"count"`
	Rows  []Row `json:"rows"`
}
