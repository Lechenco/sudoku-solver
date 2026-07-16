package regions

import (
	"Lechenco/sudoku-solver/internal/logging"
	"Lechenco/sudoku-solver/internal/models/cells"
	"fmt"
	"strings"
)

type ColumnRegion struct {
	linearRegion
}

func NewColumnRegion(cells [9]*cells.Cell) *ColumnRegion {
	res := &ColumnRegion{
		linearRegion{
			Cells: cells,
			baseRegion: &baseRegion{},
		},
	}

	res.baseRegion.Region = res
	res.baseRegion.logger = logging.LoggerFactory("region/ColumnRegion")
	return res
}

func (c *ColumnRegion) String() string {
	values := []string{}

	for _, cell := range c.GetCells() {
		values = append(values, cell.Value.String())
	}

	return fmt.Sprintf("{ColumnRegion[%v]}", strings.Join(values, ","))
}

