package a

import (
	"a/schema"
)

func foutside() {
	_ = schema.Schema{
		Computed:     true,
		ValidateFunc: func() {},
	}

	_ = map[string]*schema.Schema{
		"name": {
			Computed:     true,
			ValidateFunc: func() {},
		},
	}
}
