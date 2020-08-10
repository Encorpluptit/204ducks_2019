NAME			=			204ducks

BONUS_PATH		=			bonus/

GO_PATH			=			$(addprefix $(BONUS_PATH), go)

GO_BIN			=			204Go

HASKELL_PATH	=			$(addprefix $(BONUS_PATH), haskell)

COV_FILE		=			.coverage

PY_CACHE		=			.pytest_cache

HEADER			=			'\033[95m'

END_HEADER		=			'\033[0m'

all : prez sign

$(NAME):
	@cp 204ducks.py $(NAME)
	@-chmod u+x $(NAME)
	@-$(MAKE) -s go_bonus


prez: $(NAME) sign

go_bonus:
	@-$(MAKE) -C $(GO_PATH)
	@-cp $(GO_PATH)/$(NAME) $(GO_BIN)


tests_run:
	@printf $(HEADER)"\nLaunching Python Unit Tests\n"$(END_HEADER)
	@-python3 -m pytest -v --cov=Ducks tests/tests.py
	@-$(MAKE) -s -C $(GO_PATH) tests_run
	@printf $(HEADER)"\nLaunching Functionnals Tests\n"$(END_HEADER)
	@-./tests/jenrik tests/test_204ducks.toml

clean:
	@-$(RM) .coverage

fclean: clean
	@-rm -f $(NAME)
	@-$(MAKE) -s -C $(GO_PATH) fclean
	@-$(RM) $(GO_BIN)
	@-$(RM) $(COV_FILE) $(PY_CACHE)

re: fclean all


sign:
	@echo ""
	@echo "Bernard Damien"
	@echo "Clette Killian"
	@echo "EPITECH, 2020"
	@echo ""

.PHONY: $(NAME) clean fclean re
