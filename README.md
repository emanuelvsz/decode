# Decode Library

A Biblioteca Encode é um projeto que fornece uma coleção de funções para converter diferentes tipos de códigos para o alfabeto e vice-versa. Ela permite que você converta códigos como Código Morse, Alfabeto, Binário, Números e outros, facilitando a comunicação e a manipulação desses dados em diferentes formatos.

## Instalação

Para utilizar a Biblioteca de Conversão de Códigos em Go, siga as etapas abaixo:

1. Certifique-se de ter o Go instalado em seu sistema. Para instruções de instalação, consulte o [site oficial do Go](https://golang.org/doc/install).

2. Crie um novo diretório para o seu projeto.

3. No diretório do projeto, execute o seguinte comando para inicializar um novo módulo Go:

```shell
go mod init nome-do-modulo
```

4. Baixe a biblioteca de conversão de códigos usando o comando `go get`:

```shell
go get https://github.com/emanuelvsz/decode
```

## Uso

Após a instalação, você pode importar o pacote da biblioteca em seu projeto Go e utilizar as funções de conversão disponíveis.

```go
import "github.com/seu-usuario/decode"
```

### Conversão de Código Morse para Alfabeto

```go
texto := biblioteca_conversao_codigos.MorseParaAlfabeto("... --- ...")
fmt.Println(texto) // Saída: "SOS"
```

### Conversão de Alfabeto para Código Morse

```go
morse := biblioteca_conversao_codigos.AlfabetoParaMorse("HELLO")
fmt.Println(morse) // Saída: ".... . .-.. .-.. ---"
```

### Conversão de Binário para Números

```go
numero := biblioteca_conversao_codigos.BinarioParaNumeros("101010")
fmt.Println(numero) // Saída: 42
```

### Conversão de Números para Binário

```go
binario := biblioteca_conversao_codigos.NumerosParaBinario(42)
fmt.Println(binario) // Saída: "101010"
```

## Adicionando Novas Funcionalidades

A Biblioteca de Conversão de Códigos em Go foi projetada para ser facilmente extensível, permitindo a adição de novas funcionalidades de conversão. Para adicionar uma nova funcionalidade, siga as etapas abaixo:

1. Abra o arquivo `biblioteca_conversao_codigos.go` em um editor de texto.

2. Defina uma nova função para a conversão desejada, seguindo o mesmo padrão das funções existentes na biblioteca. Certifique-se de fornecer uma descrição adequada, parâmetros de entrada e tipo de retorno.

3. Salve o arquivo com a nova função adicionada.

Agora você pode importar o pacote da biblioteca novamente e usar a nova funcionalidade em seu projeto.

## Considerações Finais

A Biblioteca de Conversão de Códigos em Go fornece uma maneira conveniente de converter diferentes tipos de códigos para o alfabeto e vice-versa. A documentação detalhada ajuda a entender como utilizar as funções existentes e também permite

 a expansão do projeto com novas funcionalidades. Sinta-se à vontade para contribuir com a biblioteca adicionando novas conversões e melhorias.

Caso precise de mais informações ou suporte adicional, não hesite em entrar em contato com a equipe de desenvolvimento.
