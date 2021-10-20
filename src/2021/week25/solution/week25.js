const trampoline = fn => (...args) => {
    let result = fn(...args)
    while (typeof result === 'function') {
      result = result()
    }
    return result
  }

const sumBelow = (number, sum = 0) => (
    number === 0 
      ? sum
      : () => sumBelow(number - 1, sum + number)
  )

const sumBelowTrampolined = trampoline(sumBelow)

console.log(sumBelowTrampolined(100000))
