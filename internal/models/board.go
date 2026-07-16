package models

import (
	"Lechenco/sudoku-solver/internal/logging"
	"Lechenco/sudoku-solver/internal/models/cells"
	"Lechenco/sudoku-solver/internal/models/regions"
	"Lechenco/sudoku-solver/utils"
	"fmt"
	"log/slog"
	"slices"
)

// CellGrid is a abstraction from a 9x9 matrix of cells
type CellGrid [9][9]*cells.Cell

// SetValue sets value to the cell on position
func (g *CellGrid) SetValue(position cells.Position, value cells.Value) error {
	return g[position.RowNumber][position.ColumnNumber].SetValue(value)
}

type Board struct {
	Cells          CellGrid
	Columns        [9]*regions.ColumnRegion
	Rows           [9]*regions.RowRegion
	Squares        [9]*regions.SquareRegion
	AllRegions     []regions.Region
	CleanedRegions []bool
	logger         *slog.Logger
}

// Check if all regions were cleaned
func (b *Board) Finished() (res bool) {
	res = true
	for _, v := range b.CleanedRegions {
		res = res && v
		if !res {
			break
		}
	}

	return
}

func (b *Board) GetCells() CellGrid {
	return b.Cells
}

// SetValue try to set the value at position, if success it checks if the board
// is still valid and clean the value candidate from related regions.
func (b *Board) SetValue(position cells.Position, value cells.Value) (err error) {
	b.logger.Info(fmt.Sprintf("Atribuindo a posição: [%v] o valor: [%v]", position, value))
	err = b.Cells.SetValue(position, value)

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
func (b *Board) cleanFromRegions(position cells.Position, value cells.Value) {
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
func (b *Board) getRegions(position cells.Position) (arr []regions.Region) {
	arr = append(arr, b.Columns[position.ColumnNumber])
	arr = append(arr, b.Rows[position.RowNumber])

	row, col := utils.IndexSquareRegionPos(position)
	arr = append(arr, b.Squares[3*row+col])

	return
}

// Valid verify the puzzle restrictions in all board regions,
// returning a error if any of it was crossed.
func (b *Board) Valid() error {

	for _, region := range b.AllRegions {
		if err := region.Valid(); err != nil {
			return err
		}
	}

	return nil
}

// Init set all positions, regions and candidates from each cell in the board
func (b *Board) Init() {
	b.logger = logging.LoggerFactory("models/Board")
	b.logger.Info("Iniciando variáveis do Board...")
	b.initPositions()
	b.initRegions()
	b.initCandidates()
	b.logger.Info("Variáveis do Board iniciadas.")
}

// initCandidates iterate to all Cells removing a value from related regions candidates
func (b *Board) initCandidates() {
	b.logger.Info("Iniciando Candidatos do Board...")
	for cell := range CellsIterator(b.Cells) {
		b.logger.Debug("Iniciando Candidatos célula: " + cell.String())
		if !cell.IsEmpty() {
			b.cleanFromRegions(cell.Position, cell.Value)
		}
	}
	b.logger.Info("Candidatos do Board iniciados.")
}

func (b *Board) initRegions() {
	b.logger.Info("Iniciando Regiões do Board...")
	b.AllRegions = []regions.Region{}
	b.initRows()
	b.initColumns()
	b.initSquares()
	b.CleanedRegions = make([]bool, len(b.AllRegions))
	b.logger.Info("Regiões do Board iniciadas.")
}

func (b *Board) initPositions() {
	b.logger.Info("Iniciando Posições do Board...")
	for i, row := range b.Cells {
		for j, cell := range row {
			cell.Position = cells.Position{RowNumber: uint8(i), ColumnNumber: uint8(j)}
		}
	}
	b.logger.Info("Posições do Board iniciadas.")
}

func (b *Board) initRows() {
	for i := range 9 {
		b.Rows[i] = regions.NewRowsRegion(b.Cells[i])
		b.AllRegions = append(b.AllRegions, b.Rows[i])
	}
}

func (b *Board) initColumns() {
	for j := range 9 {
		cells := [9]*cells.Cell{}
		for i := range 9 {
			cells[i] = b.Cells[i][j]
		}
		b.Columns[j] = regions.NewColumnRegion(cells)
		b.AllRegions = append(b.AllRegions, b.Columns[j])
	}

}

func (b *Board) initSquares() {
	for i := range 9 {
		var cells [3][3]*cells.Cell
		for j := range 9 {
			row, col := (i/3)*3+(j/3), (i%3)*3+(j%3)
			cells[j/3][j%3] = b.Cells[row][col]
		}
		b.Squares[i] = regions.NewSquareRegion(cells)
		b.AllRegions = append(b.AllRegions, b.Squares[i])
	}
}
