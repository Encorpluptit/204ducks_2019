NAME = 204ducks

all : $(NAME)

$(NAME):
	@cp 204ducks.py $(NAME)
	@-chmod u+x $(NAME)

clean:

fclean: clean
	@-rm -f $(NAME)

re: fclean all

.PHONY: $(NAME) clean fclean re