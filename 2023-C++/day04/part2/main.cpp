#include <fstream>
#include <iostream>
#include <regex>
#include <set>
#include <sstream>
#include <string>
#include <vector>

struct Card {
	public:
		Card() { id = 0; hits = 0; }
		int id;
		int hits;
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
				card.hits++;
			}
		}

		deck.push_back(card);
	}

	return deck;
}

int getTotalCount(Deck deck) {
	int sum {};
	std::vector<int> counts(deck.size(), 1);

	for (size_t i = 0; i < deck.size(); ++i) {
		auto hits = deck[i].hits;
		auto currCount = counts[i];

		for (int j = 1; j <= hits; ++j) {
			counts[i + j] += currCount;
		}
	}

	for (auto& c : counts) {
		sum += c;
	}

	return sum;
}

int main() {
	std::ifstream input("../input/input");
	if (!input.is_open()) {
		std::cerr << "Error: Failed to open input!\n";
		return -1;
	}

	auto deck = mkDeck(input);
	auto sum = getTotalCount(deck);

	std::cout << sum << "\n";

	return 0;
}
