ANTLR_EXEC=antlr
PACKAGE_NAME=calscript_lang
SOURCE_FILE=Calscript.g4
TARGET_LANG=Go

.PHONY: build
build:
		$(ANTLR_EXEC) -o $(PACKAGE_NAME) -package $(PACKAGE_NAME) -Dlanguage=$(TARGET_LANG) $(SOURCE_FILE)
