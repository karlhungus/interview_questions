def fizzbuzz(number)
  while number > 0
    string = ''
    string += "fizz" if number % 3 == 0
    string += "buzz" if number % 5 == 0
    puts("#{number}: " + string) unless string.empty?
    number -= 1
  end
end

x = ARGV&.first&.to_i
fizzbuzz(x || 10)
