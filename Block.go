package Antarctica

type Block struct{
	epoch			uint
	network			string					// network id
	chain			string 					// chain id
	index			uint
	nonce			string
	hash			string
	previous 		string
	createdTime		uint
	closedTime		uint
	data			[] interface{}
	validator		string					// validator node id
}