![Pasted image 20240321140816.png](pasted_image_20240321140816.png)

Big Border pane, score, title, lives in top  (`HBox`?)

game grid in centre of course

score, level, `nextPiece` and piece storage on the right side, in a `VBox`

`ProgressBar` for timer at the bottom

**As for listeners:** The Game is listening for the secondary `gameBoards` to be clicked such as the one for rotation and the one for storing a piece for later.

If the rotation board is clicked, the game will be notified and the piece should be rotated.

If the "save a piece" board is clicked, the piece in that board is swapped with the one in the rotation board.


to make mine look like theirs, place the `vboxes` for the score and lives inside a `hbox` and make the `hbox` alignments left and right for score and lives respectively.

![Pasted image 20240326120432.png](pasted_image_20240326120432.png)Scores list is just a `HBox` with 2 `VBox`es inside with centre alignment. Light work ngl

**YOU FOOKIN LIAR IT WAS NOT LIGHT WORK**

![Pasted image 20240326230543.png](pasted_image_20240326230543.png)`textFlow` with `hGrow` property to always, value is bound to some `SimpleStringProperty` for the name, clears on enter/press submit, and empties the `textFlow`.

As for high score display: first off, it'll be the highest score in the high score file, if there are none or your current score beats the highest score it will then be bound to your current score.

# FIX YOUR CODEBASE IT'S ASS

![Pasted image 20240328163302.png](pasted_image_20240328163302.png)
![Pasted image 20240328163914.png](pasted_image_20240328163914.png)![Pasted image 20240329165846.png](pasted_image_20240329165846.png)
![Pasted image 20240329165942.png](pasted_image_20240329165942.png)Get rid of level and high score, replace with versus

text at the bottom is bound to the messages, and displays the most recent message
not a list of messages, but rather just a single one



![Pasted image 20240329183239.png](pasted_image_20240329183239.png)

In-Game Chat: Press T to send a chat message

![Pasted image 20240329194220.png](pasted_image_20240329194220.png)

I will have a queue that is populated by messages of added pieces

gives me producer-consumer vibes but not really? But yes really?

The aim is to produce at a fast rate, and have a large buffer of pieces ready to go, from the get-go and beyond.

The nature of the game ensures that you can't run through pieces without causing a delay on your end, so the communicator can send pieces off merrily.

25 pieces since that's the max that can be rapid-placed

faster than human speed fr, no seconds here, but milliseconds

![Pasted image 20240329195813.png](pasted_image_20240329195813.png)

whenever your score changes, send this

![Pasted image 20240329195826.png](pasted_image_20240329195826.png)

when you lose a life, send this


![Pasted image 20240329195855.png](pasted_image_20240329195855.png)

You already send this for aborting, also send this when you run out of lives

`SpawnPiece` kind of creates the "top of queue", and `nextPiece` pops that off and it becomes the `currentPiece` and creates a new "top of queue"

make an online scores scene for showing the multiplayer scores

**ACTUALLY** use `gameBlockCoordinates` and send the board thing too


My goal is to not start the game (timer included) UNTIL there are pieces in the queue.

Yes, I will need a listener on the state of the queue that kicks the game off ONLY when there are sufficient pieces in the queue.

The 'husk' - the empty shell of the game can go ahead and load, but the timer, game loop AND piece displays should NOT come to life until we have sufficient pieces.

Until now, I have neglected the importance of the pieces, they are a fundamental aspect of the game, **without them there is LITERALLY no game!** 


Timer bar: I can create the timer bar and its animation, but the animation itself will NOT go off until the game is ready.

game loop timer: I can create the `Timer` and the task, but the task will NOT start until the game is ready.

Piece boards: I can create them as empty boards that will have pieces drawn in them once the game is ready.

Game readiness: This relies on there being sufficient pieces for the game to be played while fetching more pieces.

My problem is that I have been trying to run the game and changing as little as possible from the offline game, which neglects the complexity of handling client-server communications.

You first thought this problem was out of your reach.
Had that been true, you would have found your level. 

To that I say...

# FUCK NO!
