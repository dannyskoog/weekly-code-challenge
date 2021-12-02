def get_unique_cars(carsA, carsB)
  carsADiff = carsA - carsB
  carsBDiff = carsB - carsA

  return carsADiff | carsBDiff
end

carsA = [{ brand: 'Toyota' }, { brand: 'BMW' }, { brand: 'Kia' }, { brand: 'Toyota' }]
carsB = [{ brand: 'Volkswagen' }, { brand: 'Ferrari' }, { brand: 'BMW' }]
uniqueCars = get_unique_cars(carsA, carsB)

puts "#{uniqueCars.length} unique car(s) found: #{uniqueCars}"
