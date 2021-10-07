package model

type Result struct {
	Data  interface{}
	Error error
}

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type File struct {
	Directory string `json:"directory"`
	Name      string `json:"name"`
	Size      string `json:"size"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Zip struct {
	Directory string `json:"directory"`
	Name      string `json:"name"`
}

type MultipleFilePayload struct {
	Filenames []string
}
