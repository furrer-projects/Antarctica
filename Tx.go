package Antarctica

type Tx struct{
	// data location and addressing
	network					string
	chain					string
	block 					string
	epoch					uint
	index					uint
	// data identification
	model					string
	hash					string
	createdTime				uint
	// data model specification
	incomming				[] interface{}
	outgoing				[] interface{}
	fee						uint
}