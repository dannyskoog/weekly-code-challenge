class Car
  attr_reader :brand, :hybrid

  def initialize(brand, electric, fossil_fueled)
    @brand = brand
    @electric = electric
    @fossil_fueled = fossil_fueled
    @hybrid = electric && fossil_fueled
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

hybrid_cars = cars.select &:hybrid

puts "There are #{hybrid_cars.length} hybrid cars: #{hybrid_cars.map(&:brand).join(", ")}"
