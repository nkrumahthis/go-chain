package blockchain

type TxInput struct {
	ID []byte
	Out int
	Sig string
}

type TxOutput struct {
	Value int
	PubKey string
}

type Transaction struct {
	ID []byte
	Inputs []TxInput
	Output []TxOutput
}

