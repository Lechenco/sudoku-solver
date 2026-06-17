Funcionalidade: Estratégia Naked Single

    Naked Single se refere quando um único valor válido restou para uma célula.
    Todos os valores foram preenchidos em pelo menos uma região
    Dado um tabuleiro de sudoku, ao utilizar a estratégia de Naked Single deve
    ser possível identificar a posição e valor da próxima célula a ser 
    preenchida.


    Regra: Um candidato possível
        Cenário: Naked Single na linha
            Dado a estratégia "naked_single"
            E o tabuleiro abaixo:
                  """
                  516278.34
                  2..1.5...
                  3.8....6.
                  8...6...3
                  6....3..1
                  4.......6
                  .6..5.2..
                  7..8....5
                  .......8.
                  """
            E tabuleiro válido
            Então o próximo passo é 9 na posição (0,6)

        Cenário: Naked Single na coluna
            Dado a estratégia "naked_single"
            E o tabuleiro abaixo:
                  """
                  516278934
                  2....5...
                  3.8....6.
                  8...6...3
                  6....3...
                  4.......6
                  .6..5.2..
                  72.8....5
                  1......8.
                  """
            E tabuleiro válido
            Então o próximo passo é 9 na posição (6,0)
