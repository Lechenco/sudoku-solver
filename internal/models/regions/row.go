package regions

import (
	"fmt"
	"strings"

	"github.com/Lechenco/sudoku-solver/internal/logging"
	"github.com/Lechenco/sudoku-solver/internal/models/cells"
)

type RowRegion struct {
	linearRegion
}

func NewRowsRegion(cells [9]*cells.Cell) *RowRegion {
	res := &RowRegion{
		linearRegion{
			Cells:      cells,
			baseRegion: &baseRegion{},
		},
	}
	res.baseRegion.Region = res
	res.baseRegion.logger = logging.LoggerFactory("region/RowRegion")

	return res
}

func (r *RowRegion) String() string {
	values := []string{}

	for _, c := range r.GetCells() {
		values = append(values, c.Value.String())
	}

	return fmt.Sprintf("{RowRegion[%v]}", strings.Join(values, ","))
}
