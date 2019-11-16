# Calscript Language

Calscript is a simple, intuitive and natural language to define calendar events.

## How to build
### Required dependencies

* Install `make`
* Install `antlr`
    ```
    curl https://www.antlr.org/download/antlr-4.7.2-complete.jar > antlr-4.7.2-complete.jar
    alias antlr='java -jar antlr-4.7.2-complete.jar'
    ```    
### Compiling lexer and parser
    ```
    cd grammar && make
    ```

## Language syntax