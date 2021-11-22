#/**
# * Code challenge week 47, 2021 (language: Python, playground: https://www.programiz.com/python-programming/online-compiler/)
# *
# * Description:
# *     The code below is supposed to calculate and print the nth fibonacci number. And for the sake of our own sanity we want to make sure
# *     that the algorithm function is pure (= returns the same result for the same input). Therefore we invoke the function twice with the
# *     same parameter. And "luckily" the results match. But we realise that there's one particular optimisation that we could do.
# *
# * Questions:
# *     1. What kind of optimisation do you think we are talking about?
# *     2. How could the code be adjusted to incorporate such optimisation?
# */

def fibonacci(n):
    if n in {0, 1}:
        return n
    return fibonacci(n - 1) + fibonacci(n - 2)

def main():
    n = 35

    num = fibonacci(n)
    print(f"The {n}th fibonacci number is {num}")

    num = fibonacci(n)
    print(f"The {n}th fibonacci number is {num}")

if __name__ == "__main__":
    main()
