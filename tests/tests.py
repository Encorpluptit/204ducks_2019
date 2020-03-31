#!/usr/bin/env python3
##
## EPITECH PROJECT, 2020
## 204ducks
## File description:
## pytests file
##

import pytest

try:
    module_ducks = __import__("204ducks")
    Duck = module_ducks.Duck
    main_fct = module_ducks.main
    EXIT_ERROR = 84


    @pytest.mark.parametrize(
        'param',
        [["-1"], ["3"], ["2.6"], ["a"], ["2.a"], ["-0.1"], ["1", "2"], ["-h"], [""]]
    )
    def test_bad_argv(param):
        with pytest.raises(SystemExit) as pytest_wrapped_e:
            main_fct(["204ducks"] + param)
        assert pytest_wrapped_e.type == SystemExit
        assert pytest_wrapped_e.value.code == EXIT_ERROR
except ModuleNotFoundError:
    from Ducks import Duck


@pytest.mark.parametrize(
    'a', [1, 2, 1.7, 0.3, 1, 1.4, 1.7, 2.3, 2.2, 2.65454, 2.7456, 0]
)
def test_ducks(a):
    p1 = Duck(a)
    assert p1 is not None


def calc_esperance(duck, xesp_m, xesp_s):
    esp_m, esp_s = divmod(round(duck.esp * 60), 60)
    assert esp_m == xesp_m
    assert esp_s == xesp_s


def calc_std_dev(duck, xstd_dev):
    assert round(duck.std_dev, 3) == xstd_dev


def calc_percent(duck, xper_5_m, xper_5_s, xper_9_m, xper_9_s):
    per_5_m, per_5_s = divmod(duck.time_back(duck.a, 50) * 60, 60)
    per_9_m, per_9_s = divmod(duck.time_back(duck.a, 99) * 60, 60)
    assert int(per_5_m) == xper_5_m
    assert int(per_5_s) == xper_5_s
    assert int(per_9_m) == xper_9_m
    assert int(per_9_s) == xper_9_s


def calc_time_back(duck, xtb_1, xtb_2):
    tb_1, tb_2 = duck.percent_back(duck.a, 1), duck.percent_back(duck.a, 2)
    assert xtb_1 == round(tb_1, 1)
    assert xtb_2 == round(tb_2, 1)


@pytest.mark.parametrize(
    'a, xesp_m, xesp_s',
    [
        (1.6, 1, 21),
        (1.2, 1, 12),
        (0.9, 1, 5),
        (0.2, 0, 50),
        (2.4, 1, 39),
        (2.1, 1, 32),
        (1.9, 1, 28),
        (1.5, 1, 19),
        (0.5, 0, 56),
    ]
)
def test_ducks_esperance(a, xesp_m, xesp_s):
    duck = Duck(a).run()
    calc_esperance(duck, xesp_m, xesp_s)


@pytest.mark.parametrize(
    'a, xstd_dev',
    [
        (1.6, 1.074),
        (1.2, 1.005),
        (0.9, 0.935),
        (0.2, 0.676),
        (2.4, 1.141),
        (2.1, 1.126),
        (1.9, 1.109),
        (1.5, 1.059),
        (0.5, 0.808),
    ]
)
def test_ducks_esperance(a, xstd_dev):
    duck = Duck(a).run()
    calc_std_dev(duck, xstd_dev)


@pytest.mark.parametrize(
    'a, xper_5_m, xper_5_s, xper_9_m, xper_9_s',
    [
        (0.9, 0, 49, 4, 30),
        (0.2, 0, 39, 3, 16),
        (1.5, 1, 1, 5, 0),
        (1.6, 1, 4, 5, 4),
        (1.2, 0, 55, 4, 47),
        (0.5, 0, 43, 3, 57),
        (2.4, 1, 23, 5, 28),
        (2.1, 1, 16, 5, 21),
        (1.9, 1, 11, 5, 15),
    ]
)
def test_ducks_percent(a, xper_5_m, xper_5_s, xper_9_m, xper_9_s):
    duck = Duck(a).run()
    calc_percent(duck, xper_5_m, xper_5_s, xper_9_m, xper_9_s)


@pytest.mark.parametrize(
    'a, xtb_1, xtb_2',
    [
        (0.2, 71.3, 94.2),
        (2.4, 33.0, 70.4),
        (2.1, 38.2, 73.7),
        (1.6, 46.9, 79.1),
        (1.2, 53.9, 83.4),
        (0.9, 59.1, 86.6),
        (1.9, 41.7, 75.8),
        (1.5, 48.6, 80.2),
        (0.5, 66.0, 91.0),
    ]
)
def test_ducks_time_back(a, xtb_1, xtb_2):
    duck = Duck(a).run()
    calc_time_back(duck, xtb_1, xtb_2)


@pytest.mark.parametrize(
    'a, xesp_m, xesp_s, xstd_dev, xper_5_m, xper_5_s, xper_9_m, xper_9_s, xtb_1, xtb_2',
    [
        (0.5, 0, 56, 0.808, 0, 43, 3, 57, 66.0, 91.0),
        (2.1, 1, 32, 1.126, 1, 16, 5, 21, 38.2, 73.7),
        (1.9, 1, 28, 1.109, 1, 11, 5, 15, 41.7, 75.8),
        (1.5, 1, 19, 1.059, 1, 1, 5, 0, 48.6, 80.2),
        (1.6, 1, 21, 1.074, 1, 4, 5, 4, 46.9, 79.1),
        (1.2, 1, 12, 1.005, 0, 55, 4, 47, 53.9, 83.4),
        (0.9, 1, 5, 0.935, 0, 49, 4, 30, 59.1, 86.6),
        (0.2, 0, 50, 0.676, 0, 39, 3, 16, 71.3, 94.2),
        (2.4, 1, 39, 1.141, 1, 23, 5, 28, 33.0, 70.4),
    ]
)
def test_ducks_full(a, xesp_m, xesp_s, xstd_dev, xper_5_m, xper_5_s, xper_9_m, xper_9_s, xtb_1, xtb_2):
    duck = Duck(a).run()
    calc_esperance(duck, xesp_m, xesp_s)
    calc_std_dev(duck, xstd_dev)
    calc_percent(duck, xper_5_m, xper_5_s, xper_9_m, xper_9_s)
    calc_time_back(duck, xtb_1, xtb_2)
