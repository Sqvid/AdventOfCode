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

bool isNearSymbol(Scheme scheme, size_t i, size_t j) {
	if (i < 1 || i >= scheme.size() - 1 || j < 1 || j >= scheme.size() - 1) {
		return false;
	}

	bool ans {};
	auto obj = scheme[i][j];

	if (!isdigit(obj)) {
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
		if (isSymbol(ch)) {
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
			if (isNearSymbol(scheme, i, j)) {
				auto num = grabNumber(scheme[i], j);

				// Skip past this number.
				while (isdigit(scheme[i][j + 1])) {
					++j;
				}

				sum += num;
			}
		}
	}

	std::cout << sum << "\n";

	return 0;
}
