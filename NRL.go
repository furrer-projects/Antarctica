package Antarctica

type NRL struct{
	storageModel	uint
	hash			string
	segment			uint
	mimeType		string
}

const NRL_SRC_SQL_DB			= 1
const NRL_SRC_OBJ_DB			= 2
const NRL_SRC_FILE				= 4


func parse(nrl string) NRL{
	var ret NRL
	return ret
}

func stringify(nrl NRL) string{
	var ret string
	return ret
}