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
                ┌───────┬───────┬───────┐
                │ 5 1 6 │ 2 7 8 │ _ 3 4 │ 
                │ 2 _ _ │ 1 _ 5 │ _ _ _ │ 
                │ 3 _ 8 │ _ _ _ │ _ 6 _ │ 
                ├───────┼───────┼───────┤
                │ 8 _ _ │ _ 6 _ │ _ _ 3 │ 
                │ 6 _ _ │ _ _ 3 │ _ _ 1 │ 
                │ 4 _ _ │ _ 2 _ │ _ _ 6 │ 
                ├───────┼───────┼───────┤
                │ _ 6 _ │ _ 5 _ │ 2 _ _ │ 
                │ 7 _ _ │ 8 _ 2 │ _ _ 5 │ 
                │ _ _ _ │ _ _ _ │ _ 8 _ │ 
                └───────┴───────┴───────┘
                  """
            E tabuleiro válido
            Então o próximo passo é 9 na posição (0,6)

        Cenário: Naked Single na coluna
            Dado a estratégia "naked_single"
            E o tabuleiro abaixo:
            """
                ┌───────┬───────┬───────┐
                │ 5 1 6 │ 2 7 8 │ 9 3 4 │ 
                │ 2 _ _ │ _ _ 5 │ _ _ _ │ 
                │ 3 _ 8 │ _ _ _ │ _ 6 _ │ 
                ├───────┼───────┼───────┤
                │ 8 _ _ │ _ 6 _ │ _ _ 3 │ 
                │ 6 _ _ │ _ _ 3 │ _ _ 1 │ 
                │ 4 _ _ │ _ _ _ │ _ _ 6 │ 
                ├───────┼───────┼───────┤
                │ _ 6 _ │ _ 5 _ │ 2 _ _ │ 
                │ 7 2 _ │ 8 _ _ │ _ _ 5 │ 
                │ 1 _ _ │ _ _ _ │ _ 8 _ │ 
                └───────┴───────┴───────┘
                  """
            E tabuleiro válido
            Então o próximo passo é 9 na posição (6,0)

        Cenário: Naked Single no quadrado
            Dado a estratégia "naked_single"
            E o tabuleiro abaixo:
            """
                ┌───────┬───────┬───────┐
                │ 5 1 6 │ 2 7 8 │ 9 3 4 │ 
                │ 2 _ _ │ 6 _ 5 │ _ _ _ │ 
                │ 3 _ 8 │ 1 4 9 │ _ 6 _ │ 
                ├───────┼───────┼───────┤
                │ 8 _ _ │ _ 6 _ │ _ _ 3 │ 
                │ 6 _ _ │ _ _ 3 │ _ _ 1 │ 
                │ 4 _ _ │ _ _ _ │ _ _ 6 │ 
                ├───────┼───────┼───────┤
                │ 9 6 _ │ _ 5 _ │ 2 _ _ │ 
                │ 7 2 _ │ 8 _ _ │ _ _ 5 │ 
                │ 1 _ _ │ _ _ _ │ _ 8 _ │ 
                └───────┴───────┴───────┘
                  """
            E tabuleiro válido
            Então o próximo passo é 3 na posição (1,4)

        Cenário: Naked Single na linha e coluna e quadrado
            Dado a estratégia "naked_single"
            E o tabuleiro abaixo:
            """
                ┌───────┬───────┬───────┐
                │ 5 1 6 │ 2 7 8 │ 9 3 4 │ 
                │ 2 _ _ │ _ _ 5 │ _ _ _ │ 
                │ 3 _ 8 │ _ _ _ │ _ 6 _ │ 
                ├───────┼───────┼───────┤
                │ 8 _ _ │ _ 6 _ │ _ _ 3 │ 
                │ 6 _ _ │ _ _ 3 │ 7 _ 1 │ 
                │ 4 _ _ │ 5 _ _ │ _ _ 6 │ 
                ├───────┼───────┼───────┤
                │ 9 6 _ │ _ 5 _ │ 2 _ _ │ 
                │ 7 2 _ │ 8 _ _ │ _ _ 5 │ 
                │ 1 _ _ │ _ _ _ │ _ 8 _ │ 
                └───────┴───────┴───────┘
                  """
            E tabuleiro válido
            Então o próximo passo é 8 na posição (5,6)

    Regra: Nenhum candidato possível
        Cenário: Tabuleiro não apresenta nenhum passo com naked_single
            Dado a estratégia "naked_single"
            E o tabuleiro abaixo:
            """
                ┌───────┬───────┬───────┐
                │ 5 1 6 │ 2 7 8 │ 9 3 4 │ 
                │ 2 _ _ │ _ _ 5 │ _ _ _ │ 
                │ 3 _ 8 │ _ _ _ │ _ 6 _ │ 
                ├───────┼───────┼───────┤
                │ 8 _ _ │ _ 6 _ │ _ _ 3 │ 
                │ _ _ _ │ _ _ 3 │ 7 _ 1 │ 
                │ _ _ _ │ 5 _ _ │ _ _ 6 │ 
                ├───────┼───────┼───────┤
                │ _ 6 _ │ _ 5 _ │ 2 _ _ │ 
                │ 7 2 _ │ 8 _ _ │ _ _ 5 │ 
                │ 1 _ _ │ _ _ _ │ _ 8 _ │ 
                └───────┴───────┴───────┘
                  """
            E tabuleiro válido
            Então não foi possível determinar o próximo passo
