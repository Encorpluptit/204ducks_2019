NAME = 204ducks

all : $(NAME) sign

$(NAME):
	@cp 204ducks.py $(NAME)
	@-chmod u+x $(NAME)


prez: $(NAME) sign

go_bonus:
	$(MAKE) -C bonus/go/

clean:

fclean: clean
	@-rm -f $(NAME)

re: fclean all


sign:
	@echo ""
	@echo "Bernard Damien"
	@echo "Clette Killian"
	@echo "EPITECH, 2020"
	@echo ""

.PHONY: $(NAME) clean fclean re