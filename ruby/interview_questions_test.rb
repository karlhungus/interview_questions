# frozen_string_literal: true

require "bundler/setup"
require "minitest/autorun"
require "mocha/mini_test"
require "./interview_questions.rb"

class InterviewQuestionsTest < MiniTest::Test
  def test_fizzbuzz
    [
      { test: 4, expected: "3: fizz\n" },
      { test: 15, expected: "15: fizzbuzz\n12: fizz\n10: buzz\n9: fizz\n6: fizz\n5: buzz\n3: fizz\n" }
    ].each do |hash|
      assert_equal hash[:expected], fizzbuzz(hash[:test])
    end
  end
end
