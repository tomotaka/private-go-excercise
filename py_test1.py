#!/usr/bin/python
# -*- coding: utf-8 -*-

def is_prime(n):
    i = 2
    while i < n:
        if n % i == 0:
            return False
        i += 1
    return True

count = 1
i = 2
n_max = 100000
while i <= n_max:
    if is_prime(i):
        count += 1
    i += 1

print '1-%d: %d' % (n_max, count)

