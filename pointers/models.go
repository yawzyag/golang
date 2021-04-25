package pointers

// Bitcoin type definition
type Bitcoin int

// Stringer defines a new type
type Stringer interface {
	String() string
}

// Wallet struct definition
type Wallet struct {
	balance Bitcoin
}
