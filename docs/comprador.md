# Método: AddProduct

## Descrição
O método `AddProduct` é responsável por permitir que o comprador adicione novos produtos ao sistema. O processo inclui a verificação da autenticidade da operação, a validação dos dados do produto e a persistência das informações no repositório.

## Processo de Adição de Novos Produtos

1. **Receber Requisição**
   - Endpoint: `product/add`

2. **Verificação de Autenticação**
   - Checar se o comprador está autenticado.
     - **Caso de Falha**: 
       - Se não autenticado, retornar `error - comprador não está autenticado para realizar esta operação`.

3. **Receber Dados do Produto**
   - Receber os dados do produto (ex: nome, descrição, preço de venda, desconto, categoria).
     - **Caso de Falha**: 
       - Se os dados do produto não forem fornecidos, retornar `error - dados do produto não fornecidos`.

4. **Validação dos Dados do Produto**
   - Validar os dados do produto.
     - **Caso de Falha**: 
       - Se a validação falhar (ex: preço inválido, desconto inválido), retornar `error - dados do produto inválidos`.

5. **Verificação de Produto Duplicado**
   - Verificar se já existe um produto com o mesmo nome ou identificador no sistema.
     - **Caso de Falha**: 
       - Se o produto já existir, retornar `error - produto já cadastrado no sistema`.

6. **Persistência do Produto no Repositório**
   - Persistir as informações do novo produto no repositório.
     - **Caso de Falha**: 
       - Se ocorrer um erro durante a persistência, retornar `error - falha ao adicionar o produto ao sistema`.

7. **Retorno de Sucesso**
   - Retornar sucesso com os dados do produto adicionado.

# Método: UpdateProductDescription

## Descrição
O método `UpdateProductDescription` é responsável por permitir que o comprador atualize a descrição de um produto existente no sistema. O processo inclui a verificação da autenticidade da operação, a validação dos dados do produto e a persistência das alterações no repositório.

## Processo de Atualização da Descrição do Produto

1. **Receber Requisição**
   - Endpoint: `product/updateDescription`

2. **Verificação de Autenticação**
   - Checar se o comprador está autenticado.
     - **Caso de Falha**: 
       - Se não autenticado, retornar `error - comprador não está autenticado para realizar esta operação`.

3. **Receber Dados da Atualização**
   - Receber os dados da atualização (ex: ID do produto, nova descrição).
     - **Caso de Falha**: 
       - Se os dados da atualização não forem fornecidos, retornar `error - dados da atualização não fornecidos`.

4. **Verificação de Registro do Produto**
   - Verificar se o produto está previamente registrado no sistema.
     - **Caso de Falha**: 
       - Se o produto não estiver registrado, retornar `error - produto não encontrado no sistema`.

5. **Validação da Nova Descrição**
   - Validar a nova descrição do produto.
     - **Caso de Falha**: 
       - Se a nova descrição for inválida (ex: muito curta ou vazia), retornar `error - descrição inválida`.

6. **Atualização da Descrição no Repositório**
   - Atualizar a descrição do produto no repositório.
     - **Caso de Falha**: 
       - Se ocorrer um erro durante a atualização, retornar `error - falha ao atualizar a descrição do produto`.

7. **Retorno de Sucesso**
   - Retornar sucesso com os dados do produto atualizado.

# Método: DeactivateProduct

## Descrição
O método `DeactivateProduct` é responsável por permitir que o comprador desative o status de venda de um produto existente no sistema. O processo inclui a verificação da autenticidade da operação, a validação dos dados do produto e a persistência das alterações no repositório.

## Processo de Desativação do Status de Venda do Produto

1. **Receber Requisição**
   - Endpoint: `product/deactivate`

2. **Verificação de Autenticação**
   - Checar se o comprador está autenticado.
     - **Caso de Falha**: 
       - Se não autenticado, retornar `error - comprador não está autenticado para realizar esta operação`.

3. **Receber Dados da Desativação**
   - Receber os dados da desativação (ex: ID do produto).
     - **Caso de Falha**: 
       - Se os dados da desativação não forem fornecidos, retornar `error - dados da desativação não fornecidos`.

4. **Verificação de Registro do Produto**
   - Verificar se o produto está previamente registrado no sistema.
     - **Caso de Falha**: 
       - Se o produto não estiver registrado, retornar `error - produto não encontrado no sistema`.

