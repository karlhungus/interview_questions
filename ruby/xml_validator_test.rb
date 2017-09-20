# frozen_string_literal: true

require "bundler/setup"
require "minitest/autorun"
require "mocha/mini_test"
require "./xml_validator.rb"

class XMLValidatorTest < MiniTest::Test
  def test_well_formed
    test = <<-XML
      <continent name="Europe">
        <country name="Netherlands">
        </country>
        <country name="France">
        </country>
      </continent>
    XML
    assert well_formed(test)
  end

  def test_non_well_formed
    test = <<-XML
      <continent name="Europe">
        <country name="Netherlands">
        </continent>
      </country>
    XML
    refute well_formed(test)
  end

  def test_parse_tags
    assert_equal(
      [Tag.new("a"), Tag.new("b"), Tag.new("a", false), Tag.new("b", false), Tag.new("c", false)],
      parse_tags("<a><b></a></b></c>")
    )
  end
end
