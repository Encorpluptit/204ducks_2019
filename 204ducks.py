from sys import argv
from Ducks import Duck


def print_help(ret):
    print("USAGE:\n\t./204ducks const\nDESCRIPTION:\n\tconst\tconstant\t(between 0 and 2.5)")
    exit(ret)


def main():
    bonus = "-b" in argv or "--bonus" in argv
    help_cmd = "-h" in argv or "--help" in argv
    if help_cmd or not bonus and argv.__len__() != 2:
        return print_help(84)
    try:
        const = float(argv[1])
        if not 0 < const < 2.5:
            return print_help(84)
        pgrm = Duck(const).run()
        pgrm.print()
        if bonus:
            pgrm.bonus()
    except ValueError:
        print("Error in program")
        exit(84)


if __name__ == "__main__":
    main()
