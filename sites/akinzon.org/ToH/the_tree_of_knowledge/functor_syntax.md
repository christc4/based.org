# Functor Syntax

Up: [Functors](functors)
Brother(s):
TARGET DECK

```Java
EventHandler<ActionEvent> event = new EventHandler<ActionEvent>() {  
    String message="Anonymous\n";  
    public void handle(ActionEvent actionEvent) {  
        textArea.appendText(message);  
        textArea2.appendText(message);  
        String temp = button1.getText();  
        button1.setText(button2.getText());  
        button2.setText(temp);  
    }  
};
```



































#### Why:
#### How:









