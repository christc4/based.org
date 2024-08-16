# Lambda Syntax

Up: [Lambdas](lambdas)
Brother(s):
TARGET DECK

```Java
scene.setOnKeyPressed((e) -> {  
    logger.info("Button Pressed.");  
    if (e.getCode() == KeyCode.ESCAPE) {  
        game.gameQuitProperty().set(true);  
        Multimedia.playAudio("sounds/transition.wav");  
        gameWindow.startMenu();  
    } else if(e.getCode() == KeyCode.M) {  
        Multimedia.toggleAudio();  
    } else {  
        keyPressed(e);  
    }  
});
```



































#### Why:
#### How:









