# frozen_string_literal: true

require "bundler/setup"
require "minitest/autorun"
require "mocha/mini_test"
require "./suggestor.rb"

class SuggestorTest < MiniTest::Test
  def test_suggestor
    [
      { test: "hu", expected: %w[my hi] }
    ].each do |hash|
      assert_equal hash[:expected], valid_words(hash[:test])
    end
  end
end
