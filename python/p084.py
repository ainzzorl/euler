"""
Monopoly odds
https://projecteuler.net/problem=84

We construct the transition matrix:
matrix[i, j] is the probability to get to state i from state j after one roll.

Calculating the eigenvector corresponding to the largest eigenvalue (which must be equal to 1)
gives the state probabilities.

To simplify the problem and reduce the number of states, we assume that the position on the board
is the only state. So we ignore:
- The 3 doubles in a row rule.
- Ordering CC and CH.

The result is still correct.
"""

import math
import numpy as np
from numpy import linalg as LA, mat

GO = 0
G2J = 30
CCS = [2, 17, 33]
CHS = [7, 22, 36]
JAIL = 10
C1 = 11
E3 = 24
H2 = 39
R1 = 5

def nextu(idx):
    if idx == 7:
        return 12
    elif idx == 22:
        return 28
    elif idx == 36:
        return 12
    else:
        raise("Don't know next U for " + str(idx))

def nextr(idx):
    if idx == 7:
        return 15
    elif idx == 22:
        return 25
    elif idx == 36:
        return 5
    else:
        raise Exception("Don't know next R for " + str(idx))

def nextprobs(start, sides):
    res = np.zeros(40)
    rolls = [(i + 1, j + 1) for i in range(sides) for j in range(sides)]

    for (i, j) in rolls:
        rollsum = i + j
        baseprob = 1 / (sides * sides)

        # 3 doubles in a row
        # in this version, makes the result less precise
        # if i == j:
        #     res[JAIL] += 1 / (sides * sides) * baseprob
        #     baseprob *= 1 - 1 / (sides * sides)

        nxt = (start + rollsum) % 40

        if nxt == G2J:
            res[JAIL] += baseprob
        elif nxt in CCS:
            res[GO] += baseprob * 1 / 16
            res[JAIL] += baseprob * 1 / 16
            res[nxt] += baseprob * 14 / 16
        elif nxt in CHS:
            res[GO] += baseprob * 1 / 16
            res[JAIL] += baseprob * 1 / 16
            res[C1] += baseprob * 1 / 16
            res[E3] += baseprob * 1 / 16
            res[H2] += baseprob * 1 / 16
            res[R1] += baseprob * 1 / 16
            res[nextr(nxt)] += baseprob * 2 / 16
            res[nextu(nxt)] += baseprob * 1 / 16
            res[nxt - 3] += baseprob * 1 / 16
            res[nxt] += baseprob * 6 / 16
        else:
            res[nxt] += baseprob

    return res

def probabilities(sides):
    matrix = np.array([nextprobs(i, sides) for i in range(40)]).T
    vals, vecs = LA.eig(matrix)
    assert abs(np.absolute(vals[0]) - 1) < 0.001
    vec = np.absolute(vecs[:,0])
    vec /= sum(vec)
    return vec

def assertprob(expected, actual):
    assert abs(expected - actual) < 0.01, f"Expected prob={expected:.4f}, actual={actual:.4f}"

def modalstring(probs):
    withindex = [{'idx': i, 'val': v} for i, v in enumerate(probs)]
    std = sorted(withindex, key=lambda x: -x['val'])
    mx = std[:3]
    mxi = [v['idx'] for v in mx]
    mxs = [str(v) if v >= 10 else '0' + str(v) for v in mxi]
    return ''.join(mxs)

if __name__ == "__main__":
    prob6 = probabilities(6)
    # print(f"JAIL: {prob6[JAIL]:.4f}" )
    # print(f"GO: {prob6[GO]:.4f}" )
    assert abs(sum(prob6) - 1) < 0.001

    mod6 = modalstring(prob6)
    assert mod6 == '102400', f"Expected mod6=102400, actual: {mod6}"

    prob4 = probabilities(4)
    mod4 = modalstring(prob4)
    print(mod4)