5. **Verificação do Status Atual**
   - Verificar se o produto já está desativado.
     - **Caso de Falha**: 
       - Se o produto já estiver desativado, retornar `error - o produto já está desativado`.

6. **Desativação do Produto no Repositório**
   - Atualizar o status do produto para desativado no repositório.
     - **Caso de Falha**: 
       - Se ocorrer um erro durante a desativação, retornar `error - falha ao desativar o produto`.

7. **Retorno de Sucesso**
   - Retornar sucesso com os dados do produto desativado.

# Método: PlaceOrder

## Descrição
O método `PlaceOrder` é responsável por permitir que o comprador realize pedidos de compra de produtos para abastecer o estoque. O processo inclui a verificação da autenticidade da operação, a validação dos dados do pedido e a geração de uma nota pendente no sistema.

## Processo de Realização de Pedidos de Compra

1. **Receber Requisição**
   - Endpoint: `order/place`

2. **Verificação de Autenticação**
   - Checar se o comprador está autenticado.
     - **Caso de Falha**: 
       - Se não autenticado, retornar `error - comprador não está autenticado para realizar esta operação`.

3. **Receber Dados do Pedido**
   - Receber os dados do pedido (ex: lista de produtos, quantidades, fornecedor).
     - **Caso de Falha**: 
       - Se os dados do pedido não forem fornecidos, retornar `error - dados do pedido não fornecidos`.

4. **Validação dos Dados do Pedido**
   - Validar os dados do pedido.
     - **Caso de Falha**: 
       - Se a validação falhar (ex: produtos inválidos, quantidades negativas), retornar `error - dados do pedido inválidos`.

5. **Verificação de Estoque**
   - Verificar se os produtos estão disponíveis para compra.
     - **Caso de Falha**: 
       - Se algum produto não estiver disponível, retornar `error - produto(s) não disponível(is) para compra`.

6. **Geração da Nota de Compra**
   - Gerar uma nota de compra pendente no sistema com os detalhes do pedido.
     - **Caso de Falha**: 
       - Se ocorrer um erro durante a geração da nota, retornar `error - falha ao gerar a nota de compra`.

7. **Persistência do Pedido no Repositório**
   - Persistir as informações do pedido no repositório.
     - **Caso de Falha**: 
       - Se ocorrer um erro durante a persistência, retornar `error - falha ao registrar o pedido no sistema`.

8. **Retorno de Sucesso**
   - Retornar sucesso com os dados da nota de compra gerada e do pedido realizado.


# Método: DeleteProduct

## Descrição
O método `DeleteProduct` é responsável por permitir que o comprador delete um produto do sistema. O processo inclui a verificação da autenticidade da operação, a validação dos dados do produto e a verificação de que o produto não possui registros no estoque.

## Processo de Deleção de Produto

1. **Receber Requisição**
   - Endpoint: `product/delete`

2. **Verificação de Autenticação**
   - Checar se o comprador está autenticado.
     - **Caso de Falha**: 
       - Se não autenticado, retornar `error - comprador não está autenticado para realizar esta operação`.

3. **Receber Dados da Deleção**
   - Receber os dados da deleção (ex: ID do produto).
     - **Caso de Falha**: 
       - Se os dados da deleção não forem fornecidos, retornar `error - dados da deleção não fornecidos`.

4. **Verificação de Registro do Produto**
   - Verificar se o produto está previamente registrado no sistema.
     - **Caso de Falha**: 
       - Se o produto não estiver registrado, retornar `error - produto não encontrado no sistema`.

5. **Verificação de Registros no Estoque**
   - Verificar se o produto possui registros no estoque (ex: se já foi adicionado ao estoque físico).
     - **Caso de Falha**: 
       - Se o produto tiver registros no estoque, retornar `error - não é possível deletar o produto, pois ele possui registros no estoque`.

6. **Deleção do Produto no Repositório**
   - Deletar o produto do repositório.
     - **Caso de Falha**: 
       - Se ocorrer um erro durante a deleção, retornar `error - falha ao deletar o produto`.

7. **Retorno de Sucesso**
   - Retornar sucesso com a confirmação de que o produto foi deletado.






docker container run --rm --network=host -e SONAR_HOST_URL="http://localhost:9000" -v "./internal:/usr/src" sonarsource/sonar-scanner-cli -Dsonar.projectKey=ec -Dsonar.sources=. -Dsonar.host.url=http://localhost:9000 -Dsonar.login=sqp_c1bbb5939e19fbb5518724e6d15a7c0ba31b20f6
   