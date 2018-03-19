#include "catch.hpp"
#include "StringCalculator.h"

TEST_CASE("Empty string returns 0")
{
	REQUIRE(0 == add(""));
}

TEST_CASE("Single digit string returns that number")
{
	REQUIRE(1 == add("1"));
	REQUIRE(2 == add("2"));
	REQUIRE(5 == add("5"));
}

TEST_CASE("Two numbers separated by commas are added together")
{
	REQUIRE(2 == add("1,1"));
	REQUIRE(3 == add("1,2"));
	REQUIRE(9 == add("7,2"));
}

TEST_CASE("Can handle arbitrary numbers of numbers")
{
	REQUIRE(3 == add("1,1,1"));
	REQUIRE(4 == add("1,1,1,1"));
	REQUIRE(5 == add("1,1,1,1,1"));
}

TEST_CASE("Allows newlines between numbers")
{
	REQUIRE(3 == add("1\n2"));
	REQUIRE(5 == add("2,2\n1"));
}

TEST_CASE("Allows custom delimeters")
{
	REQUIRE(4 == add("//;\n1;2;1"));
	REQUIRE(5 == add("// \n2 3"));
}
