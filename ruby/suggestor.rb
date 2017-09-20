# frozen_string_literal: true

=begin

If we know what keys are near eachother 'on a keyboard, can we use that to
spellcheck, given a dictionary of valid words?

For a given input word, produce all the valid words that are 'close' to it, as
defined by keys that are close to other keys on the keyboard.

Assume the existence of methods:

* 'is_valid_word(word:string):boolean' that will tell you if a given string is
  a valid word in the dictionary

* 'nearby_chars(c:char):list<char>' that will give you all characters 'near' a
   give charcter, e.g. (r,t,y,f,g,h,v,b) are all adjacent to 'g' on a US
   keyboard layout

# "helol" => [ "hello", ... ]
# "hu" => [ "hi", ... ]

# for each character, grab list of ajacent chars
# [h, t, y...] [ u, i, y ...]
# h u
# h i
# h y
# t u
# t i
words[] valid_words(string) {
    array = new Array[string.length]
    suggestions =
    for(i=0;i<string.length;i++){
        array[i] = nearby_chars(string[i]) + string[i]
    }

    list = permutations(string.length, array)
    for(each: list) { is valid word(each)
    }
}

=end

def valid_words(string)
  all = combinations(string)
  all.select { |e| valid_word?(e) }
end

def combinations(string)
  words = []
  first = string[0]
  rest = string[1..-1]
  chars = nearby_characters(first) << first
  return chars if rest.empty?
  chars.to_a.each do |char|
    sub_words = combinations(rest)
    words << sub_words.map { |sub_word| char + sub_word }
  end
  words.flatten
end

def nearby_characters(char)
  {
    "h" => %w[t y u g j b n m],
    "u" => %w[y h j k i]
  }[char]
end

def valid_word?(string)
  %w[hi my].include?(string)
end
