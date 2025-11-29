package plugify

// Manifest represents the structure of the .pplugin file
type Manifest struct {
	Schema       string       `json:"$schema"`
	Name         string       `json:"name"`
	Version      string       `json:"version"`
	Description  string       `json:"description,omitempty"`
	Author       string       `json:"author,omitempty"`
	Website      string       `json:"website,omitempty"`
	License      string       `json:"license,omitempty"`
	Platforms    []string     `json:"platforms,omitempty"`
	Dependencies []Dependency `json:"dependencies,omitempty"`
	Conflicts    []Conflict   `json:"conflicts,omitempty"`
	Entry        string       `json:"entry"`
	Language     string       `json:"language"`
	Methods      []Method     `json:"methods"`
}

// Method represents a single exported method
type Method struct {
	Name        string     `json:"name"`
	FuncName    string     `json:"funcName"`
	Description string     `json:"description,omitempty"`
	ParamTypes  []Property `json:"paramTypes"`
	RetType     Property   `json:"retType"`
	Group       string     `json:"group,omitempty"`
}

// EnumValue represents a single enumeration value
type EnumValue struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Value       int64  `json:"value"`
}

// EnumObject represents an enumeration
type EnumObject struct {
	Name        string      `json:"name"`
	Description string      `json:"description,omitempty"`
	Values      []EnumValue `json:"values"`
}

// Property represents a parameter type
type Property struct {
	Type        string      `json:"type"`
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	Ref         bool        `json:"ref,omitempty"`
	Prototype   *Method     `json:"prototype,omitempty"`
	Enumerator  *EnumObject `json:"enum,omitempty"`
}

// Dependency represents a plugin's dependency
type Dependency struct {
	Name        string `json:"name"`
	Constraints string `json:"constraints,omitempty"`
	Optional    bool   `json:"optional,omitempty"`
}

// Conflict represents a plugin's conflict
type Conflict struct {
	Name        string `json:"name"`
	Constraints string `json:"constraints,omitempty"`
	Reason      string `json:"reason,omitempty"`
}
