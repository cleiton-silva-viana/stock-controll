# Método: UndoSale

## Descrição
O método `UndoSale` é responsável por desfazer uma venda previamente realizada. Ele reverte as alterações feitas no estoque e remove a venda do registro, garantindo que o sistema reflita corretamente o estado atual.

## Processo de Desfazer uma Venda

1. **Receber Requisição**
   - Endpoint: `sale/undo`

2. **Verificação de Autenticação**
   - Checar se o gerente está autenticado.
     - **Caso de Falha**: 
       - Se não autenticado, retornar `error - usuário não está autenticado para realizar esta operação`.

3. **Verificação de Autorização**
   - Checar se o usuário possui a role `manager`.
     - **Caso de Falha**: 
       - Se não possuir a role `manager`, retornar `error - usuário não autorizado a desfazer vendas`.

4. **Receber Dados da Venda**
   - Receber o ID da venda a ser desfeita.
     - **Caso de Falha**: 
       - Se o ID da venda não for fornecido, retornar `error - ID da venda não fornecido`.

5. **Busca da Venda no Repositório**
   - Buscar a venda no repositório usando o ID fornecido.
     - **Caso de Falha**: 
       - Se a venda não for encontrada, retornar `error - venda não encontrada`.
6. **Verificar se a venda pode ser fesfeita**
    - Verificamos se a venda pode ser desfeita
        - Vendas com não mais de 15 dias
        - vendas mediante produto devolvido
        - vendas que não tenham sido desfeitas antes
6. **Reverter Estoque**
   - Atualizar o estoque, adicionando de volta as quantidades dos produtos vendidos.
     - **Caso de Falha**: 
       - Se ocorrer um erro ao atualizar o estoque, retornar `error - falha ao reverter estoque`.

7. **Remover Venda do Repositório**
   - Remover a venda do registro.
     - **Caso de Falha**: 
       - Se ocorrer um erro ao remover a venda, retornar `error - falha ao remover a venda`.

8. **Retorno de Sucesso**
   - Retornar sucesso com uma mensagem de confirmação.

# Método: PromoteUser

## Descrição
O método `PromoteUser` é responsável por promover um usuário a uma nova função dentro do sistema. Ele verifica a autenticidade e autorização do gerente, valida os dados da promoção e atualiza as informações do usuário.

## Processo de Promoção de Usuário

1. **Receber Requisição**
   - Endpoint: `user/promote`

2. **Verificação de Autenticação**
   - Checar se o gerente está autenticado.
     - **Caso de Falha**: 
       - Se não autenticado, retornar `error - usuário não está autenticado para realizar esta operação`.

3. **Verificação de Autorização**
   - Checar se o gerente possui a role `manager`.
     - Checar as roles do usuário enviado via JWT.
     - **Caso de Falha**: 
       - Se não possuir a role `manager`, retornar `error - usuário não autorizado a realizar esta operação`.

4. **Receber Dados da Promoção**
   - Receber os dados da promoção (ex: ID do usuário, nova role).
     - **Caso de Falha**: 
       - Se os dados da promoção não forem fornecidos, retornar `error - dados da promoção não fornecidos`.

5. **Validação dos Dados da Promoção**
   - Validar os dados da promoção.
     - **Caso de Falha**: 
       - Se a validação falhar (ex: nova role inválida), retornar `error - dados da promoção inválidos`.

6. **Busca do Usuário no Repositório**
   - Buscar o usuário no repositório usando o ID fornecido.
     - **Caso de Falha**: 
       - Se o usuário não for encontrado, retornar `error - usuário não encontrado`.

7. **Verificação da Nova Role**
   - Verificar se a nova role é diferente da atual.
     - **Caso de Falha**: 
       - Se a nova role for igual à atual, retornar `error - o usuário já possui essa role`.

8. **Aplicação da Promoção ao Usuário**
   - Aplicar a promoção ao usuário.
     - **Caso de Falha**: 
       - Se ocorrer um erro durante a aplicação da promoção, retornar `error - falha ao aplicar a promoção`.

9. **Persistência das Alterações no Repositório**
   - Persistir as alterações no repositório.
     - **Caso de Falha**: 
       - Se ocorrer um erro, retornar `error - falha ao persistir as alterações`.

10. **Retorno de Sucesso**
    - Retornar sucesso com os dados do usuário promovido.