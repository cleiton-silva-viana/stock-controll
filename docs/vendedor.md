# Método: venderProduto

## Descrição
O método `venderProduto` é responsável por realizar a venda de um produto específico. Ele verifica a disponibilidade do produto em estoque, aplica descontos se necessário e processa a transação, atualizando as informações no sistema.

## Processo de Venda

1. **Receber Requisição**
   - Endpoint: `sale/create`

2. **Verificação de Autenticação**
   - Checar se o vendedor está autenticado.
     - **Caso de Falha**: 
       - Se não autenticado, retornar `error - usuário não está autenticado para realizar esta operação`.

3. **Verificação de Autorização**
   - Checar se o vendedor possui a role `seller`.
     - Checar as roles do usuário enviado via JWT.
     - **Caso de Falha**: 
       - Se não possuir a role `seller`, retornar `error - usuário não autorizado a realizar esta operação`.

4. **Receber Dados da Venda**
   - Receber os dados da venda (ex: ID do produto, quantidade, desconto, etc.).
     - **Caso de Falha**: 
       - Se os dados da venda não forem fornecidos, retornar `error - dados da venda não fornecidos`.

5. **Validação dos Dados da Venda**
   - Validar os dados da venda.
     - **Caso de Falha**: 
       - Se a validação falhar (ex: produto não existe, quantidade inválida), retornar `error - dados da venda inválidos`.

6. **Verificação de Disponibilidade do Produto**
   - Verificar se a quantidade solicitada está disponível no estoque.
     - **Caso de Falha**: 
       - Se a quantidade solicitada não estiver disponível, retornar `error - quantidade de produto não disponível`.

7. **Aplicação de Desconto (se houver)**
   - Aplicar o desconto, se fornecido.
     - **Caso de Falha**: 
       - Se o desconto for inválido, retornar `error - desconto inválido`.

8. **Criação da Instância de Venda**
   - Criar uma instância de `Sale` com os dados fornecidos.
     - **Caso de Falha**: 
       - Se ocorrer um erro na criação da instância, retornar `error - falha ao criar a instância de venda`.

9. **Persistência da Venda no Repositório**
   - Persistir a venda no repositório.
     - **Caso de Falha**: 
       - Se ocorrer um erro, retornar `error - falha ao persistir a venda`.

10. **Atualização do Estoque**
    - Atualizar a quantidade de produtos disponíveis no estoque.
      - **Caso de Falha**: 
        - Se ocorrer um erro, retornar `error - falha ao atualizar o estoque`.

11. **Retorno de Sucesso**
    - Retornar sucesso com os dados da venda realizada.
# Método: applyDiscount

## Descrição
O método `applyDiscount` é responsável por aplicar um desconto a uma venda existente. Ele verifica a validade do desconto e atualiza os dados da venda no sistema.

## Processo de Aplicação de Desconto

1. **Receber Requisição**
   - Endpoint: `sale/apply-discount`

2. **Verificação de Autenticação**
   - Checar se o vendedor está autenticado.
     - **Caso de Falha**: 
       - Se não autenticado, retornar `error - usuário não está autenticado para realizar esta operação`.

3. **Verificação de Autorização**
   - Checar se o vendedor possui a role `seller`.
     - Checar as roles do usuário enviado via JWT.
     - **Caso de Falha**: 
       - Se não possuir a role `seller`, retornar `error - usuário não autorizado a realizar esta operação`.

4. **Receber Dados do Desconto**
   - Receber os dados do desconto (ex: ID da venda, valor do desconto, etc.).
     - **Caso de Falha**: 
       - Se os dados do desconto não forem fornecidos, retornar `error - dados do desconto não fornecidos`.

5. **Validação dos Dados do Desconto**
   - Validar os dados do desconto.
     - **Caso de Falha**: 
       - Se a validação falhar (ex: valor do desconto inválido), retornar `error - dados do desconto inválidos`.

6. **Busca da Venda no Repositório**
   - Buscar a venda no repositório usando o ID fornecido.
     - **Caso de Falha**: 
       - Se a venda não for encontrada, retornar `error - venda não encontrada`.

7. **Verificação da Aplicabilidade do Desconto**
   - Verificar se o desconto pode ser aplicado à venda.
     - **Caso de Falha**: 
       - Se o desconto for maior que o valor total da venda, retornar `error - desconto maior que o valor da venda`.

8. **Aplicação do Desconto à Venda**
   - Aplicar o desconto à venda.
     - **Caso de Falha**: 
       - Se ocorrer um erro durante a aplicação do desconto, retornar `error - falha ao aplicar desconto`.

9. **Persistência das Alterações no Repositório**
   - Persistir as alterações no repositório.
     - **Caso de Falha**: 
       - Se ocorrer um erro, retornar `error - falha ao persistir as alterações`.

10. **Retorno de Sucesso**
    - Retornar sucesso com os dados da venda atualizada.
# Método: GetSalesHistory

## Descrição
O método `GetSalesHistory` é responsável por recuperar o histórico de vendas de um vendedor autenticado. Ele retorna uma lista de vendas realizadas, incluindo detalhes como ID da venda, produtos vendidos, quantidades e valores totais.

## Processo de Recuperação do Histórico de Vendas

1. **Receber Requisição**
   - Endpoint: `sales/history`

2. **Verificação de Autenticação**
   - Checar se o vendedor está autenticado.
     - **Caso de Falha**: 
       - Se não autenticado, retornar `error - usuário não está autenticado para realizar esta operação`.

3. **Recuperação do Histórico de Vendas**
   - Buscar o histórico de vendas do vendedor no repositório.
     - **Caso de Falha**: 
       - Se ocorrer um erro ao buscar o histórico, retornar `error - falha ao recuperar histórico de vendas`.

4. **Formatação dos Dados**
   - Formatar os dados das vendas para uma apresentação clara e concisa.

5. **Retorno de Sucesso**
   - Retornar sucesso com a lista de vendas realizadas pelo vendedor.