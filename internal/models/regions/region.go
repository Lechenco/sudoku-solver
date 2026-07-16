package regions

import (
	"Lechenco/sudoku-solver/internal/models/cells"
	"fmt"
	"log/slog"
)

type Region interface {
	GetCandidates() cells.ValuesSet
	GetCell(index uint8) *cells.Cell
	GetCells() []*cells.Cell
	GetCellsCandidates() []cells.ValuesSet
	Valid() error
	RemoveCandidate(value cells.Value)
}

type baseRegion struct {
	Region
	logger *slog.Logger
}

func (s *baseRegion) GetCandidates() cells.ValuesSet {
	candidate := cells.ValuesSet(0x0)

	for _, cell := range s.GetCells() {
		candidate |= cell.Candidates
	}

	return candidate
}

func (s *baseRegion) GetCell(index uint8) *cells.Cell {
	return s.GetCells()[index]
}

// Valid iterate over the region cells searching for a broken puzzle restrition:
//   - Duplicated value in a same region
//   - Empty cell with no possible candidate
// In one of these cases, return a new error
func (s *baseRegion) Valid() error {
	s.logger.Debug(fmt.Sprintf("Validando região [%v]", s.Region))
	onRegion := cells.ValuesSet(0)

	for _, cell := range s.GetCells() {
		if cell.Value != 0 {
			if onRegion.Has(cell.Value) {
				s.logger.Error(fmt.Sprintf("Valor duplicado na mesma região: [%v]", s.Region))
				return fmt.Errorf("Valor duplicado na mesma região: [%v]", s.Region)
			}
			onRegion.Add(cell.Value)
		} else if cell.Candidates.IsEmpty() {
			s.logger.Error(fmt.Sprintf("Célula vazia com nenhum candidato possível: [%v]", cell))
			return fmt.Errorf("Célula vazia com nenhum candidato possível: [%v]", cell)
		}
	}
	return nil
}

// RemoveCandidate remove value from all region cells
func (s *baseRegion) RemoveCandidate(value cells.Value) {
	s.logger.Debug(fmt.Sprintf("Removendo candidato [%v] de todas as células da região [%v]", value, s.Region))
	for _, cell := range s.GetCells() {
		cell.Candidates.Remove(value)
	}
}

func (s *baseRegion) GetCellsCandidates() []cells.ValuesSet {
	cs := s.GetCells()
	candidates := make([]cells.ValuesSet, len(cs))

	for i, cell := range cs{
		candidates[i] = cell.Candidates
	}
	return candidates
}

type linearRegion struct {
	*baseRegion
	Cells [9]*cells.Cell
}

func (c *linearRegion) GetCells() []*cells.Cell {
	return c.Cells[:]
}

