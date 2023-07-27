package Antarctica

type Permission struct{
	resource 				string				// resource hash
	account					string				// account address
	accessMode				uint
}

const ACCESS_MODE_READ				= 1
const ACCESS_MODE_MODIFY			= 2
const ACCESS_MODE_GRANT				= 4