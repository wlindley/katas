#include <iostream>
#include <string>
#include "PrimeFactor.h"

int main(int argc, char *argv[])
{
	auto value = stoi(std::string(argv[1]));
	auto factors = primeFactors(value);
	std::cout << "Prime factors of " << value << ":";
	for (auto it = factors.begin(); it != factors.end(); it++)
		std::cout << " " << *it;
	std::cout << std::endl;
	return 0;
}
