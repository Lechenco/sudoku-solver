Funcionalidade: Identificar Sudokus inválidos

            Dado um tabuleiro de sudoku, deve-se identificar se o mesmo é válido ou não.
    Para isso checa-se se para cada região existe apenas uma única célula com cada número

    Regra: Tabuleiro válido
        Cenário: Tabuleiro vazio
            Dado o tabuleiro abaixo:
            """
                ┌───────┬───────┬───────┐
                │ _ _ _ │ _ _ _ │ _ _ _ │ 
                │ _ _ _ │ _ _ _ │ _ _ _ │ 
                │ _ _ _ │ _ _ _ │ _ _ _ │ 
                ├───────┼───────┼───────┤
                │ _ _ _ │ _ _ _ │ _ _ _ │ 
                │ _ _ _ │ _ _ _ │ _ _ _ │ 
                │ _ _ _ │ _ _ _ │ _ _ _ │ 
                ├───────┼───────┼───────┤
                │ _ _ _ │ _ _ _ │ _ _ _ │ 
                │ _ _ _ │ _ _ _ │ _ _ _ │ 
                │ _ _ _ │ _ _ _ │ _ _ _ │ 
                └───────┴───────┴───────┘
                  """
             Então o jogo deve ser válido

        Cenário: Números não ferem as restrições de região
            Dado o tabuleiro abaixo:
            """
                ┌───────┬───────┬───────┐
                │ 5 3 _ │ _ 7 _ │ _ _ _ │ 
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
             Então o jogo deve ser válido

    Regra: Valores repetidos
        Cenário: Coluna inicial com valores repetidos
            Dado o tabuleiro abaixo:
            """
                ┌───────┬───────┬───────┐
                │ 5 3 _ │ _ 7 _ │ _ _ _ │ 
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

        Cenário: Coluna com uma célula disponível
            Dado o tabuleiro abaixo:
            """
                ┌───────┬───────┬───────┐
                │ 5 _ _ │ _ 7 _ │ 3 _ 8 │ 
                │ 6 _ _ │ 1 _ 5 │ _ _ 7 │ 
                │ _ 9 8 │ _ _ _ │ _ 6 4 │ 
                ├───────┼───────┼───────┤
                │ 8 _ _ │ _ 6 _ │ _ _ 3 │ 
                │ 4 2 6 │ 7 _ 3 │ 8 5 1 │ 
                │ _ _ _ │ _ 2 _ │ _ _ 6 │ 
                ├───────┼───────┼───────┤
                │ _ 6 _ │ _ 5 _ │ 2 _ _ │ 
                │ 7 _ _ │ 8 _ 2 │ _ _ 5 │ 
                │ _ _ _ │ _ _ _ │ _ 8 9 │ 
                └───────┴───────┴───────┘
                  """
             Então o jogo deve ser inválido

        Cenário: Quadrado com uma célula disponível
            Dado o tabuleiro abaixo:
            """
                ┌───────┬───────┬───────┐
                │ 5 _ _ │ _ 7 _ │ 3 _ 8 │ 
                │ 6 _ _ │ 1 _ 5 │ _ _ 7 │ 
                │ _ 9 8 │ _ _ _ │ _ 6 4 │ 
                ├───────┼───────┼───────┤
                │ 8 _ _ │ 5 6 1 │ _ _ 3 │ 
                │ 4 2 6 │ 7 _ 3 │ 8 5 1 │ 
                │ _ _ _ │ 4 2 9 │ _ _ 6 │ 
                ├───────┼───────┼───────┤
                │ _ 6 _ │ _ 5 _ │ _ _ _ │ 
                │ 7 _ _ │ 8 _ 2 │ _ _ 5 │ 
                │ _ _ _ │ _ _ _ │ _ 8 9 │ 
                └───────┴───────┴───────┘
                  """
             Então o jogo deve ser inválido

        Cenário: Área menos óbvia
            Dado o tabuleiro abaixo:
            """
                ┌───────┬───────┬───────┐
                │ 5 _ _ │ _ 7 _ │ 3 _ 8 │ 
                │ 6 _ _ │ 1 4 5 │ _ _ 7 │ 
                │ _ 9 8 │ _ _ _ │ _ 6 4 │ 
                ├───────┼───────┼───────┤
                │ 8 _ _ │ 5 6 1 │ _ _ 3 │ 
                │ 4 2 6 │ 7 _ 3 │ _ 5 1 │ 
                │ _ _ _ │ 4 2 9 │ _ _ 6 │ 
                ├───────┼───────┼───────┤
                │ _ 6 _ │ _ 5 _ │ _ _ _ │ 
                │ 7 _ _ │ 8 _ 2 │ _ _ 5 │ 
                │ 3 _ 1 │ _ _ _ │ _ 8 9 │ 
                └───────┴───────┴───────┘
                  """
             Então o jogo deve ser inválido
