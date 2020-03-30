#!/usr/bin/env python3

from sys import argv
from Ducks import Duck


def print_help(ret):
    print("USAGE:\n\t./204ducks const\nDESCRIPTION:\n\tconst\tconstant\t(between 0 and 2.5)")
    exit(ret)


def main(args):
    if argv.__len__() != 2:
        return print_help(84)
    try:
        const = float(args[1])
        if not 0 < const < 2.5:
            return print_help(84)
        pgrm = Duck(const).run()
        pgrm.print()
    except ValueError:
        return print_help(84)


if __name__ == "__main__":
    help_cmd = "-h" in argv or "--help" in argv
    if help_cmd:
        print_help(84)
    main(argv)
