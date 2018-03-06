#include "PrimeFactor.h"

std::vector<int> primeFactors(int number)
{
	auto factors = std::vector<int>{};
	if (number >= 2 && number % 2 == 0)
	{
		factors.push_back(2);
		number /= 2;
	}
	if (2 <= number)
		factors.push_back(number);
	return factors;
}
