 # Known Issues & Edge Cases

Tested the CLI manually and found a few things to fix in a future version:

- **Raw error leak:** Entering letters (e.g. `abc`) prints the raw `strconv` error. Should replace with a cleaner user message.
- **No re-prompt on bad input:** If you type an invalid menu option like `99`, it just prints an error and exits instead of asking again.
- **No Absolute Zero check:** Entering `-300` Celsius works mathematically, but ignores the laws of physics (Absolute Zero is -273.15°C).
- **No loop:** The app exits after one conversion instead of letting the user run another.

- **