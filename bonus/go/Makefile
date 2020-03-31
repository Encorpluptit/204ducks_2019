NAME			=			204ducks

HEADER			=			'\033[95m'

END_HEADER		=			'\033[0m'

all : $(NAME)

MAIN			=			main.go

$(NAME):
	@go build -o $(NAME) $(MAIN)

tests_run:
	@printf $(HEADER)"\nLaunching Go Unit Tests\n"$(END_HEADER)
	@./gotest -v -coverpkg=./... ./tests

clean:

fclean: clean
	@-rm -f $(NAME)

re: fclean all

.PHONY: $(NAME) clean fclean re