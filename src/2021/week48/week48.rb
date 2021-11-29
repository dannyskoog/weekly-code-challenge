#/**
# * Code challenge week 48, 2021 (language: Ruby, playground: https://replit.com/languages/ruby)
# *
# * Description:
# *     The code below is supposed to print the unique cars (by brand) between two lists of cars. And a car is considered unique if it doesn't
# *     exist in the other list. So what about cars that occur multiple times in the same list? Well, that is fine with the only disclaimer
# *     of that we don't want it counted multiple times in case of it qualifying for being unique (between the lists).
# *
# * Questions:
# *     1. What would you replace the question mark (?) with in order to find out what cars that are unique?
# */

def get_unique_cars(carsA, carsB)
  uniqueCars = []
  
  # ?

  return uniqueCars
end

carsA = [{ brand: 'Toyota' }, { brand: 'BMW' }, { brand: 'Kia' }, { brand: 'Toyota' }]
carsB = [{ brand: 'Volkswagen' }, { brand: 'Ferrari' }, { brand: 'BMW' }]
uniqueCars = get_unique_cars(carsA, carsB)

puts "#{uniqueCars.length} unique car(s) found: #{uniqueCars}"
