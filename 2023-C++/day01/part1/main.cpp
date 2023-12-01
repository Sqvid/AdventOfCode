#include <fstream>
#include <iostream>
#include <string>
#include <cctype>

bool isDigitSet(char digit) {
	return digit != '\0';
}

int numFromDigits(char digit1, char digit2) {
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
		char digit1 {}, digit2 {};

		for (char &c : line) {
			if (std::isdigit(c)) {
				// Overwrite the last digit in the string.
				if (isDigitSet(digit1)) {
					digit2 = c;

				// Set the first digit in the string.
				} else {
					digit1 = c;
				}
			}
		}

		// If only one digit was found in the string, duplicate it.
		if (!isDigitSet(digit2)) {
			digit2 = digit1;
		}

		sum += numFromDigits(digit1, digit2);
	}

	std::cout << sum << "\n";

	return 0;
}
