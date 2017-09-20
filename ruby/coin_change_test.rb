# frozen_string_literal: true

require "bundler/setup"
require "minitest/autorun"
require "mocha/mini_test"
require "./coin_change.rb"

class CoinChangeTest < MiniTest::Test
  def test_make_change_greedy
    [
      { test: 33, expected: { 25 => 1, 10 => 0, 5 => 1, 1 => 3 } },
      { test: 41, expected: { 25 => 1, 10 => 1, 5 => 1, 1 => 1 } },
      { test: 91, expected: { 25 => 3, 10 => 1, 5 => 1, 1 => 1 } }
    ].each do |hash|
      assert_equal hash[:expected], make_change_greedy(hash[:test])
    end
  end

  def test_make_change_exhaustive
    [
      { test: 41, expected: { 25 => 0, 20 => 2, 10 => 0, 5 => 0, 1 => 1 } }
    ].each do |hash|
      assert_equal hash[:expected], make_change_exhaustive(hash[:test]), "for #{hash[:test]}"
    end
  end

  def test_make_change_dynamic
    [
      { test: 5, expected: { 1 => 0, 5 => 1, 10 => 0, 20 => 0, 25 => 0 } },
      { test: 6, expected: { 1 => 1, 5 => 1, 10 => 0, 20 => 0, 25 => 0 } },
      { test: 91, expected: { 25 => 2, 20 => 2, 10 => 0, 5 => 0, 1 => 1 } },
      { test: 41, expected: { 25 => 0, 20 => 2, 10 => 0, 5 => 0, 1 => 1 } }
    ].each do |hash|
      assert_equal hash[:expected], make_change_dynamic(hash[:test])
    end
  end
end
