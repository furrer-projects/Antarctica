package Antarctica

type Resource struct{
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
	objectModel				string
	location				string				// URL, NRL (Network Resource Storage Location)
	modified				uint
}