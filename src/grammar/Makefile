ANTLR_EXEC=antlr
PACKAGE_NAME=calscript_lang
SOURCE_FILE=Calscript.g4
TARGET_LANG=Go

all: clean build 

.PHONY: build
build:
		$(ANTLR_EXEC) -o ../../target/$(PACKAGE_NAME) -package $(PACKAGE_NAME) -Dlanguage=$(TARGET_LANG) $(SOURCE_FILE)

.PHONY: clean
clean:
		rm -rf ../../target
