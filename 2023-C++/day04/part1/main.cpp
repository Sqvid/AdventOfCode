#include <fstream>
#include <iostream>
#include <regex>
#include <set>
#include <sstream>
#include <string>
#include <vector>

struct Card {
	public:
		Card() { id = 0; value = 0; }
		int id;
		int value;
};

typedef std::vector<Card> Deck;

Deck mkDeck(std::ifstream& input) {
	Deck deck;
	// Captures the id, winning numbers, and pool numbers.
	std::regex cardPattern(R"(Card\s+(\d+):\s+((?:\d+\s+)+)\|\s+((?:\d+\s*)+))");
	std::string line;

	while (std::getline(input, line)) {
		std::smatch idMatches;
		if(!std::regex_match(line, idMatches, cardPattern)) {
			std::cerr << "Error: Regex match failed!\n";
		}

		Card card;
		card.id = std::stoi(idMatches[1].str());

		auto winStream = std::istringstream(idMatches[2].str());
		auto poolStream = std::istringstream(idMatches[3].str());

		std::set<std::string> wins;

		// Load winning numbers.
		std::string winStr;
		while (winStream >> winStr) {
			wins.insert(winStr);
		}

		// Load pool numbers.
		std::string poolStr;
		while (poolStream >> poolStr) {
			if (wins.count(poolStr)) {
				if (card.value == 0) {
					card.value = 1;

				} else {
					card.value *= 2;
				}
			}
		}

		deck.push_back(card);
	}

	return deck;
}

int main() {
	std::ifstream input("../input/input");
	if (!input.is_open()) {
		std::cerr << "Error: Failed to open input!\n";
		return -1;
	}

	int sum {};
	auto deck = mkDeck(input);

	for (auto& c : deck) {
		std::cout << "ID: " << c.id << "; " << "Value: " << c.value << "\n";
		sum += c.value;
	}

	std::cout << "Sum: " << sum << "\n";

	return 0;
}
