#include <iostream>
#include <map>
#include <sstream>
#include <vector>
#include <algorithm>
using namespace std;

int main() {
    map<int,int> scratchCards;
    string line;
    int index = 0;
    int sum = 0;
    while(getline(cin,line)) {
        scratchCards[index] += 1;

        string allNumbersRaw = line.substr(line.find(':') + 1, line.length());
        string winningNumbersRaw = allNumbersRaw.substr(0, allNumbersRaw.find('|') - 1);
        string matchingNumbersRaw = allNumbersRaw.substr(allNumbersRaw.find('|') + 1, allNumbersRaw.length());

        stringstream winningNumbers;
        stringstream matchingNumbers;

        vector<int> winningNumbersList;
        winningNumbers << winningNumbersRaw;
        matchingNumbers << matchingNumbersRaw;
        string temp;
        while (winningNumbers >> temp) {
            winningNumbersList.push_back(stoi(temp));
        }
        int matches = 0;

        while (matchingNumbers >> temp) {
            int num = stoi(temp);
            if(find(winningNumbersList.begin(), winningNumbersList.end(), num) != winningNumbersList.end()) {
                matches++;
                scratchCards[index + matches] += scratchCards[index];
            }
        }
        sum += scratchCards[index];
        //print matches
        cout << "Repeats : " << scratchCards[index] << endl;
        cout << "Matches : " << matches << endl;
        cout << "____________________" << endl;

        index++;
    }
    cout << "Total matches : " << sum << endl;
    return 0;
}