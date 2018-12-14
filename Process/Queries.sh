#!/bin/sh

echo What are you searching for?

read search_term
echo You are interesting in knowing more about processes involving $search_term

echo Here are the process steps involving $search_term:

grep -B 6 -A 4 $search_term * | grep process | sed -e 's/<process>/ /g'

# replace the star with a user input directory to target the search to a
# specific set of files
