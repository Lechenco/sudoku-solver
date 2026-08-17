package services

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"slices"

	"github.com/Lechenco/sudoku-solver/internal/iterators"
	"github.com/Lechenco/sudoku-solver/internal/logging"
	"github.com/Lechenco/sudoku-solver/models/board"
	"github.com/Lechenco/sudoku-solver/models/cells"
	"github.com/Lechenco/sudoku-solver/models/regions"
	"github.com/Lechenco/sudoku-solver/utils"
)

type BoardService struct {
	Board          *board.Board
	Columns        [9]*regions.ColumnRegion
	Rows           [9]*regions.RowRegion
	Squares        [9]*regions.SquareRegion
	AllRegions     []regions.Region
	CleanedRegions []bool
	logger         *slog.Logger
}

// Check if all regions were cleaned
func (b *BoardService) Finished() (res bool) {
	res = true
	for _, v := range b.CleanedRegions {
		res = res && v
		if !res {
			break
		}
	}

	return
}

func (b *BoardService) GetCells() board.CellGrid {
	return b.Board.Cells
}

// SetValue try to set the value at position, if success it checks if the board
// is still valid and clean the value candidate from related regions.
func (b *BoardService) SetValue(position cells.Position, value cells.Value) (err error) {
	b.logger.Info(fmt.Sprintf("Atribuindo a posição: [%v] o valor: [%v]", position, value))
	err = b.Board.Cells.SetValue(position, value)

	if err != nil {
		return
	}

	err = b.Valid()

	if err != nil {
		return
	}

	b.cleanFromRegions(position, value)

	return
}

// cleanFromRegions remove the value candidate from all regions which the
// position cell is related.
//
// After remove the candidates, updates if the region is cleaned.
func (b *BoardService) cleanFromRegions(position cells.Position, value cells.Value) {
	b.logger.Info(fmt.Sprintf("Removendo candidato [%v] das regiões que contem a posição: [%v]", value, position))
	for _, region := range b.getRegions(position) {
		b.logger.Debug(fmt.Sprintf("Removendo candidato [%v] da região [%v]", value, region))
		region.RemoveCandidate(value)
		if region.GetCandidates().IsEmpty() {
			b.logger.Info(fmt.Sprintf("Região [%v] não tem mais candidatos.", region))
			index := slices.IndexFunc(b.AllRegions, func(r regions.Region) bool {
				return region == r
			})
			b.CleanedRegions[index] = true
		}
	}
	b.logger.Info(fmt.Sprintf("Candidato [%v] removido das regiões", value))
}

// getRegions returns a array with all regions who overlap the position
func (b *BoardService) getRegions(position cells.Position) (arr []regions.Region) {
	arr = append(arr, b.Columns[position.ColumnNumber])
	arr = append(arr, b.Rows[position.RowNumber])

	row, col := utils.IndexSquareRegionPos(position)
	arr = append(arr, b.Squares[3*row+col])

	return
}

// Valid verify the puzzle restrictions in all board regions,
// returning a error if any of it was crossed.
func (b *BoardService) Valid() error {

	for _, region := range b.AllRegions {
		if err := region.Valid(); err != nil {
			return err
		}
	}

	return nil
}

// Init set all positions, regions and candidates from each cell in the board
func (b *BoardService) Init() {
	b.logger = logging.LoggerFactory("models/Board")
	b.logger.Info("Iniciando variáveis do Board...")
	b.initPositions()
	b.initRegions()
	b.initCandidates()
	b.logger.Info("Variáveis do Board iniciadas.")
}

// initCandidates iterate to all Cells removing a value from related regions candidates
func (b *BoardService) initCandidates() {
	b.logger.Info("Iniciando Candidatos do Board...")
	for cell := range iterators.CellsIterator(b.Board.Cells) {
		b.logger.Debug("Iniciando Candidatos célula: " + cell.String())
		if !cell.IsEmpty() {
			b.cleanFromRegions(cell.Position, cell.Value)
		}
	}
	b.logger.Info("Candidatos do Board iniciados.")
}

func (b *BoardService) initRegions() {
	b.logger.Info("Iniciando Regiões do Board...")
	b.AllRegions = []regions.Region{}
	b.initRows()
	b.initColumns()
	b.initSquares()
	b.CleanedRegions = make([]bool, len(b.AllRegions))
	b.logger.Info("Regiões do Board iniciadas.")
}

func (b *BoardService) initPositions() {
	b.logger.Info("Iniciando Posições do Board...")
	for i, row := range b.Board.Cells {
		for j, cell := range row {
			cell.Position = cells.Position{RowNumber: uint8(i), ColumnNumber: uint8(j)}
		}
	}
	b.logger.Info("Posições do Board iniciadas.")
}

func (b *BoardService) initRows() {
	for i := range 9 {
		b.Rows[i] = regions.NewRowsRegion(b.Board.Cells[i])
		b.AllRegions = append(b.AllRegions, b.Rows[i])
	}
}

func (b *BoardService) initColumns() {
	for j := range 9 {
		cells := [9]*cells.Cell{}
		for i := range 9 {
			cells[i] = b.Board.Cells[i][j]
		}
		b.Columns[j] = regions.NewColumnRegion(cells)
		b.AllRegions = append(b.AllRegions, b.Columns[j])
	}

}

func (b *BoardService) initSquares() {
	for i := range 9 {
		var cells [3][3]*cells.Cell
		for j := range 9 {
			row, col := (i/3)*3+(j/3), (i%3)*3+(j%3)
			cells[j/3][j%3] = b.Board.Cells[row][col]
		}
		b.Squares[i] = regions.NewSquareRegion(cells)
		b.AllRegions = append(b.AllRegions, b.Squares[i])
	}
}

func (b *BoardService) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Cells board.CellGrid
	}{
		Cells: b.Board.Cells,
	})
}

func (service *BoardService) UnmarshalJSON(b []byte) error {
	aux := &struct {
		Cells board.CellGrid
	}{}

	if err := json.Unmarshal(b, aux); err != nil {
		return err
	}

	service.Board.Cells = aux.Cells
	return nil
}
