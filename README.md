# Recognize golang in terminal
export PATH=$PATH:/usr/local/go/bin
go version

go install github.com/golang/mock/mockgen@v1.6.0

# Link poder executar mockgen terminal
https://stackoverflow.com/questions/72429463/zsh-command-not-found-mockgen-golang-1-18-macos-monterrey

# Criando arquivo sqlite no Linux
sqlite3 sqlite.db 

# Rodando main.go
go run main.go

# Rodando Testes
go test ./... -v

# Cobra
go get -u github.com/spf13/cobra@latest
go install github.com/spf13/cobra-cli@latest

cobra-cli init

# Atualizando/removendo dependencias do modulo
go mod tidy