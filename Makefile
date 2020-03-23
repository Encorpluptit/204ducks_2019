NAME = 204ducks

all : $(NAME)

$(NAME):
	@cp 204ducks.py $(NAME)
	@-chmod u+x $(NAME)

clean:

fclean: clean
	@-rm $(NAME)

re: fclean all

.PHONY: $(NAME) clean fclean re