    #include <stdbool.h>
    #include <stdio.h>
    #include <string.h>
    #include "location.h"
    #include "help.h"
    #include "story.h"
    bool parseAndExecute(char *input)
    {
       char *verb = strtok(input, " \n");
       char *noun = strtok(NULL, " \n");
       if (verb != NULL)
       {
          if (strcmp(verb, "quit") == 0)
          {
             return false;
          }
          else if (strcmp(verb, "look") == 0)
          {
             executeLook(noun);
          }
          else if (strcmp(verb, "go") == 0)
          {
             executeGo(noun);
          }
	  else if (strcmp(verb, "help") == 0) {
            displayHelpMenu();
	  }
	  else if (strcmp(verb, "begin") == 0) {
		  beginStory();
	  }
          else
          {
             printf("I don't know how to '%s'.\n", verb);
          }
       }
       return true;
    }

