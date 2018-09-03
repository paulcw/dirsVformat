# dirsVformat

A utility to reformat the output of "dirs -v".
It makes it easier for me to use the directory stack in Bash.
Your milage may vary.

Compile it, put it in your path, then you can define an alias like this:

	alias dl='dirs -v | dirsVformat'

## What does it do?

- it reads the output of `dirs -v`
- it sorts the output so that things closer to the current directory come first
- it colorizes the output, alternating between cyan and yellow, to make it easier to pick out which directory is which.

## Is it really worth making a repo for this?

Probably not but it makes it easier to find it next time I set up
a new environment.
