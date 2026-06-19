# Solucionador de Sudoku - TODO

## Padrões

- [ ] Interface IteratorCells
- [ ] Interface Step com método TakeStep (SetValueStep, RemoveCandidatesStep)

- [ ] Adicionar métricas ao passo (tempo gasto, comparações)

## Fundamentos

- [x] Verificar se o estado atual é válido
- [x] Definir um valor em uma célula deve remover candidatos da região
- [x] Definir um valor com uma posição deve remover candidatos de todas as regiões relacionadas
- [x] Não permitir sobrescrever o valor da célula
- [x] Não permitir definir o valor da célula como zero
- [x] Deve percorrer o tabuleiro e remover todos os candidatos inválidos
- [ ] Verificar se a solução é impossível (possui uma célula vazia com zero candidatos)

## Funcionalidades Principais

- [x] Implementar analisador sintático do quebra-cabeça Sudoku
- [ ] Criar lógica de validação de restrições

## Recursos

- [ ] Validação de entrada para o formato do quebra-cabeça
- [ ] Verificação da solução
- [ ] Compatível com quebra-cabeças de diferentes tamanhos

## Testes

- [ ] Testes unitários para o solucionador
- [ ] Testes com vários níveis de dificuldade
- [x] Testes de casos extremos (quebra-cabeças inválidos, sem solução)
- [ ] Benchmarks de desempenho

## Documentação

- [ ] Exemplos de uso
- [ ] Explicação do algoritmo
- [ ] Arquivo README com instruções de configuração

## Interface/Saída

- [x] Impressão formatada da solução
- [ ] Exibição do processo de resolução passo a passo
- [ ] Exportação dos resultados para um Arquivo
