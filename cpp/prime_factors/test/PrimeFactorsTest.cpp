#include "catch.hpp"
#include "PrimeFactor.h"
#include <vector>

TEST_CASE("Factors of 0 are nothing")
{
	REQUIRE(primeFactors(0) == std::vector<int>{});
}

TEST_CASE("Factors of 1 are nothing")
{
	REQUIRE(primeFactors(1) == std::vector<int>{});
}

TEST_CASE("Factors of 2 are 2")
{
	REQUIRE(primeFactors(2) == std::vector<int>{2});
}

TEST_CASE("Factors of 3 are 3")
{
	REQUIRE(primeFactors(3) == std::vector<int>{3});
}

TEST_CASE("Factors of 4 are 2, 2")
{
	REQUIRE(primeFactors(4) == std::vector<int>{2, 2});
}

TEST_CASE("Factors of 5 are 5")
{
	REQUIRE(primeFactors(5) == std::vector<int>{5});
}

TEST_CASE("Factors of 6 are 2, 3")
{
	REQUIRE(primeFactors(6) == std::vector<int>{2, 3});
}

TEST_CASE("Factors of 7 are 7")
{
	REQUIRE(primeFactors(7) == std::vector<int>{7});
}

TEST_CASE("Factors of 8 are 2, 2, 2")
{
	REQUIRE(primeFactors(8) == std::vector<int>{2, 2, 2});
}

TEST_CASE("Factors of 9 are 3, 3")
{
	REQUIRE(primeFactors(9) == std::vector<int>{3, 3});
}

TEST_CASE("Factors of 10 are 2, 5")
{
	REQUIRE(primeFactors(10) == std::vector<int>{2, 5});
}

TEST_CASE("Factors of 998 are 2, 499")
{
	REQUIRE(primeFactors(998) == std::vector<int>{2, 499});
}

TEST_CASE("Factors of 999 are 3, 3, 3, 37")
{
	REQUIRE(primeFactors(999) == std::vector<int>{3, 3, 3, 37});
}
