@hidden_single

Funcionalidade: Estratégia Hidden Single

    Hidden Single se refere quando um valor só é válido em uma célula de uma região.
    Todas as outras células da região não tem mais aquele valor como candidato.
    
    Dado um tabuleiro de sudoku, ao utilizar a estratégia de Hidden Single deve
    ser possível identificar a posição e valor da próxima célula a ser 
    preenchida.


    Regra: Um candidato possível
        Cenário: Hidden Single na linha
            Dado a estratégia "hidden_single"
            E o tabuleiro abaixo:
            """
                ┌───────┬───────┬───────┐
                │ 5 1 6 │ _ _ 8 │ _ 3 4 │ 
                │ 2 _ _ │ 1 _ 5 │ _ _ _ │ 
                │ 3 _ 8 │ _ _ _ │ _ 6 _ │ 
                ├───────┼───────┼───────┤
                │ _ _ _ │ _ 6 _ │ _ _ 3 │ 
                │ 6 _ _ │ _ 8 3 │ _ _ 1 │ 
                │ 4 _ _ │ _ 2 _ │ _ _ 6 │ 
                ├───────┼───────┼───────┤
                │ 8 6 _ │ _ 5 _ │ _ _ _ │ 
                │ 7 _ _ │ 8 _ 2 │ _ _ 5 │ 
                │ _ _ _ │ _ _ _ │ 2 _ _ │ 
                └───────┴───────┴───────┘
                  """
            E tabuleiro válido
            Então o próximo passo é 2 na posição (0,3)
            Então tome esse passo

        Cenário: Hidden Single na coluna
            Dado a estratégia "hidden_single"
            E o tabuleiro abaixo:
            """
                ┌───────┬───────┬───────┐
                │ 5 1 6 │ 2 7 8 │ 9 3 4 │ 
                │ 2 _ _ │ 6 3 5 │ _ _ _ │ 
                │ _ _ 8 │ _ _ _ │ _ _ _ │ 
                ├───────┼───────┼───────┤
                │ 8 _ _ │ _ 6 _ │ _ _ 3 │ 
                │ 6 _ _ │ _ _ 3 │ _ _ 1 │ 
                │ 4 3 _ │ _ _ _ │ _ _ 6 │ 
                ├───────┼───────┼───────┤
                │ _ 6 _ │ _ 5 _ │ 2 _ _ │ 
                │ 7 2 _ │ 8 _ _ │ _ _ 5 │ 
                │ _ _ 3 │ _ _ _ │ _ 8 _ │ 
                └───────┴───────┴───────┘
                  """
            E tabuleiro válido
            Então o próximo passo é 3 na posição (2,0)
            Então tome esse passo

        Cenário: Hidden Single no quadrado
            Dado a estratégia "hidden_single"
            E o tabuleiro abaixo:
            """
                ┌───────┬───────┬───────┐
                │ 5 1 6 │ 2 7 8 │ 9 3 4 │ 
                │ 2 _ _ │ 6 3 5 │ _ _ _ │ 
                │ 3 _ 8 │ 1 4 9 │ _ _ _ │ 
                ├───────┼───────┼───────┤
                │ 8 _ _ │ _ 6 _ │ _ _ 3 │ 
                │ 6 _ _ │ _ _ 3 │ _ _ 1 │ 
                │ 4 _ _ │ _ _ _ │ _ _ 6 │ 
                ├───────┼───────┼───────┤
                │ 9 6 _ │ _ 5 _ │ 2 _ _ │ 
                │ 7 2 _ │ 8 _ _ │ _ _ 5 │ 
                │ 1 _ 3 │ _ _ _ │ _ 8 _ │ 
                └───────┴───────┴───────┘
                  """
            E tabuleiro válido
            Então o próximo passo é 3 na posição (5,1)
            Então tome esse passo

    Regra: Nenhum candidato possível
        Cenário: Tabuleiro não apresenta nenhum passo com hidden_single
            Dado a estratégia "hidden_single"
            E o tabuleiro abaixo:
            """
                ┌───────┬───────┬───────┐
                │ 5 1 6 │ 2 7 8 │ 9 3 4 │ 
                │ 2 _ _ │ 6 3 5 │ _ _ 8 │ 
                │ 3 _ 8 │ _ _ _ │ _ _ 2 │ 
                ├───────┼───────┼───────┤
                │ 8 _ _ │ _ 6 _ │ _ _ 3 │ 
                │ 6 _ _ │ _ 8 _ │ 7 _ 1 │ 
                │ _ _ _ │ 5 _ _ │ 8 _ 6 │ 
                ├───────┼───────┼───────┤
                │ _ 8 _ │ _ _ _ │ 2 _ _ │ 
                │ 7 2 _ │ 8 _ _ │ _ _ 5 │ 
                │ 1 6 _ │ _ _ _ │ _ 8 _ │ 
                └───────┴───────┴───────┘
                  """
            E tabuleiro válido
            Então não foi possível determinar o próximo passo

    Regra: Finalizar tabuleiro
        Cenário: Tabuleiro faltando apenas dois números
            Dado a estratégia "hidden_single"
            E o tabuleiro abaixo:
            """
                ┌───────┬───────┬───────┐
                │ 5 1 6 │ 2 7 8 │ 9 3 4 │ 
                │ 2 9 4 │ 1 3 5 │ 6 8 7 │ 
                │ 3 7 8 │ 6 4 9 │ 1 5 2 │ 
                ├───────┼───────┼───────┤
                │ 8 4 1 │ 7 2 6 │ 5 9 3 │ 
                │ 6 5 2 │ 8 9 3 │ 7 4 1 │ 
                │ 9 3 7 │ 5 1 4 │ 8 2 6 │ 
                ├───────┼───────┼───────┤
                │ 4 6 3 │ _ 5 7 │ 2 1 8 │ 
                │ 7 2 9 │ 4 8 1 │ 3 6 5 │ 
                │ 1 8 5 │ 3 6 _ │ 4 7 9 │ 
                └───────┴───────┴───────┘
                  """
            E tabuleiro válido
            Então termine o tabuleiro
            
