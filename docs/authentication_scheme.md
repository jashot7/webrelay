Web Relay Authentication Scheme
===============================

1. Client sends HelloMessage to Relay server 

```go
type HelloMessage struct {
	UserID          string `json:"userID"`
}
```

2. Relay sends AuthChallengeMessage to Client

```go
type AuthChallengeMessage struct {
	Nonce          string `json:"nonce"`
}
```

3. Client sends AuthMessage to Relay

```go
type AuthMessage struct {
  UserID         string `json:"userID"`
  PubKey         string `json:"pubKey"`
  SignedNonce    string `json:"signedNonce"`
}
```

4. Relay validates the AuthMessage signature
5. Relay computes the SubscriptionKey from the UserID
6. Relay adds the Client to the connectedNodes array

