#include <cctype>
#include <cstddef>
#include <fstream>
#include <iostream>
#include <string>
#include <vector>

typedef std::vector<std::vector<char>> Scheme;

Scheme getScheme(std::ifstream& input) {
	Scheme scheme;
	std::string line;

	while (std::getline(input, line)) {
		std::vector<char> engineRow;
		// Insert left sentinel column.
		engineRow.push_back('\n');

		for (char& ch : line) {
			engineRow.push_back(ch);
		}

		// Insert right sentinel column.
		engineRow.push_back('\n');

		// Insert top sentinel row.
		if (scheme.empty()) {
			scheme.push_back(std::vector<char>(engineRow.size(), '\n'));
		}
		scheme.push_back(engineRow);
	}

	// Insert bottom sentinel row.
	scheme.push_back(std::vector<char>(scheme[0].size(), '\n'));

	return scheme;
}

bool isSymbol(char ch) {
	bool ans {};

	if (std::ispunct(ch) && ch != '.') {
		ans = true;
	}

	return ans;
}

bool isNearDigit(Scheme scheme, size_t i, size_t j) {
	if (i < 1 || i >= scheme.size() - 1 || j < 1 || j >= scheme.size() - 1) {
		return false;
	}

	bool ans {};
	auto obj = scheme[i][j];

	if (!isSymbol(obj)) {
		return false;
	}

	// Vector of neighbouring cells.
	std::vector<char> adj;

	adj.push_back(scheme[i - 1][j - 1]);
	adj.push_back(scheme[i - 1][j]);
	adj.push_back(scheme[i - 1][j + 1]);
	adj.push_back(scheme[i][j - 1]);
	adj.push_back(scheme[i][j + 1]);
	adj.push_back(scheme[i + 1][j - 1]);
	adj.push_back(scheme[i + 1][j]);
	adj.push_back(scheme[i + 1][j + 1]);

	for (char& ch : adj) {
		if (std::isdigit(ch)) {
			ans = true;
			break;
		}
	}

	return ans;
}

int grabNumber(std::vector<char> row, size_t j) {
	// Invalid input.
	if (!isdigit(row[j]) || j < 1 || j >= row.size() - 1) {
		return -1;
	}

	std::string numString;
	auto pos {j};

	// Not at the start of a number.
	if (isdigit(row[j - 1])) {
		// Find the beginning of the number.
		while (pos > 0) {
			if (!isdigit(row[pos - 1])) {
				break;
			}

			--pos;
		}
	}

	// Create the number string.
	while (isdigit(row[pos])) {
		numString += row[pos];
		++pos;
	}

	return std::stoi(numString);
}

int main() {
	std::ifstream input("../input/input");
	if (!input.is_open()) {
		std::cerr << "Error: Could not open input!\n";
		return 1;
	}

	auto scheme = getScheme(input);
	int sum {};

	for (size_t i = 1; i < scheme.size() - 1; ++i) {
		for (size_t j = 1; j < scheme[i].size() - 1; ++j) {
			// On a potential gear.
			if (scheme[i][j] == '*') {
				std::vector<int> adjNums;

				// Scan the candidate gear's immediate neighbours.
				for (size_t p = i - 1; p <= i + 1; ++p) {
					for (size_t q = j - 1; q <= j + 1; ++q) {
						auto adj = scheme[p][q];

						if (isdigit(adj)) {
							adjNums.push_back(grabNumber(scheme[p], q));

							// Skip past digits that are a part of the same
							// number
							while (isdigit(scheme[p][q + 1])) {
								++q;
							}
						}
					}
				}

				// Not a true gear.
				if (adjNums.size() != 2) {
					continue;
				}

				int powerLevel {1};

				for (auto& num : adjNums) {
					powerLevel *= num;
				}

				sum += powerLevel;
			}
		}
	}

	std::cout << sum << "\n";

	return 0;
}
