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
                  536278.34
                  6..195...
                  .98....6.
                  8...6...3
                  4....3..1
                  7...2...6
                  .6..5.2..
                  7..8.3..5
                  .......8.
                  """
            Então o próximo passo é 1 na posição (0,6)
