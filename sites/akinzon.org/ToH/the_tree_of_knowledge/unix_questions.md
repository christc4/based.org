# UNIX Questions

Up: [UNIX](unix)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester II::1204 Data Management::UNIX

What are 3 frequently used folders in the UNIX filesystem? #flashcard 
`/etc` stores system config files
`/var/log` stores log outputs for system programs
`/bin` and `/usr/bin` store frequently used programs
<!--ID: 1716641026139-->



What are Index Nodes in UNIX? What are its 2 benefits? #flashcard 
Index nodes (Inodes) are the data structure used to store file metadata.
Keeping metadata separate from the actual data allows for the file to be renamed or deleted even when another program is using it.
It is also easy to move files between folders.
<!--ID: 1716641026148-->


What command is used to find the current directory (absolute path)? What command can be used to switch between directories? #flashcard 
`pwd`
`cd`
<!--ID: 1716641026156-->


What command is used to list all files and folders in the current directory? What 5 arguments can be run with it? #flashcard 
`ls`
`-l` for long list
`-t` sort by modification time
`-S` sort by size
`-h` list file size in a human-readable way
`-r` reverse the order
<!--ID: 1716641026163-->



Is everything a file in UNIX? Give an example of some bizarre 'files'. #flashcard 
Yes! a directory is a file, a keyboard is a file (standard input) and the monitor is one too! (standard output.)
<!--ID: 1716641026171-->


How can you get help on UNIX commands? (2 ways) #flashcard 
`man` - for manual
`info` - for more command info
<!--ID: 1716641026179-->


How are directories made in Unix? How are they removed? How are files made in Unix? How are they removed? How are files and directories copied? How are they moved? #flashcard 
`mkdir [options] <directory name>` - make a directory
`rmdir [options] <directory name>` - remove a directory
`touch [options] <filename>` - make a file
`rm [options] <filename>` - remove a file
`cp [options] <source> <destination>` - copy a file/directory to another place
`mv [options] <source> <destination>` - copy a file/directory to another place
<!--ID: 1716641026187-->



How are files read in their entirety in Unix? How are they read page-by-page? How are the first `n` lines read? What about the last `n` lines? #flashcard 
`cat <filename>` - read whole file
`less <filename>` - read file page by page, `[space]` to go forward one page, `[b]` to go back a page
`head -n[lines] <filename>` to read the first `[lines]` lines from a file
`tail -n[lines] <filename>` to read the last `[lines]` lines from a file
<!--ID: 1716641026197-->

What are the three multi's of an operating system #flashcard 
Multi-user: supports multiple users on the same computer
Multi-processing: can use multiple [The Processor|processors](the_processor|processors) at the same time
Multi-tasking: can run multiple processes at a time
<!--ID: 1716661710077-->

What are [Version Control Systems](version_control_systems)? What 3 things do they allow you to do? #flashcard 
Version control systems are systems that record all changes made to files so that specific versions can be recalled later.
They allow you to:
roll-back some files/whole project to a specific state
compare changes over time
see who introduced something that may be causing a problem
<!--ID: 1716661915727-->



What are [Local VCSes](local_vcses)? What is a drawback? #flashcard 
Local version control systems are a local database that keeps all changes to files as patch sets. The database is stored on your computer, so it can be inefficient to use in a collaborative context.
<!--ID: 1716662171649-->


What are [Centralised VCSes](centralised_vcses)? What are its advantages and disadvantages? #flashcard 
Centralised version control systems use a single server to contain all the files of different versions, and clients take the files they need from there.
Advantage: Facilitates collaboration, easier for everyone to see what everyone else is doing
Disadvantage: Single centralised server is a single point of failure. If that goes down, **it's a stick of a bicky one still...**
<!--ID: 1716662171661-->

What are [Distributed VCSes](distributed_vcses)? What is its benefit? #flashcard 
Distributed version control systems are similar to [Centralised VCSes](centralised_vcses) but is distributed in the way that clients don't just take the files they need, but **the whole codebase** as well.
This way, if the server goes down, progress can be pieced back together by the clients.
<!--ID: 1716662293230-->


In the context of version control systems, what are branches? What are conflicts? What is reverse integration? What is forward integration? #flashcard 
Branches are versions of a project that are separate from the project that allow for testing new features in isolation, without affecting an already stable codebase.
Conflicts arise when we try to merge a branch with the main project but there are differences that cannot be solved automatically so they need to be solved manually.
Reverse integration is when branch features are integrated into the main project.
Forward integration is when features in the main project not present in branches is integrated into branches.
<!--ID: 1716662611145-->

Explain the three (+ 1) [Piping](piping) operators #flashcard 
`a | b` - the output of `a` becomes the input of `b`
`a > b` - the output of `a` is written to `b`
`a >> b` - the output of `a` is appended to `b`
`a < b` - the input of `a` is taken from `b`
<!--ID: 1716663026279-->

What does these filters do? `sort`, `wc`, `uniq` #flashcard 
`sort` - sorts the data
`wc` - can print out a count of words `-w`, lines `-l`, and characters `-m`
`uniq` - gets rid of duplicate data
<!--ID: 1716663590980-->

What does `cut -d` do? #flashcard 
`cut -d` separates on a delimiter, like the `split()` operation on [strings](strings).
<!--ID: 1716663684584-->

How are currently running **user** processes viewed in Unix? what about **all** CPU processes? How to kill these processes? What are two common options for killing processes? #flashcard 
user processes - `ps`
all processes - `top`
`kill [options] <pid>` - kill a process using its process id `pid`
`SIGTERM` - kill signal can be ignored (give him grace, that's why I made Mr. Morale)
`SIGKILL` - kill signal cannot be ignored (force kill)
<!--ID: 1716666294432-->


How does `screen` work in Unix? What are some common arguments? #flashcard 
`screen` allows you to run a process even when the computer is logged off.
`-d` - detach from existing screen (like minimise in windows)
`-list` - list existing screens
`-r` return to existing screen
`ctrl+D` - kill a screen
<!--ID: 1716666294443-->


How is `^` used in grep? how is $ used in grep? How is `{ }` used in grep? #flashcard 
`^<seq>` matches strings that begin with `<seq>`.
`[^0-9]` matches strings not in the range of `[0-9]`.
`<seq>`$ matches string that end with `<seq>`.
`<seq>{lb, ub}` matches strings where `<seq>` appears a specified number of times, between `lb` and `ub`. `ub` can be left blank to imply no upper bound.
<!--ID: 1716667201648-->

How to replace the `tunde` in `Akintunde` with `yemi`? How would you do this for all instances? How would you append `Adeoye` to `Akinyemi`? How would you delete from line `n` to the end of a file? #flashcard 
`echo "Akintunde" | sed 's/tunde/yemi/'`
`echo "Akintunde" | sed 's/tunde/yemi/g'`
`echo Akinyemi | sed '/Akinyemi/a\Adeoye'`
`sed 'n,$d' input`
<!--ID: 1716667461213-->


What are the 4 Unix wild-card symbols? #flashcard 
`*` any character 0 or more times
`?` - any character exactly once
`[]` exactly one character from a range of characters (`[a-d]` shorthand `[abcd]` longhand)
`{}` (IN UNIX) exactly one word from a range of words
<!--ID: 1716668715851-->

How are `chmod` numerical positions split up? #flashcard 
`user|group|others`, and for each partition: `read write execute` where `read` is `4 (1 0 0)`, `write` is `2 (0 1 0)` and `execute` is `1 (0 0 1)`  
<!--ID: 1716669205209-->

















#### Why:
#### How:









