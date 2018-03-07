#include "PrimeFactor.h"

std::vector<int> primeFactors(int number)
{
	auto factors = std::vector<int>{};
	for (auto i = 2; i < number; i++)
	{
		while (number >= 2 && number % i == 0)
		{
			factors.push_back(i);
			number /= i;
		}
	}
	if (number >= 2)
		factors.push_back(number);
	return factors;
}
