package tax

/*
### 1. **Identificação da Venda**

- **ID da Venda**: Um identificador único para a transação.
- **Data da Venda**: Data em que a venda foi realizada.
- **Valor Total da Venda**: O valor bruto da venda antes de impostos.
*/

/*
### 2. **Detalhes do Produto/Serviço**

- **ID do Produto/Serviço**: Identificador único do item vendido.
- **Descrição do Produto/Serviço**: Nome ou descrição do item.
- **Quantidade Vendida**: Número de unidades vendidas.
- **Preço Unitário**: Preço de cada unidade do produto/serviço.
*/

/*
### 3. **Cálculo dos Impostos**

- **Tipo de Imposto**: Identificação do tipo de imposto (ex: ICMS, IPI, ISS, etc.).
- **Alíquota**: Percentual aplicado para o cálculo do imposto.
- **Base de Cálculo**: Valor sobre o qual o imposto é calculado (pode ser o valor total da venda ou o valor de cada item).
- **Valor do Imposto**: Cálculo do imposto a ser pago (ex: Valor da Venda x Alíquota).
*/

type ICMS struct {
	value float64
	/*
		Base de Cálculo: A base de cálculo do ICMS é o valor da operação,
		que inclui o preço da mercadoria, frete, seguro e outras despesas acessórias.
		É importante observar que a legislação pode prever ajustes na base de cálculo em determinadas situações.
	*/

	aliquota int
	/*
		Alíquotas: As alíquotas do ICMS no Rio de Janeiro variam conforme
		o tipo de produto ou serviço. As alíquotas mais comuns são:

		18% para a maioria das mercadorias e serviços.
		12% para alguns produtos, como alimentos e medicamentos.
		7% para operações interestaduais, dependendo do estado de origem.
	*/

	/*
		O estado do Rio de Janeiro adota o regime de substituição tributária para diversos produtos,
		onde o ICMS é recolhido antecipadamente por um contribuinte (geralmente o fabricante ou importador)
		em vez de ser pago pelo varejista. Isso é comum em produtos como combustíveis, bebidas e eletroeletrônicos.
	*/

	/*
		Deduções e Isenções: Existem algumas situações em que o ICMS pode ser reduzido ou isento,
		como em operações com produtos da cesta básica ou em determinadas promoções.
	*/

	/*
		Recolhimento: O ICMS deve ser recolhido mensalmente,
		e as empresas devem emitir a Nota Fiscal Eletrônica (NF-e) para documentar as operações.
	*/
}
