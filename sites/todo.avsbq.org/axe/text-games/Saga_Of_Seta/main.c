#include <stdio.h>
#include <stdbool.h>

// Include your location.h header file
#include "location.h"

// Declare the parseAndExecute function (defined in parsexec.c)
extern bool parseAndExecute(char *input);

int main() {
    char input[100]; // Assuming a maximum input length of 100 characters

printf("┌──────────────────────────────────────────────┐\n");
printf("│                  The Saga of Seta            │\n");
printf("│  Type 'quit' to exit, or 'help' at any point │\n");
printf("│  in the story for basic commands.            │\n");
printf("│  You are free to look around your            │\n");
printf("│  surroundings before we begin the adventure, │\n");
printf("│  so as to familiarise yourself with the      │\n");
printf("│  Kingdom of Harkness.                        │\n");
printf("└──────────────────────────────────────────────┘\n");

    bool running = true;
    while (running) {
        printf(">_ ");
        fgets(input, sizeof(input), stdin); // Read input from the user
        running = parseAndExecute(input); // Parse and execute the input
    }

    printf("Goodbye!\n");

    return 0;
}

