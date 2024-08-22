
# Leitura de relatório CSV

Este projeto, serve para analisar um relatório ele vai fazer a leitura de um arquivo ".csv" com um padrão e depois ele grava essas informações no banco de dados, depois vai gerar número de algumas informações para verificar o andamento das tarefas. 




## Stack utilizada


**Back-end:** Golang, gocsv, gosql, postgreSQL


## Padrão do projeto

O relatório contém os seguintes campos, dentro do relatório ".csv"

card_id | historical_board | historical_workflow	| historical_section	| historical_column	| column_type	| historical_column_parent	| cycle_time (days)	| card_title	| size	| last_moved	| type	| owners


## Variáveis de Ambiente

Para rodar esse projeto, você vai precisar adicionar as seguintes variáveis de ambiente no seu .env

`NAME_DB`

`HOST_DB`

`PORT_DB`

`NAME_REPORT`

## Tarefas 
1. Projeto inicial 
    1.1. Abrir conexão com o banco de dados 
    1.2. Ler o arquivo csv 
    1.3. Gravar os dados no banco de dados 
2. Melhorias no projeto
    2.1. Adicionar as variáveis de ambiente para o banco de dados e arquivo de relatório; 
3. Adicionar a criação da tabela no banco de dados configurado; 
4. Melhorar a performance do projeto, pensar quando o relatório estiver com várias linhas
4.1. Processamento do relatório linha a linha 
4.2. Testar uma abordagem de transações agrupa vários inserts e depois commit um número "x"; 
5. Estruturar o projeto; 