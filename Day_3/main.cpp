#include <iostream>
#include <vector>
#include <map>
#include <algorithm>
#include <string>
#include <sstream>

using namespace std;

int yLen = 0;
int xLen = 0;

bool CheckAdjacentCells(char board[256][256], int yPos, int xPos) {
    for(int y = yPos - 1; y <= yPos + 1; y++) {
        for(int x = xPos - 1; x <= xPos + 1; x++) {
            bool isOutOfBounds = y < 0 || y >= yLen || x < 0 || x >= xLen;
            bool isValidCharacter = (board[y][x] < '0' || board[y][x] > '9') && board[y][x] != '.';
            if(isOutOfBounds){
                continue;
            }
            if(isValidCharacter) {
                return true;
            }
        }
    }
    return false;
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
        int temp = 0;
        bool isValidNumber = false;
        for(int x = 0; x < xLen; x++) {
            if(board[y][x] >= '0' && board[y][x] <= '9'){
                temp *= 10;
                temp += board[y][x] - '0';

                if(CheckAdjacentCells(board, y, x)){
                    isValidNumber = true;
                }
            }
            if(board[y][x] < '0' || board[y][x] > '9' || x == xLen - 1){
                if(temp > 0 && isValidNumber){
                    //remove it from board
                    for(int i = x - 1; i >= 0; i--) {
                        if(board[y][i] >= '0' && board[y][i] <= '9'){
                            board[y][i] = '.';
                        } else {
                            break;
                        }
                    }
                    result += temp;
                    isValidNumber = false;
                }
                temp = 0;

            }
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
