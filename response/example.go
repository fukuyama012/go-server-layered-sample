package response

import "github.com/fukuyama012/go-server-layered-sample/domain/model"

type (
	// Example struct
	Example struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}
)

// NewExample returns example instance
func NewExample() *Example {
	return &Example{}
}

// ConvertFromModel convert model to response
func (r *Example) ConvertFromModel(m *model.Example) *Example {
	r.ID = m.ID
	r.Name = m.Name

	return r
}
