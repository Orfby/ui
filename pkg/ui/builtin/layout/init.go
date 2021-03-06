package layout

import (
	_ "github.com/bhollier/ui/pkg/ui/builtin/layout/relative"
	"github.com/bhollier/ui/pkg/ui/element"
	"reflect"
)

// Function to register the layout types
func init() {
	// Register the layout types
	element.Register(GridLayoutTypeName,
		reflect.TypeOf((*GridLayout)(nil)).Elem(), NewGridLayout)
	element.Register(LinearLayoutTypeName,
		reflect.TypeOf((*LinearLayout)(nil)).Elem(), NewLinearLayout)
}
