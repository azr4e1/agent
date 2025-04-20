#!/usr/bin/env python3
"""
A script that decodes a ROT13-encoded message and prints it using a from-scratch implementation.
"""

def rot13_decode(encoded_string):
    """
    Decodes a ROT13-encoded string from scratch.
    
    ROT13 works by replacing each letter with the letter 13 positions after it in the alphabet,
    wrapping around if necessary. Since the English alphabet has 26 letters, applying ROT13 twice
    returns the original text.
    
    Args:
        encoded_string (str): The ROT13-encoded string to decode.
        
    Returns:
        str: The decoded string.
    """
    result = ""
    for char in encoded_string:
        # Check if character is an uppercase letter
        if 'A' <= char <= 'Z':
            # Apply ROT13 transformation (A=0, Z=25)
            # Subtract 'A' to get 0-25, add 13, mod 26 to wrap around, add 'A' to get back to ASCII
            decoded_char = chr(((ord(char) - ord('A') + 13) % 26) + ord('A'))
            result += decoded_char
        # Check if character is a lowercase letter
        elif 'a' <= char <= 'z':
            # Apply the same transformation for lowercase
            decoded_char = chr(((ord(char) - ord('a') + 13) % 26) + ord('a'))
            result += decoded_char
        else:
            # Non-alphabetic characters remain unchanged
            result += char
    
    return result

# The encoded message
encoded_message = 'Pbatenghyngvbaf ba ohvyqvat n pbqr-rqvgvat ntrag!'

# Decode and print the message
decoded_message = rot13_decode(encoded_message)
print(decoded_message)