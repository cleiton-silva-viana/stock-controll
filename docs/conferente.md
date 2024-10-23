# Método: InspectDelivery

## Descrição
O método `InspectDelivery` é responsável por permitir que o conferente inspecione uma entrega, podendo optar por recebê-la ou rejeitá-la. O processo inclui a verificação da autenticidade da operação, a validação dos dados da entrega, a verificação dos produtos e a atualização do status da entrega.

## Processo de Inspeção de Entrega

1. **Receber Requisição**
   - Endpoint: `delivery/inspect`

2. **Verificação de Autenticação**
   - Checar se o conferente está autenticado.
     - **Caso de Falha**: 
       - Se não autenticado, retornar `error - conferente não está autenticado para realizar esta operação`.

3. **Receber Dados da Entrega**
   - Receber os dados da entrega (ex: ID da entrega, status desejado - recebido ou rejeitado).
     - **Caso de Falha**: 
       - Se os dados da entrega não forem fornecidos, retornar `error - dados da entrega não fornecidos`.

4. **Verificação de Registro da Entrega**
   - Verificar se a entrega está previamente registrada no sistema de entregas a receber.
     - **Caso de Falha**: 
       - Se a entrega não estiver registrada, retornar `error - entrega não registrada no sistema`.

5. **Validação dos Dados da Entrega**
   - Validar os dados da entrega.
     - **Caso de Falha**: 
       - Se a validação falhar (ex: status inválido), retornar `error - dados da entrega inválidos`.

6. **Busca da Entrega no Repositório**
   - Buscar a entrega no repositório usando o ID fornecido.
     - **Caso de Falha**: 
       - Se a entrega não for encontrada, retornar `error - entrega não encontrada`.

7. **Verificação da Conformidade dos Produtos**
   - Validar se os produtos da entrega estão conforme os requisitos (ex: quantidade, qualidade).
     - **Caso de Falha**: 
       - Se os produtos não estiverem conformes, retornar `error - produtos não estão conformes para aceitar a entrega`.

8. **Verificação do Status Atual**
   - Verificar se o status desejado é diferente do atual.
     - **Caso de Falha**: 
       - Se o status desejado for igual ao atual, retornar `error - a entrega já possui esse status`.

9. **Aplicação da Decisão sobre a Entrega**
   - Aplicar a decisão (receber ou rejeitar) à entrega.
     - **Caso de Falha**: 
       - Se ocorrer um erro durante a aplicação da decisão, retornar `error - falha ao aplicar a decisão sobre a entrega`.

10. **Persistência das Alterações no Repositório**
    - Persistir as alterações no repositório.
      - **Caso de Falha**: 
        - Se ocorrer um erro, retornar `error - falha ao persistir as alterações`.

11. **Retorno de Sucesso**
    - Retornar sucesso com os dados da entrega após a inspeção.


# Método: TransferGoods

## Descrição
O método `TransferGoods` é responsável por permitir que o conferente realize a transferência de mercadorias. O processo inclui a verificação da autenticidade da operação, a validação dos dados da transferência e a conferência da integridade, quantidade e validade das mercadorias.

## Processo de Transferência de Mercadorias

1. **Receber Requisição**
   - Endpoint: `goods/transfer`

2. **Verificação de Autenticação**
   - Checar se o conferente está autenticado.
     - **Caso de Falha**: 
       - Se não autenticado, retornar `error - conferente não está autenticado para realizar esta operação`.

3. **Receber Dados da Transferência**
   - Receber os dados da transferência (ex: ID da transferência, lista de mercadorias).
     - **Caso de Falha**: 
       - Se os dados da transferência não forem fornecidos, retornar `error - dados da transferência não fornecidos`.

4. **Verificação de Registro da Transferência**
   - Verificar se a transferência está previamente registrada no sistema.
     - **Caso de Falha**: 
       - Se a transferência não estiver registrada, retornar `error - transferência não registrada no sistema`.

