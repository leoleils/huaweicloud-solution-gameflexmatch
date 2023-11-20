#include <iostream>
#include <stdlib.h>
#include <time.h>
using namespace std;

std::string random_string(int length)
{
    const char* charset =
        "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890";
    const int charset_size = 62;
    std::string result;
    std::srand(std::time(nullptr));
    for (int i = 0; i < length; i++)
    {
        result += charset[std::rand() % charset_size];
    }
    return result;
}
