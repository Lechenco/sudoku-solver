# Sudoku Solver - TODO

## Foundation

- [x] Verify if current state is legal
- [x] Set value in a cell should remove candidates from region
- [x] Set value with position should remove candidates from all related regions
- [x] Not allow to override cell Value
- [x] Not allow to set cell value as zero
- [x] Should scan board and clear all illegal candidates


## Core Functionality

- [x] Implement sudoku puzzle parser
- [ ] Create constraint validation logic
- [ ] Build backtracking solver algorithm
- [ ] Add constraint propagation optimization

## Features

- [ ] Input validation for puzzle format
- [ ] Multiple solver strategies (backtracking, constraint propagation)
- [ ] Solution verification
- [ ] Difficulty level detection
- [ ] Performance profiling

## Testing

- [ ] Unit tests for solver
- [ ] Test with various difficulty levels
- [x] Edge case testing (invalid puzzles, no solution)
- [ ] Performance benchmarks

## Documentation

- [ ] API documentation
- [ ] Usage examples
- [ ] Algorithm explanation
- [ ] README with setup instructions

## UI/Output

- [ ] Pretty print solution
- [ ] Display step-by-step solving process
- [ ] Export results to file
