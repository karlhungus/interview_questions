# frozen_string_literal: true

INFINITY_PROXY = 999_999_999

# Make Change for a set of denominations
def make_change_greedy(total, denominations = [25, 10, 5, 1])
  tmp_total = total
  change = {}
  denominations.each do |d|
    amount = (tmp_total / d).to_i
    tmp_total -= (d * amount)
    change[d] = amount
  end
  change
end

# rubocop:disable Metrics/AbcSize, Metrics/MethodLength
def make_change_exhaustive(total, denominations = [25, 20, 10, 5, 1])
  solution = {}
  denominations.each { |den| solution[den] = 0 }
  return solution if total.zero?
  solution[denominations.first] = INFINITY_PROXY + total
  denominations.each do |den|
    next unless total >= den
    sub_solution = make_change_exhaustive((total - den), denominations)
    if sub_solution.sum { |_, v| v } + 1 < solution.sum { |_, v| v }
      solution = sub_solution
      solution[den] = solution[den] + 1
    end
  end
  solution
end

# stolen/fixed from http://www.cs.ubc.ca/~liorma/cpsc320/files/coin-changing-1x2.pdf (second min coins assignment was to
# -0
def make_change_dynamic(total, denominations = [1, 5, 10, 20, 25])
  min_coins = Array.new(total + 1)
  best_choice = Array.new(total + 1)
  min_coins[0] = 0
  best_choice[0] = -1
  (1..total).each do |m|
    min_coins[m] = INFINITY_PROXY + total
    denominations.each_with_index do |den, i|
      next unless m >= den
      if min_coins[m - den] + 1 < min_coins[m]
        min_coins[m] = min_coins[m - den] + 1
        best_choice[m] = i
      end
    end
  end

  determine_solution(best_choice, total, denominations)
end

def determine_solution(best_choice, total, denominations)
  solution = {}
  denominations.each { |den| solution[den] = 0 }
  tmp_total = total
  while tmp_total.positive?
    denom = denominations[best_choice[tmp_total]]
    solution[denom] = solution[denom] + 1
    tmp_total -= denom
  end
  solution
end
