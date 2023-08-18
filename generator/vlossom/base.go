package vlossom

type Component interface {
	_Component()
}

func NewBaseComponent(component string, option BaseComponentOptions) BaseComponent {
	base := BaseComponent{
		Component:    component,
		PropertyName: option.PropertyName,
		Label:        option.Label,
		Placeholder:  option.Placeholder,
		Required:     option.Required,
		State:        option.State,
		Width:        option.Width,
	}
	if option.LG != 0 {
		base.Grid = &Grid{LG: option.LG}
	}
	return base
}

type BaseComponentOptions struct {
	PropertyName string
	Label        string
	Placeholder  string
	Required     bool
	State        string

	Width string
	LG    int
}

type BaseComponent struct {
	Component string `json:"component"`

	PropertyName string `json:"propertyName"`

	Label string `json:"label"`

	Placeholder string `json:"placeholder"`

	Required bool `json:"required"`

	State string `json:"state,omitempty"`

	// Width of component conflict with grid
	Width string `json:"width,omitempty"`
	// Grid of component conflict with width
	Grid *Grid `json:"grid,omitempty"`

	/**
	 * Not useful options
	 */
	// No Message? IDK
	NoMsg *bool `json:"no-msg,omitempty"`
	// Disable component
	Disabled *bool `json:"disabled,omitempty"`
	// Component is visible or not
	Visible *bool `json:"visible,omitempty"`
	// Cannot modify input form
	Readonly *bool `json:"readonly,omitempty"`
	// No Label? IDK, It doesn't work to my expected
	NoLabel *bool `json:"no-label,omitempty"`
	// No Clear? IDK
	NoClear *bool `json:"no-clear,omitempty"`
}

type Grid struct {
	LG int `json:"lg"`
}

func (c *BaseComponent) _Component() {}

type VlossomOptionsOptions interface {
	_VlossomOptionsOptions()
}

type SimpleOptions []string

func (o SimpleOptions) _VlossomOptionsOptions() {}

type LabeledOptions []LabeledOption

func (o LabeledOptions) _VlossomOptionsOptions() {}

type VlossomOptions struct {
	Options     VlossomOptionsOptions `json:"options,omitempty"`
	OptionLabel string                `json:"option-label,omitempty"`
	OptionValue string                `json:"option-value,omitempty"`
}

func NewVlossomOptions() *VlossomOptions {
	return &VlossomOptions{}
}

func (o *VlossomOptions) SetLabeledOptions(options LabeledOptions) {
	o.Options = options
	o.OptionValue = "value"
	o.OptionLabel = "label"
}

func (o *VlossomOptions) SetSimpleOptions(options SimpleOptions) {
	o.Options = options
	o.OptionValue = ""
	o.OptionLabel = ""
}

type LabeledOption struct {
	Value any    `json:"value"`
	Label string `json:"label"`
}