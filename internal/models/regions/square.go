package regions

import (
	"Lechenco/sudoku-solver/internal/logging"
	"Lechenco/sudoku-solver/internal/models/cells"
	"fmt"
	"slices"
	"strings"
)

type SquareRegion struct {
	*baseRegion
	Cells [3][3]*cells.Cell
}

func NewSquareRegion(cells [3][3]*cells.Cell) *SquareRegion {
	res := &SquareRegion{
		Cells: cells,
		baseRegion: &baseRegion{},
	}

	res.baseRegion.Region = res
	res.baseRegion.logger = logging.LoggerFactory("region/SquareRegion")

	return res
}

func (r *SquareRegion) GetCells() []*cells.Cell {
	return slices.Concat(
		r.Cells[0][:],
		r.Cells[1][:],
		r.Cells[2][:],
	)
}

func (s *SquareRegion) String() string {
	values := []string{}

	for _, c := range s.GetCells() {
		values = append(values, c.Value.String())
	}

	return fmt.Sprintf("{SquareRegion[%v]}", strings.Join(values, ","))
}

