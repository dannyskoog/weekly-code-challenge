#/**
# * Code challenge week 46, 2021 (language: Python, playground: https://www.programiz.com/python-programming/online-compiler/)
# *
# * Questions:
# *     1. What would you replace the question mark (?) with in order to calculate and return the most common age of the population?
# */

class Person:
    def __init__(self, name, age):
        self.name = name
        self.age = age

class Population:
    def __init__(self, people):
        self.people = people

    def get_most_common_age(self):
        most_common_age = -1

        try:
            # ?

            return most_common_age
        except:
            print("An exception occurred")
            return most_common_age

def main():
    people = [
        Person("John", 56),
        Person("Jane", 10),
        Person("Eric", 37),
        Person("Sylvia", 23),
        Person("Rebecca", 24),
        Person("Michael", 43),
        Person("Tory", 37),
        Person("Adam", 61),
        Person("Elisabeth", 43),
        Person("Anna", 37),
    ]
    population = Population(people)
    most_common_age = population.get_most_common_age()

    print(most_common_age)

if __name__ == "__main__":
    main()
