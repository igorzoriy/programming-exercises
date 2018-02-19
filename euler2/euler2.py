fibList = [0, 1]
sum = 0

while fibList[-1] < 4000000:
    fibList.append(fibList[-2] + fibList[-1])
    if (fibList[-1] % 2 == 0):
        sum = sum + fibList[-1]

print(sum)
