"""
Odd period square roots
https://projecteuler.net/problem=64
"""

import math


def gcd(x, y):
    while y != 0:
        x, y = y, x % y
    return x


def continued_fraction(x):
    """
    Returns a tuple of (pre_period, period)
    """

    a, b, c = 1, 0, 1

    chain = []
    seen = {}

    i = 0
    while True:
        # (a * sqrt(x) + b) / c -> r + 1 / ((a_next * sqrt(x) + b_next) / c_next)
        # Iterate until we find the period.

        r = int(math.floor((a * math.sqrt(x) + b) / c))
        a, b, c = a * c, r * c * c - b * c, a * a * x - (b - r * c)**2
        if c == 0:
            assert x == r * r
            return [r], []
        d = gcd(gcd(a, b), c)
        a, b, c = int(a / d), int(b / d), int(c / d)

        #print(f"r={r} a={a} b={b} c={c}")
        chain.append(r)
        i += 1

        key = f"{a}.{b}.{c}"
        if key in seen:
            period_start = seen[key]
            return chain[:period_start], chain[period_start:]
        else:
            seen[key] = i


if __name__ == "__main__":
    assert continued_fraction(2) == ([1], [2])
    assert continued_fraction(3) == ([1], [1, 2])
    assert continued_fraction(4) == ([2], [])
    assert continued_fraction(5) == ([2], [4])
    assert continued_fraction(6) == ([2], [2, 4])
    assert continued_fraction(7) == ([2], [1, 1, 1, 4])
    assert continued_fraction(8) == ([2], [1, 4])
    assert continued_fraction(9) == ([3], [])
    assert continued_fraction(10) == ([3], [6])
    assert continued_fraction(11) == ([3], [3, 6])
    assert continued_fraction(12) == ([3], [2, 6])
    assert continued_fraction(13) == ([3], [1, 1, 1, 1, 6])
    assert continued_fraction(23) == ([4], [1, 3, 1, 8])

    result = 0
    for i in range(2, 10000 + 1):
        _, period = continued_fraction(i)
        if len(period) % 2 == 1:
            result += 1
    print(result)
