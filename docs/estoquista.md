# Método: ReportExpiredProducts

## 1. Receber Requisição
- Endpoint: `inventory/report-expired`

## 2. Autenticação do Estoquista
- Verificar se o estoquista está autenticado:
  - **Se não autenticado**: Retornar `error - usuário não está autenticado para realizar esta operação`.

## 3. Verificação de Permissões
- Verificar se o estoquista possui a role `estoquista`:
  - Checar as roles do usuário via JWT.
  - **Se não possuir a role `estoquista`**: Retornar `error - usuário não autorizado a realizar esta operação`.

## 4. Receber Dados dos Produtos
- Receber lista de IDs de produtos vencidos.
  - **Se os dados não forem fornecidos**: Retornar `error - dados dos produtos não fornecidos`.

## 5. Busca de Produtos
- Buscar produtos no repositório usando os IDs fornecidos.
  - **Se algum produto não for encontrado**: Retornar `error - produto não encontrado`.

## 6. Verificação de Vencimento
- Verificar se os produtos estão vencidos.
  - **Se algum produto não estiver vencido**: Retornar `error - um ou mais produtos não estão vencidos`.

## 7. Geração do Relatório
- Gerar o relatório de produtos vencidos.
  - **Se ocorrer um erro durante a geração**: Retornar o erro ao cliente.

## 8. Persistência do Relatório
- Persistir o relatório no repositório.
  - **Se ocorrer um erro**: Retornar o erro ao usuário.

## 9. Retorno de Sucesso
- Retornar sucesso com os dados do relatório de produtos vencidos.

# Método: ReportShortExpiryProducts

## Descrição
O método `ReportShortExpiryProducts` é responsável por permitir que o estoquista relatar produtos com validade curta no sistema. O processo inclui a verificação da autenticidade e autorização do estoquista, a validação dos dados dos produtos e a geração e persistência do relatório.

## Processo de Relato de Produtos com Validade Curta

1. **Receber Requisição**
   - Endpoint: `inventory/reportShortExpiry`

2. **Verificação de Autenticação**
   - Checar se o estoquista está autenticado.
     - **Caso de Falha**: 
       - Se não autenticado, retornar `error - usuário não está autenticado para realizar esta operação`.

3. **Verificação de Autorização**
   - Checar se o estoquista possui a role `estoquista`.
     - Checar as roles do usuário enviado via JWT.
     - **Caso de Falha**: 
       - Se o estoquista não possuir a role `estoquista`, retornar `error - usuário não autorizado a realizar esta operação`.

4. **Receber Dados dos Produtos**
   - Receber os dados dos produtos a serem relatados (ex: lista de IDs de produtos com validade curta).
     - **Caso de Falha**: 
       - Se os dados dos produtos não forem fornecidos, retornar `error - dados dos produtos não fornecidos`.

5. **Busca de Produtos no Repositório**
   - Buscar os produtos no repositório usando os IDs fornecidos.
     - **Caso de Falha**: 
       - Se algum produto não for encontrado, retornar `error - produto não encontrado`.

6. **Verificação de Validade dos Produtos**
   - Verificar se os produtos têm validade curta.
     - **Caso de Falha**: 
       - Se algum produto não tiver validade curta, retornar `error - um ou mais produtos não têm validade curta`.

7. **Geração do Relatório**
   - Gerar o relatório de produtos com validade curta.
     - **Caso de Falha**: 
       - Se ocorrer um erro durante a geração do relatório, retornar o erro ao cliente.

8. **Persistência do Relatório no Repositório**
   - Persistir o relatório no repositório.
     - **Caso de Falha**: 
       - Se ocorrer um erro, retornar o erro ao usuário.

9. **Retorno de Sucesso**
   - Retornar sucesso com os dados do relatório de produtos com validade curta.




Estoquista faz relatório de:
	produtos vencidos
	produtos com baixa validade
	produtos com baixo volume
	produtos avariados
	produtos em falta - Será que é necessário?