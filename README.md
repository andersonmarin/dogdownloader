# Dog Downloader

É uma ferramenta para download de imagens de cachorros usando a API https://dog.ceo/dog-api/documentation. 
O objetivo principal é a demonstração da utilização de *goroutines* e *channels* em Go.

## Como usar

Para fazer o download das imagens na pasta padrão execute `go run main.go` ou `go build && ./dogdownloader`.

Para fazer o download das imagens em outra pasta execute `go run main.go -d /caminho/da/pasta` ou `go build && ./dogdownloader -d /caminho/da/pasta`.
