Funcionalidade: Identificar Sudokus inválidos

            Dado um tabuleiro de sudoku, deve-se identificar se o mesmo é válido ou não.
    Para isso checa-se se para cada região existe apenas uma única célula com cada número

    Regra: Valores repetidos

        Cenário: Coluna inicial com valores repetidos
            Dado o tabuleiro abaixo:
            """
                ┌───────┬───────┬───────┐
                │ 5 3 _ │ _ 7 _ │ 3 _ _ │ 
                │ 6 _ _ │ 1 9 5 │ _ _ _ │ 
                │ _ 9 8 │ _ _ _ │ _ 6 _ │ 
                ├───────┼───────┼───────┤
                │ 8 _ _ │ _ 6 _ │ _ _ 3 │ 
                │ 4 _ _ │ _ _ 3 │ _ _ 1 │ 
                │ 7 _ _ │ _ 2 _ │ _ _ 6 │ 
                ├───────┼───────┼───────┤
                │ _ 6 _ │ _ 5 _ │ 2 _ _ │ 
                │ 7 _ _ │ 8 _ 2 │ _ _ 5 │ 
                │ _ _ _ │ _ _ _ │ _ 8 _ │ 
                └───────┴───────┴───────┘
                  """
             Então o jogo deve ser inválido

        Cenário: Linha inicial com valores repetidos
            Dado o tabuleiro abaixo:
            """
                ┌───────┬───────┬───────┐
                │ 5 3 _ │ _ 7 _ │ 3 _ _ │ 
                │ 6 _ _ │ 1 9 5 │ _ _ _ │ 
                │ _ 9 8 │ _ _ _ │ _ 6 _ │ 
                ├───────┼───────┼───────┤
                │ 8 _ _ │ _ 6 _ │ _ _ 3 │ 
                │ 4 _ _ │ _ _ 3 │ _ _ 1 │ 
                │ _ _ _ │ _ 2 _ │ _ _ 6 │ 
                ├───────┼───────┼───────┤
                │ _ 6 _ │ _ 5 _ │ 2 _ _ │ 
                │ 7 _ _ │ 8 _ 2 │ _ _ 5 │ 
                │ _ _ _ │ _ _ _ │ _ 8 _ │ 
                └───────┴───────┴───────┘
                  """
             Então o jogo deve ser inválido

        Cenário: Quadrado inicial com valores repetidos
            Dado o tabuleiro abaixo:
            """
                ┌───────┬───────┬───────┐
                │ 5 _ _ │ _ 7 _ │ 3 _ _ │ 
                │ 6 _ _ │ 1 9 5 │ _ _ _ │ 
                │ _ 9 6 │ _ _ _ │ _ 6 _ │ 
                ├───────┼───────┼───────┤
                │ 8 _ _ │ _ 6 _ │ _ _ 3 │ 
                │ 4 _ _ │ _ _ 3 │ _ _ 1 │ 
                │ _ _ _ │ _ 2 _ │ _ _ 6 │ 
                ├───────┼───────┼───────┤
                │ _ 6 _ │ _ 5 _ │ 2 _ _ │ 
                │ 7 _ _ │ 8 _ 2 │ _ _ 5 │ 
                │ _ _ _ │ _ _ _ │ _ 8 _ │ 
                └───────┴───────┴───────┘
                  """
             Então o jogo deve ser inválido

    Regra: Solução impossível quando célula vazia não tem candidatos possíveis
        Cenário: Linha com uma célula disponível
            Dado o tabuleiro abaixo:
            """
                ┌───────┬───────┬───────┐
                │ 5 _ _ │ _ 7 _ │ 3 _ _ │ 
                │ 6 _ _ │ 1 9 5 │ _ _ _ │ 
                │ _ 9 8 │ _ _ _ │ _ 6 _ │ 
                ├───────┼───────┼───────┤
                │ 8 _ _ │ _ 6 _ │ _ _ 3 │ 
                │ 4 2 6 │ 7 _ 3 │ 8 5 1 │ 
                │ _ _ _ │ _ 2 _ │ _ _ 6 │ 
                ├───────┼───────┼───────┤
                │ _ 6 _ │ _ 5 _ │ 2 _ _ │ 
                │ 7 _ _ │ 8 _ 2 │ _ _ 5 │ 
                │ _ _ _ │ _ _ _ │ _ 8 _ │ 
                └───────┴───────┴───────┘
                  """
             Então o jogo deve ser inválido
