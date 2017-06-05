deps:
	@printf "Installing all dependencies\n\n"
	@go get -u -v github.com/githubnemo/CompileDaemon
	@printf "Installing glide dependencies\n\n"
	@glide install

install:
	@printf "Installing MySlackBot\n\n"
	go install -v github.com/disiqueira/MySlackBot/cmd/msb

local:
	@printf "Installing Pre-commit\n\n"
	@brew install pre-commit

	@printf "Installing Pre-commit hooks\n\n"
	@pre-commit install
