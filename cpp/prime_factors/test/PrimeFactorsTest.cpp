#include "catch.hpp"
#include "PrimeFactor.h"
#include <vector>

void verify(int number, std::vector<int> factors)
{
	REQUIRE(primeFactors(number) == factors);
}

TEST_CASE("Factors of 0 are nothing")
{
	verify(0, {});
}

TEST_CASE("Factors of 1 are nothing")
{
	verify(1, {});
}

TEST_CASE("Factors of 2 are 2")
{
	verify(2, {2});
}

TEST_CASE("Factors of 3 are 3")
{
	verify(3, {3});
}

TEST_CASE("Factors of 4 are 2, 2")
{
	verify(4, {2, 2});
}

TEST_CASE("Factors of 5 are 5")
{
	verify(5, {5});
}

TEST_CASE("Factors of 6 are 2, 3")
{
	verify(6, {2, 3});
}

TEST_CASE("Factors of 7 are 7")
{
	verify(7, {7});
}

TEST_CASE("Factors of 8 are 2, 2, 2")
{
	verify(8, {2, 2, 2});
}

TEST_CASE("Factors of 9 are 3, 3")
{
	verify(9, {3, 3});
}

TEST_CASE("Factors of 10 are 2, 5")
{
	verify(10, {2, 5});
}

TEST_CASE("Factors of 998 are 2, 499")
{
	verify(998, {2, 499});
}

TEST_CASE("Factors of 999 are 3, 3, 3, 37")
{
	verify(999, {3, 3, 3, 37});
}
