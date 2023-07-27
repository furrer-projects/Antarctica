package Antarctica

type Feed struct{
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
	modifiedTime			uint
	// data model specification
	title					string
	text					string
	images					[] string			// image url, network based file storage, nft hash
}