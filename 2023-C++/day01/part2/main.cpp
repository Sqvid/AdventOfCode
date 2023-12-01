#include <fstream>
#include <ios>
#include <iostream>
#include <string>
#include <cctype>

char strToDigit(std::string str) {
	char digit;

	if (str == "zero") {
		digit = '0';

	} else if (str == "one") {
		digit = '1';

	} else if (str == "two") {
		digit = '2';

	} else if (str == "three") {
		digit = '3';

	} else if (str == "four") {
		digit = '4';

	} else if (str == "five") {
		digit = '5';

	} else if (str == "six") {
		digit = '6';

	} else if (str == "seven") {
		digit = '7';

	} else if (str == "eight") {
		digit = '8';

	} else if (str == "nine") {
		digit = '9';

	} else {
		digit = '\0';
	}

	return digit;
}

int getFirstLastDigit(std::string str) {
	std::string digitStr {};

	for (int i = 0; i < str.length(); ++i) {
		char c = str[i];

		if (std::isdigit(c)) {
			digitStr += c;

		} else {
			// All digit strings: "one", "two", ..., "nine" are 3, 4, or 5
			// characters in length. Therefore search these many subsequent
			// characters and look for a match, and skip the index ahead when a
			// match is found.
			char candidate1 = strToDigit(str.substr(i, 3));
			char candidate2 = strToDigit(str.substr(i, 4));
			char candidate3 = strToDigit(str.substr(i, 5));

			if (std::isdigit(candidate1)) {
				digitStr += candidate1;
				i += 1;
			} else if (std::isdigit(candidate2)) {
				digitStr += candidate2;
				i += 2;
			} else if (std::isdigit(candidate3)) {
				digitStr += candidate3;
				i += 3;
			}
		}
	}

	char digit1 {}, digit2 {};
	digit1 = digitStr[0];

	if (digitStr.length() == 1) {
		digit2 = digit1;

	} else {
		digit2 = digitStr[digitStr.length() - 1];
	}

	return std::stoi(std::string({digit1, digit2}));
}

int main(void) {
	std::ifstream input("../input/input");

	if (!input.is_open()) {
		std::cerr << "Failed to open input file\n";
		return 1;
	}

	std::string line;
	int sum {};

	while (std::getline(input, line)) {
		sum += getFirstLastDigit(line);
	}

	std::cout << sum << "\n";

	return 0;
}
