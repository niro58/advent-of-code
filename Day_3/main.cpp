#include <iostream>
#include <vector>
#include <map>
#include <algorithm>
#include <string>
#include <sstream>

using namespace std;

int yLen = 0;
int xLen = 0;
struct Vector2Int{
    int y;
    int x;
    bool operator==(const Vector2Int& other) const {
        return x == other.x && y == other.y;
    }
    bool operator<(const Vector2Int& other) const {
        if(x == other.x) {
            return y < other.y;
        }
        return x < other.x;
    }
};
int GetAdjacentNumbers(char board[256][256], int yPos, int xPos) {
    int firstNumber = 0;
    int secondNumber = 0;
    map<Vector2Int, bool> visited;
    for(int y = yPos - 1; y <= yPos + 1; y++) {
        for(int x = xPos - 1; x <= xPos + 1; x++) {
            Vector2Int pos{y, x};
            if(visited.find(pos) != visited.end()) {
                continue;
            }
            if(y == yPos && x == xPos) {
                continue;
            }
            if(y < 0 || y >= yLen || x < 0 || x >= xLen) {
                continue;
            }
            if(board[y][x] < '0' || board[y][x] > '9') {
                continue;
            }

            int tempPos = x;
            int number = 0;
            while (board[y][tempPos] >= '0' && board[y][tempPos] <= '9') {
                tempPos--;
            }
            tempPos++;
            while (board[y][tempPos] >= '0' && board[y][tempPos] <= '9') {
                Vector2Int pos2{y, tempPos};
                visited[pos2] = true;
                number *= 10;
                number += board[y][tempPos] - '0';
                tempPos++;
            }
            cout << number << endl;
            if(number > 0){
                if(firstNumber == 0) {
                    firstNumber = number;
                } else if(secondNumber == 0) {
                    secondNumber = number;
                }else{
                    return 0;
                }
            }
            visited[pos] = true;

        }
    }
    if(firstNumber == 0 || secondNumber == 0) {
        return 0;
    }
    return firstNumber * secondNumber;
}
int main() {
    char board[256][256];
    string line;

    while (getline(cin, line)) {
        for(int i = 0; i < line.length(); i++) {
            board[yLen][i] = line[i];
            if(i > xLen){
                xLen = i;
            }
        }
        yLen++;
    }
    int result = 0;
    for(int y = 0; y<  yLen; y++) {
        for(int x = 0; x < xLen; x++) {
            if(board[y][x] != '*'){
                continue;
            }
            result += GetAdjacentNumbers(board, y, x);
        }
    }
    for(int y = 0; y < yLen; y++) {
        for(int x = 0; x < xLen; x++) {
            cout << board[y][x];
        }
        cout << endl;
    }
    cout << result << endl;
}
