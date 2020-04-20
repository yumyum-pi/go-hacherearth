#!/bin/bash
# This scripts creates formated folder (date-name) of the problem with the given name and make the main.go file

NAME="" #name of the folder
folder=$(pwd)


FILENAME="main.go"

date=$(date '+%Y%m%d')

get_name() {
    #get the folder name
    echo "> Enter the folder name"
    read NAME

    #input formating
    #make the input lower case
    NAME=$( tr '[:upper:]' '[:lower:]' <<< $NAME) 
    #remove the white space
    NAME="${NAME// /_}" 
}

if [ -z "$1" ]
  then
    echo "No argument supplied"
fi

URL=$1

get_name

FOLDER="${date}-${NAME}"
echo -e "> The final folder name is: $FOLDER"

#make the file directory
FILEDIR="$folder/$FOLDER/$FILENAME"

#check if the directory does not exist
if [ ! -d "$FILEDIR" ]; then
    #make the folder
    mkdir "$folder/$FOLDER"
fi

# replace the add the url in the file
FILEDATA="// URL: $URL\npackage main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Happy Solving\")\n\n\tvar n int\t\t\t// first input\n\tfmt.Scanf(\"%d\", &n) // get the input value\n\tfmt.Println(n)\n}"
echo -e $FILEDATA > $FILEDIR #write the data to the file