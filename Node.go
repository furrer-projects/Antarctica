package Antarctica

type Node struct{
	id			string
	name		string
	created 	uint
	modified 	uint
	account 	string
	networks 	[] string
	connections	Connection
	parent 		interface{}
	children	[] interface{}
	version		float64
}

var nodeList []Node