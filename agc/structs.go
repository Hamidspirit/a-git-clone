package agc

// i actually dont use most of this structs currently
// GitSruct is my main repo handler
type GitStruct struct {
	HEAD    string
	Config  string
	Objects map[string]interface{}
	Refs    Refs
}

type Refs struct {
	Heads map[string]interface{}
}

type HashedObject struct {
	FPath    string
	ObjectID string
	Name     string
}

type TreeEntrie struct {
	Mode string
	Path string
	Hash string
}
