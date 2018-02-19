def isPalindrome(number):
    number = str(number)
    length = len(number)
    halfLen = int(length / 2)

    if (length == 1):
        return False

    i = 0
    while i < halfLen:
        if number[i] != number[length - 1 - i]:
            return False
        i = i + 1
    return True

max = 999
result = 0

for i in range(max, 1, -1):
    for j in range(max, i, -1):
        r = i * j
        if isPalindrome(r) and r > result:
            result = r

print(result)
