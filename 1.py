def main():
    num = int(input())
    print(num, end='')
    while 1:
        if num == 1:
            break
        elif num % 2 == 0:
            num //= 2
        else:
            num = num * 3 + 1
        print(' ', num, sep='', end='')
    print()


if __name__ == "__main__":
    main()