package tax

type Product struct {
	Name     string
	Price    float64
	Category string
}

type shoppingCart struct {
	Products []Product
}

/*
ICMS (Imposto sobre Circulação de Mercadorias e Serviços): 
	É um imposto estadual que incide sobre a circulação de mercadorias, 
	incluindo produtos alimentícios. As alíquotas variam de estado para estado
	 e podem ser diferentes para produtos essenciais.
PIS 
	(Programa de Integração Social): É um imposto federal que incide sobre a receita bruta das empresas. 
	O PIS pode ter alíquotas diferentes dependendo do regime de tributação da empresa
	(cumulativo ou não cumulativo).
COFINS 
	(Contribuição para o Financiamento da Seguridade Social): 
	Assim como o PIS, a COFINS é um imposto federal que incide sobre a receita bruta das empresas. 
	Também possui alíquotas que variam conforme o regime de tributação.
ISS 
	(Imposto sobre Serviços): Embora não incida diretamente sobre a venda de produtos, 
	o ISS pode ser aplicado a serviços prestados por supermercados, como entrega em domicílio.
Imposto de Importação (II): 
	Para produtos importados, o imposto de importação pode ser aplicado, 
	dependendo da categoria do produto.
Impostos Municipais: 
	Além do ISS, alguns municípios podem ter outros impostos que podem incidir 
	sobre a operação de supermercados.
Taxas e Contribuições: 
	Além dos impostos, os supermercados podem estar sujeitos a taxas e contribuições específicas, 
	como taxas de licenciamento e contribuições para entidades de classe.
*/

type Taxes map[string]func(product Product) float64

func icms(product Product) float64 {
	return product.Price * 0.18
}

func ipi(product Product) float64 {
	return product.Price * 0.10
}

// Estrutura para representar um imposto
type Tax struct {
	Name string
	Rate float64
}

func NewTax(name string, rate float64) Tax {
	return Tax{Name: name, Rate: rate}
}

func CalculateTaxes(cart shoppingCart) map[string]float64 {
	return nil
}
