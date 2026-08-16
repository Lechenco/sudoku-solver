@io

Funcionalidade: Salvamento e recuperação de arquivos

    Regra: Salvar steps em arquivo JSON
        Cenário: Salvar tabuleiro em JSON
            Dado o tabuleiro abaixo:
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
                │ 4 6 3 │ 9 5 7 │ 2 1 8 │ 
                │ 7 2 9 │ 4 8 1 │ 3 6 5 │ 
                │ 1 8 5 │ 3 6 2 │ 4 7 9 │ 
                └───────┴───────┴───────┘
                  """
             Então o jogo deve ser válido
             Então salve o jogo em "save.json"
            
        Cenário: Salvar tabuleiro e passos em JSON
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
            Então salve o jogo em "save.json"
            
    Regra: Recuperar steps de um arquivo JSON
        Cenário: Recuperar tabuleiro em JSON
            Dado o tabuleiro abaixo:
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
                │ 4 6 3 │ 9 5 7 │ 2 1 8 │ 
                │ 7 2 9 │ 4 8 1 │ 3 6 5 │ 
                │ 1 8 5 │ 3 6 2 │ 4 7 9 │ 
                └───────┴───────┴───────┘
                  """
             Então o jogo deve ser válido
             Então salve o jogo em "save.json"
             Então recupere o jogo de "save.json"
            
        Cenário: Recuperar tabuleiro e passos em JSON
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
            Então salve o jogo em "save.json"
            Então recupere o jogo de "save.json"
