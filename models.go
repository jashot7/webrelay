package main

import "encoding/json"

type TypedMessage struct{
	Type string
	Data json.RawMessage
}

type SubscribeMessage struct {
	UserID          string `json:"userID"`
	PubKey 			string `json:"pubKey"`
	SignedNonce		string `json:"signedNonce"`
}

type HelloMessage struct {
	UserID          string `json:"userID"`
}

type EncryptedMessage struct {
	ID        string `json:"id"`
	Message   string `json:"encryptedMessage"`
	Recipient string `json:"recipient"`
}

type AckMessage struct {
	MessageID string `json:"messageID"`
}
