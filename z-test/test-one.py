# Vulnerable: Processes user input directly without sanitization
user_input = input("Enter a Python expression: ")
eval(user_input)
eval(user_input)