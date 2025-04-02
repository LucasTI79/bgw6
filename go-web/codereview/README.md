handler (controller)
    /users

    pegar os dados da request, e retornar a resposta no formato que quem solicitou espera

service ->
   
    processar
    verificacoes de regra de negocio
        caso precise de dados para processar e validar, buscar dados na repository

repository -> pegar dados

=============================

Buscar veículos por cor e ano

BDD - Behavior Driven Development (Desenvolvimento orientado a comportamento)

Como: Usuário da API.

Quero: Listar veículos por cor e ano.

Para: Filtrar veículos com base nessas especificações.

Endpoint: GET /vehicles/color/{color}/year/{year}

Respostas:

200 OK: Retorna uma lista de veículos que atendem aos critérios.

404 Not Found: Nenhum veículo encontrado com esses critérios.
