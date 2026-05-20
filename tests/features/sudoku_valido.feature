
Funcionalidade: Identificar Sudokus inválidos

            Dado um tabuleiro de sudoku, deve-se identificar se o mesmo é válido ou não.
    Para isso checa-se se para cada região existe apenas uma única célula com cada número

    Regra: Valores repetidos

        Cenário: Coluna inicial com valores repetidos
            Dado o tabuleiro abaixo:
                  """
                  53..7....
                  6..195...
                  .98....6.
                  8...6...3
                  4....3..1
                  7...2...6
                  .6..5.2..
                  7..8.3..5
                  .9.....6.
                  """
             Então o jogo deve ser inválido
