Web Relay Authentication Scheme
===============================

1. Client sends HelloMessage to Relay server 

```go
type HelloMessage struct {
	UserID          string `json:"userID"`  // Client peer id
}
```

2. Relay sends AuthChallengeMessage to Client. 

```go
type AuthChallengeMessage struct {
	Nonce          string `json:"nonce"`  // 24 bit random nonce
}
```

3. Client sends AuthMessage to Relay

```go
type AuthMessage struct {
  UserID         string `json:"userID"`   // client peer id
  PubKey         string `json:"pubKey"`   // client's identity pubkey
  SignedNonce    string `json:"signedNonce"`  // signed 24 bit nonce
}
```

4. Relay validates the AuthMessage signature
5. Relay computes the SubscriptionKey from the UserID
6. Relay adds the Client to the connectedNodes array

