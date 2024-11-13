# Go User Password Profiler

**GuPP** (Go User Password Profiler) is a utility to generate custom wordlists for password cracking based on common information. Inspired by the [CUPP](https://github.com/Mebus/cupp) project, it allows you to create wordlists based on a user’s first name, last name, pet name, company name, and more!

## 📸 Demo
[Demo](https://github.com/user-attachments/assets/9a548211-9ce4-4746-81ab-d3c78b53231b)


## 🚀 Features

- **Custom Wordlist Generation:** Generate password wordlists based on personal data (names, nicknames, companies, etc.).
- **Concurrency:** Uses Go's powerful goroutines to speed up wordlist generation, making it faster and more efficient.
- **Character Replacements:** Automatically generates variations using common character replacements like `a → @`, `o → 0`, `s → $`, and more.
- **Flexible Custom Keywords:** Add custom keywords to the wordlist for more variations.
- **Unique Entries:** Ensures no duplicate entries in the generated wordlist.

## 📥 Installation

### Prerequisites

- Go 1.18+ (or higher) installed on your system.

### Install
    git clone https://github.com/yourusername/gupp.git
    cd gupp
    go build -o gupp
    sudo mv gupp /usr/local/bin

## 🔧 Usage

1. **Personal Information Input:**  
   You'll be prompted to enter various information regarding the "target", including:
    - First Name
    - Last Name
    - Nickname
    - Partner's First Name
    - Child's Name
    - Pet's Name
    - Company Name
    - Custom Keywords (Optional)
