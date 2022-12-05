Output demo.gif

# Set up a 1200x600 terminal with 46px font.
Set FontSize 24
Set Width 1200
Set Height 600

# Type a command in the terminal.
Type "cd presentations/demo_folder/"
Sleep 500ms
Enter

Type "dignore list"
Sleep 500ms
Enter
Sleep 2s
Enter

Type 'echo "exclude_dir/" >> .dockerignore'
Sleep 200ms
Enter
Type "dignore list"
Sleep 500ms
Enter
Sleep 2s
Enter

Type 'echo "!**/*include*" >> .dockerignore'
Sleep 200ms
Enter
Type "dignore list"
Sleep 500ms
Enter
Sleep 2s
Enter

Type "dignore list --excluded"
Sleep 500ms
Enter

# Admire the output for a bit.
Sleep 8s
