NAME			=			204ducks

BONUS_PATH		=			bonus/

GO_PATH			=			$(addprefix $(BONUS_PATH), go)

HASKELL_PATH	=			$(addprefix $(BONUS_PATH), haskell)

HEADER			=			'\033[95m'

END_HEADER		=			'\033[0m'

all : $(NAME) sign

$(NAME):
	@cp 204ducks.py $(NAME)
	@-chmod u+x $(NAME)
	@-$(MAKE) -s go_bonus
	@-$(MAKE) -C haskell_bonus



prez: $(NAME) go_bonus sign

go_bonus:
	@-$(MAKE) -C $(GO_PATH)

haskell_bonus:
	@-$(MAKE) -C $(HASKELL_PATH)


tests_run:
	@printf $(HEADER)"\nLaunching Python Unit Tests\n"$(END_HEADER)
	@-python3 -m pytest -v --cov=Ducks tests/tests.py
	@-$(MAKE) -s -C $(GO_PATH) tests_run
	@printf $(HEADER)"\nLaunching Functionnals Tests\n"$(END_HEADER)
	@-./tests/jenrik tests/test_204ducks.toml

clean:

fclean: clean
	@-rm -f $(NAME)
	@-$(MAKE) -s -C $(GO_PATH) fclean

re: fclean all


sign:
	@echo ""
	@echo "Bernard Damien"
	@echo "Clette Killian"
	@echo "EPITECH, 2020"
	@echo ""

.PHONY: $(NAME) clean fclean re
