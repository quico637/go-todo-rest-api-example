package model

type Project struct {
	Title    string `json:"title"`
	Description string `json:"description"`
	Archived bool   `json:"archived"`

}

func (p *Project) Archive() {
	p.Archived = true
}

func (p *Project) Restore() {
	p.Archived = false
}

