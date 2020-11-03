package rosetta

type NetworkIdentifier struct {
	Blockchain string `json:"blockchain,omitempty"`
	Network    string `json:"network,omitempty"`
}

type RosettaRequest struct {
	NetworkIdentifier     NetworkIdentifier     `json:"network_identifier,omitempty"`
	BlockIdentifier       BlockIdentifier       `json:"block_identifier,omitempty"`
	TransactionIdentifier TransactionIdentifier `json:"transaction_identifier,omitempty"`
}

type BlockIdentifier struct {
	Index int64  `bson:"index" json:"index,omitempty"`
	Hash  string `bson:"hash" json:"hash,omitempty"`
}
type Peer struct {
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	ID       string                 `json:"peer_id,omitempty"`
}

type NetworkStatus struct {
	CurrentBlockIdentifier BlockIdentifier `json:"current_block_identifier,omitempty"`
	CurrentBlockTimestamp  int64           `json:"current_block_timestamp,omitempty"`
	GenesisBlockIdentifier BlockIdentifier `json:"genesis_block_identifier,omitempty"`
	Peers                  []Peer          `json:"peers,omitempty"`
}

type Block struct {
	BlockIdentifier       BlockIdentifier `json:"block_identifier,omitempty"`
	ParentBlockIdentifier BlockIdentifier `json:"parent_block_identifier,omitempty"`
}
type AccountIdentifier struct {
	Address string `json:"address,omitempty"`
}
type OperationIdentifier struct {
	Index        int64 `json:"index,omitempty"`
	NetworkIndex int64 `json:"network_index,omitempty"`
}

type Operation struct {
	OperationIdentifier OperationIdentifier `json:"operation_identifier,omitempty"`
	Account             *AccountIdentifier  `json:"account,omitempty"`
}
type Metadata map[string]interface{}

type Transaction struct {
	TransactionIdentifier TransactionIdentifier `json:"transaction_identifier,omitempty"`
	Operations            []Operation           `json:"operations,omitempty"`
	Metadata              Metadata              `json:"metadata,omitempty"`
}
type TransactionIdentifier struct {
	Hash string `json:"hash,omitempty"`
}

type BlockResponse struct {
	Block             Block                   `json:"block,omitempty"`
	OtherTransactions []TransactionIdentifier `json:"other_transactions,omitempty"`
}
type BlockTransactionResponse struct {
	Transaction Transaction `json:"transaction,omitempty"`
}
