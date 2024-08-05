
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

