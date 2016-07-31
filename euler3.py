from math import sqrt


def isPrime(number):
    iterator = range(2, number)
    for i in iterator:
        if number % i == 0:
            return False
    return True

n = 600851475143
factor = int(sqrt(n)) + 1

while True:
    if n % factor == 0 and isPrime(factor):
        break
    factor = factor - 1

print(factor)
