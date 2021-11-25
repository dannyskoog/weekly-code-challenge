def memoize(func):
    cache = {}
    def wrapper(n):
        if n not in cache:
            cache[n] = func(n)
        return cache[n]
    return wrapper

@memoize
def fibonacci(n):
    def fib(n):
        if n in {0, 1}:
            return n
        return fib(n - 1) + fib(n - 2)
    return fib(n)

def main():
    n = 35

    num = fibonacci(n)
    print(f"The {n}th fibonacci number is {num}")

    num = fibonacci(n)
    print(f"The {n}th fibonacci number is {num}")

if __name__ == "__main__":
    main()
