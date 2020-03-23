#!/usr/bin/env python3

import sys
from math import exp, sqrt
from decimal import *

INTERVAL = 1000


def proba_density(a, t):
    return a * exp(-t) + (4 - 3 * a) * exp(-2 * t) + (2 * a - 4) * exp(-4 * t)


def time_back(a, p):
    res = 0
    for t in frange(0, 1000, 0.01):
        res += proba_density(a, t)
        if res >= p:
            return t
    raise ValueError


def percent_back(const: float, time_snd: int):
    return sum(proba_density(const, i / INTERVAL) for i in range(time_snd * INTERVAL)) / 10


def variance(esp: float, a: float, t: float):
    return pow((t - esp), 2) * (proba_density(a, t) / 10)


def esperance(const: float, interval):
    return sum(time * (proba_density(const, time) / 10) for time in interval) / interval[-1]


def standard_deviation(const: float, esp: float, interval):
    return sqrt(sum(variance(esp, const, time) for time in interval) / interval[-1])


def frange(start: float, end: float = None, inc: float = 1.0):
    if end is None:
        end = start
        start = 0.0

    range_list: list[float] = []
    # value = start # + range_list.__len__() * inc == 0
    # while (inc > 0 and value >= end) or (inc < 0 and value <= end):
    #     range_list.append(value)
    #     value = start + range_list.__len__() * inc
    while 1:
        value = start + range_list.__len__() * inc
        if inc > 0 and value >= end or inc < 0 and value <= end:
            break
        range_list.append(value)
    return range_list


def ducks(const):
    interval = frange(0, 1000, 0.001)
    esp = esperance(const, interval)
    std_dev = standard_deviation(const, esp, interval)
    print("Average return time: %0.0fm %0.02ds" % divmod(esp * 60, 60))
    print("Standard deviation: %.3f" % std_dev)
    print("Time after which 50%% of the ducks are back: %dm %02ds" % divmod(time_back(const, 50) * 60, 60))
    print("Time after which 99%% of the ducks are back: %dm %02ds" % divmod(time_back(const, 99) * 60, 60))
    print("Percentage of ducks back after 1 minute: %.1f%%" % (percent_back(const, 1)))
    print("Percentage of ducks back after 2 minutes: %.1f%%" % percent_back(const, 2))


def print_help(ret):
    print("USAGE:\n\t./204ducks const\nDESCRIPTION:\n\tconst\tconstant (between 0 and 2.5)")
    exit(ret)


def main():
    if "-h" in sys.argv or "--help" in sys.argv or sys.argv.__len__() != 2:
        return print_help(84)
    try:
        const = float(sys.argv[1])
        if not 0 < const < 2.5:
            return print_help(84)
        ducks(const)
    except ValueError:
        print("Error in program")
        exit(84)


if __name__ == "__main__":
    main()
