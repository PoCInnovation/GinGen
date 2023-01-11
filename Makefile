
TARGET 		= GinGen
SRC 		= main.go

GROUP_ID 	= $$(id -g)
USER_ID 	= $$(id -u)

RED		=	\033[1;31m
GREEN	=	\033[1;32m
BLUE	=	\033[1;34m
CYAN	=	\033[1;36m
NC		=	\033[0m

.PHONY: all
all: compile

.PHONY: help
help:
	@echo -e "$(BLUE)All Rules:$(NC)"
	@perl -nle'print $& if m{^[a-zA-Z_-\d]+:.*?## .*$$}' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "$(GREEN)%-10s$(CYAN) ➜ %s\n", $$1, $$2}'

.PHONY: compile
compile: ## Compile the project
	@go build -o $(TARGET) $(SRC)

.PHONY: docker
docker: ## Builds a docker image from source
	@docker build -t gingen \
		--build-arg USER_ID=$(USER_ID) \
		--build-arg GROUP_ID=$(GROUP_ID) \
		.

.PHONY: clean
clean: ## Cleans up the project
	@rm -f $(TARGET)
