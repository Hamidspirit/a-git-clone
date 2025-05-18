package agc

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

type Tree struct {
	TreeHash string
	Entries  []TreeLeaf
}

type TreeLeaf struct {
	ObjectType string
	OID        string
	Name       string
}

type HashedObject struct {
	FPath    string
	ObjectID string
	Name     string
}
