#include "StringCalculator.h"
#include <vector>
#include <list>
#include <iostream>
#include <sstream>
#include <utility>

typedef std::string Token;
typedef std::list<Token> TokenList;
typedef std::vector<char> DelimeterList;

const DelimeterList defaultDelimeters = {',', '\n'};
TokenList tokenize(std::string input, std::vector<char> delimeters);
int firstIndex(std::string input, std::vector<char> delimeters);
std::pair<DelimeterList, std::string> extractDelimeters(std::string input);

int add(std::string input)
{
	DelimeterList delimeters;
	std::tie(delimeters, input) = extractDelimeters(input);
	auto numbers = tokenize(input, delimeters);
	auto total = 0;
	for (auto cur = numbers.begin(); cur != numbers.end(); cur++)
	{
		if (cur->size() > 0)
			total += stoi(*cur);
	}
	return total;
}

std::pair<DelimeterList, std::string> extractDelimeters(std::string input)
{
	if (0 != input.find("//"))
		return std::make_pair(defaultDelimeters, input);
	auto lineSep = input.find('\n');
	auto delimeters = DelimeterList{input[2]};
	return std::make_pair(delimeters, input.substr(lineSep+1));
}

int firstIndex(std::string input, std::vector<char> delimeters)
{
	auto index = input.size();
	for (auto delimeter = delimeters.begin(); delimeter != delimeters.end(); delimeter++)
	{
		auto newIndex = input.find(*delimeter);
		if (std::string::npos != newIndex && newIndex < index)
			index = newIndex;
	}
	return index;
}

TokenList tokenize(std::string input, std::vector<char> delimeters)
{
	auto index = firstIndex(input, delimeters);
	if (input.size() == index)
		return TokenList{input};
	auto rest = tokenize(input.substr(index + 1), delimeters);
	rest.emplace_front(input.substr(0, index));
	return rest;
}
