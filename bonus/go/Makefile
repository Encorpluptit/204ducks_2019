NAME = 204ducks

all : $(NAME)

MAIN			=			main.go

$(NAME):
	go build $(MAIN)

clean:

fclean: clean
	@-rm -f $(NAME)

re: fclean all

.PHONY: $(NAME) clean fclean re