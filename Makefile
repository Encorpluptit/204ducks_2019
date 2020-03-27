NAME			=			204ducks

BONUS_PATH		=			bonus/

GO_PATH			=			$(addprefix $(BONUS_PATH), go)

HASKELL_PATH	=			$(addprefix $(BONUS_PATH), haskell)

all : $(NAME) sign

$(NAME):
	@cp 204ducks.py $(NAME)
	@-chmod u+x $(NAME)


prez: $(NAME) go_bonus sign

go_bonus:
	@$(MAKE) -C $(GO_PATH)

haskell_bonus:
	@$(MAKE) -C $(HASKELL_PATH)

clean:
	@$(MAKE)

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