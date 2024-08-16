<pre><code>
#!/bin/sh

# Define the directory path
directory="/var/www/werc/sites"

# Navigate to the directory
cd "$directory" || exit 1

# Execute git commands
git add . && git commit -m "New pages written" && git push -u origin main

exit 0
</code></pre>
