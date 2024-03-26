Aplicação de Armazenamento de Logs em Arquitetura Hexagonal
====================================================================

Esta aplicação em Golang implementa uma arquitetura hexagonal para armazenar logs de atividade em um object storage Parquet usando o cliente DuckDB. A comunicação é realizada através de protocolo REST, com rotas definidas para inserção e consulta de logs.

Funcionalidades
---------------

### 1\. Inserir Novos Logs de Atividade

*   **Método HTTP:** POST
*   **Rota:** `/v1/logs`
*   **Descrição:** Esta rota permite inserir novos logs de atividade na aplicação.

### 2\. Consultar Logs de Atividades com Filtros

*   **Método HTTP:** GET
*   **Rota:** `/v1/logs`
*   **Descrição:** Esta rota permite consultar logs de atividades inserindo filtros específicos.

### JSON de Filtro para Consulta


```
{
    "query": {
        "expressions": {
            "fieldComparisons": [
                {
                    "field": "tableName",
                    "operator": "=",
                    "value": "Beneficiario"
                }
            ]
        }
    },
    "page": 1,
    "pageSize": 1
}
```

Resposta da Consulta GET
------------------------

A resposta do GET é um array de objetos JSON, onde cada objeto representa um log de atividade encontrado, com os seguintes campos:

*   **id:** Identificador único do log.
*   **name:** Nome associado ao log (opcional).
*   **dataObject:** Objeto de dados associado ao log.
*   **description:** Descrição do log (opcional).
*   **email:** Email relacionado ao log.
*   **eventDate:** Data e hora do evento no formato "DD/MM/AAAA HH:MM:SS".
*   **eventTicks:** Valor de tempo do evento em ticks.
*   **idClient:** Identificador do cliente.
*   **idProfile:** Identificador do perfil.
*   **identifier:** Identificador único associado ao evento.
*   **operation:** Operação realizada (ex: Alteração, Inserção, Exclusão).
*   **processName:** Nome do processo relacionado ao evento.
*   **tableName:** Nome da tabela associada ao evento.
*   **userId:** Identificador do usuário.
*   **username:** Nome de usuário associado ao evento.

Exemplo de resposta:


```
[
    {
        "id": "f5fd3b07-c6a8-404e-93e0-342734ffb9ba",
        "name": "",
        "dataObject": "{any object here}",
        "description": "",
        "email": "email-here",
        "eventDate": "26/02/2024 22:41:12",
        "eventTicks": 638445840723170700,
        "idClient": 5,
        "idProfile": 1,
        "identifier": 10707,
        "operation": "Alteração",
        "processName": "w3wp",
        "tableName": "Beneficiario",
        "userId": 10707,
        "username": "username-example-here"
    }
]

```

Executando a Aplicação
----------------------

Para executar a aplicação, siga os passos abaixo:

1.  Clone o repositório do projeto.
2.  Certifique-se de ter o Golang instalado em seu sistema.
3.  Instale as dependências necessárias.
4.  Execute o comando para iniciar a aplicação.

Contribuindo
------------

Contribuições são bem-vindas! Sinta-se à vontade para enviar pull requests ou abrir issues para reportar bugs ou propor novas funcionalidades.

Licença
-------

Este projeto está licenciado sob a Licença MIT.

