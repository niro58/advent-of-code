#include <iostream>
#include <vector>
#include <cstring>

using namespace std;
bool CheckAdjacentCellsValidity(vector<char*> board, int xPos, int yPos, int offset, int sum) {
    //cout << "--------------" << sum << "--------------" << endl;
    //cout << offset << endl;
    //cout << xPos - offset << " " << xPos + 1 << endl;
    for(int y = yPos - 1; y <= yPos + 1; y++){
        if(y < 0 || y >= board.size()){
            continue;
        }
        for(int x = xPos - offset - 1; x <= xPos + 1; x++){
            if(x < 0 || x >= strlen(board[y])){
                continue;
            }
            cout << "Checking " << y << " " << x << " Char: " << board[y][x] <<endl;
            if(board[y][x] != '.' && !isdigit(board[y][x])){
                cout << sum << " ! Valid at " << y << " " << x << " Char: " << board[y][x] <<endl;
                return true;
            }
        }
    }

    return false;
}
int main() {
    vector<char*> board; // Using a vector to store rows
    char tempLine[1000]; // Temporary array to store each line

    while (cin.getline(tempLine, 1000)) {
        int lineLength = strlen(tempLine);
        char* row = new char[lineLength + 1]; // +1 for the null terminator
        strcpy(row, tempLine);
        board.push_back(row);
    }
    int sum = 0;
    for(int y = 0; y < board.size(); y++){
        int tempNumber = 0;
        int length = 0;
        for(int x = 0; x < strlen(board[y]); x++){
            char cellValue = board[y][x];
            if(tempNumber != 0){
                length++;
            }
            if(isdigit(cellValue)){
                tempNumber *= 10;
                tempNumber += cellValue - '0';
            }
            if(tempNumber != 0 && (!isdigit(cellValue) || x == strlen (board[y]) - 1)){
                // check adjacent validity, input : board ,x, y and length of tempnumber
                if(CheckAdjacentCellsValidity(board, x, y, length, tempNumber)){
                    sum += tempNumber;
                }
                length = 0;
                tempNumber = 0;
            }
        }
    }
    cout << "--------------------------" << endl;
    for (char* row : board) {
        cout << row << endl;
    }
    for (char* row : board) {
        delete[] row;
    }
    cout << sum << endl;
    return 0;
}