5. **Validação dos Dados da Transferência**
   - Validar os dados da transferência.
     - **Caso de Falha**: 
       - Se a validação falhar (ex: dados inválidos), retornar `error - dados da transferência inválidos`.

6. **Busca das Mercadorias no Repositório**
   - Buscar as mercadorias no repositório usando os IDs fornecidos.
     - **Caso de Falha**: 
       - Se alguma mercadoria não for encontrada, retornar `error - mercadoria não encontrada`.

7. **Conferência da Integridade, Quantidade e Validade**
   - Conferir se as mercadorias estão em conformidade com os requisitos (ex: integridade, quantidade, validade).
     - **Caso de Falha**: 
       - Se alguma mercadoria não estiver conforme, retornar `error - mercadorias não estão em conformidade para transferência`.

8. **Aplicação da Transferência**
   - Aplicar a transferência das mercadorias.
     - **Caso de Falha**: 
       - Se ocorrer um erro durante a aplicação da transferência, retornar `error - falha ao aplicar a transferência`.

9. **Persistência das Alterações no Repositório**
   - Persistir as alterações no repositório.
     - **Caso de Falha**: 
       - Se ocorrer um erro, retornar `error - falha ao persistir as alterações`.

10. **Retorno de Sucesso**
    - Retornar sucesso com os dados da transferência realizada.

# Método: ValidateDisposalOrExchange

## Descrição
O método `ValidateDisposalOrExchange` é responsável por permitir que o conferente valide o descarte ou a troca de mercadorias. O processo inclui a verificação da autenticidade da operação, a validação dos dados da transação e a conferência das razões para o descarte ou troca (validade ou avaria).

## Processo de Validação de Descarte ou Troca de Mercadorias

1. **Receber Requisição**
   - Endpoint: `goods/validateDisposalOrExchange`

2. **Verificação de Autenticação**
   - Checar se o conferente está autenticado.
     - **Caso de Falha**: 
       - Se não autenticado, retornar `error - conferente não está autenticado para realizar esta operação`.

3. **Receber Dados da Transação**
   - Receber os dados da transação (ex: ID da mercadoria, motivo - validade ou avaria, tipo - descarte ou troca).
     - **Caso de Falha**: 
       - Se os dados da transação não forem fornecidos, retornar `error - dados da transação não fornecidos`.

4. **Verificação de Registro da Mercadoria**
   - Verificar se a mercadoria está previamente registrada no sistema.
     - **Caso de Falha**: 
       - Se a mercadoria não estiver registrada, retornar `error - mercadoria não registrada no sistema`.

5. **Validação dos Dados da Transação**
   - Validar os dados da transação.
     - **Caso de Falha**: 
       - Se a validação falhar (ex: dados inválidos), retornar `error - dados da transação inválidos`.

6. **Busca da Mercadoria no Repositório**
   - Buscar a mercadoria no repositório usando o ID fornecido.
     - **Caso de Falha**: 
       - Se a mercadoria não for encontrada, retornar `error - mercadoria não encontrada`.

7. **Conferência do Motivo de Descarte ou Troca**
   - Conferir se o motivo para o descarte ou troca é válido (ex: validade expirada ou avaria).
     - **Caso de Falha**: 
       - Se o motivo não for válido, retornar `error - motivo para descarte ou troca inválido`.

8. **Aplicação do Descarte ou Troca**
   - Aplicar o descarte ou troca da mercadoria.
     - **Caso de Falha**: 
       - Se ocorrer um erro durante a aplicação da transação, retornar `error - falha ao aplicar o descarte ou troca`.

9. **Persistência das Alterações no Repositório**
   - Persistir as alterações no repositório.
     - **Caso de Falha**: 
       - Se ocorrer um erro, retornar `error - falha ao persistir as alterações`.

10. **Retorno de Sucesso**
    - Retornar sucesso com os dados da transação de descarte ou troca realizada.
