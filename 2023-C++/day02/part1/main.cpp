#include <fstream>
#include <iostream>
#include <sstream>
#include <string>
#include <vector>

struct Draw {
	int red;
	int green;
	int blue;
};

struct GameInfo {
	int id;
	std::vector<Draw> draws;
};

GameInfo parseGameData(std::string data) {
	GameInfo info;

	auto colonPos = data.find(":");
	// Offset is set to 5 to skip the substring "Game ".
	info.id = std::stoi(data.substr(5, colonPos - 5));

	std::stringstream gameStream(data.substr(colonPos + 2));
	std::string setString;

	while (std::getline(gameStream, setString, ';')) {
		std::stringstream cubeStream(setString);
		std::string cubeString;
		Draw draw {};

		while (std::getline(cubeStream, cubeString, ',')) {
			int count {};
			std::string colour {};

			std::stringstream(cubeString) >> count >> colour;

			if (colour == "red") {
				draw.red += count;

			} else if (colour == "green") {
				draw.green += count;

			} else if (colour == "blue") {
				draw.blue += count;

			} else {
				std::cerr << "Parse Error! Expected a colour, got: " << colour
					<< "\n";
			}

		}

		info.draws.push_back(draw);
	}

	return info;
}

int main(void) {
	std::ifstream input("../input/input");

	if (!input.is_open()) {
		std::cerr << "Failed to open input file\n";
		return 1;
	}

	std::string line;
	int sum {};
	const int redLimit = 12;
	const int greenLimit = 13;
	const int blueLimit = 14;

	while (std::getline(input, line)) {
		auto info = parseGameData(line);
		bool possible = true;

		for (auto& draw : info.draws) {
			if (draw.red > redLimit || draw.green > greenLimit
				|| draw.blue > blueLimit) {

				possible = false;
				break;
			}
		}

		if (possible) {
			sum += info.id;
		}
	}

	std::cout << sum << "\n";

	return 0;
}
