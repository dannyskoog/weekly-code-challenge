#/**
# * Code challenge week 49, 2021 (language: Ruby, playground: https://replit.com/languages/ruby)
# *
# * Description:
# *     The code below is supposed to filter out hybrid cars and print them out. However, the implementation is unfinished.
# *
# * Questions:
# *     1. What would you replace the question mark (?) with in order to find out which cars that are hybrids?
# */

class Car
  attr_reader :brand, :hybrid

  def initialize(brand, electric, fossil_fueled)
    @brand = brand
    @electric = electric
    @fossil_fueled = fossil_fueled
    @hybrid = electric and fossil_fueled
  end
end

cars = [
  Car.new("Volvo", true, false),
  Car.new("Saab", true, true),
  Car.new("Audi", false, true),
  Car.new("Volkswagen", false, true),
  Car.new("BMW", true, true),
  Car.new("Kia", true, false),
]

hybrid_cars = []

# ?

puts "There are #{hybrid_cars.length} hybrid cars: #{hybrid_cars.map(&:brand).join(", ")}"
