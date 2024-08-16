#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "story.h"

#define CHAPTERS_DIR "text/"

void beginStory() {
    FILE *chapterFile;
    char chapterFilename[50];
    char buffer[256];

    int chapterNumber = 1;
    int paragraphNumber = 1;

    sprintf(chapterFilename, "%s%c%d.txt", CHAPTERS_DIR, 'c', chapterNumber);

    chapterFile = fopen(chapterFilename, "r");

    if (chapterFile == NULL) {
        printf("Chapter %d not found.\n", chapterNumber);
        return;
    }

    while (fgets(buffer, sizeof(buffer), chapterFile) != NULL) {
        if (strcmp(buffer, "\n") == 0) {
            printf("\n");
            paragraphNumber++;
            continue;
        }
        printf("%s", buffer);

        printf("\nEnter... ");
        fgets(buffer, sizeof(buffer), stdin);
        if (strcmp(buffer, "quit\n") == 0) {
            fclose(chapterFile);
            return;
        }
    }

    fclose(chapterFile);
}

