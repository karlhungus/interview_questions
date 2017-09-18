# frozen_string_literal: true

def fizzbuzz(number)
  result = ""
  while number.positive?
    string = ""
    string += "fizz" if (number % 3).zero?
    string += "buzz" if (number % 5).zero?
    result += "#{number}: #{string}\n" unless string.empty?
    number -= 1
  end
  result
end

x = ARGV&.first&.to_i
exec = ARGV&.last == "-x"
puts fizzbuzz(x || 10) if exec
