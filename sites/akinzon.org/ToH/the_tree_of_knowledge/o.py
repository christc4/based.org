import os
import re

# Get the current working directory
directory = os.getcwd()

# Regular expression pattern to match the Obsidian links
obsidian_link_pattern = re.compile(r'\[\[([^\]]+)\]\]')

# Function to convert Obsidian links to traditional markdown links
def convert_link(match):
    original_text = match.group(1)
    # Convert to lowercase, replace spaces with underscores
    link_text = original_text.lower().replace(' ', '_')
    print(f"Converting: [[{original_text}]] to [{original_text}]({link_text})")  # Debugging output
    return f'[{original_text}]({link_text})'

# Iterate through all markdown files in the current directory
for filename in os.listdir(directory):
    if filename.endswith('.md'):
        filepath = os.path.join(directory, filename)
        print(f"Processing file: {filename}")  # Debugging output

        with open(filepath, 'r') as file:
            content = file.read()

        # Replace Obsidian links with traditional markdown links
        new_content = re.sub(obsidian_link_pattern, convert_link, content)

        if new_content != content:
            print(f"Changes detected in file: {filename}, writing changes...")  # Debugging output
            # Write the modified content back to the file
            with open(filepath, 'w') as file:
                file.write(new_content)
        else:
            print(f"No changes needed for file: {filename}")  # Debugging output

print("Conversion complete.")

