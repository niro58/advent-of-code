#include <iostream>
#include <algorithm>
#include <sstream>

using namespace std;
int GCD(int a, int b) {
    if (b == 0) return a;
    return GCD(b, a%b);
}
void FirstPart(const string& line, int &sum) {
    int gameNumber = -1;
    sscanf(line.c_str(), "Game %d", &gameNumber);


    int redCube = 0, greenCube = 0, blueCube = 0;

    string substr = line.substr(line.find(':') + 2);
    // get word by word
    istringstream iss(substr);
    int amount = 0;
    string temp;
    cout << "-------" << endl;
    while (iss >> temp) {
        if(isdigit(temp[0])){
            amount = stoi(temp);
            continue;
        }

        int *colorCubePtr = nullptr;
        int maxAmount = 0;

        switch (temp[0]) {
            case 'r':
                colorCubePtr = &redCube;
                maxAmount = 12;
                break;
            case 'g':
                colorCubePtr = &greenCube;
                maxAmount = 13;
                break;
            case 'b':
                colorCubePtr = &blueCube;
                maxAmount = 14;
                break;
            default:
                // Handle unexpected color
                break;
        }
        if(amount > maxAmount){
            redCube = 0;
            break;
        }
        if (colorCubePtr) {
            *colorCubePtr = amount;
        }

    }
    cout << "Game " << gameNumber << " : " << redCube << " red cubes, " << greenCube << " green cubes, " << blueCube << " blue cubes" << endl;
    if(redCube != 0 && greenCube != 0 && blueCube != 0){
        sum += gameNumber;
    }
}
void SecondPart(const string& line, int &sum) {
    int gameNumber = -1;
    sscanf(line.c_str(), "Game %d", &gameNumber);


    int redCube = 0, greenCube = 0, blueCube = 0;

    string substr = line.substr(line.find(':') + 2);
    // get word by word
    istringstream iss(substr);
    int amount = 0;
    string temp;
    cout << "-------" << endl;
    while (iss >> temp) {
        if(isdigit(temp[0])){
            amount = stoi(temp);
            continue;
        }

        int *colorCubePtr = nullptr;

        switch (temp[0]) {
            case 'r':
                colorCubePtr = &redCube;
                break;
            case 'g':
                colorCubePtr = &greenCube;
                break;
            case 'b':
                colorCubePtr = &blueCube;
                break;
            default:
                // Handle unexpected color
                break;
        }
        if (colorCubePtr) {
            if(*colorCubePtr < amount){
                *colorCubePtr = amount;
            }
        }

    }
    cout << "Game " << gameNumber << " : " << redCube << " red cubes, " << greenCube << " green cubes, " << blueCube << " blue cubes" << endl;
    sum += redCube * greenCube * blueCube;
}
int main() {
    // getline
    int sum = 0;
    string line;
    // lf 12 red cubes, 13 green cubes, 14 blue cubes;
    while (getline(cin, line)) {
        SecondPart(line, sum);
    }
    cout << "Sum : " << sum << endl;

    return 0;
}
