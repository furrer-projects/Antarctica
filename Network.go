package Antarctica

type Network struct{
	id				string
	name 			string
	epoch			uint
	chains			[]string 				// supported chains ids
	plugins			[]string				// supported data plugins names
	currentNode		interface{}
	version 		float64
}