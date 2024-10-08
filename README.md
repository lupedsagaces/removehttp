# Removehttp

RemoveHTTP é uma ferramenta simples em Go que remove os prefixos http:// e https:// de uma lista de domínios em um arquivo .txt. A ferramenta lê cada linha do arquivo e remove esses prefixos, retornando a lista formatada sem os protocolos.

## Instalação:
Você pode instalar esta ferramenta diretamente usando o comando go install:

```bash
go install github.com/lupedsagaces/removehttp@latest
```
Isso irá baixar e instalar o binário da ferramenta no seu $GOPATH/bin.

## Uso:
Com um arquivo .txt contendo uma lista de domínios, por exemplo:

```
https://chatr.google.com
http://chat2.google.com
https://developers.google.com
```
Para rodar a ferramenta, use o seguinte comando:
```bash

cat dominios.txt | removehttp
```

Isso irá remover http:// e https:// de cada linha do arquivo e retornar o resultado.

Exemplo
Entrada:

```
https://chatr.google.com
http://chat2.google.com
https://developers.google.com
```

Saída:

```
chatr.stg.google.com
chat2.google.com
developers.google.com
```
