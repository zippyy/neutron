package backend

type User struct {
	ID string
	Name string
	NotificationEmail string
	Signature string
	NumMessagePerPage int
	UsedSpace int
	Notify int
	AutoSaveContacts int
	Language string
	LogAuth int
	ComposerMode int
	MessageButtons int
	ShowImages int
	ViewMode int
	ViewLayout int
	SwipeLeft int
	SwipeRight int
	Theme string
	Currency string
	Credit int
	DisplayName string
	MaxSpace int
	MaxUpload int
	Role int
	Private int
	Subscribed int
	Deliquent int
	Addresses []*Address
	PublicKey string
	EncPrivateKey string
}

type Address struct {
	ID string
	DomainID string
	Email string
	Send int
	Receive int
	Status int
	Type int
	DisplayName string
	Signature string
	HashKeys int
	Keys []*Keypair
}

type Keypair struct {
	ID string
	PublicKey string
	PrivateKey string
}