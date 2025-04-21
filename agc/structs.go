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
