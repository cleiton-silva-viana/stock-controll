# Método: PromoteUser

## Descrição
O método `PromoteUser` é responsável por promover um usuário a uma nova função dentro do sistema. Ele verifica a autenticidade e autorização do usuário do RH, valida os dados da promoção e atualiza as informações do usuário.

## Processo de Promoção de Usuário

1. **Receber Requisição**
   - Endpoint: `user/promote`

2. **Verificação de Autenticação**
   - Checar se o usuário do RH está autenticado.
     - **Caso de Falha**: 
       - Se não autenticado, retornar `error - usuário não está autenticado para realizar esta operação`.

3. **Verificação de Autorização**
   - Checar se o usuário possui a role `HR`.
     - Checar as roles do usuário enviado via JWT.
     - **Caso de Falha**: 
       - Se não possuir a role `HR`, retornar `error - usuário não autorizado a realizar esta operação`.

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



# Método: RevokeUser

## Descrição
O método `RevokeUser` é responsável por revogar o cargo e as permissões de um usuário dentro do sistema. Ele verifica a autenticidade e autorização do usuário do RH, valida os dados da revogação e atualiza as informações do usuário.

## Processo de Revogação de Usuário

1. **Receber Requisição**
   - Endpoint: `user/revoke`

2. **Verificação de Autenticação**
   - Checar se o usuário do RH está autenticado.
     - **Caso de Falha**: 
       - Se não autenticado, retornar `error - usuário não está autenticado para realizar esta operação`.

3. **Verificação de Autorização**
   - Checar se o usuário possui a role `HR`.
     - Checar as roles do usuário enviado via JWT.
     - **Caso de Falha**: 
       - Se não possuir a role `HR`, retornar `error - usuário não autorizado a realizar esta operação`.

4. **Receber Dados da Revogação**
   - Receber os dados da revogação (ex: ID do usuário, role a ser revogada).
     - **Caso de Falha**: 
       - Se os dados da revogação não forem fornecidos, retornar `error - dados da revogação não fornecidos`.

5. **Validação dos Dados da Revogação**
   - Validar os dados da revogação.
     - **Caso de Falha**: 
       - Se a validação falhar (ex: role inválida), retornar `error - dados da revogação inválidos`.

6. **Busca do Usuário no Repositório**
   - Buscar o usuário no repositório usando o ID fornecido.
     - **Caso de Falha**: 
       - Se o usuário não for encontrado, retornar `error - usuário não encontrado`.

7. **Verificação da Role Atual**
   - Verificar se a role a ser revogada é diferente da atual.
     - **Caso de Falha**: 
       - Se a role a ser revogada não estiver atribuída ao usuário, retornar `error - o usuário não possui essa role`.

8. **Aplicação da Revogação ao Usuário**
   - Aplicar a revogação ao usuário.
     - **Caso de Falha**: 
       - Se ocorrer um erro durante a aplicação da revogação, retornar `error - falha ao aplicar a revogação`.

9. **Persistência das Alterações no Repositório**
   - Persistir as alterações no repositório.
     - **Caso de Falha**: 
       - Se ocorrer um erro, retornar `error - falha ao persistir as alterações`.

10. **Retorno de Sucesso**
    - Retornar sucesso com os dados do usuário após a revogação.
