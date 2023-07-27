package Antarctica

type Block struct{
	// data location
	network			string					// network id
	chain			string 					// chain id
	epoch			uint
	index			uint
	// block data
	nonce			string
	hash			string
	previous 		string
	createdTime		uint
	closedTime		uint
	data			[] interface{}
	validator		string					// validator node id
}